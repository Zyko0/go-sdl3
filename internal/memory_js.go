//go:build js

package internal

import (
	"slices"
	"strconv"
	"strings"
	"sync"
	"syscall/js"
	"unsafe"
)

var (
	heapU8 = js.Global().Get("Module").Get("HEAPU8")

	stackAlloc   = js.Global().Get("Module").Get("stackAlloc")
	stackString  = js.Global().Get("Module").Get("stringToUTF8OnStack")
	utf8String   = js.Global().Get("Module").Get("UTF8ToString")
	getValue     = js.Global().Get("getValue")
	setValue     = js.Global().Get("setValue")
	stackSave    = js.Global().Get("stackSave")
	stackRestore = js.Global().Get("stackRestore")

	bigInt  = js.Global().Get("BigInt")
	boolean = js.Global().Get("Boolean")

	lastStackPtr js.Value

	lock sync.Mutex
)

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

func CloneOnStackGoToJS[T any](obj *T) js.Value {
	if obj == nil {
		return js.ValueOf(0)
	}
	size := unsafe.Sizeof(*obj)
	ptr := stackAlloc.Invoke(int32(size))
	arr := heapU8.Call("subarray", ptr.Int(), ptr.Int()+int(size))
	js.CopyBytesToJS(arr, unsafe.Slice((*byte)(unsafe.Pointer(obj)), size))

	return ptr
}

func CloneOnStackJSToGo[T any](obj *T, ptr js.Value) {
	size := unsafe.Sizeof(*obj)
	arr := heapU8.Call("slice", ptr.Int(), ptr.Int()+int(size))
	js.CopyBytesToGo(unsafe.Slice((*byte)(unsafe.Pointer(obj)), size), arr)
}

func CloneSliceOnHeapGoToJS(s []byte) js.Value {
	ptr := js.Global().Call("_malloc", len(s))
	arr := heapU8.Call("subarray", ptr.Int(), ptr.Int()+len(s))
	js.CopyBytesToJS(arr, s)

	return ptr
}

func StringOnStackGoToJS(str string) js.Value {
	return stackString.Invoke(str)
}

func UTF8ToStringJSToGo(ptr js.Value) string {
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

func GetJSPointerFromUintptr(ptr uintptr) (js.Value, bool) {
	v, ok := jsPtrsByObject[ptr]
	if ok {
		return v.value, true
	}
	return js.Value{}, false
}

func GetJSPointer[T any](obj *T) (js.Value, bool) {
	v, ok := jsPtrsByObject[uintptr(unsafe.Pointer(obj))]
	if ok {
		return v.value, true
	}
	return js.Value{}, false
}

func NewPointer[T any](ptr js.Value) *T {
	var t T

	obj := pool.Get().(*object)
	obj.Map(max(unsafe.Sizeof(t), 1))
	obj.value = ptr
	obj.ExtractSelf()
	objPtr := unsafe.Pointer(&obj.data[0])
	obj.finalizers = append(obj.finalizers, func() {
		delete(jsPtrsByObject, uintptr(objPtr))
		pool.Put(obj)
	})
	jsPtrsByObject[uintptr(objPtr)] = obj
	ret := (*T)(objPtr)

	return ret
}

func DeletePointerReference(ptr uintptr) {
	obj, ok := jsPtrsByObject[ptr]
	if !ok {
		return
	}
	for _, fn := range obj.finalizers {
		fn()
	}
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
