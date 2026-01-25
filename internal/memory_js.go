//go:build js

package internal

import (
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
	"time"
	"unsafe"
)

type Pointer = int32

var (
	heapU8 js.Value

	stackAlloc   js.Value
	stackString  js.Value
	utf8String   js.Value
	getValue     js.Value
	setValue     js.Value
	stackSave    js.Value
	stackRestore js.Value

	bigInt  js.Value
	boolean js.Value

	lastStackPtr js.Value

	lock sync.Mutex
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

	heapU8 = js.Global().Get("HEAPU8")

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

type object struct {
	data       []byte
	value      js.Value
	finalizers []func()
}

func (o *object) Map(size uintptr) {
	o.data = slices.Grow(o.data, int(size))[:size]
	o.finalizers = nil
	o.value = js.Value{}
	for i := range size {
		o.data[i] = 0
	}
}

func (o *object) ExtractSelf() {
	arr := heapU8.Call("slice", o.value.Int(), o.value.Int()+len(o.data))
	js.CopyBytesToGo(o.data, arr)
}

var (
	jsPtrsByObject = map[uintptr]*object{}
	pool           = sync.Pool{
		New: func() any {
			return new(object)
		},
	}
)

func StackSave() {
	lock.Lock()
	lastStackPtr = stackSave.Invoke()
}

func StackRestore() {
	stackRestore.Invoke(lastStackPtr)
	lock.Unlock()
}

func StackAlloc(n int) js.Value {
	return stackAlloc.Invoke(int32(n))
}

func CloneObjectToJSStack[T any](obj *T) js.Value {
	if obj == nil {
		return js.Null()
	}
	size := unsafe.Sizeof(*obj)
	ptr := stackAlloc.Invoke(int32(size))
	arr := heapU8.Call("subarray", ptr.Int(), ptr.Int()+int(size))
	js.CopyBytesToJS(arr, unsafe.Slice((*byte)(unsafe.Pointer(obj)), size))

	return ptr
}

func CopyJSToObject[T any](obj *T, ptr js.Value) {
	size := unsafe.Sizeof(*obj)
	arr := heapU8.Call("slice", ptr.Int(), ptr.Int()+int(size))
	js.CopyBytesToGo(unsafe.Slice((*byte)(unsafe.Pointer(obj)), size), arr)
}

func CloneByteSliceToJSHeap(s []byte) js.Value {
	ptr := js.Global().Call("_malloc", len(s))
	arr := heapU8.Call("subarray", ptr.Int(), ptr.Int()+len(s))
	js.CopyBytesToJS(arr, s)

	return ptr
}

func ClonePtrArrayToJSHeap[T any](ptr *T, count int) (js.Value, func()) {
	if ptr == nil || count <= 0 {
		return js.Null(), func() {}
	}
	size := int(count) * int(unsafe.Sizeof(*ptr))
	s := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), size)
	arr := CloneByteSliceToJSHeap(s)
	runtime.KeepAlive(s)
	return arr, func() {
		js.Global().Call("_free", arr)
	}
}

func StringOnJSStack(str string) js.Value {
	return stackString.Invoke(str)
}

func UTF8JSToString(ptr js.Value) string {
	return utf8String.Invoke(ptr).String()
}

func GetValue(ptr js.Value, typ string) js.Value {
	if strings.HasPrefix(typ, "*") {
		typ = "*"
	}
	switch typ {
	case "i1", "i8", "i16", "i32", "*", "float", "double":
		return getValue.Invoke(ptr, typ)
	default:
		panic("unknown incoming js type")
	}
}

func SetValue(ptr js.Value, value js.Value, typ string) {
	if strings.HasPrefix(typ, "*") {
		typ = "*"
	}
	switch typ {
	case "i1", "i8", "i16", "i32", "*", "float", "double":
		setValue.Invoke(ptr, value, typ)
	default:
		panic("unknown incoming js type")
	}
}

func GetJSPointerFromUintptr(ptr uintptr) (js.Value, bool) {
	v, ok := jsPtrsByObject[ptr]
	if ok {
		return v.value, true
	}
	return js.Null(), false
}

func GetJSPointer[T any](obj *T) (js.Value, bool) {
	v, ok := jsPtrsByObject[uintptr(unsafe.Pointer(obj))]
	if ok {
		return v.value, true
	}
	return js.Null(), false
}

func NewObject[T any](ptr js.Value) *T {
	var t T

	obj := pool.Get().(*object)
	obj.Map(max(unsafe.Sizeof(t), 1))
	obj.value = ptr
	obj.ExtractSelf()
	objPtr := unsafe.Pointer(&obj.data[0])
	// If an object already exists for this js pointer, return it
	// and do not add finalizer
	if _, ok := jsPtrsByObject[uintptr(objPtr)]; ok {
		pool.Put(obj)
		return (*T)(objPtr)
	}

	// Add finalizer to the object and return the js pointer as a Go one
	obj.finalizers = append(obj.finalizers, func() {
		delete(jsPtrsByObject, uintptr(objPtr))
		pool.Put(obj)
	})
	runtime.SetFinalizer(obj, func(o *object) {
		for _, fn := range o.finalizers {
			fn()
		}
	})

	jsPtrsByObject[uintptr(objPtr)] = obj
	ret := (*T)(objPtr)

	return ret
}

func DeleteJSPointer(ptr uintptr) {
	_, ok := jsPtrsByObject[ptr]
	if !ok {
		return
	}
	delete(jsPtrsByObject, ptr)
}

func AttachFinalizer[T any](obj *T, fn func()) {
	v, ok := jsPtrsByObject[uintptr(unsafe.Pointer(obj))]
	if !ok {
		return
	}
	v.finalizers = append(v.finalizers, fn)
}

func NewBigInt(n any) js.Value {
	return bigInt.Invoke(n)
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

func GetByteSliceFromJSPtr(ptr js.Value, count int) []byte {
	s := make([]byte, count)
	arr := heapU8.Call("slice", ptr.Int(), ptr.Int()+count)
	js.CopyBytesToGo(s, arr)

	return s
}

type Numeric interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func GetNumericSliceFromJSPtr[T Numeric](ptr js.Value, count int) []T {
	s := make([]T, count)
	bytesCount := count * int(unsafe.Sizeof(T(0)))
	arr := heapU8.Call("slice", ptr.Int(), ptr.Int()+bytesCount)
	js.CopyBytesToGo(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), bytesCount), arr)

	return s
}

func GetObjectSliceFromJSPtr[T any](ptr js.Value, count int) []*T {
	ptrs := GetNumericSliceFromJSPtr[int32](ptr, count)
	ret := make([]*T, count)
	for i := range ptrs {
		var obj T
		CopyJSToObject(&obj, js.ValueOf(ptrs[i]))
		ret[i] = &obj
	}

	return ret
}
