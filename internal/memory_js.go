//go:build js

package internal

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall/js"
	"time"
	"unsafe"
)

type Pointer int32

var (
	stackAlloc   js.Value
	stackString  js.Value
	utf8String   js.Value
	getValue     js.Value
	setValue     js.Value
	stackSave    js.Value
	stackRestore js.Value

	bigInt  js.Value
	boolean js.Value

	// SDL callbacks run on the goroutine of the call that triggered them,
	// so stack frames nest instead of needing a lock.
	savedStackPtrs []js.Value
)

func init() {
	// Wait for runtime initialization
	for js.Global().Get("runtimeInitialized").IsUndefined() {
		time.Sleep(1 * time.Millisecond)
	}
	for !js.Global().Get("runtimeInitialized").Bool() {
		time.Sleep(1 * time.Millisecond)
	}
	// TODO: find a cleaner way than loop+sleep?
	for js.Global().Get("HEAPU8").IsUndefined() {
		time.Sleep(1 * time.Millisecond)
	}

	stackAlloc = js.Global().Get("stackAlloc")
	stackString = js.Global().Get("stringToUTF8OnStack")
	utf8String = js.Global().Get("UTF8ToString")
	getValue = js.Global().Get("getValue")
	setValue = js.Global().Get("setValue")
	stackSave = js.Global().Get("stackSave")
	stackRestore = js.Global().Get("stackRestore")

	bigInt = js.Global().Get("BigInt")
	boolean = js.Global().Get("Boolean")
}

// heapU8 must not be cached: memory growth replaces HEAPU8 and detaches
// the previous view.
func heapU8() js.Value {
	return js.Global().Get("HEAPU8")
}

// On js, Go pointers and uintptr values always hold Go addresses: Go and
// emscripten have separate linear memories, and a heap address stored in a
// Go pointer would confuse the garbage collector. Heap memory is reached
// through a registered Go stand-in instead.
type object struct {
	// ptr is the Go stand-in, kept alive by the registry.
	ptr        unsafe.Pointer
	key        heapKey
	finalizers []func()
}

type heapKey struct {
	addr int
	typ  reflect.Type
}

var (
	objectsByGoAddr   = map[uintptr]*object{}
	objectsByHeapAddr = map[heapKey]*object{}
	// pinned maps Go memory to a heap copy freed when the Go memory is
	// collected. Keys are not references, so pins do not keep it alive.
	pinned = map[uintptr]int{}
)

func lookupHeap(goAddr uintptr) (int, bool) {
	if obj, ok := objectsByGoAddr[goAddr]; ok {
		return obj.key.addr, true
	}
	addr, ok := pinned[goAddr]
	return addr, ok
}

// newObject returns the Go stand-in of the t value at addr, refreshed from
// the heap. The same addr and t always yield the same Go pointer.
func newObject(t reflect.Type, addr int) unsafe.Pointer {
	if addr == 0 {
		return nil
	}
	key := heapKey{addr, t}
	if obj, ok := objectsByHeapAddr[key]; ok {
		decodeFromHeap(obj.ptr, t, addr)
		return obj.ptr
	}
	var p unsafe.Pointer
	if t.Size() == 0 {
		// Zero-size allocations share one address.
		p = unsafe.Pointer(new(byte))
	} else {
		p = reflect.New(t).UnsafePointer()
	}
	obj := &object{ptr: p, key: key}
	// Registered before decoding so that cyclic C data terminates.
	objectsByGoAddr[uintptr(p)] = obj
	objectsByHeapAddr[key] = obj
	decodeFromHeap(p, t, addr)
	return p
}

func pin[T any](c []T) {
	t := reflect.TypeFor[T]()
	l := layoutOf(t)
	allocs := &heapAllocs{}
	buf := make([]byte, int(l.size)*len(c))
	for i := range c {
		encodeValue(buf[uintptr(i)*l.size:], unsafe.Pointer(&c[i]), t, allocs.alloc)
	}
	addr := allocs.alloc(len(buf))
	writeHeap(addr, buf)
	key := uintptr(unsafe.Pointer(&c[0]))
	pinned[key] = addr
	runtime.SetFinalizer(&c[0], func(*T) {
		delete(pinned, key)
		allocs.free()
	})
}

// pinnable returns a slice of length n that owns its allocation, as
// runtime.SetFinalizer requires; the tiny allocator would share it.
func pinnable[T any](n int) []T {
	var t T
	return make([]T, n, max(n, 16/max(int(unsafe.Sizeof(t)), 1)))
}

// SliceData returns a pointer to the first element of s that C can
// dereference: a Go copy whose heap mirror lives as long as the copy.
func SliceData[T any](s []T) *T {
	if len(s) == 0 {
		return nil
	}
	c := pinnable[T](len(s))
	copy(c, s)
	pin(c)
	return &c[0]
}

func StackSave() {
	savedStackPtrs = append(savedStackPtrs, stackSave.Invoke())
}

func StackRestore() {
	last := len(savedStackPtrs) - 1
	stackRestore.Invoke(savedStackPtrs[last])
	savedStackPtrs = savedStackPtrs[:last]
}

func StackAlloc(n int) js.Value {
	return stackAlloc.Invoke(int32(n))
}

// CloneObjectToJSStack copies obj to the stack in C layout. It must be
// called between StackSave and StackRestore.
func CloneObjectToJSStack[T any](obj *T) js.Value {
	if obj == nil {
		return js.Null()
	}
	if addr, ok := lookupHeap(uintptr(unsafe.Pointer(obj))); ok {
		return js.ValueOf(addr)
	}
	return js.ValueOf(int(encodePointer(unsafe.Pointer(obj), reflect.TypeFor[T](), stackAllocator)))
}

// CopyJSToObject reads the C value at ptr into obj.
func CopyJSToObject[T any](obj *T, ptr js.Value) {
	decodeFromHeap(unsafe.Pointer(obj), reflect.TypeFor[T](), ptr.Int())
}

func CloneByteSliceToJSHeap(s []byte) js.Value {
	ptr := js.Global().Call("_malloc", len(s))
	writeHeap(ptr.Int(), s)

	return ptr
}

// ClonePtrArrayToJSHeap copies count elements starting at ptr to the heap
// in C layout. The returned function frees the copy.
func ClonePtrArrayToJSHeap[T any](ptr *T, count int) (js.Value, func()) {
	if ptr == nil || count <= 0 {
		return js.Null(), func() {}
	}
	t := reflect.TypeFor[T]()
	l := layoutOf(t)
	allocs := &heapAllocs{}
	buf := make([]byte, int(l.size)*count)
	for i, v := range unsafe.Slice(ptr, count) {
		encodeValue(buf[uintptr(i)*l.size:], unsafe.Pointer(&v), t, allocs.alloc)
	}
	addr := allocs.alloc(len(buf))
	writeHeap(addr, buf)

	return js.ValueOf(addr), allocs.free
}

func StringOnJSStack(str string) js.Value {
	return stackString.Invoke(str)
}

func UTF8JSToString(ptr js.Value) string {
	return utf8String.Invoke(ptr).String()
}

func jsValueType(typ string) string {
	if strings.HasPrefix(typ, "*") {
		return "*"
	}
	switch typ {
	case "i1", "i8", "i16", "i32", "i64", "float", "double":
		return typ
	case "f32":
		return "float"
	case "f64":
		return "double"
	default:
		panic("unknown incoming js type")
	}
}

func GetValue(ptr js.Value, typ string) js.Value {
	return getValue.Invoke(ptr, jsValueType(typ))
}

func SetValue(ptr js.Value, value js.Value, typ string) {
	setValue.Invoke(ptr, value, jsValueType(typ))
}

func GetJSPointerFromUintptr(ptr uintptr) (js.Value, bool) {
	if addr, ok := lookupHeap(ptr); ok {
		return js.ValueOf(addr), true
	}
	return js.Null(), false
}

func GetJSPointer[T any](obj *T) (js.Value, bool) {
	return GetJSPointerFromUintptr(uintptr(unsafe.Pointer(obj)))
}

// NewObject returns the Go stand-in for the T at ptr, or nil for NULL.
func NewObject[T any](ptr js.Value) *T {
	if ptr.IsNull() || ptr.IsUndefined() {
		return nil
	}
	return (*T)(newObject(reflect.TypeFor[T](), ptr.Int()))
}

// NewUintptr returns the Go stand-in address for the C pointer ptr, or 0
// for NULL.
func NewUintptr(ptr js.Value) uintptr {
	return uintptr(unsafe.Pointer(NewObject[byte](ptr)))
}

func DeleteJSPointer(ptr uintptr) {
	obj, ok := objectsByGoAddr[ptr]
	if !ok {
		return
	}
	delete(objectsByGoAddr, ptr)
	delete(objectsByHeapAddr, obj.key)
	for _, fn := range obj.finalizers {
		fn()
	}
}

// AttachFinalizer runs fn when obj is unregistered by DeleteJSPointer.
func AttachFinalizer[T any](obj *T, fn func()) {
	v, ok := objectsByGoAddr[uintptr(unsafe.Pointer(obj))]
	if !ok {
		return
	}
	v.finalizers = append(v.finalizers, fn)
}

// NewBigInt converts n to a BigInt through its decimal form: a float64
// Number loses precision above 2^53.
func NewBigInt(n any) js.Value {
	switch v := n.(type) {
	case int:
		return bigInt.Invoke(strconv.FormatInt(int64(v), 10))
	case int32:
		return bigInt.Invoke(strconv.FormatInt(int64(v), 10))
	case int64:
		return bigInt.Invoke(strconv.FormatInt(v, 10))
	case uint:
		return bigInt.Invoke(strconv.FormatUint(uint64(v), 10))
	case uint32:
		return bigInt.Invoke(strconv.FormatUint(uint64(v), 10))
	case uint64:
		return bigInt.Invoke(strconv.FormatUint(v, 10))
	case uintptr:
		return bigInt.Invoke(strconv.FormatUint(uint64(v), 10))
	default:
		return bigInt.Invoke(n)
	}
}

func NewBoolean(b bool) js.Value {
	return boolean.Invoke(b)
}

func GetBool(v js.Value) bool {
	return boolean.Invoke(v).Bool()
}

func GetInt64(v js.Value) int64 {
	str := js.Global().Get("String").Invoke(v).String()
	n, _ := strconv.ParseInt(str, 10, 64)

	return n
}

// GetUint64 also accepts negative values: wasm passes Uint64 results as
// signed i64 BigInts.
func GetUint64(v js.Value) uint64 {
	str := js.Global().Get("String").Invoke(v).String()
	if n, err := strconv.ParseInt(str, 10, 64); err == nil {
		return uint64(n)
	}
	n, _ := strconv.ParseUint(str, 10, 64)

	return n
}

func GetByteSliceFromJSPtr(ptr js.Value, count int) []byte {
	s := make([]byte, count)
	readHeap(ptr.Int(), s)

	return s
}

type Numeric interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func GetNumericSliceFromJSPtr[T Numeric](ptr js.Value, count int) []T {
	s := make([]T, count)
	bytesCount := count * int(unsafe.Sizeof(T(0)))
	readHeap(ptr.Int(), unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), bytesCount))

	return s
}

// SetNumericSliceToJSPtr writes s to the heap at ptr.
func SetNumericSliceToJSPtr[T Numeric](ptr js.Value, s []T) {
	bytesCount := len(s) * int(unsafe.Sizeof(T(0)))
	writeHeap(ptr.Int(), unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), bytesCount))
}

// GetObjectSliceFromJSPtr reads count T pointers at ptr.
func GetObjectSliceFromJSPtr[T any](ptr js.Value, count int) []*T {
	ptrs := GetNumericSliceFromJSPtr[int32](ptr, count)
	ret := make([]*T, count)
	for i := range ptrs {
		ret[i] = NewObject[T](js.ValueOf(ptrs[i]))
	}

	return ret
}

// NewCallback registers fn in the wasm function table and returns its
// index, which C uses as a function pointer. sig uses emscripten's
// addFunction encoding: v void, i int32/bool, j int64 (BigInt), f float,
// d double, p pointer; return type first.
func NewCallback(sig string, fn func(args []js.Value) any) uintptr {
	f := js.FuncOf(func(_ js.Value, args []js.Value) any {
		return fn(args)
	})
	return uintptr(js.Global().Get("Module").Call("addFunction", f, sig).Int())
}
