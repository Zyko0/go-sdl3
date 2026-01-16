//go:build js

package sdl

import (
	"runtime"
	js "syscall/js"
	"unsafe"

	internal "github.com/Zyko0/go-sdl3/internal"
)

func initialize() {
	ifree = func(mem uintptr) {
		_mem, ok := internal.GetJSPointerFromUintptr(mem)
		if !ok {
			panic("free invalid pointer target")
		}
		js.Global().Get("Module").Call(
			"_SDL_free",
			_mem,
		)
	}

	iGetVersion = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetVersion",
		)
		return int32(ret.Int())
	}

	iAsyncIOFromFile = func(file string, mode string) *AsyncIO {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_mode := internal.StringOnJSStack(mode)
		ret := js.Global().Get("Module").Call(
			"_SDL_AsyncIOFromFile",
			_file,
			_mode,
		)
		_ = ret

		_obj := &AsyncIO{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iGetAsyncIOSize = func(asyncio *AsyncIO) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_asyncio, ok := internal.GetJSPointer(asyncio)
		if !ok {
			_asyncio = internal.StackAlloc(int(unsafe.Sizeof(*asyncio)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAsyncIOSize",
			_asyncio,
		)

		return int64(internal.GetInt64(ret))
	}

	iReadAsyncIO = func(asyncio *AsyncIO, ptr uintptr, offset uint64, size uint64, queue *AsyncIOQueue, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_asyncio, ok := internal.GetJSPointer(asyncio)
		if !ok {
			_asyncio = internal.StackAlloc(int(unsafe.Sizeof(*asyncio)))
		}
		_ptr := internal.NewBigInt(ptr)
		_offset := internal.NewBigInt(offset)
		_size := internal.NewBigInt(size)
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadAsyncIO",
			_asyncio,
			_ptr,
			_offset,
			_size,
			_queue,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iWriteAsyncIO = func(asyncio *AsyncIO, ptr uintptr, offset uint64, size uint64, queue *AsyncIOQueue, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_asyncio, ok := internal.GetJSPointer(asyncio)
		if !ok {
			_asyncio = internal.StackAlloc(int(unsafe.Sizeof(*asyncio)))
		}
		_ptr := internal.NewBigInt(ptr)
		_offset := internal.NewBigInt(offset)
		_size := internal.NewBigInt(size)
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteAsyncIO",
			_asyncio,
			_ptr,
			_offset,
			_size,
			_queue,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iCloseAsyncIO = func(asyncio *AsyncIO, flush bool, queue *AsyncIOQueue, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_asyncio, ok := internal.GetJSPointer(asyncio)
		if !ok {
			_asyncio = internal.StackAlloc(int(unsafe.Sizeof(*asyncio)))
		}
		_flush := internal.NewBoolean(flush)
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_CloseAsyncIO",
			_asyncio,
			_flush,
			_queue,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iCreateAsyncIOQueue = func() *AsyncIOQueue {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateAsyncIOQueue",
		)
		_ = ret

		_obj := &AsyncIOQueue{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iDestroyAsyncIOQueue = func(queue *AsyncIOQueue) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyAsyncIOQueue",
			_queue,
		)
	}

	iGetAsyncIOResult = func(queue *AsyncIOQueue, outcome *AsyncIOOutcome) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		_outcome, ok := internal.GetJSPointer(outcome)
		if !ok {
			_outcome = internal.StackAlloc(int(unsafe.Sizeof(*outcome)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAsyncIOResult",
			_queue,
			_outcome,
		)

		return internal.GetBool(ret)
	}

	iWaitAsyncIOResult = func(queue *AsyncIOQueue, outcome *AsyncIOOutcome, timeoutMS int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		_outcome, ok := internal.GetJSPointer(outcome)
		if !ok {
			_outcome = internal.StackAlloc(int(unsafe.Sizeof(*outcome)))
		}
		_timeoutMS := int32(timeoutMS)
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitAsyncIOResult",
			_queue,
			_outcome,
			_timeoutMS,
		)

		return internal.GetBool(ret)
	}

	iSignalAsyncIOQueue = func(queue *AsyncIOQueue) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		js.Global().Get("Module").Call(
			"_SDL_SignalAsyncIOQueue",
			_queue,
		)
	}

	iLoadFileAsync = func(file string, queue *AsyncIOQueue, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_queue, ok := internal.GetJSPointer(queue)
		if !ok {
			_queue = internal.StackAlloc(int(unsafe.Sizeof(*queue)))
		}
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadFileAsync",
			_file,
			_queue,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iTryLockSpinlock = func(lock *SpinLock) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_lock, ok := internal.GetJSPointer(lock)
		if !ok {
			_lock = internal.StackAlloc(int(unsafe.Sizeof(*lock)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TryLockSpinlock",
			_lock,
		)

		return internal.GetBool(ret)
	}

	iLockSpinlock = func(lock *SpinLock) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_lock, ok := internal.GetJSPointer(lock)
		if !ok {
			_lock = internal.StackAlloc(int(unsafe.Sizeof(*lock)))
		}
		js.Global().Get("Module").Call(
			"_SDL_LockSpinlock",
			_lock,
		)
	}

	iUnlockSpinlock = func(lock *SpinLock) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_lock, ok := internal.GetJSPointer(lock)
		if !ok {
			_lock = internal.StackAlloc(int(unsafe.Sizeof(*lock)))
		}
		js.Global().Get("Module").Call(
			"_SDL_UnlockSpinlock",
			_lock,
		)
	}

	iMemoryBarrierReleaseFunction = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_MemoryBarrierReleaseFunction",
		)
	}

	iMemoryBarrierAcquireFunction = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_MemoryBarrierAcquireFunction",
		)
	}

	iCompareAndSwapAtomicInt = func(a *AtomicInt, oldval int32, newval int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_oldval := int32(oldval)
		_newval := int32(newval)
		ret := js.Global().Get("Module").Call(
			"_SDL_CompareAndSwapAtomicInt",
			_a,
			_oldval,
			_newval,
		)

		return internal.GetBool(ret)
	}

	iSetAtomicInt = func(a *AtomicInt, v int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_v := int32(v)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAtomicInt",
			_a,
			_v,
		)

		return int32(ret.Int())
	}

	iGetAtomicInt = func(a *AtomicInt) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAtomicInt",
			_a,
		)

		return int32(ret.Int())
	}

	iAddAtomicInt = func(a *AtomicInt, v int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_v := int32(v)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddAtomicInt",
			_a,
			_v,
		)

		return int32(ret.Int())
	}

	iCompareAndSwapAtomicU32 = func(a *AtomicU32, oldval uint32, newval uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_oldval := int32(oldval)
		_newval := int32(newval)
		ret := js.Global().Get("Module").Call(
			"_SDL_CompareAndSwapAtomicU32",
			_a,
			_oldval,
			_newval,
		)

		return internal.GetBool(ret)
	}

	iSetAtomicU32 = func(a *AtomicU32, v uint32) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_v := int32(v)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAtomicU32",
			_a,
			_v,
		)

		return uint32(ret.Int())
	}

	iGetAtomicU32 = func(a *AtomicU32) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAtomicU32",
			_a,
		)

		return uint32(ret.Int())
	}

	iCompareAndSwapAtomicPointer = func(a *uintptr, oldval uintptr, newval uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_oldval := internal.NewBigInt(oldval)
		_newval := internal.NewBigInt(newval)
		ret := js.Global().Get("Module").Call(
			"_SDL_CompareAndSwapAtomicPointer",
			_a,
			_oldval,
			_newval,
		)

		return internal.GetBool(ret)
	}

	iSetAtomicPointer = func(a *uintptr, v uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		_v := internal.NewBigInt(v)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAtomicPointer",
			_a,
			_v,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetAtomicPointer = func(a *uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAtomicPointer",
			_a,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iSetError = func(fmt string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_fmt := internal.StringOnJSStack(fmt)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetError",
			_fmt,
		)

		return internal.GetBool(ret)
	}

	iSetErrorV = func(fmt string, ap va_list) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_fmt := internal.StringOnJSStack(fmt)
		_ap := int32(ap)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetErrorV",
			_fmt,
			_ap,
		)

		return internal.GetBool(ret)
	}

	iOutOfMemory = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_OutOfMemory",
		)

		return internal.GetBool(ret)
	}

	iGetError = func() string {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetError",
		)

		return internal.UTF8JSToString(ret)
	}

	iClearError = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_ClearError",
		)

		return internal.GetBool(ret)
	}

	iGetGlobalProperties = func() PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlobalProperties",
		)

		return PropertiesID(ret.Int())
	}

	iCreateProperties = func() PropertiesID {
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateProperties",
		)

		return PropertiesID(ret.Int())
	}

	iCopyProperties = func(src PropertiesID, dst PropertiesID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src := int32(src)
		_dst := int32(dst)
		ret := js.Global().Get("Module").Call(
			"_SDL_CopyProperties",
			_src,
			_dst,
		)

		return internal.GetBool(ret)
	}

	iLockProperties = func(props PropertiesID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_LockProperties",
			_props,
		)

		return internal.GetBool(ret)
	}

	iUnlockProperties = func(props PropertiesID) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		js.Global().Get("Module").Call(
			"_SDL_UnlockProperties",
			_props,
		)
	}

	/*iSetPointerPropertyWithCleanup = func(props PropertiesID, name string, value uintptr, cleanup CleanupPropertyCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnStackGoToJS(name)
		_value := internal.NewBigInt(value)
		_cleanup := int32(cleanup)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetPointerPropertyWithCleanup",
			_props,
			_name,
			_value,
			_cleanup,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	iSetPointerProperty = func(props PropertiesID, name string, value uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_value := internal.NewBigInt(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetPointerProperty",
			_props,
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iSetStringProperty = func(props PropertiesID, name string, value string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_value := internal.StringOnJSStack(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetStringProperty",
			_props,
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iSetNumberProperty = func(props PropertiesID, name string, value int64) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_value := internal.NewBigInt(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetNumberProperty",
			_props,
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iSetFloatProperty = func(props PropertiesID, name string, value float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFloatProperty",
			_props,
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iSetBooleanProperty = func(props PropertiesID, name string, value bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_value := internal.NewBoolean(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetBooleanProperty",
			_props,
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iHasProperty = func(props PropertiesID, name string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_HasProperty",
			_props,
			_name,
		)

		return internal.GetBool(ret)
	}

	iGetPropertyType = func(props PropertiesID, name string) PropertyType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPropertyType",
			_props,
			_name,
		)

		return PropertyType(ret.Int())
	}

	iGetPointerProperty = func(props PropertiesID, name string, default_value uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_default_value := internal.NewBigInt(default_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPointerProperty",
			_props,
			_name,
			_default_value,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetStringProperty = func(props PropertiesID, name string, default_value string) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_default_value := internal.StringOnJSStack(default_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetStringProperty",
			_props,
			_name,
			_default_value,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetNumberProperty = func(props PropertiesID, name string, default_value int64) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_default_value := internal.NewBigInt(default_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumberProperty",
			_props,
			_name,
			_default_value,
		)

		return int64(internal.GetInt64(ret))
	}

	iGetFloatProperty = func(props PropertiesID, name string, default_value float32) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_default_value := int32(default_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFloatProperty",
			_props,
			_name,
			_default_value,
		)

		return float32(ret.Int())
	}

	iGetBooleanProperty = func(props PropertiesID, name string, default_value bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		_default_value := internal.NewBoolean(default_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetBooleanProperty",
			_props,
			_name,
			_default_value,
		)

		return internal.GetBool(ret)
	}

	iClearProperty = func(props PropertiesID, name string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_ClearProperty",
			_props,
			_name,
		)

		return internal.GetBool(ret)
	}

	iEnumerateProperties = func(props PropertiesID, callback EnumeratePropertiesCallback, userdata uintptr) bool {
		_props := int32(props)
		_callback := int32(callback)
		_userdata := int32(0) // Note: ignored

		ret := js.Global().Get("Module").Call(
			"_SDL_EnumerateProperties",
			_props,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iDestroyProperties = func(props PropertiesID) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		js.Global().Get("Module").Call(
			"_SDL_DestroyProperties",
			_props,
		)
	}

	iGetThreadName = func(thread *Thread) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_thread, ok := internal.GetJSPointer(thread)
		if !ok {
			_thread = internal.StackAlloc(int(unsafe.Sizeof(*thread)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetThreadName",
			_thread,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetCurrentThreadID = func() ThreadID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentThreadID",
		)

		return ThreadID(ret.Int())
	}

	iGetThreadID = func(thread *Thread) ThreadID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_thread, ok := internal.GetJSPointer(thread)
		if !ok {
			_thread = internal.StackAlloc(int(unsafe.Sizeof(*thread)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetThreadID",
			_thread,
		)

		return ThreadID(ret.Int())
	}

	iSetCurrentThreadPriority = func(priority ThreadPriority) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_priority := int32(priority)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetCurrentThreadPriority",
			_priority,
		)

		return internal.GetBool(ret)
	}

	iWaitThread = func(thread *Thread, status *int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_thread, ok := internal.GetJSPointer(thread)
		if !ok {
			_thread = internal.StackAlloc(int(unsafe.Sizeof(*thread)))
		}
		_status, ok := internal.GetJSPointer(status)
		if !ok {
			_status = internal.StackAlloc(int(unsafe.Sizeof(*status)))
		}
		js.Global().Get("Module").Call(
			"_SDL_WaitThread",
			_thread,
			_status,
		)
	}

	iGetThreadState = func(thread *Thread) ThreadState {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_thread, ok := internal.GetJSPointer(thread)
		if !ok {
			_thread = internal.StackAlloc(int(unsafe.Sizeof(*thread)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetThreadState",
			_thread,
		)

		return ThreadState(ret.Int())
	}

	iDetachThread = func(thread *Thread) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_thread, ok := internal.GetJSPointer(thread)
		if !ok {
			_thread = internal.StackAlloc(int(unsafe.Sizeof(*thread)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DetachThread",
			_thread,
		)
	}

	iGetTLS = func(id *TLSID) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_id, ok := internal.GetJSPointer(id)
		if !ok {
			_id = internal.StackAlloc(int(unsafe.Sizeof(*id)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTLS",
			_id,
		)

		return uintptr(internal.GetInt64(ret))
	}

	/*iSetTLS = func(id *TLSID, value uintptr, destructor TLSDestructorCallback) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_id, ok := internal.GetJSPointer(id)
		if !ok {
			_id = internal.StackAlloc(int(unsafe.Sizeof(*id)))
		}
		_value := internal.NewBigInt(value)
		_destructor := int32(destructor)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTLS",
			_id,
			_value,
			_destructor,
		)

		return internal.GetBool(ret)
	}*/

	iCleanupTLS = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_CleanupTLS",
		)
	}

	iCreateMutex = func() *Mutex {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateMutex",
		)
		_ = ret

		_obj := &Mutex{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iLockMutex = func(mutex *Mutex) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mutex, ok := internal.GetJSPointer(mutex)
		if !ok {
			_mutex = internal.StackAlloc(int(unsafe.Sizeof(*mutex)))
		}
		js.Global().Get("Module").Call(
			"_SDL_LockMutex",
			_mutex,
		)
	}

	iTryLockMutex = func(mutex *Mutex) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mutex, ok := internal.GetJSPointer(mutex)
		if !ok {
			_mutex = internal.StackAlloc(int(unsafe.Sizeof(*mutex)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TryLockMutex",
			_mutex,
		)

		return internal.GetBool(ret)
	}

	iUnlockMutex = func(mutex *Mutex) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mutex, ok := internal.GetJSPointer(mutex)
		if !ok {
			_mutex = internal.StackAlloc(int(unsafe.Sizeof(*mutex)))
		}
		js.Global().Get("Module").Call(
			"_SDL_UnlockMutex",
			_mutex,
		)
	}

	iDestroyMutex = func(mutex *Mutex) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mutex, ok := internal.GetJSPointer(mutex)
		if !ok {
			_mutex = internal.StackAlloc(int(unsafe.Sizeof(*mutex)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyMutex",
			_mutex,
		)
	}

	iCreateRWLock = func() *RWLock {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateRWLock",
		)
		_ = ret

		_obj := &RWLock{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iLockRWLockForReading = func(rwlock *RWLock) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rwlock, ok := internal.GetJSPointer(rwlock)
		if !ok {
			_rwlock = internal.StackAlloc(int(unsafe.Sizeof(*rwlock)))
		}
		js.Global().Get("Module").Call(
			"_SDL_LockRWLockForReading",
			_rwlock,
		)
	}

	iLockRWLockForWriting = func(rwlock *RWLock) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rwlock, ok := internal.GetJSPointer(rwlock)
		if !ok {
			_rwlock = internal.StackAlloc(int(unsafe.Sizeof(*rwlock)))
		}
		js.Global().Get("Module").Call(
			"_SDL_LockRWLockForWriting",
			_rwlock,
		)
	}

	iTryLockRWLockForReading = func(rwlock *RWLock) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rwlock, ok := internal.GetJSPointer(rwlock)
		if !ok {
			_rwlock = internal.StackAlloc(int(unsafe.Sizeof(*rwlock)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TryLockRWLockForReading",
			_rwlock,
		)

		return internal.GetBool(ret)
	}

	iTryLockRWLockForWriting = func(rwlock *RWLock) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rwlock, ok := internal.GetJSPointer(rwlock)
		if !ok {
			_rwlock = internal.StackAlloc(int(unsafe.Sizeof(*rwlock)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TryLockRWLockForWriting",
			_rwlock,
		)

		return internal.GetBool(ret)
	}

	iUnlockRWLock = func(rwlock *RWLock) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rwlock, ok := internal.GetJSPointer(rwlock)
		if !ok {
			_rwlock = internal.StackAlloc(int(unsafe.Sizeof(*rwlock)))
		}
		js.Global().Get("Module").Call(
			"_SDL_UnlockRWLock",
			_rwlock,
		)
	}

	iDestroyRWLock = func(rwlock *RWLock) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rwlock, ok := internal.GetJSPointer(rwlock)
		if !ok {
			_rwlock = internal.StackAlloc(int(unsafe.Sizeof(*rwlock)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyRWLock",
			_rwlock,
		)
	}

	iCreateSemaphore = func(initial_value uint32) *Semaphore {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_initial_value := int32(initial_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSemaphore",
			_initial_value,
		)
		_ = ret

		_obj := &Semaphore{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iDestroySemaphore = func(sem *Semaphore) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sem, ok := internal.GetJSPointer(sem)
		if !ok {
			_sem = internal.StackAlloc(int(unsafe.Sizeof(*sem)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroySemaphore",
			_sem,
		)
	}

	iWaitSemaphore = func(sem *Semaphore) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sem, ok := internal.GetJSPointer(sem)
		if !ok {
			_sem = internal.StackAlloc(int(unsafe.Sizeof(*sem)))
		}
		js.Global().Get("Module").Call(
			"_SDL_WaitSemaphore",
			_sem,
		)
	}

	iTryWaitSemaphore = func(sem *Semaphore) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sem, ok := internal.GetJSPointer(sem)
		if !ok {
			_sem = internal.StackAlloc(int(unsafe.Sizeof(*sem)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TryWaitSemaphore",
			_sem,
		)

		return internal.GetBool(ret)
	}

	iWaitSemaphoreTimeout = func(sem *Semaphore, timeoutMS int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sem, ok := internal.GetJSPointer(sem)
		if !ok {
			_sem = internal.StackAlloc(int(unsafe.Sizeof(*sem)))
		}
		_timeoutMS := int32(timeoutMS)
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitSemaphoreTimeout",
			_sem,
			_timeoutMS,
		)

		return internal.GetBool(ret)
	}

	iSignalSemaphore = func(sem *Semaphore) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sem, ok := internal.GetJSPointer(sem)
		if !ok {
			_sem = internal.StackAlloc(int(unsafe.Sizeof(*sem)))
		}
		js.Global().Get("Module").Call(
			"_SDL_SignalSemaphore",
			_sem,
		)
	}

	iGetSemaphoreValue = func(sem *Semaphore) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sem, ok := internal.GetJSPointer(sem)
		if !ok {
			_sem = internal.StackAlloc(int(unsafe.Sizeof(*sem)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSemaphoreValue",
			_sem,
		)

		return uint32(ret.Int())
	}

	iCreateCondition = func() *Condition {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateCondition",
		)
		_ = ret

		_obj := &Condition{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iDestroyCondition = func(cond *Condition) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cond, ok := internal.GetJSPointer(cond)
		if !ok {
			_cond = internal.StackAlloc(int(unsafe.Sizeof(*cond)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyCondition",
			_cond,
		)
	}

	iSignalCondition = func(cond *Condition) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cond, ok := internal.GetJSPointer(cond)
		if !ok {
			_cond = internal.StackAlloc(int(unsafe.Sizeof(*cond)))
		}
		js.Global().Get("Module").Call(
			"_SDL_SignalCondition",
			_cond,
		)
	}

	iBroadcastCondition = func(cond *Condition) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cond, ok := internal.GetJSPointer(cond)
		if !ok {
			_cond = internal.StackAlloc(int(unsafe.Sizeof(*cond)))
		}
		js.Global().Get("Module").Call(
			"_SDL_BroadcastCondition",
			_cond,
		)
	}

	iWaitCondition = func(cond *Condition, mutex *Mutex) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cond, ok := internal.GetJSPointer(cond)
		if !ok {
			_cond = internal.StackAlloc(int(unsafe.Sizeof(*cond)))
		}
		_mutex, ok := internal.GetJSPointer(mutex)
		if !ok {
			_mutex = internal.StackAlloc(int(unsafe.Sizeof(*mutex)))
		}
		js.Global().Get("Module").Call(
			"_SDL_WaitCondition",
			_cond,
			_mutex,
		)
	}

	iWaitConditionTimeout = func(cond *Condition, mutex *Mutex, timeoutMS int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cond, ok := internal.GetJSPointer(cond)
		if !ok {
			_cond = internal.StackAlloc(int(unsafe.Sizeof(*cond)))
		}
		_mutex, ok := internal.GetJSPointer(mutex)
		if !ok {
			_mutex = internal.StackAlloc(int(unsafe.Sizeof(*mutex)))
		}
		_timeoutMS := int32(timeoutMS)
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitConditionTimeout",
			_cond,
			_mutex,
			_timeoutMS,
		)

		return internal.GetBool(ret)
	}

	iShouldInit = func(state *InitState) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_state, ok := internal.GetJSPointer(state)
		if !ok {
			_state = internal.StackAlloc(int(unsafe.Sizeof(*state)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ShouldInit",
			_state,
		)

		return internal.GetBool(ret)
	}

	iShouldQuit = func(state *InitState) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_state, ok := internal.GetJSPointer(state)
		if !ok {
			_state = internal.StackAlloc(int(unsafe.Sizeof(*state)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ShouldQuit",
			_state,
		)

		return internal.GetBool(ret)
	}

	iSetInitialized = func(state *InitState, initialized bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_state, ok := internal.GetJSPointer(state)
		if !ok {
			_state = internal.StackAlloc(int(unsafe.Sizeof(*state)))
		}
		_initialized := internal.NewBoolean(initialized)
		js.Global().Get("Module").Call(
			"_SDL_SetInitialized",
			_state,
			_initialized,
		)
	}

	iIOFromFile = func(file string, mode string) *IOStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_mode := internal.StringOnJSStack(mode)
		ret := js.Global().Get("Module").Call(
			"_SDL_IOFromFile",
			_file,
			_mode,
		)
		_ = ret

		_obj := &IOStream{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iIOFromMem = func(mem uintptr, size uintptr) *IOStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mem := internal.NewBigInt(mem)
		_size := internal.NewBigInt(size)
		ret := js.Global().Get("Module").Call(
			"_SDL_IOFromMem",
			_mem,
			_size,
		)
		_ = ret

		_obj := &IOStream{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iIOFromConstMem = func(mem, size uintptr) *IOStream {
		_mem := internal.CloneByteSliceToJSHeap(
			unsafe.Slice(*(**byte)(unsafe.Pointer(&mem)), int(size)),
		)
		_size := int32(size)

		ret := js.Global().Get("Module").Call(
			"_SDL_IOFromConstMem",
			_mem,
			_size,
		)

		_obj := internal.NewObject[IOStream](ret)
		internal.AttachFinalizer(_obj, func() {
			js.Global().Call("_free", _mem)
		})

		return _obj
	}

	iIOFromDynamicMem = func() *IOStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_IOFromDynamicMem",
		)
		_ = ret

		_obj := &IOStream{}
		//internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iOpenIO = func(iface *IOStreamInterface, userdata uintptr) *IOStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_iface, ok := internal.GetJSPointer(iface)
		if !ok {
			_iface = internal.StackAlloc(int(unsafe.Sizeof(*iface)))
		}
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenIO",
			_iface,
			_userdata,
		)
		_ = ret

		_obj := &IOStream{}
		//internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iCloseIO = func(context *IOStream) bool {
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			panic("nil context")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CloseIO",
			_context,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(context)))

		return internal.GetBool(ret)
	}

	iGetIOProperties = func(context *IOStream) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetIOProperties",
			_context,
		)

		return PropertiesID(ret.Int())
	}

	iGetIOStatus = func(context *IOStream) IOStatus {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetIOStatus",
			_context,
		)

		return IOStatus(ret.Int())
	}

	iGetIOSize = func(context *IOStream) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetIOSize",
			_context,
		)

		return int64(internal.GetInt64(ret))
	}

	iSeekIO = func(context *IOStream, offset int64, whence IOWhence) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		_offset := internal.NewBigInt(offset)
		_whence := int32(whence)
		ret := js.Global().Get("Module").Call(
			"_SDL_SeekIO",
			_context,
			_offset,
			_whence,
		)

		return int64(internal.GetInt64(ret))
	}

	iTellIO = func(context *IOStream) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TellIO",
			_context,
		)

		return int64(internal.GetInt64(ret))
	}

	iReadIO = func(context *IOStream, ptr uintptr, size uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		_ptr := internal.NewBigInt(ptr)
		_size := internal.NewBigInt(size)
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadIO",
			_context,
			_ptr,
			_size,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iWriteIO = func(context *IOStream, ptr uintptr, size uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		_ptr := internal.NewBigInt(ptr)
		_size := internal.NewBigInt(size)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteIO",
			_context,
			_ptr,
			_size,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iIOprintf = func(context *IOStream, fmt string) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		_fmt := internal.StringOnJSStack(fmt)
		ret := js.Global().Get("Module").Call(
			"_SDL_IOprintf",
			_context,
			_fmt,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iIOvprintf = func(context *IOStream, fmt string, ap va_list) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		_fmt := internal.StringOnJSStack(fmt)
		_ap := int32(ap)
		ret := js.Global().Get("Module").Call(
			"_SDL_IOvprintf",
			_context,
			_fmt,
			_ap,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iFlushIO = func(context *IOStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context, ok := internal.GetJSPointer(context)
		if !ok {
			_context = internal.StackAlloc(int(unsafe.Sizeof(*context)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_FlushIO",
			_context,
		)

		return internal.GetBool(ret)
	}

	iLoadFile_IO = func(src *IOStream, datasize *uintptr, closeio bool) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_datasize, ok := internal.GetJSPointer(datasize)
		if !ok {
			_datasize = internal.StackAlloc(int(unsafe.Sizeof(*datasize)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadFile_IO",
			_src,
			_datasize,
			_closeio,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iLoadFile = func(file string, datasize *uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_datasize, ok := internal.GetJSPointer(datasize)
		if !ok {
			_datasize = internal.StackAlloc(int(unsafe.Sizeof(*datasize)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadFile",
			_file,
			_datasize,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iSaveFile_IO = func(src *IOStream, data uintptr, datasize uintptr, closeio bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_data := internal.NewBigInt(data)
		_datasize := internal.NewBigInt(datasize)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_SaveFile_IO",
			_src,
			_data,
			_datasize,
			_closeio,
		)

		return internal.GetBool(ret)
	}

	iSaveFile = func(file string, data uintptr, datasize uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_data := internal.NewBigInt(data)
		_datasize := internal.NewBigInt(datasize)
		ret := js.Global().Get("Module").Call(
			"_SDL_SaveFile",
			_file,
			_data,
			_datasize,
		)

		return internal.GetBool(ret)
	}

	iReadU8 = func(src *IOStream, value *uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU8",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS8 = func(src *IOStream, value *int8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS8",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadU16LE = func(src *IOStream, value *uint16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU16LE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS16LE = func(src *IOStream, value *int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS16LE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadU16BE = func(src *IOStream, value *uint16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU16BE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS16BE = func(src *IOStream, value *int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS16BE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadU32LE = func(src *IOStream, value *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU32LE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS32LE = func(src *IOStream, value *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS32LE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadU32BE = func(src *IOStream, value *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU32BE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS32BE = func(src *IOStream, value *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS32BE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadU64LE = func(src *IOStream, value *uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU64LE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS64LE = func(src *IOStream, value *int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS64LE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadU64BE = func(src *IOStream, value *uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadU64BE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iReadS64BE = func(src *IOStream, value *int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadS64BE",
			_src,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU8 = func(dst *IOStream, value uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU8",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS8 = func(dst *IOStream, value int8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS8",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU16LE = func(dst *IOStream, value uint16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU16LE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS16LE = func(dst *IOStream, value int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS16LE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU16BE = func(dst *IOStream, value uint16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU16BE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS16BE = func(dst *IOStream, value int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS16BE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU32LE = func(dst *IOStream, value uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU32LE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS32LE = func(dst *IOStream, value int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS32LE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU32BE = func(dst *IOStream, value uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU32BE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS32BE = func(dst *IOStream, value int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS32BE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU64LE = func(dst *IOStream, value uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := internal.NewBigInt(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU64LE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS64LE = func(dst *IOStream, value int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := internal.NewBigInt(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS64LE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteU64BE = func(dst *IOStream, value uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := internal.NewBigInt(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteU64BE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iWriteS64BE = func(dst *IOStream, value int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_value := internal.NewBigInt(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteS64BE",
			_dst,
			_value,
		)

		return internal.GetBool(ret)
	}

	iGetNumAudioDrivers = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumAudioDrivers",
		)

		return int32(ret.Int())
	}

	iGetAudioDriver = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioDriver",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetCurrentAudioDriver = func() string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentAudioDriver",
		)

		return internal.UTF8JSToString(ret)
	}

	iGetAudioPlaybackDevices = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioPlaybackDevices",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetAudioRecordingDevices = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioRecordingDevices",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetAudioDeviceName = func(devid AudioDeviceID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioDeviceName",
			_devid,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetAudioDeviceFormat = func(devid AudioDeviceID, spec *AudioSpec, sample_frames *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		_sample_frames, ok := internal.GetJSPointer(sample_frames)
		if !ok {
			_sample_frames = internal.StackAlloc(int(unsafe.Sizeof(*sample_frames)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioDeviceFormat",
			_devid,
			_spec,
			_sample_frames,
		)

		return internal.GetBool(ret)
	}

	iGetAudioDeviceChannelMap = func(devid AudioDeviceID, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioDeviceChannelMap",
			_devid,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iOpenAudioDevice = func(devid AudioDeviceID, spec *AudioSpec) AudioDeviceID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenAudioDevice",
			_devid,
			_spec,
		)

		return AudioDeviceID(ret.Int())
	}

	iIsAudioDevicePhysical = func(devid AudioDeviceID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_IsAudioDevicePhysical",
			_devid,
		)

		return internal.GetBool(ret)
	}

	iIsAudioDevicePlayback = func(devid AudioDeviceID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_IsAudioDevicePlayback",
			_devid,
		)

		return internal.GetBool(ret)
	}

	iPauseAudioDevice = func(devid AudioDeviceID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_PauseAudioDevice",
			_devid,
		)

		return internal.GetBool(ret)
	}

	iResumeAudioDevice = func(devid AudioDeviceID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_ResumeAudioDevice",
			_devid,
		)

		return internal.GetBool(ret)
	}

	iAudioDevicePaused = func(devid AudioDeviceID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_AudioDevicePaused",
			_devid,
		)

		return internal.GetBool(ret)
	}

	iGetAudioDeviceGain = func(devid AudioDeviceID) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioDeviceGain",
			_devid,
		)

		return float32(ret.Int())
	}

	iSetAudioDeviceGain = func(devid AudioDeviceID, gain float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_gain := int32(gain)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioDeviceGain",
			_devid,
			_gain,
		)

		return internal.GetBool(ret)
	}

	iCloseAudioDevice = func(devid AudioDeviceID) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		js.Global().Get("Module").Call(
			"_SDL_CloseAudioDevice",
			_devid,
		)
	}

	iBindAudioStreams = func(devid AudioDeviceID, streams **AudioStream, num_streams int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_streams, ok := internal.GetJSPointer(streams)
		if !ok {
			_streams = internal.StackAlloc(4)
		}
		_num_streams := int32(num_streams)
		ret := js.Global().Get("Module").Call(
			"_SDL_BindAudioStreams",
			_devid,
			_streams,
			_num_streams,
		)

		return internal.GetBool(ret)
	}

	iBindAudioStream = func(devid AudioDeviceID, stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_BindAudioStream",
			_devid,
			_stream,
		)

		return internal.GetBool(ret)
	}

	iUnbindAudioStreams = func(streams **AudioStream, num_streams int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_streams, ok := internal.GetJSPointer(streams)
		if !ok {
			_streams = internal.StackAlloc(4)
		}
		_num_streams := int32(num_streams)
		js.Global().Get("Module").Call(
			"_SDL_UnbindAudioStreams",
			_streams,
			_num_streams,
		)
	}

	iUnbindAudioStream = func(stream *AudioStream) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		js.Global().Get("Module").Call(
			"_SDL_UnbindAudioStream",
			_stream,
		)
	}

	iGetAudioStreamDevice = func(stream *AudioStream) AudioDeviceID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamDevice",
			_stream,
		)

		return AudioDeviceID(ret.Int())
	}

	iCreateAudioStream = func(src_spec *AudioSpec, dst_spec *AudioSpec) *AudioStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src_spec, ok := internal.GetJSPointer(src_spec)
		if !ok {
			_src_spec = internal.StackAlloc(int(unsafe.Sizeof(*src_spec)))
		}
		_dst_spec, ok := internal.GetJSPointer(dst_spec)
		if !ok {
			_dst_spec = internal.StackAlloc(int(unsafe.Sizeof(*dst_spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateAudioStream",
			_src_spec,
			_dst_spec,
		)

		_obj := &AudioStream{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetAudioStreamProperties = func(stream *AudioStream) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamProperties",
			_stream,
		)

		return PropertiesID(ret.Int())
	}

	iGetAudioStreamFormat = func(stream *AudioStream, src_spec *AudioSpec, dst_spec *AudioSpec) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_src_spec, ok := internal.GetJSPointer(src_spec)
		if !ok {
			_src_spec = internal.StackAlloc(int(unsafe.Sizeof(*src_spec)))
		}
		_dst_spec, ok := internal.GetJSPointer(dst_spec)
		if !ok {
			_dst_spec = internal.StackAlloc(int(unsafe.Sizeof(*dst_spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamFormat",
			_stream,
			_src_spec,
			_dst_spec,
		)

		return internal.GetBool(ret)
	}

	iSetAudioStreamFormat = func(stream *AudioStream, src_spec *AudioSpec, dst_spec *AudioSpec) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_src_spec, ok := internal.GetJSPointer(src_spec)
		if !ok {
			_src_spec = internal.StackAlloc(int(unsafe.Sizeof(*src_spec)))
		}
		_dst_spec, ok := internal.GetJSPointer(dst_spec)
		if !ok {
			_dst_spec = internal.StackAlloc(int(unsafe.Sizeof(*dst_spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamFormat",
			_stream,
			_src_spec,
			_dst_spec,
		)

		return internal.GetBool(ret)
	}

	iGetAudioStreamFrequencyRatio = func(stream *AudioStream) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamFrequencyRatio",
			_stream,
		)

		return float32(ret.Int())
	}

	iSetAudioStreamFrequencyRatio = func(stream *AudioStream, ratio float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_ratio := int32(ratio)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamFrequencyRatio",
			_stream,
			_ratio,
		)

		return internal.GetBool(ret)
	}

	iGetAudioStreamGain = func(stream *AudioStream) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamGain",
			_stream,
		)

		return float32(ret.Int())
	}

	iSetAudioStreamGain = func(stream *AudioStream, gain float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_gain := int32(gain)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamGain",
			_stream,
			_gain,
		)

		return internal.GetBool(ret)
	}

	iGetAudioStreamInputChannelMap = func(stream *AudioStream, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamInputChannelMap",
			_stream,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetAudioStreamOutputChannelMap = func(stream *AudioStream, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamOutputChannelMap",
			_stream,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iSetAudioStreamInputChannelMap = func(stream *AudioStream, chmap *int32, count int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_chmap, ok := internal.GetJSPointer(chmap)
		if !ok {
			_chmap = internal.StackAlloc(int(unsafe.Sizeof(*chmap)))
		}
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamInputChannelMap",
			_stream,
			_chmap,
			_count,
		)

		return internal.GetBool(ret)
	}

	iSetAudioStreamOutputChannelMap = func(stream *AudioStream, chmap *int32, count int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_chmap, ok := internal.GetJSPointer(chmap)
		if !ok {
			_chmap = internal.StackAlloc(int(unsafe.Sizeof(*chmap)))
		}
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamOutputChannelMap",
			_stream,
			_chmap,
			_count,
		)

		return internal.GetBool(ret)
	}

	iPutAudioStreamData = func(stream *AudioStream, buf uintptr, len int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_buf := internal.NewBigInt(buf)
		_len := int32(len)
		ret := js.Global().Get("Module").Call(
			"_SDL_PutAudioStreamData",
			_stream,
			_buf,
			_len,
		)

		return internal.GetBool(ret)
	}

	iGetAudioStreamData = func(stream *AudioStream, buf uintptr, len int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_buf := internal.NewBigInt(buf)
		_len := int32(len)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamData",
			_stream,
			_buf,
			_len,
		)

		return int32(ret.Int())
	}

	iGetAudioStreamAvailable = func(stream *AudioStream) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamAvailable",
			_stream,
		)

		return int32(ret.Int())
	}

	iGetAudioStreamQueued = func(stream *AudioStream) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioStreamQueued",
			_stream,
		)

		return int32(ret.Int())
	}

	iFlushAudioStream = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_FlushAudioStream",
			_stream,
		)

		return internal.GetBool(ret)
	}

	iClearAudioStream = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ClearAudioStream",
			_stream,
		)

		return internal.GetBool(ret)
	}

	iPauseAudioStreamDevice = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_PauseAudioStreamDevice",
			_stream,
		)

		return internal.GetBool(ret)
	}

	iResumeAudioStreamDevice = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ResumeAudioStreamDevice",
			_stream,
		)

		return internal.GetBool(ret)
	}

	iAudioStreamDevicePaused = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_AudioStreamDevicePaused",
			_stream,
		)

		return internal.GetBool(ret)
	}

	iLockAudioStream = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LockAudioStream",
			_stream,
		)

		return internal.GetBool(ret)
	}

	iUnlockAudioStream = func(stream *AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_UnlockAudioStream",
			_stream,
		)

		return internal.GetBool(ret)
	}

	/*iSetAudioStreamGetCallback = func(stream *AudioStream, callback AudioStreamCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamGetCallback",
			_stream,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	/*iSetAudioStreamPutCallback = func(stream *AudioStream, callback AudioStreamCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioStreamPutCallback",
			_stream,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	iDestroyAudioStream = func(stream *AudioStream) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyAudioStream",
			_stream,
		)
	}

	/*iOpenAudioDeviceStream = func(devid AudioDeviceID, spec *AudioSpec, callback AudioStreamCallback, userdata uintptr) *AudioStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenAudioDeviceStream",
			_devid,
			_spec,
			_callback,
			_userdata,
		)

		_obj := &AudioStream{}
		internal.StoreJSPointer(_obj, ret)
		return _obj
	}*/

	/*iSetAudioPostmixCallback = func(devid AudioDeviceID, callback AudioPostmixCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAudioPostmixCallback",
			_devid,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	iLoadWAV_IO = func(src *IOStream, closeio bool, spec *AudioSpec, audio_buf **uint8, audio_len *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		_audio_buf, ok := internal.GetJSPointer(audio_buf)
		if !ok {
			_audio_buf = internal.StackAlloc(4)
		}
		_audio_len, ok := internal.GetJSPointer(audio_len)
		if !ok {
			_audio_len = internal.StackAlloc(int(unsafe.Sizeof(*audio_len)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadWAV_IO",
			_src,
			_closeio,
			_spec,
			_audio_buf,
			_audio_len,
		)

		return internal.GetBool(ret)
	}

	iLoadWAV = func(path string, spec *AudioSpec, audio_buf **uint8, audio_len *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		_audio_buf, ok := internal.GetJSPointer(audio_buf)
		if !ok {
			_audio_buf = internal.StackAlloc(4)
		}
		_audio_len, ok := internal.GetJSPointer(audio_len)
		if !ok {
			_audio_len = internal.StackAlloc(int(unsafe.Sizeof(*audio_len)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadWAV",
			_path,
			_spec,
			_audio_buf,
			_audio_len,
		)

		return internal.GetBool(ret)
	}

	iMixAudio = func(dst *uint8, src *uint8, format AudioFormat, len uint32, volume float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_format := int32(format)
		_len := int32(len)
		_volume := int32(volume)
		ret := js.Global().Get("Module").Call(
			"_SDL_MixAudio",
			_dst,
			_src,
			_format,
			_len,
			_volume,
		)

		return internal.GetBool(ret)
	}

	iConvertAudioSamples = func(src_spec *AudioSpec, src_data *uint8, src_len int32, dst_spec *AudioSpec, dst_data **uint8, dst_len *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src_spec, ok := internal.GetJSPointer(src_spec)
		if !ok {
			_src_spec = internal.StackAlloc(int(unsafe.Sizeof(*src_spec)))
		}
		_src_data, ok := internal.GetJSPointer(src_data)
		if !ok {
			_src_data = internal.StackAlloc(int(unsafe.Sizeof(*src_data)))
		}
		_src_len := int32(src_len)
		_dst_spec, ok := internal.GetJSPointer(dst_spec)
		if !ok {
			_dst_spec = internal.StackAlloc(int(unsafe.Sizeof(*dst_spec)))
		}
		_dst_data, ok := internal.GetJSPointer(dst_data)
		if !ok {
			_dst_data = internal.StackAlloc(4)
		}
		_dst_len, ok := internal.GetJSPointer(dst_len)
		if !ok {
			_dst_len = internal.StackAlloc(int(unsafe.Sizeof(*dst_len)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ConvertAudioSamples",
			_src_spec,
			_src_data,
			_src_len,
			_dst_spec,
			_dst_data,
			_dst_len,
		)

		return internal.GetBool(ret)
	}

	iGetAudioFormatName = func(format AudioFormat) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_format := int32(format)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAudioFormatName",
			_format,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetSilenceValueForFormat = func(format AudioFormat) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_format := int32(format)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSilenceValueForFormat",
			_format,
		)

		return int32(ret.Int())
	}

	iComposeCustomBlendMode = func(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_srcColorFactor := int32(srcColorFactor)
		_dstColorFactor := int32(dstColorFactor)
		_colorOperation := int32(colorOperation)
		_srcAlphaFactor := int32(srcAlphaFactor)
		_dstAlphaFactor := int32(dstAlphaFactor)
		_alphaOperation := int32(alphaOperation)
		ret := js.Global().Get("Module").Call(
			"_SDL_ComposeCustomBlendMode",
			_srcColorFactor,
			_dstColorFactor,
			_colorOperation,
			_srcAlphaFactor,
			_dstAlphaFactor,
			_alphaOperation,
		)

		return BlendMode(ret.Int())
	}

	iGetPixelFormatName = func(format PixelFormat) string {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPixelFormatName",
			int32(format),
		)

		return internal.UTF8JSToString(ret)
	}

	iGetMasksForPixelFormat = func(format PixelFormat, bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_format := int32(format)
		_bpp, ok := internal.GetJSPointer(bpp)
		if !ok {
			_bpp = internal.StackAlloc(int(unsafe.Sizeof(*bpp)))
		}
		_Rmask, ok := internal.GetJSPointer(Rmask)
		if !ok {
			_Rmask = internal.StackAlloc(int(unsafe.Sizeof(*Rmask)))
		}
		_Gmask, ok := internal.GetJSPointer(Gmask)
		if !ok {
			_Gmask = internal.StackAlloc(int(unsafe.Sizeof(*Gmask)))
		}
		_Bmask, ok := internal.GetJSPointer(Bmask)
		if !ok {
			_Bmask = internal.StackAlloc(int(unsafe.Sizeof(*Bmask)))
		}
		_Amask, ok := internal.GetJSPointer(Amask)
		if !ok {
			_Amask = internal.StackAlloc(int(unsafe.Sizeof(*Amask)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMasksForPixelFormat",
			_format,
			_bpp,
			_Rmask,
			_Gmask,
			_Bmask,
			_Amask,
		)

		return internal.GetBool(ret)
	}

	iGetPixelFormatForMasks = func(bpp int32, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) PixelFormat {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPixelFormatForMasks",
			bpp,
			int32(Rmask),
			int32(Gmask),
			int32(Bmask),
			int32(Amask),
		)

		return PixelFormat(ret.Int())
	}

	iGetPixelFormatDetails = func(format PixelFormat) *PixelFormatDetails {
		_format := int32(format)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPixelFormatDetails",
			_format,
		)

		_obj := internal.NewObject[PixelFormatDetails](ret)

		return _obj
	}

	iCreatePalette = func(ncolors int32) *Palette {
		ret := js.Global().Get("Module").Call(
			"_SDL_CreatePalette",
			ncolors,
		)

		return internal.NewObject[Palette](ret)
	}

	iSetPaletteColors = func(palette *Palette, colors *Color, firstcolor int32, ncolors int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			panic("nil palette")
		}
		_colors, freeColors := internal.ClonePtrArrayToJSHeap(colors, int(ncolors))
		defer freeColors()
		ret := js.Global().Get("Module").Call(
			"_SDL_SetPaletteColors",
			_palette,
			_colors,
			firstcolor,
			ncolors,
		)

		return internal.GetBool(ret)
	}

	iDestroyPalette = func(palette *Palette) {
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			panic("nil palette")
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyPalette",
			_palette,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(palette)))
	}

	iMapRGB = func(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8) uint32 {
		internal.StackSave()
		defer internal.StackRestore()

		_format, ok := internal.GetJSPointer(format)
		if !ok {
			_format = internal.CloneObjectToJSStack(format)
		}
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			_palette = internal.CloneObjectToJSStack(palette)
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		ret := js.Global().Get("Module").Call(
			"_SDL_MapRGB",
			_format,
			_palette,
			_r,
			_g,
			_b,
		)

		return uint32(ret.Int())
	}

	iMapRGBA = func(format *PixelFormatDetails, palette *Palette, r uint8, g uint8, b uint8, a uint8) uint32 {
		internal.StackSave()
		defer internal.StackRestore()
		_format := internal.CloneObjectToJSStack(format)
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			_palette = js.Null()
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_MapRGBA",
			_format,
			_palette,
			int32(r),
			int32(g),
			int32(b),
			int32(a),
		)

		return uint32(ret.Int())
	}

	iGetRGB = func(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8) {
		internal.StackSave()
		defer internal.StackRestore()
		_format := internal.CloneObjectToJSStack(format)
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			_palette = js.Null()
		}
		_r := internal.StackAlloc(4)
		_g := internal.StackAlloc(4)
		_b := internal.StackAlloc(4)
		js.Global().Get("Module").Call(
			"_SDL_GetRGB",
			int32(pixel),
			_format,
			_palette,
			_r,
			_g,
			_b,
		)
		if r != nil {
			*r = uint8(internal.GetValue(_r, "i8").Int())
		}
		if g != nil {
			*g = uint8(internal.GetValue(_g, "i8").Int())
		}
		if b != nil {
			*b = uint8(internal.GetValue(_b, "i8").Int())
		}
	}

	iGetRGBA = func(pixel uint32, format *PixelFormatDetails, palette *Palette, r *uint8, g *uint8, b *uint8, a *uint8) {
		internal.StackSave()
		defer internal.StackRestore()
		_format := internal.CloneObjectToJSStack(format)
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			_palette = js.Null()
		}
		_r := internal.StackAlloc(4)
		_g := internal.StackAlloc(4)
		_b := internal.StackAlloc(4)
		_a := internal.StackAlloc(4)
		js.Global().Get("Module").Call(
			"_SDL_GetRGBA",
			int32(pixel),
			_format,
			_palette,
			_r,
			_g,
			_b,
			_a,
		)
		if r != nil {
			*r = uint8(internal.GetValue(_r, "i8").Int())
		}
		if g != nil {
			*g = uint8(internal.GetValue(_g, "i8").Int())
		}
		if b != nil {
			*b = uint8(internal.GetValue(_b, "i8").Int())
		}
		if a != nil {
			*a = uint8(internal.GetValue(_a, "i8").Int())
		}
	}

	iHasRectIntersection = func(A *Rect, B *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_A, ok := internal.GetJSPointer(A)
		if !ok {
			_A = internal.StackAlloc(int(unsafe.Sizeof(*A)))
		}
		_B, ok := internal.GetJSPointer(B)
		if !ok {
			_B = internal.StackAlloc(int(unsafe.Sizeof(*B)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_HasRectIntersection",
			_A,
			_B,
		)

		return internal.GetBool(ret)
	}

	iGetRectIntersection = func(A *Rect, B *Rect, result *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_A, ok := internal.GetJSPointer(A)
		if !ok {
			_A = internal.StackAlloc(int(unsafe.Sizeof(*A)))
		}
		_B, ok := internal.GetJSPointer(B)
		if !ok {
			_B = internal.StackAlloc(int(unsafe.Sizeof(*B)))
		}
		_result, ok := internal.GetJSPointer(result)
		if !ok {
			_result = internal.StackAlloc(int(unsafe.Sizeof(*result)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectIntersection",
			_A,
			_B,
			_result,
		)

		return internal.GetBool(ret)
	}

	iGetRectUnion = func(A *Rect, B *Rect, result *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_A, ok := internal.GetJSPointer(A)
		if !ok {
			_A = internal.StackAlloc(int(unsafe.Sizeof(*A)))
		}
		_B, ok := internal.GetJSPointer(B)
		if !ok {
			_B = internal.StackAlloc(int(unsafe.Sizeof(*B)))
		}
		_result, ok := internal.GetJSPointer(result)
		if !ok {
			_result = internal.StackAlloc(int(unsafe.Sizeof(*result)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectUnion",
			_A,
			_B,
			_result,
		)

		return internal.GetBool(ret)
	}

	iGetRectEnclosingPoints = func(points *Point, count int32, clip *Rect, result *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_points, ok := internal.GetJSPointer(points)
		if !ok {
			_points = internal.StackAlloc(int(unsafe.Sizeof(*points)))
		}
		_count := int32(count)
		_clip, ok := internal.GetJSPointer(clip)
		if !ok {
			_clip = internal.StackAlloc(int(unsafe.Sizeof(*clip)))
		}
		_result, ok := internal.GetJSPointer(result)
		if !ok {
			_result = internal.StackAlloc(int(unsafe.Sizeof(*result)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectEnclosingPoints",
			_points,
			_count,
			_clip,
			_result,
		)

		return internal.GetBool(ret)
	}

	iGetRectAndLineIntersection = func(rect *Rect, X1 *int32, Y1 *int32, X2 *int32, Y2 *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		_X1, ok := internal.GetJSPointer(X1)
		if !ok {
			_X1 = internal.StackAlloc(int(unsafe.Sizeof(*X1)))
		}
		_Y1, ok := internal.GetJSPointer(Y1)
		if !ok {
			_Y1 = internal.StackAlloc(int(unsafe.Sizeof(*Y1)))
		}
		_X2, ok := internal.GetJSPointer(X2)
		if !ok {
			_X2 = internal.StackAlloc(int(unsafe.Sizeof(*X2)))
		}
		_Y2, ok := internal.GetJSPointer(Y2)
		if !ok {
			_Y2 = internal.StackAlloc(int(unsafe.Sizeof(*Y2)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectAndLineIntersection",
			_rect,
			_X1,
			_Y1,
			_X2,
			_Y2,
		)

		return internal.GetBool(ret)
	}

	iHasRectIntersectionFloat = func(A *FRect, B *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_A, ok := internal.GetJSPointer(A)
		if !ok {
			_A = internal.StackAlloc(int(unsafe.Sizeof(*A)))
		}
		_B, ok := internal.GetJSPointer(B)
		if !ok {
			_B = internal.StackAlloc(int(unsafe.Sizeof(*B)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_HasRectIntersectionFloat",
			_A,
			_B,
		)

		return internal.GetBool(ret)
	}

	iGetRectIntersectionFloat = func(A *FRect, B *FRect, result *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_A, ok := internal.GetJSPointer(A)
		if !ok {
			_A = internal.StackAlloc(int(unsafe.Sizeof(*A)))
		}
		_B, ok := internal.GetJSPointer(B)
		if !ok {
			_B = internal.StackAlloc(int(unsafe.Sizeof(*B)))
		}
		_result, ok := internal.GetJSPointer(result)
		if !ok {
			_result = internal.StackAlloc(int(unsafe.Sizeof(*result)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectIntersectionFloat",
			_A,
			_B,
			_result,
		)

		return internal.GetBool(ret)
	}

	iGetRectUnionFloat = func(A *FRect, B *FRect, result *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_A, ok := internal.GetJSPointer(A)
		if !ok {
			_A = internal.StackAlloc(int(unsafe.Sizeof(*A)))
		}
		_B, ok := internal.GetJSPointer(B)
		if !ok {
			_B = internal.StackAlloc(int(unsafe.Sizeof(*B)))
		}
		_result, ok := internal.GetJSPointer(result)
		if !ok {
			_result = internal.StackAlloc(int(unsafe.Sizeof(*result)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectUnionFloat",
			_A,
			_B,
			_result,
		)

		return internal.GetBool(ret)
	}

	iGetRectEnclosingPointsFloat = func(points *FPoint, count int32, clip *FRect, result *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_points, ok := internal.GetJSPointer(points)
		if !ok {
			_points = internal.StackAlloc(int(unsafe.Sizeof(*points)))
		}
		_count := int32(count)
		_clip, ok := internal.GetJSPointer(clip)
		if !ok {
			_clip = internal.StackAlloc(int(unsafe.Sizeof(*clip)))
		}
		_result, ok := internal.GetJSPointer(result)
		if !ok {
			_result = internal.StackAlloc(int(unsafe.Sizeof(*result)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectEnclosingPointsFloat",
			_points,
			_count,
			_clip,
			_result,
		)

		return internal.GetBool(ret)
	}

	iGetRectAndLineIntersectionFloat = func(rect *FRect, X1 *float32, Y1 *float32, X2 *float32, Y2 *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		_X1, ok := internal.GetJSPointer(X1)
		if !ok {
			_X1 = internal.StackAlloc(int(unsafe.Sizeof(*X1)))
		}
		_Y1, ok := internal.GetJSPointer(Y1)
		if !ok {
			_Y1 = internal.StackAlloc(int(unsafe.Sizeof(*Y1)))
		}
		_X2, ok := internal.GetJSPointer(X2)
		if !ok {
			_X2 = internal.StackAlloc(int(unsafe.Sizeof(*X2)))
		}
		_Y2, ok := internal.GetJSPointer(Y2)
		if !ok {
			_Y2 = internal.StackAlloc(int(unsafe.Sizeof(*Y2)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRectAndLineIntersectionFloat",
			_rect,
			_X1,
			_Y1,
			_X2,
			_Y2,
		)

		return internal.GetBool(ret)
	}

	iCreateSurface = func(width int32, height int32, format PixelFormat) *Surface {
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSurface",
			width,
			height,
			int32(format),
		)

		return internal.NewObject[Surface](ret)
	}

	iCreateSurfaceFrom = func(width int32, height int32, format PixelFormat, pixels uintptr, pitch int32) *Surface {
		_pixels := internal.NewBigInt(uint64(pixels))
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSurfaceFrom",
			width,
			height,
			int32(format),
			_pixels,
			pitch,
		)

		return internal.NewObject[Surface](ret)
	}

	iDestroySurface = func(surface *Surface) {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroySurface",
			_surface,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(surface)))
	}

	iGetSurfaceProperties = func(surface *Surface) PropertiesID {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceProperties",
			_surface,
		)

		return PropertiesID(ret.Int())
	}

	iSetSurfaceColorspace = func(surface *Surface, colorspace Colorspace) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_colorspace := int32(colorspace)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceColorspace",
			_surface,
			_colorspace,
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceColorspace = func(surface *Surface) Colorspace {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceColorspace",
			_surface,
		)

		return Colorspace(ret.Int())
	}

	iCreateSurfacePalette = func(surface *Surface) *Palette {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSurfacePalette",
			_surface,
		)

		return internal.NewObject[Palette](ret)
	}

	iSetSurfacePalette = func(surface *Surface, palette *Palette) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			_palette = internal.StackAlloc(int(unsafe.Sizeof(*palette)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfacePalette",
			_surface,
			_palette,
		)

		return internal.GetBool(ret)
	}

	iGetSurfacePalette = func(surface *Surface) *Palette {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfacePalette",
			_surface,
		)

		return internal.NewObject[Palette](ret)
	}

	iAddSurfaceAlternateImage = func(surface *Surface, image *Surface) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_image, ok := internal.GetJSPointer(image)
		if !ok {
			_image = internal.StackAlloc(int(unsafe.Sizeof(*image)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_AddSurfaceAlternateImage",
			_surface,
			_image,
		)

		return internal.GetBool(ret)
	}

	iSurfaceHasAlternateImages = func(surface *Surface) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SurfaceHasAlternateImages",
			_surface,
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceImages = func(surface *Surface, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_count := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceImages",
			_surface,
			_count,
		)
		*count = int32(internal.GetValue(_count, "i32").Int())

		//

		return uintptr(internal.GetInt64(ret))
	}

	iRemoveSurfaceAlternateImages = func(surface *Surface) {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		js.Global().Get("Module").Call(
			"_SDL_RemoveSurfaceAlternateImages",
			_surface,
		)
	}

	iLockSurface = func(surface *Surface) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LockSurface",
			_surface,
		)

		return internal.GetBool(ret)
	}

	iUnlockSurface = func(surface *Surface) {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		js.Global().Get("Module").Call(
			"_SDL_UnlockSurface",
			_surface,
		)
	}

	iLoadBMP_IO = func(src *IOStream, closeio bool) *Surface {
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil IOStream")
		}
		_closeio := int32(0)
		if closeio {
			_closeio = 1
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadBMP_IO",
			_src,
			_closeio,
		)

		_obj := internal.NewObject[Surface](ret)

		return _obj
	}

	iLoadBMP = func(file string) *Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadBMP",
			_file,
		)

		return internal.NewObject[Surface](ret)
	}

	iSaveBMP_IO = func(surface *Surface, dst *IOStream, closeio bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_SaveBMP_IO",
			_surface,
			_dst,
			_closeio,
		)

		return internal.GetBool(ret)
	}

	iSaveBMP = func(surface *Surface, file string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_SDL_SaveBMP",
			_surface,
			_file,
		)

		return internal.GetBool(ret)
	}

	iSetSurfaceRLE = func(surface *Surface, enabled bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceRLE",
			_surface,
			_enabled,
		)

		return internal.GetBool(ret)
	}

	iSurfaceHasRLE = func(surface *Surface) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SurfaceHasRLE",
			_surface,
		)

		return internal.GetBool(ret)
	}

	iSetSurfaceColorKey = func(surface *Surface, enabled bool, key uint32) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceColorKey",
			_surface,
			_enabled,
			int32(key),
		)

		return internal.GetBool(ret)
	}

	iSurfaceHasColorKey = func(surface *Surface) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SurfaceHasColorKey",
			_surface,
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceColorKey = func(surface *Surface, key *uint32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_key := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceColorKey",
			_surface,
			_key,
		)
		*key = uint32(internal.GetValue(_key, "i32").Int())

		return internal.GetBool(ret)
	}

	iSetSurfaceColorMod = func(surface *Surface, r uint8, g uint8, b uint8) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceColorMod",
			_surface,
			int32(r),
			int32(g),
			int32(b),
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceColorMod = func(surface *Surface, r *uint8, g *uint8, b *uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceColorMod",
			_surface,
			_r,
			_g,
			_b,
		)

		return internal.GetBool(ret)
	}

	iSetSurfaceAlphaMod = func(surface *Surface, alpha uint8) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceAlphaMod",
			_surface,
			int32(alpha),
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceAlphaMod = func(surface *Surface, alpha *uint8) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_alpha := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceAlphaMod",
			_surface,
			_alpha,
		)
		*alpha = uint8(internal.GetValue(_alpha, "i8").Int())

		return internal.GetBool(ret)
	}

	iSetSurfaceBlendMode = func(surface *Surface, blendMode BlendMode) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceBlendMode",
			_surface,
			int32(blendMode),
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceBlendMode = func(surface *Surface, blendMode *BlendMode) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_blendMode := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceBlendMode",
			_surface,
			_blendMode,
		)
		*blendMode = BlendMode(internal.GetValue(_blendMode, "i32").Int())

		return internal.GetBool(ret)
	}

	iSetSurfaceClipRect = func(surface *Surface, rect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSurfaceClipRect",
			_surface,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iGetSurfaceClipRect = func(surface *Surface, rect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_rect := internal.StackAlloc(int(unsafe.Sizeof(Rect{})))
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSurfaceClipRect",
			_surface,
			_rect,
		)
		internal.CopyJSToObject(rect, _rect)

		return internal.GetBool(ret)
	}

	iFlipSurface = func(surface *Surface, flip FlipMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_flip := int32(flip)
		ret := js.Global().Get("Module").Call(
			"_SDL_FlipSurface",
			_surface,
			_flip,
		)

		return internal.GetBool(ret)
	}

	iDuplicateSurface = func(surface *Surface) *Surface {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_DuplicateSurface",
			_surface,
		)

		return internal.NewObject[Surface](ret)
	}

	iScaleSurface = func(surface *Surface, width int32, height int32, scaleMode ScaleMode) *Surface {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ScaleSurface",
			_surface,
			width,
			height,
			int32(scaleMode),
		)

		return internal.NewObject[Surface](ret)
	}

	iConvertSurface = func(surface *Surface, format PixelFormat) *Surface {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_format := int32(format)
		ret := js.Global().Get("Module").Call(
			"_SDL_ConvertSurface",
			_surface,
			_format,
		)

		_obj := internal.NewObject[Surface](ret)

		return _obj
	}

	iConvertSurfaceAndColorspace = func(surface *Surface, format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) *Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_format := int32(format)
		_palette, ok := internal.GetJSPointer(palette)
		if !ok {
			_palette = internal.StackAlloc(int(unsafe.Sizeof(*palette)))
		}
		_colorspace := int32(colorspace)
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_ConvertSurfaceAndColorspace",
			_surface,
			_format,
			_palette,
			_colorspace,
			_props,
		)

		_obj := &Surface{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iConvertPixels = func(width int32, height int32, src_format PixelFormat, src uintptr, src_pitch int32, dst_format PixelFormat, dst uintptr, dst_pitch int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_width := int32(width)
		_height := int32(height)
		_src_format := int32(src_format)
		_src := internal.NewBigInt(src)
		_src_pitch := int32(src_pitch)
		_dst_format := int32(dst_format)
		_dst := internal.NewBigInt(dst)
		_dst_pitch := int32(dst_pitch)
		ret := js.Global().Get("Module").Call(
			"_SDL_ConvertPixels",
			_width,
			_height,
			_src_format,
			_src,
			_src_pitch,
			_dst_format,
			_dst,
			_dst_pitch,
		)

		return internal.GetBool(ret)
	}

	iConvertPixelsAndColorspace = func(width int32, height int32, src_format PixelFormat, src_colorspace Colorspace, src_properties PropertiesID, src uintptr, src_pitch int32, dst_format PixelFormat, dst_colorspace Colorspace, dst_properties PropertiesID, dst uintptr, dst_pitch int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_width := int32(width)
		_height := int32(height)
		_src_format := int32(src_format)
		_src_colorspace := int32(src_colorspace)
		_src_properties := int32(src_properties)
		_src := internal.NewBigInt(src)
		_src_pitch := int32(src_pitch)
		_dst_format := int32(dst_format)
		_dst_colorspace := int32(dst_colorspace)
		_dst_properties := int32(dst_properties)
		_dst := internal.NewBigInt(dst)
		_dst_pitch := int32(dst_pitch)
		ret := js.Global().Get("Module").Call(
			"_SDL_ConvertPixelsAndColorspace",
			_width,
			_height,
			_src_format,
			_src_colorspace,
			_src_properties,
			_src,
			_src_pitch,
			_dst_format,
			_dst_colorspace,
			_dst_properties,
			_dst,
			_dst_pitch,
		)

		return internal.GetBool(ret)
	}

	iPremultiplyAlpha = func(width int32, height int32, src_format PixelFormat, src uintptr, src_pitch int32, dst_format PixelFormat, dst uintptr, dst_pitch int32, linear bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_width := int32(width)
		_height := int32(height)
		_src_format := int32(src_format)
		_src := internal.NewBigInt(src)
		_src_pitch := int32(src_pitch)
		_dst_format := int32(dst_format)
		_dst := internal.NewBigInt(dst)
		_dst_pitch := int32(dst_pitch)
		_linear := internal.NewBoolean(linear)
		ret := js.Global().Get("Module").Call(
			"_SDL_PremultiplyAlpha",
			_width,
			_height,
			_src_format,
			_src,
			_src_pitch,
			_dst_format,
			_dst,
			_dst_pitch,
			_linear,
		)

		return internal.GetBool(ret)
	}

	iPremultiplySurfaceAlpha = func(surface *Surface, linear bool) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_linear := internal.NewBoolean(linear)
		ret := js.Global().Get("Module").Call(
			"_SDL_PremultiplySurfaceAlpha",
			_surface,
			_linear,
		)

		return internal.GetBool(ret)
	}

	iClearSurface = func(surface *Surface, r float32, g float32, b float32, a float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		_a := int32(a)
		ret := js.Global().Get("Module").Call(
			"_SDL_ClearSurface",
			_surface,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iFillSurfaceRect = func(dst *Surface, rect *Rect, color uint32) bool {
		internal.StackSave()
		defer internal.StackRestore()

		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil surface")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_FillSurfaceRect",
			_dst,
			_rect,
			color,
		)

		return internal.GetBool(ret)
	}

	iFillSurfaceRects = func(dst *Surface, rects *Rect, count int32, color uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_rects, ok := internal.GetJSPointer(rects)
		if !ok {
			_rects = internal.StackAlloc(int(unsafe.Sizeof(*rects)))
		}
		_count := int32(count)
		_color := int32(color)
		ret := js.Global().Get("Module").Call(
			"_SDL_FillSurfaceRects",
			_dst,
			_rects,
			_count,
			_color,
		)

		return internal.GetBool(ret)
	}

	iBlitSurface = func(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurface",
			_src,
			_srcrect,
			_dst,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iBlitSurfaceUnchecked = func(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurfaceUnchecked",
			_src,
			_srcrect,
			_dst,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iBlitSurfaceScaled = func(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurfaceScaled",
			_src,
			_srcrect,
			_dst,
			_dstrect,
			int32(scaleMode),
		)

		return internal.GetBool(ret)
	}

	iBlitSurfaceUncheckedScaled = func(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurfaceUncheckedScaled",
			_src,
			_srcrect,
			_dst,
			_dstrect,
			int32(scaleMode),
		)

		return internal.GetBool(ret)
	}

	iBlitSurfaceTiled = func(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurfaceTiled",
			_src,
			_srcrect,
			_dst,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iBlitSurfaceTiledWithScale = func(src *Surface, srcrect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_scale := scale
		_scaleMode := int32(scaleMode)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurfaceTiledWithScale",
			_src,
			_srcrect,
			_scale,
			_scaleMode,
			_dst,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iBlitSurface9Grid = func(src *Surface, srcrect *Rect, left_width int32, right_width int32, top_height int32, bottom_height int32, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil src surface")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_left_width := left_width
		_right_width := right_width
		_top_height := top_height
		_bottom_height := bottom_height
		_scale := scale
		_scaleMode := int32(scaleMode)
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			panic("nil dst surface")
		}
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_BlitSurface9Grid",
			_src,
			_srcrect,
			_left_width,
			_right_width,
			_top_height,
			_bottom_height,
			_scale,
			_scaleMode,
			_dst,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iMapSurfaceRGB = func(surface *Surface, r uint8, g uint8, b uint8) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		ret := js.Global().Get("Module").Call(
			"_SDL_MapSurfaceRGB",
			_surface,
			_r,
			_g,
			_b,
		)

		return uint32(ret.Int())
	}

	iMapSurfaceRGBA = func(surface *Surface, r uint8, g uint8, b uint8, a uint8) uint32 {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_MapSurfaceRGBA",
			_surface,
			int32(r),
			int32(g),
			int32(b),
			int32(a),
		)

		return uint32(ret.Int())
	}

	iReadSurfacePixel = func(surface *Surface, x int32, y int32, r *uint8, g *uint8, b *uint8, a *uint8) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_x := int32(x)
		_y := int32(y)
		_r := internal.StackAlloc(4)
		_g := internal.StackAlloc(4)
		_b := internal.StackAlloc(4)
		_a := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadSurfacePixel",
			_surface,
			_x,
			_y,
			_r,
			_g,
			_b,
			_a,
		)
		*r = uint8(internal.GetValue(_r, "i8").Int())
		*g = uint8(internal.GetValue(_r, "i8").Int())
		*b = uint8(internal.GetValue(_r, "i8").Int())
		*a = uint8(internal.GetValue(_r, "i8").Int())

		return internal.GetBool(ret)
	}

	iReadSurfacePixelFloat = func(surface *Surface, x int32, y int32, r *float32, g *float32, b *float32, a *float32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		_r := internal.StackAlloc(4)
		_g := internal.StackAlloc(4)
		_b := internal.StackAlloc(4)
		_a := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadSurfacePixelFloat",
			_surface,
			x,
			y,
			_r,
			_g,
			_b,
			_a,
		)
		*r = float32(internal.GetValue(_r, "f32").Float())
		*g = float32(internal.GetValue(_g, "f32").Float())
		*b = float32(internal.GetValue(_b, "f32").Float())
		*a = float32(internal.GetValue(_a, "f32").Float())

		return internal.GetBool(ret)
	}

	iWriteSurfacePixel = func(surface *Surface, x int32, y int32, r uint8, g uint8, b uint8, a uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_x := int32(x)
		_y := int32(y)
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		_a := int32(a)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteSurfacePixel",
			_surface,
			_x,
			_y,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iWriteSurfacePixelFloat = func(surface *Surface, x int32, y int32, r float32, g float32, b float32, a float32) bool {
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteSurfacePixelFloat",
			_surface,
			x,
			y,
			r,
			g,
			b,
			a,
		)

		return internal.GetBool(ret)
	}

	iGetNumCameraDrivers = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumCameraDrivers",
		)

		return int32(ret.Int())
	}

	iGetCameraDriver = func(index int32) string {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraDriver",
			index,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetCurrentCameraDriver = func() string {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentCameraDriver",
		)

		return internal.UTF8JSToString(ret)
	}

	iGetCameras = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameras",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetCameraSupportedFormats = func(instance_id CameraID, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraSupportedFormats",
			_instance_id,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetCameraName = func(instance_id CameraID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraName",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetCameraPosition = func(instance_id CameraID) CameraPosition {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraPosition",
			int32(instance_id),
		)

		return CameraPosition(ret.Int())
	}

	iOpenCamera = func(instance_id CameraID, spec *CameraSpec) *Camera {
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		_spec := internal.CloneObjectToJSStack(spec)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenCamera",
			_instance_id,
			_spec,
		)

		return internal.NewObject[Camera](ret)
	}

	iGetCameraPermissionState = func(camera *Camera) CameraPermissionState {
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraPermissionState",
			_camera,
		)

		return CameraPermissionState(ret.Int())
	}

	iGetCameraID = func(camera *Camera) CameraID {
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraID",
			_camera,
		)

		return CameraID(ret.Int())
	}

	iGetCameraProperties = func(camera *Camera) PropertiesID {
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraProperties",
			_camera,
		)

		return PropertiesID(ret.Int())
	}

	iGetCameraFormat = func(camera *Camera, spec *CameraSpec) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		_spec := internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCameraFormat",
			_camera,
			_spec,
		)
		internal.CopyJSToObject(spec, _spec)

		return internal.GetBool(ret)
	}

	iAcquireCameraFrame = func(camera *Camera, timestampNS *uint64) *Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		_timestampNS := internal.StackAlloc(int(unsafe.Sizeof(*timestampNS)))
		ret := js.Global().Get("Module").Call(
			"_SDL_AcquireCameraFrame",
			_camera,
			_timestampNS,
		)
		if timestampNS != nil {
			*timestampNS = uint64(internal.GetValue(_timestampNS, "i64").Int())
		}

		return internal.NewObject[Surface](ret)
	}

	iReleaseCameraFrame = func(camera *Camera, frame *Surface) {
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		_frame, ok := internal.GetJSPointer(frame)
		if !ok {
			panic("nil frame")
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseCameraFrame",
			_camera,
			_frame,
		)
	}

	iCloseCamera = func(camera *Camera) {
		_camera, ok := internal.GetJSPointer(camera)
		if !ok {
			panic("nil camera")
		}
		js.Global().Get("Module").Call(
			"_SDL_CloseCamera",
			_camera,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(camera)))
	}

	iSetClipboardText = func(text string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_text := internal.StringOnJSStack(text)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetClipboardText",
			_text,
		)

		return internal.GetBool(ret)
	}

	iGetClipboardText = func() uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetClipboardText",
		)

		return uintptr(internal.GetInt64(ret))
	}

	iHasClipboardText = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasClipboardText",
		)

		return internal.GetBool(ret)
	}

	iSetPrimarySelectionText = func(text string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text := internal.StringOnJSStack(text)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetPrimarySelectionText",
			_text,
		)

		return internal.GetBool(ret)
	}

	iGetPrimarySelectionText = func() uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPrimarySelectionText",
		)

		return uintptr(internal.GetInt64(ret))
	}

	iHasPrimarySelectionText = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasPrimarySelectionText",
		)

		return internal.GetBool(ret)
	}

	/*iSetClipboardData = func(callback ClipboardDataCallback, cleanup ClipboardCleanupCallback, userdata uintptr, mime_types *string, num_mime_types uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_cleanup := int32(cleanup)
		_userdata := internal.NewBigInt(userdata)
		_mime_types, ok := internal.GetJSPointer(mime_types)
		if !ok {
			_mime_types = internal.StackAlloc()
		}
		_num_mime_types := internal.NewBigInt(num_mime_types)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetClipboardData",
			_callback,
			_cleanup,
			_userdata,
			_mime_types,
			_num_mime_types,
		)

		return internal.GetBool(ret)
	}*/

	iClearClipboardData = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_ClearClipboardData",
		)

		return internal.GetBool(ret)
	}

	iGetClipboardData = func(mime_type string, size *uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mime_type := internal.StringOnJSStack(mime_type)
		_size, ok := internal.GetJSPointer(size)
		if !ok {
			_size = internal.StackAlloc(int(unsafe.Sizeof(*size)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetClipboardData",
			_mime_type,
			_size,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iHasClipboardData = func(mime_type string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mime_type := internal.StringOnJSStack(mime_type)
		ret := js.Global().Get("Module").Call(
			"_SDL_HasClipboardData",
			_mime_type,
		)

		return internal.GetBool(ret)
	}

	iGetClipboardMimeTypes = func(num_mime_types *uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_num_mime_types, ok := internal.GetJSPointer(num_mime_types)
		if !ok {
			_num_mime_types = internal.StackAlloc(int(unsafe.Sizeof(*num_mime_types)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetClipboardMimeTypes",
			_num_mime_types,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetNumLogicalCPUCores = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumLogicalCPUCores",
		)

		return int32(ret.Int())
	}

	iGetCPUCacheLineSize = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCPUCacheLineSize",
		)

		return int32(ret.Int())
	}

	iHasAltiVec = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasAltiVec",
		)

		return internal.GetBool(ret)
	}

	iHasMMX = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasMMX",
		)

		return internal.GetBool(ret)
	}

	iHasSSE = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasSSE",
		)

		return internal.GetBool(ret)
	}

	iHasSSE2 = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasSSE2",
		)

		return internal.GetBool(ret)
	}

	iHasSSE3 = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasSSE3",
		)

		return internal.GetBool(ret)
	}

	iHasSSE41 = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasSSE41",
		)

		return internal.GetBool(ret)
	}

	iHasSSE42 = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasSSE42",
		)

		return internal.GetBool(ret)
	}

	iHasAVX = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasAVX",
		)

		return internal.GetBool(ret)
	}

	iHasAVX2 = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasAVX2",
		)

		return internal.GetBool(ret)
	}

	iHasAVX512F = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasAVX512F",
		)

		return internal.GetBool(ret)
	}

	iHasARMSIMD = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasARMSIMD",
		)

		return internal.GetBool(ret)
	}

	iHasNEON = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasNEON",
		)

		return internal.GetBool(ret)
	}

	iHasLSX = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasLSX",
		)

		return internal.GetBool(ret)
	}

	iHasLASX = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasLASX",
		)

		return internal.GetBool(ret)
	}

	iGetSystemRAM = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSystemRAM",
		)

		return int32(ret.Int())
	}

	iGetSIMDAlignment = func() uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSIMDAlignment",
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetNumVideoDrivers = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumVideoDrivers",
		)

		return int32(ret.Int())
	}

	iGetVideoDriver = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetVideoDriver",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetCurrentVideoDriver = func() string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentVideoDriver",
		)

		return internal.UTF8JSToString(ret)
	}

	iGetSystemTheme = func() SystemTheme {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSystemTheme",
		)

		return SystemTheme(ret.Int())
	}

	iGetDisplays = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplays",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetPrimaryDisplay = func() DisplayID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPrimaryDisplay",
		)

		return DisplayID(ret.Int())
	}

	iGetDisplayProperties = func(displayID DisplayID) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayProperties",
			_displayID,
		)

		return PropertiesID(ret.Int())
	}

	iGetDisplayName = func(displayID DisplayID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayName",
			_displayID,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetDisplayBounds = func(displayID DisplayID, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayBounds",
			_displayID,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iGetDisplayUsableBounds = func(displayID DisplayID, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayUsableBounds",
			_displayID,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iGetNaturalDisplayOrientation = func(displayID DisplayID) DisplayOrientation {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNaturalDisplayOrientation",
			_displayID,
		)

		return DisplayOrientation(ret.Int())
	}

	iGetCurrentDisplayOrientation = func(displayID DisplayID) DisplayOrientation {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentDisplayOrientation",
			_displayID,
		)

		return DisplayOrientation(ret.Int())
	}

	iGetDisplayContentScale = func(displayID DisplayID) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayContentScale",
			_displayID,
		)

		return float32(ret.Int())
	}

	iGetFullscreenDisplayModes = func(displayID DisplayID, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFullscreenDisplayModes",
			_displayID,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetClosestFullscreenDisplayMode = func(displayID DisplayID, w int32, h int32, refresh_rate float32, include_high_density_modes bool, closest *DisplayMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		_w := int32(w)
		_h := int32(h)
		_refresh_rate := int32(refresh_rate)
		_include_high_density_modes := internal.NewBoolean(include_high_density_modes)
		_closest, ok := internal.GetJSPointer(closest)
		if !ok {
			_closest = internal.StackAlloc(int(unsafe.Sizeof(*closest)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetClosestFullscreenDisplayMode",
			_displayID,
			_w,
			_h,
			_refresh_rate,
			_include_high_density_modes,
			_closest,
		)

		return internal.GetBool(ret)
	}

	iGetDesktopDisplayMode = func(displayID DisplayID) *DisplayMode {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDesktopDisplayMode",
			_displayID,
		)

		_obj := &DisplayMode{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetCurrentDisplayMode = func(displayID DisplayID) *DisplayMode {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_displayID := int32(displayID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentDisplayMode",
			_displayID,
		)

		_obj := &DisplayMode{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetDisplayForPoint = func(point *Point) DisplayID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_point, ok := internal.GetJSPointer(point)
		if !ok {
			_point = internal.StackAlloc(int(unsafe.Sizeof(*point)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayForPoint",
			_point,
		)

		return DisplayID(ret.Int())
	}

	iGetDisplayForRect = func(rect *Rect) DisplayID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayForRect",
			_rect,
		)

		return DisplayID(ret.Int())
	}

	iGetDisplayForWindow = func(window *Window) DisplayID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDisplayForWindow",
			_window,
		)

		return DisplayID(ret.Int())
	}

	iGetWindowPixelDensity = func(window *Window) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowPixelDensity",
			_window,
		)

		return float32(ret.Int())
	}

	iGetWindowDisplayScale = func(window *Window) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowDisplayScale",
			_window,
		)

		return float32(ret.Int())
	}

	iSetWindowFullscreenMode = func(window *Window, mode *DisplayMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_mode, ok := internal.GetJSPointer(mode)
		if !ok {
			_mode = internal.StackAlloc(int(unsafe.Sizeof(*mode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowFullscreenMode",
			_window,
			_mode,
		)

		return internal.GetBool(ret)
	}

	iGetWindowFullscreenMode = func(window *Window) *DisplayMode {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowFullscreenMode",
			_window,
		)

		_obj := &DisplayMode{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetWindowICCProfile = func(window *Window, size *uintptr) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_size, ok := internal.GetJSPointer(size)
		if !ok {
			_size = internal.StackAlloc(int(unsafe.Sizeof(*size)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowICCProfile",
			_window,
			_size,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetWindowPixelFormat = func(window *Window) PixelFormat {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowPixelFormat",
			_window,
		)

		return PixelFormat(ret.Int())
	}

	iGetWindows = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindows",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iCreateWindow = func(title string, w int32, h int32, flags WindowFlags) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_title := internal.StringOnJSStack(title)
		_w := int32(w)
		_h := int32(h)
		_flags := int32(flags)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateWindow",
			_title,
			_w,
			_h,
			_flags,
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreatePopupWindow = func(parent *Window, offset_x int32, offset_y int32, w int32, h int32, flags WindowFlags) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_parent, ok := internal.GetJSPointer(parent)
		if !ok {
			_parent = internal.StackAlloc(int(unsafe.Sizeof(*parent)))
		}
		_offset_x := int32(offset_x)
		_offset_y := int32(offset_y)
		_w := int32(w)
		_h := int32(h)
		_flags := int32(flags)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreatePopupWindow",
			_parent,
			_offset_x,
			_offset_y,
			_w,
			_h,
			_flags,
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateWindowWithProperties = func(props PropertiesID) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateWindowWithProperties",
			_props,
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetWindowID = func(window *Window) WindowID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowID",
			_window,
		)

		return WindowID(ret.Int())
	}

	iGetWindowFromID = func(id WindowID) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_id := int32(id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowFromID",
			_id,
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetWindowParent = func(window *Window) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowParent",
			_window,
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetWindowProperties = func(window *Window) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowProperties",
			_window,
		)

		return PropertiesID(ret.Int())
	}

	iGetWindowFlags = func(window *Window) WindowFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowFlags",
			_window,
		)

		return WindowFlags(ret.Int())
	}

	iSetWindowTitle = func(window *Window, title string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_title := internal.StringOnJSStack(title)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowTitle",
			_window,
			_title,
		)

		return internal.GetBool(ret)
	}

	iGetWindowTitle = func(window *Window) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowTitle",
			_window,
		)

		return internal.UTF8JSToString(ret)
	}

	iSetWindowIcon = func(window *Window, icon *Surface) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_icon, ok := internal.GetJSPointer(icon)
		if !ok {
			_icon = internal.StackAlloc(int(unsafe.Sizeof(*icon)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowIcon",
			_window,
			_icon,
		)

		return internal.GetBool(ret)
	}

	iSetWindowPosition = func(window *Window, x int32, y int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_x := int32(x)
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowPosition",
			_window,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iGetWindowPosition = func(window *Window, x *int32, y *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_x, ok := internal.GetJSPointer(x)
		if !ok {
			_x = internal.StackAlloc(int(unsafe.Sizeof(*x)))
		}
		_y, ok := internal.GetJSPointer(y)
		if !ok {
			_y = internal.StackAlloc(int(unsafe.Sizeof(*y)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowPosition",
			_window,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iSetWindowSize = func(window *Window, w int32, h int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_w := int32(w)
		_h := int32(h)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowSize",
			_window,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iGetWindowSize = func(window *Window, w *int32, h *int32) bool {
		internal.StackSave()
		defer internal.StackRestore()

		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		_w := internal.StackAlloc(4)
		_h := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowSize",
			_window,
			_w,
			_h,
		)
		*w = int32(internal.GetValue(_w, "i32").Int())
		*h = int32(internal.GetValue(_h, "i32").Int())

		return internal.GetBool(ret)
	}

	iGetWindowSafeArea = func(window *Window, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowSafeArea",
			_window,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iSetWindowAspectRatio = func(window *Window, min_aspect float32, max_aspect float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_min_aspect := int32(min_aspect)
		_max_aspect := int32(max_aspect)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowAspectRatio",
			_window,
			_min_aspect,
			_max_aspect,
		)

		return internal.GetBool(ret)
	}

	iGetWindowAspectRatio = func(window *Window, min_aspect *float32, max_aspect *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_min_aspect, ok := internal.GetJSPointer(min_aspect)
		if !ok {
			_min_aspect = internal.StackAlloc(int(unsafe.Sizeof(*min_aspect)))
		}
		_max_aspect, ok := internal.GetJSPointer(max_aspect)
		if !ok {
			_max_aspect = internal.StackAlloc(int(unsafe.Sizeof(*max_aspect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowAspectRatio",
			_window,
			_min_aspect,
			_max_aspect,
		)

		return internal.GetBool(ret)
	}

	iGetWindowBordersSize = func(window *Window, top *int32, left *int32, bottom *int32, right *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_top, ok := internal.GetJSPointer(top)
		if !ok {
			_top = internal.StackAlloc(int(unsafe.Sizeof(*top)))
		}
		_left, ok := internal.GetJSPointer(left)
		if !ok {
			_left = internal.StackAlloc(int(unsafe.Sizeof(*left)))
		}
		_bottom, ok := internal.GetJSPointer(bottom)
		if !ok {
			_bottom = internal.StackAlloc(int(unsafe.Sizeof(*bottom)))
		}
		_right, ok := internal.GetJSPointer(right)
		if !ok {
			_right = internal.StackAlloc(int(unsafe.Sizeof(*right)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowBordersSize",
			_window,
			_top,
			_left,
			_bottom,
			_right,
		)

		return internal.GetBool(ret)
	}

	iGetWindowSizeInPixels = func(window *Window, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowSizeInPixels",
			_window,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iSetWindowMinimumSize = func(window *Window, min_w int32, min_h int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_min_w := int32(min_w)
		_min_h := int32(min_h)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowMinimumSize",
			_window,
			_min_w,
			_min_h,
		)

		return internal.GetBool(ret)
	}

	iGetWindowMinimumSize = func(window *Window, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowMinimumSize",
			_window,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iSetWindowMaximumSize = func(window *Window, max_w int32, max_h int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_max_w := int32(max_w)
		_max_h := int32(max_h)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowMaximumSize",
			_window,
			_max_w,
			_max_h,
		)

		return internal.GetBool(ret)
	}

	iGetWindowMaximumSize = func(window *Window, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowMaximumSize",
			_window,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iSetWindowBordered = func(window *Window, bordered bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_bordered := internal.NewBoolean(bordered)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowBordered",
			_window,
			_bordered,
		)

		return internal.GetBool(ret)
	}

	iSetWindowResizable = func(window *Window, resizable bool) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		_resizable := internal.NewBoolean(resizable)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowResizable",
			_window,
			_resizable,
		)

		return internal.GetBool(ret)
	}

	iSetWindowAlwaysOnTop = func(window *Window, on_top bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_on_top := internal.NewBoolean(on_top)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowAlwaysOnTop",
			_window,
			_on_top,
		)

		return internal.GetBool(ret)
	}

	iShowWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ShowWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iHideWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_HideWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iRaiseWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RaiseWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iMaximizeWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_MaximizeWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iMinimizeWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_MinimizeWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iRestoreWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RestoreWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iSetWindowFullscreen = func(window *Window, fullscreen bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_fullscreen := internal.NewBoolean(fullscreen)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowFullscreen",
			_window,
			_fullscreen,
		)

		return internal.GetBool(ret)
	}

	iSyncWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SyncWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	iWindowHasSurface = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WindowHasSurface",
			_window,
		)

		return internal.GetBool(ret)
	}

	iGetWindowSurface = func(window *Window) *Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowSurface",
			_window,
		)

		_obj := &Surface{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iSetWindowSurfaceVSync = func(window *Window, vsync int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_vsync := int32(vsync)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowSurfaceVSync",
			_window,
			_vsync,
		)

		return internal.GetBool(ret)
	}

	iGetWindowSurfaceVSync = func(window *Window, vsync *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_vsync, ok := internal.GetJSPointer(vsync)
		if !ok {
			_vsync = internal.StackAlloc(int(unsafe.Sizeof(*vsync)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowSurfaceVSync",
			_window,
			_vsync,
		)

		return internal.GetBool(ret)
	}

	iUpdateWindowSurface = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateWindowSurface",
			_window,
		)

		return internal.GetBool(ret)
	}

	iUpdateWindowSurfaceRects = func(window *Window, rects *Rect, numrects int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_rects, ok := internal.GetJSPointer(rects)
		if !ok {
			_rects = internal.StackAlloc(int(unsafe.Sizeof(*rects)))
		}
		_numrects := int32(numrects)
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateWindowSurfaceRects",
			_window,
			_rects,
			_numrects,
		)

		return internal.GetBool(ret)
	}

	iDestroyWindowSurface = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_DestroyWindowSurface",
			_window,
		)

		return internal.GetBool(ret)
	}

	iSetWindowKeyboardGrab = func(window *Window, grabbed bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_grabbed := internal.NewBoolean(grabbed)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowKeyboardGrab",
			_window,
			_grabbed,
		)

		return internal.GetBool(ret)
	}

	iSetWindowMouseGrab = func(window *Window, grabbed bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_grabbed := internal.NewBoolean(grabbed)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowMouseGrab",
			_window,
			_grabbed,
		)

		return internal.GetBool(ret)
	}

	iGetWindowKeyboardGrab = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowKeyboardGrab",
			_window,
		)

		return internal.GetBool(ret)
	}

	iGetWindowMouseGrab = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowMouseGrab",
			_window,
		)

		return internal.GetBool(ret)
	}

	iGetGrabbedWindow = func() *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGrabbedWindow",
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iSetWindowMouseRect = func(window *Window, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowMouseRect",
			_window,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iGetWindowMouseRect = func(window *Window) *Rect {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowMouseRect",
			_window,
		)

		_obj := &Rect{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iSetWindowOpacity = func(window *Window, opacity float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_opacity := int32(opacity)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowOpacity",
			_window,
			_opacity,
		)

		return internal.GetBool(ret)
	}

	iGetWindowOpacity = func(window *Window) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowOpacity",
			_window,
		)

		return float32(ret.Int())
	}

	iSetWindowParent = func(window *Window, parent *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_parent, ok := internal.GetJSPointer(parent)
		if !ok {
			_parent = internal.StackAlloc(int(unsafe.Sizeof(*parent)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowParent",
			_window,
			_parent,
		)

		return internal.GetBool(ret)
	}

	iSetWindowModal = func(window *Window, modal bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_modal := internal.NewBoolean(modal)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowModal",
			_window,
			_modal,
		)

		return internal.GetBool(ret)
	}

	iSetWindowFocusable = func(window *Window, focusable bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_focusable := internal.NewBoolean(focusable)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowFocusable",
			_window,
			_focusable,
		)

		return internal.GetBool(ret)
	}

	iShowWindowSystemMenu = func(window *Window, x int32, y int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_x := int32(x)
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_SDL_ShowWindowSystemMenu",
			_window,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iSetWindowHitTest = func(window *Window, callback HitTest, callback_data uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_callback := int32(callback)
		_callback_data := internal.NewBigInt(callback_data)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowHitTest",
			_window,
			_callback,
			_callback_data,
		)

		return internal.GetBool(ret)
	}

	iSetWindowShape = func(window *Window, shape *Surface) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_shape, ok := internal.GetJSPointer(shape)
		if !ok {
			_shape = internal.StackAlloc(int(unsafe.Sizeof(*shape)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowShape",
			_window,
			_shape,
		)

		return internal.GetBool(ret)
	}

	iFlashWindow = func(window *Window, operation FlashOperation) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_operation := int32(operation)
		ret := js.Global().Get("Module").Call(
			"_SDL_FlashWindow",
			_window,
			_operation,
		)

		return internal.GetBool(ret)
	}

	iDestroyWindow = func(window *Window) {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			return
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyWindow",
			_window,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(window)))
	}

	iScreenSaverEnabled = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_ScreenSaverEnabled",
		)

		return internal.GetBool(ret)
	}

	iEnableScreenSaver = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_EnableScreenSaver",
		)

		return internal.GetBool(ret)
	}

	iDisableScreenSaver = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_DisableScreenSaver",
		)

		return internal.GetBool(ret)
	}

	iGL_LoadLibrary = func(path string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_LoadLibrary",
			_path,
		)

		return internal.GetBool(ret)
	}

	/*iGL_GetProcAddress = func(proc string) FunctionPointer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_proc := internal.StringOnStackGoToJS(proc)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_GetProcAddress",
			_proc,
		)

		return FunctionPointer(ret.Int())
	}*/

	/*iEGL_GetProcAddress = func(proc string) FunctionPointer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_proc := internal.StringOnStackGoToJS(proc)
		ret := js.Global().Get("Module").Call(
			"_SDL_EGL_GetProcAddress",
			_proc,
		)

		return FunctionPointer(ret.Int())
	}*/

	iGL_UnloadLibrary = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_GL_UnloadLibrary",
		)
	}

	iGL_ExtensionSupported = func(extension string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_extension := internal.StringOnJSStack(extension)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_ExtensionSupported",
			_extension,
		)

		return internal.GetBool(ret)
	}

	iGL_ResetAttributes = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_GL_ResetAttributes",
		)
	}

	iGL_SetAttribute = func(attr GLAttr, value int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_attr := int32(attr)
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_SetAttribute",
			_attr,
			_value,
		)

		return internal.GetBool(ret)
	}

	iGL_GetAttribute = func(attr GLAttr, value *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_attr := int32(attr)
		_value, ok := internal.GetJSPointer(value)
		if !ok {
			_value = internal.StackAlloc(int(unsafe.Sizeof(*value)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_GetAttribute",
			_attr,
			_value,
		)

		return internal.GetBool(ret)
	}

	/*iGL_CreateContext = func(window *Window) GLContext {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_CreateContext",
			_window,
		)

		return GLContext(ret.Int())
	}*/

	/*iGL_MakeCurrent = func(window *Window, context GLContext) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_context := int32(context)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_MakeCurrent",
			_window,
			_context,
		)

		return internal.GetBool(ret)
	}*/

	iGL_GetCurrentWindow = func() *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_GetCurrentWindow",
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	/*iGL_GetCurrentContext = func() GLContext {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_GetCurrentContext",
		)

		return GLContext(ret.Int())
	}*/

	iEGL_GetCurrentDisplay = func() EGLDisplay {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_EGL_GetCurrentDisplay",
		)

		return EGLDisplay(ret.Int())
	}

	iEGL_GetCurrentConfig = func() EGLConfig {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_EGL_GetCurrentConfig",
		)

		return EGLConfig(ret.Int())
	}

	iEGL_GetWindowSurface = func(window *Window) EGLSurface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_EGL_GetWindowSurface",
			_window,
		)

		return EGLSurface(ret.Int())
	}

	/*iEGL_SetAttributeCallbacks = func(platformAttribCallback EGLAttribArrayCallback, surfaceAttribCallback EGLIntArrayCallback, contextAttribCallback EGLIntArrayCallback, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_platformAttribCallback := int32(platformAttribCallback)
		_surfaceAttribCallback := int32(surfaceAttribCallback)
		_contextAttribCallback := int32(contextAttribCallback)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_EGL_SetAttributeCallbacks",
			_platformAttribCallback,
			_surfaceAttribCallback,
			_contextAttribCallback,
			_userdata,
		)
	}*/

	iGL_SetSwapInterval = func(interval int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_interval := int32(interval)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_SetSwapInterval",
			_interval,
		)

		return internal.GetBool(ret)
	}

	iGL_GetSwapInterval = func(interval *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_interval, ok := internal.GetJSPointer(interval)
		if !ok {
			_interval = internal.StackAlloc(int(unsafe.Sizeof(*interval)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_GetSwapInterval",
			_interval,
		)

		return internal.GetBool(ret)
	}

	iGL_SwapWindow = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_SwapWindow",
			_window,
		)

		return internal.GetBool(ret)
	}

	/*iGL_DestroyContext = func(context GLContext) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_context := int32(context)
		ret := js.Global().Get("Module").Call(
			"_SDL_GL_DestroyContext",
			_context,
		)

		return internal.GetBool(ret)
	}*/

	/*iShowOpenFileDialog = func(callback DialogFileCallback, userdata uintptr, window *Window, filters *DialogFileFilter, nfilters int32, default_location string, allow_many bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_filters, ok := internal.GetJSPointer(filters)
		if !ok {
			_filters = internal.StackAlloc(int(unsafe.Sizeof(*filters)))
		}
		_nfilters := int32(nfilters)
		_default_location := internal.StringOnStackGoToJS(default_location)
		_allow_many := internal.NewBoolean(allow_many)
		js.Global().Get("Module").Call(
			"_SDL_ShowOpenFileDialog",
			_callback,
			_userdata,
			_window,
			_filters,
			_nfilters,
			_default_location,
			_allow_many,
		)
	}*/

	/*iShowSaveFileDialog = func(callback DialogFileCallback, userdata uintptr, window *Window, filters *DialogFileFilter, nfilters int32, default_location string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_filters, ok := internal.GetJSPointer(filters)
		if !ok {
			_filters = internal.StackAlloc(int(unsafe.Sizeof(*filters)))
		}
		_nfilters := int32(nfilters)
		_default_location := internal.StringOnStackGoToJS(default_location)
		js.Global().Get("Module").Call(
			"_SDL_ShowSaveFileDialog",
			_callback,
			_userdata,
			_window,
			_filters,
			_nfilters,
			_default_location,
		)
	}*/

	/*iShowOpenFolderDialog = func(callback DialogFileCallback, userdata uintptr, window *Window, default_location string, allow_many bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_default_location := internal.StringOnStackGoToJS(default_location)
		_allow_many := internal.NewBoolean(allow_many)
		js.Global().Get("Module").Call(
			"_SDL_ShowOpenFolderDialog",
			_callback,
			_userdata,
			_window,
			_default_location,
			_allow_many,
		)
	}*/

	/*iShowFileDialogWithProperties = func(typ FileDialogType, callback DialogFileCallback, userdata uintptr, props PropertiesID) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		_props := int32(props)
		js.Global().Get("Module").Call(
			"_SDL_ShowFileDialogWithProperties",
			_typ,
			_callback,
			_userdata,
			_props,
		)
	}*/

	/*iGUIDToString = func(guid GUID, pszGUID string, cbGUID int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_guid := int32(guid)
		_pszGUID := internal.StringOnStackGoToJS(pszGUID)
		_cbGUID := int32(cbGUID)
		js.Global().Get("Module").Call(
			"_SDL_GUIDToString",
			_guid,
			_pszGUID,
			_cbGUID,
		)
	}*/

	/*iStringToGUID = func(pchGUID string) GUID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_pchGUID := internal.StringOnStackGoToJS(pchGUID)
		ret := js.Global().Get("Module").Call(
			"_SDL_StringToGUID",
			_pchGUID,
		)

		return GUID(ret.Int())
	}*/

	iGetPowerInfo = func(seconds *int32, percent *int32) PowerState {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_seconds, ok := internal.GetJSPointer(seconds)
		if !ok {
			_seconds = internal.StackAlloc(int(unsafe.Sizeof(*seconds)))
		}
		_percent, ok := internal.GetJSPointer(percent)
		if !ok {
			_percent = internal.StackAlloc(int(unsafe.Sizeof(*percent)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPowerInfo",
			_seconds,
			_percent,
		)

		return PowerState(ret.Int())
	}

	iGetSensors = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensors",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetSensorNameForID = func(instance_id SensorID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorNameForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetSensorTypeForID = func(instance_id SensorID) SensorType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorTypeForID",
			_instance_id,
		)

		return SensorType(ret.Int())
	}

	iGetSensorNonPortableTypeForID = func(instance_id SensorID) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorNonPortableTypeForID",
			_instance_id,
		)

		return int32(ret.Int())
	}

	iOpenSensor = func(instance_id SensorID) *Sensor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenSensor",
			_instance_id,
		)

		_obj := &Sensor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetSensorFromID = func(instance_id SensorID) *Sensor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorFromID",
			_instance_id,
		)

		_obj := &Sensor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetSensorProperties = func(sensor *Sensor) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorProperties",
			_sensor,
		)

		return PropertiesID(ret.Int())
	}

	iGetSensorName = func(sensor *Sensor) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorName",
			_sensor,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetSensorType = func(sensor *Sensor) SensorType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorType",
			_sensor,
		)

		return SensorType(ret.Int())
	}

	iGetSensorNonPortableType = func(sensor *Sensor) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorNonPortableType",
			_sensor,
		)

		return int32(ret.Int())
	}

	iGetSensorID = func(sensor *Sensor) SensorID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorID",
			_sensor,
		)

		return SensorID(ret.Int())
	}

	iGetSensorData = func(sensor *Sensor, data *float32, num_values int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		_data, ok := internal.GetJSPointer(data)
		if !ok {
			_data = internal.StackAlloc(int(unsafe.Sizeof(*data)))
		}
		_num_values := int32(num_values)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSensorData",
			_sensor,
			_data,
			_num_values,
		)

		return internal.GetBool(ret)
	}

	iCloseSensor = func(sensor *Sensor) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sensor, ok := internal.GetJSPointer(sensor)
		if !ok {
			_sensor = internal.StackAlloc(int(unsafe.Sizeof(*sensor)))
		}
		js.Global().Get("Module").Call(
			"_SDL_CloseSensor",
			_sensor,
		)
	}

	iUpdateSensors = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_UpdateSensors",
		)
	}

	iLockJoysticks = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_LockJoysticks",
		)
	}

	iUnlockJoysticks = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_UnlockJoysticks",
		)
	}

	iHasJoystick = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasJoystick",
		)

		return internal.GetBool(ret)
	}

	iGetJoysticks = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoysticks",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetJoystickNameForID = func(instance_id JoystickID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickNameForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetJoystickPathForID = func(instance_id JoystickID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickPathForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetJoystickPlayerIndexForID = func(instance_id JoystickID) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickPlayerIndexForID",
			_instance_id,
		)

		return int32(ret.Int())
	}

	/*iGetJoystickGUIDForID = func(instance_id JoystickID) GUID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickGUIDForID",
			_instance_id,
		)

		return GUID(ret.Int())
	}*/

	iGetJoystickVendorForID = func(instance_id JoystickID) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickVendorForID",
			_instance_id,
		)

		return uint16(ret.Int())
	}

	iGetJoystickProductForID = func(instance_id JoystickID) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickProductForID",
			_instance_id,
		)

		return uint16(ret.Int())
	}

	iGetJoystickProductVersionForID = func(instance_id JoystickID) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickProductVersionForID",
			_instance_id,
		)

		return uint16(ret.Int())
	}

	iGetJoystickTypeForID = func(instance_id JoystickID) JoystickType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickTypeForID",
			_instance_id,
		)

		return JoystickType(ret.Int())
	}

	iOpenJoystick = func(instance_id JoystickID) *Joystick {
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenJoystick",
			_instance_id,
		)

		_obj := internal.NewObject[Joystick](ret)

		return _obj
	}

	iGetJoystickFromID = func(instance_id JoystickID) *Joystick {
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickFromID",
			_instance_id,
		)

		// TODO: Add a getPointerFromJSValue, pretty sure the same
		// joystick is returned for the same id
		_obj := internal.NewObject[Joystick](ret)

		return _obj
	}

	iGetJoystickFromPlayerIndex = func(player_index int32) *Joystick {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_player_index := int32(player_index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickFromPlayerIndex",
			_player_index,
		)

		_obj := &Joystick{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iAttachVirtualJoystick = func(desc *virtualJoystickDesc) JoystickID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_desc, ok := internal.GetJSPointer(desc)
		if !ok {
			_desc = internal.StackAlloc(int(unsafe.Sizeof(*desc)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_AttachVirtualJoystick",
			_desc,
		)

		return JoystickID(ret.Int())
	}

	iDetachVirtualJoystick = func(instance_id JoystickID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_DetachVirtualJoystick",
			_instance_id,
		)

		return internal.GetBool(ret)
	}

	iIsJoystickVirtual = func(instance_id JoystickID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_IsJoystickVirtual",
			_instance_id,
		)

		return internal.GetBool(ret)
	}

	iSetJoystickVirtualAxis = func(joystick *Joystick, axis int32, value int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_axis := int32(axis)
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickVirtualAxis",
			_joystick,
			_axis,
			_value,
		)

		return internal.GetBool(ret)
	}

	iSetJoystickVirtualBall = func(joystick *Joystick, ball int32, xrel int16, yrel int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_ball := int32(ball)
		_xrel := int32(xrel)
		_yrel := int32(yrel)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickVirtualBall",
			_joystick,
			_ball,
			_xrel,
			_yrel,
		)

		return internal.GetBool(ret)
	}

	iSetJoystickVirtualButton = func(joystick *Joystick, button int32, down bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_button := int32(button)
		_down := internal.NewBoolean(down)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickVirtualButton",
			_joystick,
			_button,
			_down,
		)

		return internal.GetBool(ret)
	}

	iSetJoystickVirtualHat = func(joystick *Joystick, hat int32, value uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_hat := int32(hat)
		_value := int32(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickVirtualHat",
			_joystick,
			_hat,
			_value,
		)

		return internal.GetBool(ret)
	}

	iSetJoystickVirtualTouchpad = func(joystick *Joystick, touchpad int32, finger int32, down bool, x float32, y float32, pressure float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_touchpad := int32(touchpad)
		_finger := int32(finger)
		_down := internal.NewBoolean(down)
		_x := int32(x)
		_y := int32(y)
		_pressure := int32(pressure)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickVirtualTouchpad",
			_joystick,
			_touchpad,
			_finger,
			_down,
			_x,
			_y,
			_pressure,
		)

		return internal.GetBool(ret)
	}

	iSendJoystickVirtualSensorData = func(joystick *Joystick, typ SensorType, sensor_timestamp uint64, data *float32, num_values int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_typ := int32(typ)
		_sensor_timestamp := internal.NewBigInt(sensor_timestamp)
		_data, ok := internal.GetJSPointer(data)
		if !ok {
			_data = internal.StackAlloc(int(unsafe.Sizeof(*data)))
		}
		_num_values := int32(num_values)
		ret := js.Global().Get("Module").Call(
			"_SDL_SendJoystickVirtualSensorData",
			_joystick,
			_typ,
			_sensor_timestamp,
			_data,
			_num_values,
		)

		return internal.GetBool(ret)
	}

	iGetJoystickProperties = func(joystick *Joystick) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickProperties",
			_joystick,
		)

		return PropertiesID(ret.Int())
	}

	iGetJoystickName = func(joystick *Joystick) string {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickName",
			_joystick,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetJoystickPath = func(joystick *Joystick) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickPath",
			_joystick,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetJoystickPlayerIndex = func(joystick *Joystick) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickPlayerIndex",
			_joystick,
		)

		return int32(ret.Int())
	}

	iSetJoystickPlayerIndex = func(joystick *Joystick, player_index int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_player_index := int32(player_index)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickPlayerIndex",
			_joystick,
			_player_index,
		)

		return internal.GetBool(ret)
	}

	/*iGetJoystickGUID = func(joystick *Joystick) GUID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickGUID",
			_joystick,
		)

		return GUID(ret.Int())
	}*/

	iGetJoystickVendor = func(joystick *Joystick) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickVendor",
			_joystick,
		)

		return uint16(ret.Int())
	}

	iGetJoystickProduct = func(joystick *Joystick) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickProduct",
			_joystick,
		)

		return uint16(ret.Int())
	}

	iGetJoystickProductVersion = func(joystick *Joystick) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickProductVersion",
			_joystick,
		)

		return uint16(ret.Int())
	}

	iGetJoystickFirmwareVersion = func(joystick *Joystick) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickFirmwareVersion",
			_joystick,
		)

		return uint16(ret.Int())
	}

	iGetJoystickSerial = func(joystick *Joystick) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickSerial",
			_joystick,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetJoystickType = func(joystick *Joystick) JoystickType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickType",
			_joystick,
		)

		return JoystickType(ret.Int())
	}

	/*iGetJoystickGUIDInfo = func(guid GUID, vendor *uint16, product *uint16, version *uint16, crc16 *uint16) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_guid := int32(guid)
		_vendor, ok := internal.GetJSPointer(vendor)
		if !ok {
			_vendor = internal.StackAlloc(int(unsafe.Sizeof(*vendor)))
		}
		_product, ok := internal.GetJSPointer(product)
		if !ok {
			_product = internal.StackAlloc(int(unsafe.Sizeof(*product)))
		}
		_version, ok := internal.GetJSPointer(version)
		if !ok {
			_version = internal.StackAlloc(int(unsafe.Sizeof(*version)))
		}
		_crc16, ok := internal.GetJSPointer(crc16)
		if !ok {
			_crc16 = internal.StackAlloc(int(unsafe.Sizeof(*crc16)))
		}
		js.Global().Get("Module").Call(
			"_SDL_GetJoystickGUIDInfo",
			_guid,
			_vendor,
			_product,
			_version,
			_crc16,
		)
	}*/

	iJoystickConnected = func(joystick *Joystick) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_JoystickConnected",
			_joystick,
		)

		return internal.GetBool(ret)
	}

	iGetJoystickID = func(joystick *Joystick) JoystickID {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickID",
			_joystick,
		)

		return JoystickID(ret.Int())
	}

	iGetNumJoystickAxes = func(joystick *Joystick) int32 {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumJoystickAxes",
			_joystick,
		)

		return int32(ret.Int())
	}

	iGetNumJoystickBalls = func(joystick *Joystick) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumJoystickBalls",
			_joystick,
		)

		return int32(ret.Int())
	}

	iGetNumJoystickHats = func(joystick *Joystick) int32 {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumJoystickHats",
			_joystick,
		)

		return int32(ret.Int())
	}

	iGetNumJoystickButtons = func(joystick *Joystick) int32 {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumJoystickButtons",
			_joystick,
		)

		return int32(ret.Int())
	}

	iSetJoystickEventsEnabled = func(enabled bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_enabled := internal.NewBoolean(enabled)
		js.Global().Get("Module").Call(
			"_SDL_SetJoystickEventsEnabled",
			_enabled,
		)
	}

	iJoystickEventsEnabled = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_JoystickEventsEnabled",
		)

		return internal.GetBool(ret)
	}

	iUpdateJoysticks = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_UpdateJoysticks",
		)
	}

	iGetJoystickAxis = func(joystick *Joystick, axis int32) int16 {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		_axis := int32(axis)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickAxis",
			_joystick,
			_axis,
		)

		return int16(ret.Int())
	}

	iGetJoystickAxisInitialState = func(joystick *Joystick, axis int32, state *int16) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_axis := int32(axis)
		_state, ok := internal.GetJSPointer(state)
		if !ok {
			_state = internal.StackAlloc(int(unsafe.Sizeof(*state)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickAxisInitialState",
			_joystick,
			_axis,
			_state,
		)

		return internal.GetBool(ret)
	}

	iGetJoystickBall = func(joystick *Joystick, ball int32, dx *int32, dy *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_ball := int32(ball)
		_dx, ok := internal.GetJSPointer(dx)
		if !ok {
			_dx = internal.StackAlloc(int(unsafe.Sizeof(*dx)))
		}
		_dy, ok := internal.GetJSPointer(dy)
		if !ok {
			_dy = internal.StackAlloc(int(unsafe.Sizeof(*dy)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickBall",
			_joystick,
			_ball,
			_dx,
			_dy,
		)

		return internal.GetBool(ret)
	}

	iGetJoystickHat = func(joystick *Joystick, hat int32) uint8 {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		_hat := int32(hat)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickHat",
			_joystick,
			_hat,
		)

		return uint8(ret.Int())
	}

	iGetJoystickButton = func(joystick *Joystick, button int32) bool {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickButton",
			_joystick,
			_button,
		)

		return internal.GetBool(ret)
	}

	iRumbleJoystick = func(joystick *Joystick, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_low_frequency_rumble := int32(low_frequency_rumble)
		_high_frequency_rumble := int32(high_frequency_rumble)
		_duration_ms := int32(duration_ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_RumbleJoystick",
			_joystick,
			_low_frequency_rumble,
			_high_frequency_rumble,
			_duration_ms,
		)

		return internal.GetBool(ret)
	}

	iRumbleJoystickTriggers = func(joystick *Joystick, left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_left_rumble := int32(left_rumble)
		_right_rumble := int32(right_rumble)
		_duration_ms := int32(duration_ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_RumbleJoystickTriggers",
			_joystick,
			_left_rumble,
			_right_rumble,
			_duration_ms,
		)

		return internal.GetBool(ret)
	}

	iSetJoystickLED = func(joystick *Joystick, red uint8, green uint8, blue uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_red := int32(red)
		_green := int32(green)
		_blue := int32(blue)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetJoystickLED",
			_joystick,
			_red,
			_green,
			_blue,
		)

		return internal.GetBool(ret)
	}

	iSendJoystickEffect = func(joystick *Joystick, data uintptr, size int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_data := internal.NewBigInt(data)
		_size := int32(size)
		ret := js.Global().Get("Module").Call(
			"_SDL_SendJoystickEffect",
			_joystick,
			_data,
			_size,
		)

		return internal.GetBool(ret)
	}

	iCloseJoystick = func(joystick *Joystick) {
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			panic("nil joystick")
		}
		js.Global().Get("Module").Call(
			"_SDL_CloseJoystick",
			_joystick,
		)

		internal.DeleteJSPointer(uintptr(unsafe.Pointer(joystick)))
	}

	iGetJoystickConnectionState = func(joystick *Joystick) JoystickConnectionState {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickConnectionState",
			_joystick,
		)

		return JoystickConnectionState(ret.Int())
	}

	iGetJoystickPowerInfo = func(joystick *Joystick, percent *int32) PowerState {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		_percent, ok := internal.GetJSPointer(percent)
		if !ok {
			_percent = internal.StackAlloc(int(unsafe.Sizeof(*percent)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetJoystickPowerInfo",
			_joystick,
			_percent,
		)

		return PowerState(ret.Int())
	}

	iAddGamepadMapping = func(mapping string) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mapping := internal.StringOnJSStack(mapping)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddGamepadMapping",
			_mapping,
		)

		return int32(ret.Int())
	}

	iAddGamepadMappingsFromIO = func(src *IOStream, closeio bool) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddGamepadMappingsFromIO",
			_src,
			_closeio,
		)

		return int32(ret.Int())
	}

	iAddGamepadMappingsFromFile = func(file string) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddGamepadMappingsFromFile",
			_file,
		)

		return int32(ret.Int())
	}

	iReloadGamepadMappings = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_ReloadGamepadMappings",
		)

		return internal.GetBool(ret)
	}

	iGetGamepadMappings = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadMappings",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	/*iGetGamepadMappingForGUID = func(guid GUID) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_guid := int32(guid)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadMappingForGUID",
			_guid,
		)

		return uintptr(internal.GetInt64(ret))
	}*/

	iGetGamepadMapping = func(gamepad *Gamepad) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadMapping",
			_gamepad,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iSetGamepadMapping = func(instance_id JoystickID, mapping *byte) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		_mapping := internal.StringOnJSStack(internal.PtrToString(uintptr(unsafe.Pointer(mapping))))
		ret := js.Global().Get("Module").Call(
			"_SDL_SetGamepadMapping",
			_instance_id,
			_mapping,
		)

		runtime.KeepAlive(mapping)

		return internal.GetBool(ret)
	}

	iHasGamepad = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasGamepad",
		)

		return internal.GetBool(ret)
	}

	iGetGamepads = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepads",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iIsGamepad = func(instance_id JoystickID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_IsGamepad",
			_instance_id,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadNameForID = func(instance_id JoystickID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadNameForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadPathForID = func(instance_id JoystickID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadPathForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadPlayerIndexForID = func(instance_id JoystickID) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadPlayerIndexForID",
			_instance_id,
		)

		return int32(ret.Int())
	}

	/*iGetGamepadGUIDForID = func(instance_id JoystickID) GUID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadGUIDForID",
			_instance_id,
		)

		return GUID(ret.Int())
	}*/

	iGetGamepadVendorForID = func(instance_id JoystickID) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadVendorForID",
			_instance_id,
		)

		return uint16(ret.Int())
	}

	iGetGamepadProductForID = func(instance_id JoystickID) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadProductForID",
			_instance_id,
		)

		return uint16(ret.Int())
	}

	iGetGamepadProductVersionForID = func(instance_id JoystickID) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadProductVersionForID",
			_instance_id,
		)

		return uint16(ret.Int())
	}

	iGetGamepadTypeForID = func(instance_id JoystickID) GamepadType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadTypeForID",
			_instance_id,
		)

		return GamepadType(ret.Int())
	}

	iGetRealGamepadTypeForID = func(instance_id JoystickID) GamepadType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRealGamepadTypeForID",
			_instance_id,
		)

		return GamepadType(ret.Int())
	}

	iGetGamepadMappingForID = func(instance_id JoystickID) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadMappingForID",
			_instance_id,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iOpenGamepad = func(instance_id JoystickID) *Gamepad {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenGamepad",
			_instance_id,
		)

		_obj := &Gamepad{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetGamepadFromID = func(instance_id JoystickID) *Gamepad {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadFromID",
			_instance_id,
		)

		_obj := &Gamepad{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetGamepadFromPlayerIndex = func(player_index int32) *Gamepad {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_player_index := int32(player_index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadFromPlayerIndex",
			_player_index,
		)

		_obj := &Gamepad{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetGamepadProperties = func(gamepad *Gamepad) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadProperties",
			_gamepad,
		)

		return PropertiesID(ret.Int())
	}

	iGetGamepadID = func(gamepad *Gamepad) JoystickID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadID",
			_gamepad,
		)

		return JoystickID(ret.Int())
	}

	iGetGamepadName = func(gamepad *Gamepad) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadName",
			_gamepad,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadPath = func(gamepad *Gamepad) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadPath",
			_gamepad,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadType = func(gamepad *Gamepad) GamepadType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadType",
			_gamepad,
		)

		return GamepadType(ret.Int())
	}

	iGetRealGamepadType = func(gamepad *Gamepad) GamepadType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRealGamepadType",
			_gamepad,
		)

		return GamepadType(ret.Int())
	}

	iGetGamepadPlayerIndex = func(gamepad *Gamepad) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadPlayerIndex",
			_gamepad,
		)

		return int32(ret.Int())
	}

	iSetGamepadPlayerIndex = func(gamepad *Gamepad, player_index int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_player_index := int32(player_index)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetGamepadPlayerIndex",
			_gamepad,
			_player_index,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadVendor = func(gamepad *Gamepad) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadVendor",
			_gamepad,
		)

		return uint16(ret.Int())
	}

	iGetGamepadProduct = func(gamepad *Gamepad) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadProduct",
			_gamepad,
		)

		return uint16(ret.Int())
	}

	iGetGamepadProductVersion = func(gamepad *Gamepad) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadProductVersion",
			_gamepad,
		)

		return uint16(ret.Int())
	}

	iGetGamepadFirmwareVersion = func(gamepad *Gamepad) uint16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadFirmwareVersion",
			_gamepad,
		)

		return uint16(ret.Int())
	}

	iGetGamepadSerial = func(gamepad *Gamepad) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadSerial",
			_gamepad,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadSteamHandle = func(gamepad *Gamepad) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadSteamHandle",
			_gamepad,
		)

		return uint64(internal.GetInt64(ret))
	}

	iGetGamepadConnectionState = func(gamepad *Gamepad) JoystickConnectionState {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadConnectionState",
			_gamepad,
		)

		return JoystickConnectionState(ret.Int())
	}

	iGetGamepadPowerInfo = func(gamepad *Gamepad, percent *int32) PowerState {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_percent, ok := internal.GetJSPointer(percent)
		if !ok {
			_percent = internal.StackAlloc(int(unsafe.Sizeof(*percent)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadPowerInfo",
			_gamepad,
			_percent,
		)

		return PowerState(ret.Int())
	}

	iGamepadConnected = func(gamepad *Gamepad) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GamepadConnected",
			_gamepad,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadJoystick = func(gamepad *Gamepad) *Joystick {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadJoystick",
			_gamepad,
		)

		_obj := &Joystick{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iSetGamepadEventsEnabled = func(enabled bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_enabled := internal.NewBoolean(enabled)
		js.Global().Get("Module").Call(
			"_SDL_SetGamepadEventsEnabled",
			_enabled,
		)
	}

	iGamepadEventsEnabled = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GamepadEventsEnabled",
		)

		return internal.GetBool(ret)
	}

	iGetGamepadBindings = func(gamepad *Gamepad, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadBindings",
			_gamepad,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iUpdateGamepads = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_UpdateGamepads",
		)
	}

	iGetGamepadTypeFromString = func(str string) GamepadType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_str := internal.StringOnJSStack(str)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadTypeFromString",
			_str,
		)

		return GamepadType(ret.Int())
	}

	iGetGamepadStringForType = func(typ GamepadType) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadStringForType",
			_typ,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadAxisFromString = func(str string) GamepadAxis {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_str := internal.StringOnJSStack(str)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadAxisFromString",
			_str,
		)

		return GamepadAxis(ret.Int())
	}

	iGetGamepadStringForAxis = func(axis GamepadAxis) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_axis := int32(axis)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadStringForAxis",
			_axis,
		)

		return internal.UTF8JSToString(ret)
	}

	iGamepadHasAxis = func(gamepad *Gamepad, axis GamepadAxis) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_axis := int32(axis)
		ret := js.Global().Get("Module").Call(
			"_SDL_GamepadHasAxis",
			_gamepad,
			_axis,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadAxis = func(gamepad *Gamepad, axis GamepadAxis) int16 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_axis := int32(axis)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadAxis",
			_gamepad,
			_axis,
		)

		return int16(ret.Int())
	}

	iGetGamepadButtonFromString = func(str string) GamepadButton {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_str := internal.StringOnJSStack(str)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadButtonFromString",
			_str,
		)

		return GamepadButton(ret.Int())
	}

	iGetGamepadStringForButton = func(button GamepadButton) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadStringForButton",
			_button,
		)

		return internal.UTF8JSToString(ret)
	}

	iGamepadHasButton = func(gamepad *Gamepad, button GamepadButton) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GamepadHasButton",
			_gamepad,
			_button,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadButton = func(gamepad *Gamepad, button GamepadButton) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadButton",
			_gamepad,
			_button,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadButtonLabelForType = func(typ GamepadType, button GamepadButton) GamepadButtonLabel {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadButtonLabelForType",
			_typ,
			_button,
		)

		return GamepadButtonLabel(ret.Int())
	}

	iGetGamepadButtonLabel = func(gamepad *Gamepad, button GamepadButton) GamepadButtonLabel {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadButtonLabel",
			_gamepad,
			_button,
		)

		return GamepadButtonLabel(ret.Int())
	}

	iGetNumGamepadTouchpads = func(gamepad *Gamepad) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumGamepadTouchpads",
			_gamepad,
		)

		return int32(ret.Int())
	}

	iGetNumGamepadTouchpadFingers = func(gamepad *Gamepad, touchpad int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_touchpad := int32(touchpad)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumGamepadTouchpadFingers",
			_gamepad,
			_touchpad,
		)

		return int32(ret.Int())
	}

	iGetGamepadTouchpadFinger = func(gamepad *Gamepad, touchpad int32, finger int32, down *bool, x *float32, y *float32, pressure *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_touchpad := int32(touchpad)
		_finger := int32(finger)
		_down, ok := internal.GetJSPointer(down)
		if !ok {
			_down = internal.StackAlloc(int(unsafe.Sizeof(*down)))
		}
		_x, ok := internal.GetJSPointer(x)
		if !ok {
			_x = internal.StackAlloc(int(unsafe.Sizeof(*x)))
		}
		_y, ok := internal.GetJSPointer(y)
		if !ok {
			_y = internal.StackAlloc(int(unsafe.Sizeof(*y)))
		}
		_pressure, ok := internal.GetJSPointer(pressure)
		if !ok {
			_pressure = internal.StackAlloc(int(unsafe.Sizeof(*pressure)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadTouchpadFinger",
			_gamepad,
			_touchpad,
			_finger,
			_down,
			_x,
			_y,
			_pressure,
		)

		return internal.GetBool(ret)
	}

	iGamepadHasSensor = func(gamepad *Gamepad, typ SensorType) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_typ := int32(typ)
		ret := js.Global().Get("Module").Call(
			"_SDL_GamepadHasSensor",
			_gamepad,
			_typ,
		)

		return internal.GetBool(ret)
	}

	iSetGamepadSensorEnabled = func(gamepad *Gamepad, typ SensorType, enabled bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_typ := int32(typ)
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetGamepadSensorEnabled",
			_gamepad,
			_typ,
			_enabled,
		)

		return internal.GetBool(ret)
	}

	iGamepadSensorEnabled = func(gamepad *Gamepad, typ SensorType) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_typ := int32(typ)
		ret := js.Global().Get("Module").Call(
			"_SDL_GamepadSensorEnabled",
			_gamepad,
			_typ,
		)

		return internal.GetBool(ret)
	}

	iGetGamepadSensorDataRate = func(gamepad *Gamepad, typ SensorType) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_typ := int32(typ)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadSensorDataRate",
			_gamepad,
			_typ,
		)

		return float32(ret.Int())
	}

	iGetGamepadSensorData = func(gamepad *Gamepad, typ SensorType, data *float32, num_values int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_typ := int32(typ)
		_data, ok := internal.GetJSPointer(data)
		if !ok {
			_data = internal.StackAlloc(int(unsafe.Sizeof(*data)))
		}
		_num_values := int32(num_values)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadSensorData",
			_gamepad,
			_typ,
			_data,
			_num_values,
		)

		return internal.GetBool(ret)
	}

	iRumbleGamepad = func(gamepad *Gamepad, low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_low_frequency_rumble := int32(low_frequency_rumble)
		_high_frequency_rumble := int32(high_frequency_rumble)
		_duration_ms := int32(duration_ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_RumbleGamepad",
			_gamepad,
			_low_frequency_rumble,
			_high_frequency_rumble,
			_duration_ms,
		)

		return internal.GetBool(ret)
	}

	iRumbleGamepadTriggers = func(gamepad *Gamepad, left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_left_rumble := int32(left_rumble)
		_right_rumble := int32(right_rumble)
		_duration_ms := int32(duration_ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_RumbleGamepadTriggers",
			_gamepad,
			_left_rumble,
			_right_rumble,
			_duration_ms,
		)

		return internal.GetBool(ret)
	}

	iSetGamepadLED = func(gamepad *Gamepad, red uint8, green uint8, blue uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_red := int32(red)
		_green := int32(green)
		_blue := int32(blue)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetGamepadLED",
			_gamepad,
			_red,
			_green,
			_blue,
		)

		return internal.GetBool(ret)
	}

	iSendGamepadEffect = func(gamepad *Gamepad, data uintptr, size int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_data := internal.NewBigInt(data)
		_size := int32(size)
		ret := js.Global().Get("Module").Call(
			"_SDL_SendGamepadEffect",
			_gamepad,
			_data,
			_size,
		)

		return internal.GetBool(ret)
	}

	iCloseGamepad = func(gamepad *Gamepad) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		js.Global().Get("Module").Call(
			"_SDL_CloseGamepad",
			_gamepad,
		)
	}

	iGetGamepadAppleSFSymbolsNameForButton = func(gamepad *Gamepad, button GamepadButton) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_button := int32(button)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadAppleSFSymbolsNameForButton",
			_gamepad,
			_button,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGamepadAppleSFSymbolsNameForAxis = func(gamepad *Gamepad, axis GamepadAxis) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_gamepad, ok := internal.GetJSPointer(gamepad)
		if !ok {
			_gamepad = internal.StackAlloc(int(unsafe.Sizeof(*gamepad)))
		}
		_axis := int32(axis)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGamepadAppleSFSymbolsNameForAxis",
			_gamepad,
			_axis,
		)

		return internal.UTF8JSToString(ret)
	}

	iHasKeyboard = func() bool {
		ret := js.Global().Get("Module").Call(
			"_SDL_HasKeyboard",
		)

		return internal.GetBool(ret)
	}

	iGetKeyboards = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyboards",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetKeyboardNameForID = func(instance_id KeyboardID) string {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyboardNameForID",
			int32(instance_id),
		)

		return internal.UTF8JSToString(ret)
	}

	iGetKeyboardFocus = func() *Window {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyboardFocus",
		)

		return internal.NewObject[Window](ret)
	}

	/*iGetKeyboardState = func(numkeys *int32) *bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_numkeys, ok := internal.GetJSPointer(numkeys)
		if !ok {
			_numkeys = internal.StackAlloc(int(unsafe.Sizeof(*numkeys)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyboardState",
			_numkeys,
		)

		_obj := &bool{}
		internal.StoreJSPointer(_obj, ret)
		return _obj
	}*/

	iResetKeyboard = func() {
		js.Global().Get("Module").Call(
			"_SDL_ResetKeyboard",
		)
	}

	iGetModState = func() Keymod {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetModState",
		)

		return Keymod(ret.Int())
	}

	iSetModState = func(modstate Keymod) {
		js.Global().Get("Module").Call(
			"_SDL_SetModState",
			int32(modstate),
		)
	}

	iGetKeyFromScancode = func(scancode Scancode, modstate Keymod, key_event bool) Keycode {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_scancode := int32(scancode)
		_modstate := int32(modstate)
		_key_event := internal.NewBoolean(key_event)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyFromScancode",
			_scancode,
			_modstate,
			_key_event,
		)

		return Keycode(ret.Int())
	}

	iGetScancodeFromKey = func(key Keycode, modstate *Keymod) Scancode {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_key := int32(key)
		_modstate, ok := internal.GetJSPointer(modstate)
		if !ok {
			_modstate = internal.StackAlloc(int(unsafe.Sizeof(*modstate)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetScancodeFromKey",
			_key,
			_modstate,
		)

		return Scancode(ret.Int())
	}

	iSetScancodeName = func(scancode Scancode, name string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetScancodeName",
			int32(scancode),
			_name,
		)

		return internal.GetBool(ret)
	}

	iGetScancodeName = func(scancode Scancode) string {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetScancodeName",
			int32(scancode),
		)

		return internal.UTF8JSToString(ret)
	}

	iGetScancodeFromName = func(name string) Scancode {
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetScancodeFromName",
			_name,
		)

		return Scancode(ret.Int())
	}

	iGetKeyName = func(key Keycode) string {
		_key := int32(key)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyName",
			_key,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetKeyFromName = func(name string) Keycode {
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetKeyFromName",
			_name,
		)

		return Keycode(ret.Int())
	}

	iStartTextInput = func(window *Window) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_StartTextInput",
			_window,
		)

		return internal.GetBool(ret)
	}

	iStartTextInputWithProperties = func(window *Window, props PropertiesID) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_StartTextInputWithProperties",
			_window,
			int32(props),
		)

		return internal.GetBool(ret)
	}

	iTextInputActive = func(window *Window) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TextInputActive",
			_window,
		)

		return internal.GetBool(ret)
	}

	iStopTextInput = func(window *Window) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_StopTextInput",
			_window,
		)

		return internal.GetBool(ret)
	}

	iClearComposition = func(window *Window) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ClearComposition",
			_window,
		)

		return internal.GetBool(ret)
	}

	iSetTextInputArea = func(window *Window, rect *Rect, cursor int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextInputArea",
			_window,
			_rect,
			cursor,
		)

		return internal.GetBool(ret)
	}

	iGetTextInputArea = func(window *Window, rect *Rect, cursor *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		_cursor, ok := internal.GetJSPointer(cursor)
		if !ok {
			_cursor = internal.StackAlloc(int(unsafe.Sizeof(*cursor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextInputArea",
			_window,
			_rect,
			_cursor,
		)

		return internal.GetBool(ret)
	}

	iHasScreenKeyboardSupport = func() bool {
		ret := js.Global().Get("Module").Call(
			"_SDL_HasScreenKeyboardSupport",
		)

		return internal.GetBool(ret)
	}

	iScreenKeyboardShown = func(window *Window) bool {
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			panic("nil window")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ScreenKeyboardShown",
			_window,
		)

		return internal.GetBool(ret)
	}

	iHasMouse = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HasMouse",
		)

		return internal.GetBool(ret)
	}

	iGetMice = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMice",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetMouseNameForID = func(instance_id MouseID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMouseNameForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetMouseFocus = func() *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMouseFocus",
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetMouseState = func(x *float32, y *float32) MouseButtonFlags {
		internal.StackSave()
		defer internal.StackRestore()
		_x := internal.StackAlloc(int(unsafe.Sizeof(float32(0))))
		_y := internal.StackAlloc(int(unsafe.Sizeof(float32(0))))
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMouseState",
			_x,
			_y,
		)
		*x = float32(internal.GetValue(_x, "float").Float())
		*y = float32(internal.GetValue(_y, "float").Float())

		return MouseButtonFlags(ret.Int())
	}

	iGetGlobalMouseState = func(x *float32, y *float32) MouseButtonFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_x, ok := internal.GetJSPointer(x)
		if !ok {
			_x = internal.StackAlloc(int(unsafe.Sizeof(*x)))
		}
		_y, ok := internal.GetJSPointer(y)
		if !ok {
			_y = internal.StackAlloc(int(unsafe.Sizeof(*y)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlobalMouseState",
			_x,
			_y,
		)

		return MouseButtonFlags(ret.Int())
	}

	iGetRelativeMouseState = func(x *float32, y *float32) MouseButtonFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_x, ok := internal.GetJSPointer(x)
		if !ok {
			_x = internal.StackAlloc(int(unsafe.Sizeof(*x)))
		}
		_y, ok := internal.GetJSPointer(y)
		if !ok {
			_y = internal.StackAlloc(int(unsafe.Sizeof(*y)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRelativeMouseState",
			_x,
			_y,
		)

		return MouseButtonFlags(ret.Int())
	}

	iWarpMouseInWindow = func(window *Window, x float32, y float32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_x := int32(x)
		_y := int32(y)
		js.Global().Get("Module").Call(
			"_SDL_WarpMouseInWindow",
			_window,
			_x,
			_y,
		)
	}

	iWarpMouseGlobal = func(x float32, y float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_x := int32(x)
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_SDL_WarpMouseGlobal",
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iSetWindowRelativeMouseMode = func(window *Window, enabled bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetWindowRelativeMouseMode",
			_window,
			_enabled,
		)

		return internal.GetBool(ret)
	}

	iGetWindowRelativeMouseMode = func(window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowRelativeMouseMode",
			_window,
		)

		return internal.GetBool(ret)
	}

	iCaptureMouse = func(enabled bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_SDL_CaptureMouse",
			_enabled,
		)

		return internal.GetBool(ret)
	}

	iCreateCursor = func(data *uint8, mask *uint8, w int32, h int32, hot_x int32, hot_y int32) *Cursor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_data, ok := internal.GetJSPointer(data)
		if !ok {
			_data = internal.StackAlloc(int(unsafe.Sizeof(*data)))
		}
		_mask, ok := internal.GetJSPointer(mask)
		if !ok {
			_mask = internal.StackAlloc(int(unsafe.Sizeof(*mask)))
		}
		_w := int32(w)
		_h := int32(h)
		_hot_x := int32(hot_x)
		_hot_y := int32(hot_y)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateCursor",
			_data,
			_mask,
			_w,
			_h,
			_hot_x,
			_hot_y,
		)

		_obj := &Cursor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateColorCursor = func(surface *Surface, hot_x int32, hot_y int32) *Cursor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_hot_x := int32(hot_x)
		_hot_y := int32(hot_y)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateColorCursor",
			_surface,
			_hot_x,
			_hot_y,
		)

		_obj := &Cursor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateSystemCursor = func(id SystemCursor) *Cursor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_id := int32(id)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSystemCursor",
			_id,
		)

		_obj := &Cursor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iSetCursor = func(cursor *Cursor) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cursor, ok := internal.GetJSPointer(cursor)
		if !ok {
			_cursor = internal.StackAlloc(int(unsafe.Sizeof(*cursor)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetCursor",
			_cursor,
		)

		return internal.GetBool(ret)
	}

	iGetCursor = func() *Cursor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCursor",
		)

		_obj := &Cursor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetDefaultCursor = func() *Cursor {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDefaultCursor",
		)

		_obj := &Cursor{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iDestroyCursor = func(cursor *Cursor) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_cursor, ok := internal.GetJSPointer(cursor)
		if !ok {
			_cursor = internal.StackAlloc(int(unsafe.Sizeof(*cursor)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyCursor",
			_cursor,
		)
	}

	iShowCursor = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_ShowCursor",
		)

		return internal.GetBool(ret)
	}

	iHideCursor = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_HideCursor",
		)

		return internal.GetBool(ret)
	}

	iCursorVisible = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_CursorVisible",
		)

		return internal.GetBool(ret)
	}

	iGetTouchDevices = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTouchDevices",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetTouchDeviceName = func(touchID TouchID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_touchID := int32(touchID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTouchDeviceName",
			_touchID,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetTouchDeviceType = func(touchID TouchID) TouchDeviceType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_touchID := int32(touchID)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTouchDeviceType",
			_touchID,
		)

		return TouchDeviceType(ret.Int())
	}

	iGetTouchFingers = func(touchID TouchID, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_touchID := int32(touchID)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTouchFingers",
			_touchID,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iPumpEvents = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_PumpEvents",
		)
	}

	iPeepEvents = func(events *Event, numevents int32, action EventAction, minType uint32, maxType uint32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_events, ok := internal.GetJSPointer(events)
		if !ok {
			_events = internal.StackAlloc(int(unsafe.Sizeof(*events)))
		}
		_numevents := int32(numevents)
		_action := int32(action)
		_minType := int32(minType)
		_maxType := int32(maxType)
		ret := js.Global().Get("Module").Call(
			"_SDL_PeepEvents",
			_events,
			_numevents,
			_action,
			_minType,
			_maxType,
		)

		return int32(ret.Int())
	}

	iHasEvent = func(typ uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		ret := js.Global().Get("Module").Call(
			"_SDL_HasEvent",
			_typ,
		)

		return internal.GetBool(ret)
	}

	iHasEvents = func(minType uint32, maxType uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_minType := int32(minType)
		_maxType := int32(maxType)
		ret := js.Global().Get("Module").Call(
			"_SDL_HasEvents",
			_minType,
			_maxType,
		)

		return internal.GetBool(ret)
	}

	iFlushEvent = func(typ uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		js.Global().Get("Module").Call(
			"_SDL_FlushEvent",
			_typ,
		)
	}

	iFlushEvents = func(minType uint32, maxType uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_minType := int32(minType)
		_maxType := int32(maxType)
		js.Global().Get("Module").Call(
			"_SDL_FlushEvents",
			_minType,
			_maxType,
		)
	}

	iPollEvent = func(event *Event) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_event := internal.StackAlloc(int(unsafe.Sizeof(*event)))
		ret := js.Global().Get("Module").Call(
			"_SDL_PollEvent",
			_event,
		)
		internal.CopyJSToObject(event, _event)

		return internal.GetBool(ret)
	}

	iWaitEvent = func(event *Event) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_event, ok := internal.GetJSPointer(event)
		if !ok {
			_event = internal.StackAlloc(int(unsafe.Sizeof(*event)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitEvent",
			_event,
		)

		return internal.GetBool(ret)
	}

	iWaitEventTimeout = func(event *Event, timeoutMS int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_event, ok := internal.GetJSPointer(event)
		if !ok {
			_event = internal.StackAlloc(int(unsafe.Sizeof(*event)))
		}
		_timeoutMS := int32(timeoutMS)
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitEventTimeout",
			_event,
			_timeoutMS,
		)

		return internal.GetBool(ret)
	}

	iPushEvent = func(event *Event) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_event, ok := internal.GetJSPointer(event)
		if !ok {
			_event = internal.StackAlloc(int(unsafe.Sizeof(*event)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_PushEvent",
			_event,
		)

		return internal.GetBool(ret)
	}

	/*iSetEventFilter = func(filter EventFilter, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_filter := int32(filter)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_SetEventFilter",
			_filter,
			_userdata,
		)
	}*/

	iGetEventFilter = func(filter *EventFilter, userdata *uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_filter, ok := internal.GetJSPointer(filter)
		if !ok {
			_filter = internal.StackAlloc(int(unsafe.Sizeof(*filter)))
		}
		_userdata, ok := internal.GetJSPointer(userdata)
		if !ok {
			_userdata = internal.StackAlloc(int(unsafe.Sizeof(*userdata)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetEventFilter",
			_filter,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	/*iAddEventWatch = func(filter EventFilter, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_filter := int32(filter)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddEventWatch",
			_filter,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	/*iRemoveEventWatch = func(filter EventFilter, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_filter := int32(filter)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_RemoveEventWatch",
			_filter,
			_userdata,
		)
	}*/

	/*iFilterEvents = func(filter EventFilter, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_filter := int32(filter)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_FilterEvents",
			_filter,
			_userdata,
		)
	}*/

	iSetEventEnabled = func(typ uint32, enabled bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		_enabled := internal.NewBoolean(enabled)
		js.Global().Get("Module").Call(
			"_SDL_SetEventEnabled",
			_typ,
			_enabled,
		)
	}

	iEventEnabled = func(typ uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_typ := int32(typ)
		ret := js.Global().Get("Module").Call(
			"_SDL_EventEnabled",
			_typ,
		)

		return internal.GetBool(ret)
	}

	iRegisterEvents = func(numevents int32) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_numevents := int32(numevents)
		ret := js.Global().Get("Module").Call(
			"_SDL_RegisterEvents",
			_numevents,
		)

		return uint32(ret.Int())
	}

	iGetWindowFromEvent = func(event *Event) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_event, ok := internal.GetJSPointer(event)
		if !ok {
			_event = internal.StackAlloc(int(unsafe.Sizeof(*event)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetWindowFromEvent",
			_event,
		)

		_obj := &Window{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetBasePath = func() string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetBasePath",
		)

		return internal.UTF8JSToString(ret)
	}

	iGetPrefPath = func(org string, app string) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_org := internal.StringOnJSStack(org)
		_app := internal.StringOnJSStack(app)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPrefPath",
			_org,
			_app,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetUserFolder = func(folder Folder) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_folder := int32(folder)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetUserFolder",
			_folder,
		)

		return internal.UTF8JSToString(ret)
	}

	iCreateDirectory = func(path string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateDirectory",
			_path,
		)

		return internal.GetBool(ret)
	}

	/*iEnumerateDirectory = func(path string, callback EnumerateDirectoryCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnStackGoToJS(path)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_EnumerateDirectory",
			_path,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	iRemovePath = func(path string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_RemovePath",
			_path,
		)

		return internal.GetBool(ret)
	}

	iRenamePath = func(oldpath string, newpath string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_oldpath := internal.StringOnJSStack(oldpath)
		_newpath := internal.StringOnJSStack(newpath)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenamePath",
			_oldpath,
			_newpath,
		)

		return internal.GetBool(ret)
	}

	iCopyFile = func(oldpath string, newpath string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_oldpath := internal.StringOnJSStack(oldpath)
		_newpath := internal.StringOnJSStack(newpath)
		ret := js.Global().Get("Module").Call(
			"_SDL_CopyFile",
			_oldpath,
			_newpath,
		)

		return internal.GetBool(ret)
	}

	iGetPathInfo = func(path string, info *PathInfo) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		_info, ok := internal.GetJSPointer(info)
		if !ok {
			_info = internal.StackAlloc(int(unsafe.Sizeof(*info)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPathInfo",
			_path,
			_info,
		)

		return internal.GetBool(ret)
	}

	iGlobDirectory = func(path string, pattern string, flags GlobFlags, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		_pattern := internal.StringOnJSStack(pattern)
		_flags := int32(flags)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GlobDirectory",
			_path,
			_pattern,
			_flags,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetCurrentDirectory = func() uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentDirectory",
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGPUSupportsShaderFormats = func(format_flags GPUShaderFormat, name *byte) bool {
		panic("not implemented on js")
		/*internal.StackSave()
		defer internal.StackRestore()
		_format_flags := int32(format_flags)
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_GPUSupportsShaderFormats",
			_format_flags,
			_name,
		)

		return internal.GetBool(ret)*/
	}

	iGPUSupportsProperties = func(props PropertiesID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_GPUSupportsProperties",
			_props,
		)

		return internal.GetBool(ret)
	}

	iCreateGPUDevice = func(format_flags GPUShaderFormat, debug_mode bool, name *byte) *GPUDevice {
		panic("not implemented on js")
		/*internal.StackSave()
		defer internal.StackRestore()
		_format_flags := int32(format_flags)
		_debug_mode := internal.NewBoolean(debug_mode)
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUDevice",
			_format_flags,
			_debug_mode,
			_name,
		)

		_obj := &GPUDevice{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj*/
	}

	iCreateGPUDeviceWithProperties = func(props PropertiesID) *GPUDevice {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUDeviceWithProperties",
			_props,
		)

		_obj := &GPUDevice{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iDestroyGPUDevice = func(device *GPUDevice) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyGPUDevice",
			_device,
		)
	}

	iGetNumGPUDrivers = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumGPUDrivers",
		)

		return int32(ret.Int())
	}

	iGetGPUDriver = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGPUDriver",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGPUDeviceDriver = func(device *GPUDevice) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGPUDeviceDriver",
			_device,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetGPUShaderFormats = func(device *GPUDevice) GPUShaderFormat {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGPUShaderFormats",
			_device,
		)

		return GPUShaderFormat(ret.Int())
	}

	iCreateGPUComputePipeline = func(device *GPUDevice, createinfo *gpuComputePipelineCreateInfo) *GPUComputePipeline {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUComputePipeline",
			_device,
			_createinfo,
		)

		_obj := &GPUComputePipeline{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateGPUGraphicsPipeline = func(device *GPUDevice, createinfo *gpuGraphicsPipelineCreateInfo) *GPUGraphicsPipeline {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUGraphicsPipeline",
			_device,
			_createinfo,
		)

		_obj := &GPUGraphicsPipeline{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateGPUSampler = func(device *GPUDevice, createinfo *GPUSamplerCreateInfo) *GPUSampler {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUSampler",
			_device,
			_createinfo,
		)

		_obj := &GPUSampler{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	/*iCreateGPUShader = func(device *GPUDevice, createinfo *GPUShaderCreateInfo) *GPUShader {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUShader",
			_device,
			_createinfo,
		)

		_obj := &GPUShader{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}*/

	iCreateGPUTexture = func(device *GPUDevice, createinfo *GPUTextureCreateInfo) *GPUTexture {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUTexture",
			_device,
			_createinfo,
		)

		_obj := &GPUTexture{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateGPUBuffer = func(device *GPUDevice, createinfo *GPUBufferCreateInfo) *GPUBuffer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUBuffer",
			_device,
			_createinfo,
		)

		_obj := &GPUBuffer{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCreateGPUTransferBuffer = func(device *GPUDevice, createinfo *GPUTransferBufferCreateInfo) *GPUTransferBuffer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_createinfo, ok := internal.GetJSPointer(createinfo)
		if !ok {
			_createinfo = internal.StackAlloc(int(unsafe.Sizeof(*createinfo)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUTransferBuffer",
			_device,
			_createinfo,
		)

		_obj := &GPUTransferBuffer{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iSetGPUBufferName = func(device *GPUDevice, buffer *GPUBuffer, text string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_buffer, ok := internal.GetJSPointer(buffer)
		if !ok {
			_buffer = internal.StackAlloc(int(unsafe.Sizeof(*buffer)))
		}
		_text := internal.StringOnJSStack(text)
		js.Global().Get("Module").Call(
			"_SDL_SetGPUBufferName",
			_device,
			_buffer,
			_text,
		)
	}

	iSetGPUTextureName = func(device *GPUDevice, texture *GPUTexture, text string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_text := internal.StringOnJSStack(text)
		js.Global().Get("Module").Call(
			"_SDL_SetGPUTextureName",
			_device,
			_texture,
			_text,
		)
	}

	iInsertGPUDebugLabel = func(command_buffer *GPUCommandBuffer, text string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_text := internal.StringOnJSStack(text)
		js.Global().Get("Module").Call(
			"_SDL_InsertGPUDebugLabel",
			_command_buffer,
			_text,
		)
	}

	iPushGPUDebugGroup = func(command_buffer *GPUCommandBuffer, name string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_name := internal.StringOnJSStack(name)
		js.Global().Get("Module").Call(
			"_SDL_PushGPUDebugGroup",
			_command_buffer,
			_name,
		)
	}

	iPopGPUDebugGroup = func(command_buffer *GPUCommandBuffer) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		js.Global().Get("Module").Call(
			"_SDL_PopGPUDebugGroup",
			_command_buffer,
		)
	}

	iReleaseGPUTexture = func(device *GPUDevice, texture *GPUTexture) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUTexture",
			_device,
			_texture,
		)
	}

	iReleaseGPUSampler = func(device *GPUDevice, sampler *GPUSampler) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_sampler, ok := internal.GetJSPointer(sampler)
		if !ok {
			_sampler = internal.StackAlloc(int(unsafe.Sizeof(*sampler)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUSampler",
			_device,
			_sampler,
		)
	}

	iReleaseGPUBuffer = func(device *GPUDevice, buffer *GPUBuffer) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_buffer, ok := internal.GetJSPointer(buffer)
		if !ok {
			_buffer = internal.StackAlloc(int(unsafe.Sizeof(*buffer)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUBuffer",
			_device,
			_buffer,
		)
	}

	iReleaseGPUTransferBuffer = func(device *GPUDevice, transfer_buffer *GPUTransferBuffer) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_transfer_buffer, ok := internal.GetJSPointer(transfer_buffer)
		if !ok {
			_transfer_buffer = internal.StackAlloc(int(unsafe.Sizeof(*transfer_buffer)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUTransferBuffer",
			_device,
			_transfer_buffer,
		)
	}

	iReleaseGPUComputePipeline = func(device *GPUDevice, compute_pipeline *GPUComputePipeline) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_compute_pipeline, ok := internal.GetJSPointer(compute_pipeline)
		if !ok {
			_compute_pipeline = internal.StackAlloc(int(unsafe.Sizeof(*compute_pipeline)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUComputePipeline",
			_device,
			_compute_pipeline,
		)
	}

	iReleaseGPUShader = func(device *GPUDevice, shader *GPUShader) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_shader, ok := internal.GetJSPointer(shader)
		if !ok {
			_shader = internal.StackAlloc(int(unsafe.Sizeof(*shader)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUShader",
			_device,
			_shader,
		)
	}

	iReleaseGPUGraphicsPipeline = func(device *GPUDevice, graphics_pipeline *GPUGraphicsPipeline) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_graphics_pipeline, ok := internal.GetJSPointer(graphics_pipeline)
		if !ok {
			_graphics_pipeline = internal.StackAlloc(int(unsafe.Sizeof(*graphics_pipeline)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUGraphicsPipeline",
			_device,
			_graphics_pipeline,
		)
	}

	iAcquireGPUCommandBuffer = func(device *GPUDevice) *GPUCommandBuffer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_AcquireGPUCommandBuffer",
			_device,
		)

		_obj := &GPUCommandBuffer{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iPushGPUVertexUniformData = func(command_buffer *GPUCommandBuffer, slot_index uint32, data uintptr, length uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_slot_index := int32(slot_index)
		_data := internal.NewBigInt(data)
		_length := int32(length)
		js.Global().Get("Module").Call(
			"_SDL_PushGPUVertexUniformData",
			_command_buffer,
			_slot_index,
			_data,
			_length,
		)
	}

	iPushGPUFragmentUniformData = func(command_buffer *GPUCommandBuffer, slot_index uint32, data uintptr, length uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_slot_index := int32(slot_index)
		_data := internal.NewBigInt(data)
		_length := int32(length)
		js.Global().Get("Module").Call(
			"_SDL_PushGPUFragmentUniformData",
			_command_buffer,
			_slot_index,
			_data,
			_length,
		)
	}

	iPushGPUComputeUniformData = func(command_buffer *GPUCommandBuffer, slot_index uint32, data uintptr, length uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_slot_index := int32(slot_index)
		_data := internal.NewBigInt(data)
		_length := int32(length)
		js.Global().Get("Module").Call(
			"_SDL_PushGPUComputeUniformData",
			_command_buffer,
			_slot_index,
			_data,
			_length,
		)
	}

	iBeginGPURenderPass = func(command_buffer *GPUCommandBuffer, color_target_infos *GPUColorTargetInfo, num_color_targets uint32, depth_stencil_target_info *GPUDepthStencilTargetInfo) *GPURenderPass {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_color_target_infos, ok := internal.GetJSPointer(color_target_infos)
		if !ok {
			_color_target_infos = internal.StackAlloc(int(unsafe.Sizeof(*color_target_infos)))
		}
		_num_color_targets := int32(num_color_targets)
		_depth_stencil_target_info, ok := internal.GetJSPointer(depth_stencil_target_info)
		if !ok {
			_depth_stencil_target_info = internal.StackAlloc(int(unsafe.Sizeof(*depth_stencil_target_info)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_BeginGPURenderPass",
			_command_buffer,
			_color_target_infos,
			_num_color_targets,
			_depth_stencil_target_info,
		)

		_obj := &GPURenderPass{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iBindGPUGraphicsPipeline = func(render_pass *GPURenderPass, graphics_pipeline *GPUGraphicsPipeline) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_graphics_pipeline, ok := internal.GetJSPointer(graphics_pipeline)
		if !ok {
			_graphics_pipeline = internal.StackAlloc(int(unsafe.Sizeof(*graphics_pipeline)))
		}
		js.Global().Get("Module").Call(
			"_SDL_BindGPUGraphicsPipeline",
			_render_pass,
			_graphics_pipeline,
		)
	}

	iSetGPUViewport = func(render_pass *GPURenderPass, viewport *GPUViewport) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_viewport, ok := internal.GetJSPointer(viewport)
		if !ok {
			_viewport = internal.StackAlloc(int(unsafe.Sizeof(*viewport)))
		}
		js.Global().Get("Module").Call(
			"_SDL_SetGPUViewport",
			_render_pass,
			_viewport,
		)
	}

	iSetGPUScissor = func(render_pass *GPURenderPass, scissor *Rect) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_scissor, ok := internal.GetJSPointer(scissor)
		if !ok {
			_scissor = internal.StackAlloc(int(unsafe.Sizeof(*scissor)))
		}
		js.Global().Get("Module").Call(
			"_SDL_SetGPUScissor",
			_render_pass,
			_scissor,
		)
	}

	iSetGPUStencilReference = func(render_pass *GPURenderPass, reference uint8) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_reference := int32(reference)
		js.Global().Get("Module").Call(
			"_SDL_SetGPUStencilReference",
			_render_pass,
			_reference,
		)
	}

	iBindGPUVertexBuffers = func(render_pass *GPURenderPass, first_slot uint32, bindings *GPUBufferBinding, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_bindings, ok := internal.GetJSPointer(bindings)
		if !ok {
			_bindings = internal.StackAlloc(int(unsafe.Sizeof(*bindings)))
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUVertexBuffers",
			_render_pass,
			_first_slot,
			_bindings,
			_num_bindings,
		)
	}

	iBindGPUIndexBuffer = func(render_pass *GPURenderPass, binding *GPUBufferBinding, index_element_size GPUIndexElementSize) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_binding, ok := internal.GetJSPointer(binding)
		if !ok {
			_binding = internal.StackAlloc(int(unsafe.Sizeof(*binding)))
		}
		_index_element_size := int32(index_element_size)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUIndexBuffer",
			_render_pass,
			_binding,
			_index_element_size,
		)
	}

	iBindGPUVertexSamplers = func(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_texture_sampler_bindings, ok := internal.GetJSPointer(texture_sampler_bindings)
		if !ok {
			_texture_sampler_bindings = internal.StackAlloc(int(unsafe.Sizeof(*texture_sampler_bindings)))
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUVertexSamplers",
			_render_pass,
			_first_slot,
			_texture_sampler_bindings,
			_num_bindings,
		)
	}

	iBindGPUVertexStorageTextures = func(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_storage_textures, ok := internal.GetJSPointer(storage_textures)
		if !ok {
			_storage_textures = internal.StackAlloc(4)
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUVertexStorageTextures",
			_render_pass,
			_first_slot,
			_storage_textures,
			_num_bindings,
		)
	}

	iBindGPUVertexStorageBuffers = func(render_pass *GPURenderPass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_storage_buffers, ok := internal.GetJSPointer(storage_buffers)
		if !ok {
			_storage_buffers = internal.StackAlloc(4)
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUVertexStorageBuffers",
			_render_pass,
			_first_slot,
			_storage_buffers,
			_num_bindings,
		)
	}

	iBindGPUFragmentSamplers = func(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_texture_sampler_bindings, ok := internal.GetJSPointer(texture_sampler_bindings)
		if !ok {
			_texture_sampler_bindings = internal.StackAlloc(int(unsafe.Sizeof(*texture_sampler_bindings)))
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUFragmentSamplers",
			_render_pass,
			_first_slot,
			_texture_sampler_bindings,
			_num_bindings,
		)
	}

	iBindGPUFragmentStorageTextures = func(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_storage_textures, ok := internal.GetJSPointer(storage_textures)
		if !ok {
			_storage_textures = internal.StackAlloc(4)
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUFragmentStorageTextures",
			_render_pass,
			_first_slot,
			_storage_textures,
			_num_bindings,
		)
	}

	iBindGPUFragmentStorageBuffers = func(render_pass *GPURenderPass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_first_slot := int32(first_slot)
		_storage_buffers, ok := internal.GetJSPointer(storage_buffers)
		if !ok {
			_storage_buffers = internal.StackAlloc(4)
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUFragmentStorageBuffers",
			_render_pass,
			_first_slot,
			_storage_buffers,
			_num_bindings,
		)
	}

	iDrawGPUIndexedPrimitives = func(render_pass *GPURenderPass, num_indices uint32, num_instances uint32, first_index uint32, vertex_offset int32, first_instance uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_num_indices := int32(num_indices)
		_num_instances := int32(num_instances)
		_first_index := int32(first_index)
		_vertex_offset := int32(vertex_offset)
		_first_instance := int32(first_instance)
		js.Global().Get("Module").Call(
			"_SDL_DrawGPUIndexedPrimitives",
			_render_pass,
			_num_indices,
			_num_instances,
			_first_index,
			_vertex_offset,
			_first_instance,
		)
	}

	iDrawGPUPrimitives = func(render_pass *GPURenderPass, num_vertices uint32, num_instances uint32, first_vertex uint32, first_instance uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_num_vertices := int32(num_vertices)
		_num_instances := int32(num_instances)
		_first_vertex := int32(first_vertex)
		_first_instance := int32(first_instance)
		js.Global().Get("Module").Call(
			"_SDL_DrawGPUPrimitives",
			_render_pass,
			_num_vertices,
			_num_instances,
			_first_vertex,
			_first_instance,
		)
	}

	iDrawGPUPrimitivesIndirect = func(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_buffer, ok := internal.GetJSPointer(buffer)
		if !ok {
			_buffer = internal.StackAlloc(int(unsafe.Sizeof(*buffer)))
		}
		_offset := int32(offset)
		_draw_count := int32(draw_count)
		js.Global().Get("Module").Call(
			"_SDL_DrawGPUPrimitivesIndirect",
			_render_pass,
			_buffer,
			_offset,
			_draw_count,
		)
	}

	iDrawGPUIndexedPrimitivesIndirect = func(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		_buffer, ok := internal.GetJSPointer(buffer)
		if !ok {
			_buffer = internal.StackAlloc(int(unsafe.Sizeof(*buffer)))
		}
		_offset := int32(offset)
		_draw_count := int32(draw_count)
		js.Global().Get("Module").Call(
			"_SDL_DrawGPUIndexedPrimitivesIndirect",
			_render_pass,
			_buffer,
			_offset,
			_draw_count,
		)
	}

	iEndGPURenderPass = func(render_pass *GPURenderPass) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_render_pass, ok := internal.GetJSPointer(render_pass)
		if !ok {
			_render_pass = internal.StackAlloc(int(unsafe.Sizeof(*render_pass)))
		}
		js.Global().Get("Module").Call(
			"_SDL_EndGPURenderPass",
			_render_pass,
		)
	}

	iBeginGPUComputePass = func(command_buffer *GPUCommandBuffer, storage_texture_bindings *GPUStorageTextureReadWriteBinding, num_storage_texture_bindings uint32, storage_buffer_bindings *GPUStorageBufferReadWriteBinding, num_storage_buffer_bindings uint32) *GPUComputePass {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_storage_texture_bindings, ok := internal.GetJSPointer(storage_texture_bindings)
		if !ok {
			_storage_texture_bindings = internal.StackAlloc(int(unsafe.Sizeof(*storage_texture_bindings)))
		}
		_num_storage_texture_bindings := int32(num_storage_texture_bindings)
		_storage_buffer_bindings, ok := internal.GetJSPointer(storage_buffer_bindings)
		if !ok {
			_storage_buffer_bindings = internal.StackAlloc(int(unsafe.Sizeof(*storage_buffer_bindings)))
		}
		_num_storage_buffer_bindings := int32(num_storage_buffer_bindings)
		ret := js.Global().Get("Module").Call(
			"_SDL_BeginGPUComputePass",
			_command_buffer,
			_storage_texture_bindings,
			_num_storage_texture_bindings,
			_storage_buffer_bindings,
			_num_storage_buffer_bindings,
		)

		_obj := &GPUComputePass{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iBindGPUComputePipeline = func(compute_pass *GPUComputePass, compute_pipeline *GPUComputePipeline) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		_compute_pipeline, ok := internal.GetJSPointer(compute_pipeline)
		if !ok {
			_compute_pipeline = internal.StackAlloc(int(unsafe.Sizeof(*compute_pipeline)))
		}
		js.Global().Get("Module").Call(
			"_SDL_BindGPUComputePipeline",
			_compute_pass,
			_compute_pipeline,
		)
	}

	iBindGPUComputeSamplers = func(compute_pass *GPUComputePass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		_first_slot := int32(first_slot)
		_texture_sampler_bindings, ok := internal.GetJSPointer(texture_sampler_bindings)
		if !ok {
			_texture_sampler_bindings = internal.StackAlloc(int(unsafe.Sizeof(*texture_sampler_bindings)))
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUComputeSamplers",
			_compute_pass,
			_first_slot,
			_texture_sampler_bindings,
			_num_bindings,
		)
	}

	iBindGPUComputeStorageTextures = func(compute_pass *GPUComputePass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		_first_slot := int32(first_slot)
		_storage_textures, ok := internal.GetJSPointer(storage_textures)
		if !ok {
			_storage_textures = internal.StackAlloc(4)
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUComputeStorageTextures",
			_compute_pass,
			_first_slot,
			_storage_textures,
			_num_bindings,
		)
	}

	iBindGPUComputeStorageBuffers = func(compute_pass *GPUComputePass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		_first_slot := int32(first_slot)
		_storage_buffers, ok := internal.GetJSPointer(storage_buffers)
		if !ok {
			_storage_buffers = internal.StackAlloc(4)
		}
		_num_bindings := int32(num_bindings)
		js.Global().Get("Module").Call(
			"_SDL_BindGPUComputeStorageBuffers",
			_compute_pass,
			_first_slot,
			_storage_buffers,
			_num_bindings,
		)
	}

	iDispatchGPUCompute = func(compute_pass *GPUComputePass, groupcount_x uint32, groupcount_y uint32, groupcount_z uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		_groupcount_x := int32(groupcount_x)
		_groupcount_y := int32(groupcount_y)
		_groupcount_z := int32(groupcount_z)
		js.Global().Get("Module").Call(
			"_SDL_DispatchGPUCompute",
			_compute_pass,
			_groupcount_x,
			_groupcount_y,
			_groupcount_z,
		)
	}

	iDispatchGPUComputeIndirect = func(compute_pass *GPUComputePass, buffer *GPUBuffer, offset uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		_buffer, ok := internal.GetJSPointer(buffer)
		if !ok {
			_buffer = internal.StackAlloc(int(unsafe.Sizeof(*buffer)))
		}
		_offset := int32(offset)
		js.Global().Get("Module").Call(
			"_SDL_DispatchGPUComputeIndirect",
			_compute_pass,
			_buffer,
			_offset,
		)
	}

	iEndGPUComputePass = func(compute_pass *GPUComputePass) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_compute_pass, ok := internal.GetJSPointer(compute_pass)
		if !ok {
			_compute_pass = internal.StackAlloc(int(unsafe.Sizeof(*compute_pass)))
		}
		js.Global().Get("Module").Call(
			"_SDL_EndGPUComputePass",
			_compute_pass,
		)
	}

	iMapGPUTransferBuffer = func(device *GPUDevice, transfer_buffer *GPUTransferBuffer, cycle bool) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_transfer_buffer, ok := internal.GetJSPointer(transfer_buffer)
		if !ok {
			_transfer_buffer = internal.StackAlloc(int(unsafe.Sizeof(*transfer_buffer)))
		}
		_cycle := internal.NewBoolean(cycle)
		ret := js.Global().Get("Module").Call(
			"_SDL_MapGPUTransferBuffer",
			_device,
			_transfer_buffer,
			_cycle,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iUnmapGPUTransferBuffer = func(device *GPUDevice, transfer_buffer *GPUTransferBuffer) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_transfer_buffer, ok := internal.GetJSPointer(transfer_buffer)
		if !ok {
			_transfer_buffer = internal.StackAlloc(int(unsafe.Sizeof(*transfer_buffer)))
		}
		js.Global().Get("Module").Call(
			"_SDL_UnmapGPUTransferBuffer",
			_device,
			_transfer_buffer,
		)
	}

	iBeginGPUCopyPass = func(command_buffer *GPUCommandBuffer) *GPUCopyPass {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_BeginGPUCopyPass",
			_command_buffer,
		)

		_obj := &GPUCopyPass{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iUploadToGPUTexture = func(copy_pass *GPUCopyPass, source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		_source, ok := internal.GetJSPointer(source)
		if !ok {
			_source = internal.StackAlloc(int(unsafe.Sizeof(*source)))
		}
		_destination, ok := internal.GetJSPointer(destination)
		if !ok {
			_destination = internal.StackAlloc(int(unsafe.Sizeof(*destination)))
		}
		_cycle := internal.NewBoolean(cycle)
		js.Global().Get("Module").Call(
			"_SDL_UploadToGPUTexture",
			_copy_pass,
			_source,
			_destination,
			_cycle,
		)
	}

	iUploadToGPUBuffer = func(copy_pass *GPUCopyPass, source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		_source, ok := internal.GetJSPointer(source)
		if !ok {
			_source = internal.StackAlloc(int(unsafe.Sizeof(*source)))
		}
		_destination, ok := internal.GetJSPointer(destination)
		if !ok {
			_destination = internal.StackAlloc(int(unsafe.Sizeof(*destination)))
		}
		_cycle := internal.NewBoolean(cycle)
		js.Global().Get("Module").Call(
			"_SDL_UploadToGPUBuffer",
			_copy_pass,
			_source,
			_destination,
			_cycle,
		)
	}

	iCopyGPUTextureToTexture = func(copy_pass *GPUCopyPass, source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		_source, ok := internal.GetJSPointer(source)
		if !ok {
			_source = internal.StackAlloc(int(unsafe.Sizeof(*source)))
		}
		_destination, ok := internal.GetJSPointer(destination)
		if !ok {
			_destination = internal.StackAlloc(int(unsafe.Sizeof(*destination)))
		}
		_w := int32(w)
		_h := int32(h)
		_d := int32(d)
		_cycle := internal.NewBoolean(cycle)
		js.Global().Get("Module").Call(
			"_SDL_CopyGPUTextureToTexture",
			_copy_pass,
			_source,
			_destination,
			_w,
			_h,
			_d,
			_cycle,
		)
	}

	iCopyGPUBufferToBuffer = func(copy_pass *GPUCopyPass, source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		_source, ok := internal.GetJSPointer(source)
		if !ok {
			_source = internal.StackAlloc(int(unsafe.Sizeof(*source)))
		}
		_destination, ok := internal.GetJSPointer(destination)
		if !ok {
			_destination = internal.StackAlloc(int(unsafe.Sizeof(*destination)))
		}
		_size := int32(size)
		_cycle := internal.NewBoolean(cycle)
		js.Global().Get("Module").Call(
			"_SDL_CopyGPUBufferToBuffer",
			_copy_pass,
			_source,
			_destination,
			_size,
			_cycle,
		)
	}

	iDownloadFromGPUTexture = func(copy_pass *GPUCopyPass, source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		_source, ok := internal.GetJSPointer(source)
		if !ok {
			_source = internal.StackAlloc(int(unsafe.Sizeof(*source)))
		}
		_destination, ok := internal.GetJSPointer(destination)
		if !ok {
			_destination = internal.StackAlloc(int(unsafe.Sizeof(*destination)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DownloadFromGPUTexture",
			_copy_pass,
			_source,
			_destination,
		)
	}

	iDownloadFromGPUBuffer = func(copy_pass *GPUCopyPass, source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		_source, ok := internal.GetJSPointer(source)
		if !ok {
			_source = internal.StackAlloc(int(unsafe.Sizeof(*source)))
		}
		_destination, ok := internal.GetJSPointer(destination)
		if !ok {
			_destination = internal.StackAlloc(int(unsafe.Sizeof(*destination)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DownloadFromGPUBuffer",
			_copy_pass,
			_source,
			_destination,
		)
	}

	iEndGPUCopyPass = func(copy_pass *GPUCopyPass) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_copy_pass, ok := internal.GetJSPointer(copy_pass)
		if !ok {
			_copy_pass = internal.StackAlloc(int(unsafe.Sizeof(*copy_pass)))
		}
		js.Global().Get("Module").Call(
			"_SDL_EndGPUCopyPass",
			_copy_pass,
		)
	}

	iGenerateMipmapsForGPUTexture = func(command_buffer *GPUCommandBuffer, texture *GPUTexture) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		js.Global().Get("Module").Call(
			"_SDL_GenerateMipmapsForGPUTexture",
			_command_buffer,
			_texture,
		)
	}

	iBlitGPUTexture = func(command_buffer *GPUCommandBuffer, info *GPUBlitInfo) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_info, ok := internal.GetJSPointer(info)
		if !ok {
			_info = internal.StackAlloc(int(unsafe.Sizeof(*info)))
		}
		js.Global().Get("Module").Call(
			"_SDL_BlitGPUTexture",
			_command_buffer,
			_info,
		)
	}

	iWindowSupportsGPUSwapchainComposition = func(device *GPUDevice, window *Window, swapchain_composition GPUSwapchainComposition) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_swapchain_composition := int32(swapchain_composition)
		ret := js.Global().Get("Module").Call(
			"_SDL_WindowSupportsGPUSwapchainComposition",
			_device,
			_window,
			_swapchain_composition,
		)

		return internal.GetBool(ret)
	}

	iWindowSupportsGPUPresentMode = func(device *GPUDevice, window *Window, present_mode GPUPresentMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_present_mode := int32(present_mode)
		ret := js.Global().Get("Module").Call(
			"_SDL_WindowSupportsGPUPresentMode",
			_device,
			_window,
			_present_mode,
		)

		return internal.GetBool(ret)
	}

	iClaimWindowForGPUDevice = func(device *GPUDevice, window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ClaimWindowForGPUDevice",
			_device,
			_window,
		)

		return internal.GetBool(ret)
	}

	iReleaseWindowFromGPUDevice = func(device *GPUDevice, window *Window) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseWindowFromGPUDevice",
			_device,
			_window,
		)
	}

	iSetGPUSwapchainParameters = func(device *GPUDevice, window *Window, swapchain_composition GPUSwapchainComposition, present_mode GPUPresentMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_swapchain_composition := int32(swapchain_composition)
		_present_mode := int32(present_mode)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetGPUSwapchainParameters",
			_device,
			_window,
			_swapchain_composition,
			_present_mode,
		)

		return internal.GetBool(ret)
	}

	iSetGPUAllowedFramesInFlight = func(device *GPUDevice, allowed_frames_in_flight uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_allowed_frames_in_flight := int32(allowed_frames_in_flight)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetGPUAllowedFramesInFlight",
			_device,
			_allowed_frames_in_flight,
		)

		return internal.GetBool(ret)
	}

	iGetGPUSwapchainTextureFormat = func(device *GPUDevice, window *Window) GPUTextureFormat {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGPUSwapchainTextureFormat",
			_device,
			_window,
		)

		return GPUTextureFormat(ret.Int())
	}

	iAcquireGPUSwapchainTexture = func(command_buffer *GPUCommandBuffer, window *Window, swapchain_texture **GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_swapchain_texture, ok := internal.GetJSPointer(swapchain_texture)
		if !ok {
			_swapchain_texture = internal.StackAlloc(4)
		}
		_swapchain_texture_width, ok := internal.GetJSPointer(swapchain_texture_width)
		if !ok {
			_swapchain_texture_width = internal.StackAlloc(int(unsafe.Sizeof(*swapchain_texture_width)))
		}
		_swapchain_texture_height, ok := internal.GetJSPointer(swapchain_texture_height)
		if !ok {
			_swapchain_texture_height = internal.StackAlloc(int(unsafe.Sizeof(*swapchain_texture_height)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_AcquireGPUSwapchainTexture",
			_command_buffer,
			_window,
			_swapchain_texture,
			_swapchain_texture_width,
			_swapchain_texture_height,
		)

		return internal.GetBool(ret)
	}

	iWaitForGPUSwapchain = func(device *GPUDevice, window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitForGPUSwapchain",
			_device,
			_window,
		)

		return internal.GetBool(ret)
	}

	iWaitAndAcquireGPUSwapchainTexture = func(command_buffer *GPUCommandBuffer, window *Window, swapchain_texture **GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_swapchain_texture, ok := internal.GetJSPointer(swapchain_texture)
		if !ok {
			_swapchain_texture = internal.StackAlloc(4)
		}
		_swapchain_texture_width, ok := internal.GetJSPointer(swapchain_texture_width)
		if !ok {
			_swapchain_texture_width = internal.StackAlloc(int(unsafe.Sizeof(*swapchain_texture_width)))
		}
		_swapchain_texture_height, ok := internal.GetJSPointer(swapchain_texture_height)
		if !ok {
			_swapchain_texture_height = internal.StackAlloc(int(unsafe.Sizeof(*swapchain_texture_height)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitAndAcquireGPUSwapchainTexture",
			_command_buffer,
			_window,
			_swapchain_texture,
			_swapchain_texture_width,
			_swapchain_texture_height,
		)

		return internal.GetBool(ret)
	}

	iSubmitGPUCommandBuffer = func(command_buffer *GPUCommandBuffer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SubmitGPUCommandBuffer",
			_command_buffer,
		)

		return internal.GetBool(ret)
	}

	iSubmitGPUCommandBufferAndAcquireFence = func(command_buffer *GPUCommandBuffer) *GPUFence {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SubmitGPUCommandBufferAndAcquireFence",
			_command_buffer,
		)

		_obj := &GPUFence{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCancelGPUCommandBuffer = func(command_buffer *GPUCommandBuffer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_command_buffer, ok := internal.GetJSPointer(command_buffer)
		if !ok {
			_command_buffer = internal.StackAlloc(int(unsafe.Sizeof(*command_buffer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CancelGPUCommandBuffer",
			_command_buffer,
		)

		return internal.GetBool(ret)
	}

	iWaitForGPUIdle = func(device *GPUDevice) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitForGPUIdle",
			_device,
		)

		return internal.GetBool(ret)
	}

	iWaitForGPUFences = func(device *GPUDevice, wait_all bool, fences **GPUFence, num_fences uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_wait_all := internal.NewBoolean(wait_all)
		_fences, ok := internal.GetJSPointer(fences)
		if !ok {
			_fences = internal.StackAlloc(4)
		}
		_num_fences := int32(num_fences)
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitForGPUFences",
			_device,
			_wait_all,
			_fences,
			_num_fences,
		)

		return internal.GetBool(ret)
	}

	iQueryGPUFence = func(device *GPUDevice, fence *GPUFence) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_fence, ok := internal.GetJSPointer(fence)
		if !ok {
			_fence = internal.StackAlloc(int(unsafe.Sizeof(*fence)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_QueryGPUFence",
			_device,
			_fence,
		)

		return internal.GetBool(ret)
	}

	iReleaseGPUFence = func(device *GPUDevice, fence *GPUFence) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_fence, ok := internal.GetJSPointer(fence)
		if !ok {
			_fence = internal.StackAlloc(int(unsafe.Sizeof(*fence)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ReleaseGPUFence",
			_device,
			_fence,
		)
	}

	iGPUTextureFormatTexelBlockSize = func(format GPUTextureFormat) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_format := int32(format)
		ret := js.Global().Get("Module").Call(
			"_SDL_GPUTextureFormatTexelBlockSize",
			_format,
		)

		return uint32(ret.Int())
	}

	iGPUTextureSupportsFormat = func(device *GPUDevice, format GPUTextureFormat, typ GPUTextureType, usage GPUTextureUsageFlags) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_format := int32(format)
		_typ := int32(typ)
		_usage := int32(usage)
		ret := js.Global().Get("Module").Call(
			"_SDL_GPUTextureSupportsFormat",
			_device,
			_format,
			_typ,
			_usage,
		)

		return internal.GetBool(ret)
	}

	iGPUTextureSupportsSampleCount = func(device *GPUDevice, format GPUTextureFormat, sample_count GPUSampleCount) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		_format := int32(format)
		_sample_count := int32(sample_count)
		ret := js.Global().Get("Module").Call(
			"_SDL_GPUTextureSupportsSampleCount",
			_device,
			_format,
			_sample_count,
		)

		return internal.GetBool(ret)
	}

	iCalculateGPUTextureFormatSize = func(format GPUTextureFormat, width uint32, height uint32, depth_or_layer_count uint32) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_format := int32(format)
		_width := int32(width)
		_height := int32(height)
		_depth_or_layer_count := int32(depth_or_layer_count)
		ret := js.Global().Get("Module").Call(
			"_SDL_CalculateGPUTextureFormatSize",
			_format,
			_width,
			_height,
			_depth_or_layer_count,
		)

		return uint32(ret.Int())
	}

	iGetHaptics = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHaptics",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetHapticNameForID = func(instance_id HapticID) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHapticNameForID",
			_instance_id,
		)

		return internal.UTF8JSToString(ret)
	}

	iOpenHaptic = func(instance_id HapticID) *Haptic {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenHaptic",
			_instance_id,
		)

		_obj := &Haptic{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetHapticFromID = func(instance_id HapticID) *Haptic {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_instance_id := int32(instance_id)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHapticFromID",
			_instance_id,
		)

		_obj := &Haptic{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetHapticID = func(haptic *Haptic) HapticID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHapticID",
			_haptic,
		)

		return HapticID(ret.Int())
	}

	iGetHapticName = func(haptic *Haptic) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHapticName",
			_haptic,
		)

		return internal.UTF8JSToString(ret)
	}

	iIsMouseHaptic = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_IsMouseHaptic",
		)

		return internal.GetBool(ret)
	}

	iOpenHapticFromMouse = func() *Haptic {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenHapticFromMouse",
		)

		_obj := &Haptic{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iIsJoystickHaptic = func(joystick *Joystick) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_IsJoystickHaptic",
			_joystick,
		)

		return internal.GetBool(ret)
	}

	iOpenHapticFromJoystick = func(joystick *Joystick) *Haptic {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_joystick, ok := internal.GetJSPointer(joystick)
		if !ok {
			_joystick = internal.StackAlloc(int(unsafe.Sizeof(*joystick)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenHapticFromJoystick",
			_joystick,
		)

		_obj := &Haptic{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iCloseHaptic = func(haptic *Haptic) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		js.Global().Get("Module").Call(
			"_SDL_CloseHaptic",
			_haptic,
		)
	}

	iGetMaxHapticEffects = func(haptic *Haptic) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMaxHapticEffects",
			_haptic,
		)

		return int32(ret.Int())
	}

	iGetMaxHapticEffectsPlaying = func(haptic *Haptic) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMaxHapticEffectsPlaying",
			_haptic,
		)

		return int32(ret.Int())
	}

	iGetHapticFeatures = func(haptic *Haptic) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHapticFeatures",
			_haptic,
		)

		return uint32(ret.Int())
	}

	iGetNumHapticAxes = func(haptic *Haptic) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumHapticAxes",
			_haptic,
		)

		return int32(ret.Int())
	}

	iHapticEffectSupported = func(haptic *Haptic, effect *HapticEffect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect, ok := internal.GetJSPointer(effect)
		if !ok {
			_effect = internal.StackAlloc(int(unsafe.Sizeof(*effect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_HapticEffectSupported",
			_haptic,
			_effect,
		)

		return internal.GetBool(ret)
	}

	iCreateHapticEffect = func(haptic *Haptic, effect *HapticEffect) HapticEffectID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect, ok := internal.GetJSPointer(effect)
		if !ok {
			_effect = internal.StackAlloc(int(unsafe.Sizeof(*effect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateHapticEffect",
			_haptic,
			_effect,
		)

		return HapticEffectID(ret.Int())
	}

	iUpdateHapticEffect = func(haptic *Haptic, effect HapticEffectID, data *HapticEffect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect := int32(effect)
		_data, ok := internal.GetJSPointer(data)
		if !ok {
			_data = internal.StackAlloc(int(unsafe.Sizeof(*data)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateHapticEffect",
			_haptic,
			_effect,
			_data,
		)

		return internal.GetBool(ret)
	}

	iRunHapticEffect = func(haptic *Haptic, effect HapticEffectID, iterations uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect := int32(effect)
		_iterations := int32(iterations)
		ret := js.Global().Get("Module").Call(
			"_SDL_RunHapticEffect",
			_haptic,
			_effect,
			_iterations,
		)

		return internal.GetBool(ret)
	}

	iStopHapticEffect = func(haptic *Haptic, effect HapticEffectID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect := int32(effect)
		ret := js.Global().Get("Module").Call(
			"_SDL_StopHapticEffect",
			_haptic,
			_effect,
		)

		return internal.GetBool(ret)
	}

	iDestroyHapticEffect = func(haptic *Haptic, effect HapticEffectID) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect := int32(effect)
		js.Global().Get("Module").Call(
			"_SDL_DestroyHapticEffect",
			_haptic,
			_effect,
		)
	}

	iGetHapticEffectStatus = func(haptic *Haptic, effect HapticEffectID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_effect := int32(effect)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHapticEffectStatus",
			_haptic,
			_effect,
		)

		return internal.GetBool(ret)
	}

	iSetHapticGain = func(haptic *Haptic, gain int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_gain := int32(gain)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetHapticGain",
			_haptic,
			_gain,
		)

		return internal.GetBool(ret)
	}

	iSetHapticAutocenter = func(haptic *Haptic, autocenter int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_autocenter := int32(autocenter)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetHapticAutocenter",
			_haptic,
			_autocenter,
		)

		return internal.GetBool(ret)
	}

	iPauseHaptic = func(haptic *Haptic) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_PauseHaptic",
			_haptic,
		)

		return internal.GetBool(ret)
	}

	iResumeHaptic = func(haptic *Haptic) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ResumeHaptic",
			_haptic,
		)

		return internal.GetBool(ret)
	}

	iStopHapticEffects = func(haptic *Haptic) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_StopHapticEffects",
			_haptic,
		)

		return internal.GetBool(ret)
	}

	iHapticRumbleSupported = func(haptic *Haptic) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_HapticRumbleSupported",
			_haptic,
		)

		return internal.GetBool(ret)
	}

	iInitHapticRumble = func(haptic *Haptic) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_InitHapticRumble",
			_haptic,
		)

		return internal.GetBool(ret)
	}

	iPlayHapticRumble = func(haptic *Haptic, strength float32, length uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		_strength := int32(strength)
		_length := int32(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_PlayHapticRumble",
			_haptic,
			_strength,
			_length,
		)

		return internal.GetBool(ret)
	}

	iStopHapticRumble = func(haptic *Haptic) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_haptic, ok := internal.GetJSPointer(haptic)
		if !ok {
			_haptic = internal.StackAlloc(int(unsafe.Sizeof(*haptic)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_StopHapticRumble",
			_haptic,
		)

		return internal.GetBool(ret)
	}

	iSetHintWithPriority = func(name string, value string, priority HintPriority) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		_value := internal.StringOnJSStack(value)
		_priority := int32(priority)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetHintWithPriority",
			_name,
			_value,
			_priority,
		)

		return internal.GetBool(ret)
	}

	iSetHint = func(name string, value string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		_value := internal.StringOnJSStack(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetHint",
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iResetHint = func(name string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_ResetHint",
			_name,
		)

		return internal.GetBool(ret)
	}

	iResetHints = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_ResetHints",
		)
	}

	iGetHint = func(name string) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHint",
			_name,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetHintBoolean = func(name string, default_value bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		_default_value := internal.NewBoolean(default_value)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetHintBoolean",
			_name,
			_default_value,
		)

		return internal.GetBool(ret)
	}

	/*iAddHintCallback = func(name string, callback HintCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnStackGoToJS(name)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddHintCallback",
			_name,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	/*iRemoveHintCallback = func(name string, callback HintCallback, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnStackGoToJS(name)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_RemoveHintCallback",
			_name,
			_callback,
			_userdata,
		)
	}*/

	iInit = func(flags InitFlags) bool {
		_flags := uint32(flags)
		ret := js.Global().Get("Module").Call(
			"_SDL_Init",
			_flags,
		)

		return internal.GetBool(ret)
	}

	iInitSubSystem = func(flags InitFlags) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_flags := int32(flags)
		ret := js.Global().Get("Module").Call(
			"_SDL_InitSubSystem",
			_flags,
		)

		return internal.GetBool(ret)
	}

	iQuitSubSystem = func(flags InitFlags) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_flags := int32(flags)
		js.Global().Get("Module").Call(
			"_SDL_QuitSubSystem",
			_flags,
		)
	}

	iWasInit = func(flags InitFlags) InitFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_flags := int32(flags)
		ret := js.Global().Get("Module").Call(
			"_SDL_WasInit",
			_flags,
		)

		return InitFlags(ret.Int())
	}

	iQuit = func() {
		js.Global().Get("Module").Call("_SDL_Quit")
	}

	iIsMainThread = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_IsMainThread",
		)

		return internal.GetBool(ret)
	}

	/*iRunOnMainThread = func(callback MainThreadCallback, userdata uintptr, wait_complete bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		_wait_complete := internal.NewBoolean(wait_complete)
		ret := js.Global().Get("Module").Call(
			"_SDL_RunOnMainThread",
			_callback,
			_userdata,
			_wait_complete,
		)

		return internal.GetBool(ret)
	}*/

	iSetAppMetadata = func(appname string, appversion string, appidentifier string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_appname := internal.StringOnJSStack(appname)
		_appversion := internal.StringOnJSStack(appversion)
		_appidentifier := internal.StringOnJSStack(appidentifier)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAppMetadata",
			_appname,
			_appversion,
			_appidentifier,
		)

		return internal.GetBool(ret)
	}

	iSetAppMetadataProperty = func(name string, value string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		_value := internal.StringOnJSStack(value)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetAppMetadataProperty",
			_name,
			_value,
		)

		return internal.GetBool(ret)
	}

	iGetAppMetadataProperty = func(name string) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetAppMetadataProperty",
			_name,
		)

		return internal.UTF8JSToString(ret)
	}

	iLoadObject = func(sofile string) *SharedObject {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sofile := internal.StringOnJSStack(sofile)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadObject",
			_sofile,
		)

		_obj := &SharedObject{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	/*iLoadFunction = func(handle *SharedObject, name string) FunctionPointer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_handle, ok := internal.GetJSPointer(handle)
		if !ok {
			_handle = internal.StackAlloc(int(unsafe.Sizeof(*handle)))
		}
		_name := internal.StringOnStackGoToJS(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadFunction",
			_handle,
			_name,
		)

		return FunctionPointer(ret.Int())
	}*/

	iUnloadObject = func(handle *SharedObject) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_handle, ok := internal.GetJSPointer(handle)
		if !ok {
			_handle = internal.StackAlloc(int(unsafe.Sizeof(*handle)))
		}
		js.Global().Get("Module").Call(
			"_SDL_UnloadObject",
			_handle,
		)
	}

	iGetPreferredLocales = func(count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPreferredLocales",
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iSetLogPriorities = func(priority LogPriority) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_priority := int32(priority)
		js.Global().Get("Module").Call(
			"_SDL_SetLogPriorities",
			_priority,
		)
	}

	iSetLogPriority = func(category int32, priority LogPriority) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_priority := int32(priority)
		js.Global().Get("Module").Call(
			"_SDL_SetLogPriority",
			_category,
			_priority,
		)
	}

	iGetLogPriority = func(category int32) LogPriority {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetLogPriority",
			_category,
		)

		return LogPriority(ret.Int())
	}

	iResetLogPriorities = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_ResetLogPriorities",
		)
	}

	iSetLogPriorityPrefix = func(priority LogPriority, prefix string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_priority := int32(priority)
		_prefix := internal.StringOnJSStack(prefix)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetLogPriorityPrefix",
			_priority,
			_prefix,
		)

		return internal.GetBool(ret)
	}

	iLog = func(fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_Log",
			_fmt,
		)
	}

	iLogTrace = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogTrace",
			_category,
			_fmt,
		)
	}

	iLogVerbose = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogVerbose",
			_category,
			_fmt,
		)
	}

	iLogDebug = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogDebug",
			_category,
			_fmt,
		)
	}

	iLogInfo = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogInfo",
			_category,
			_fmt,
		)
	}

	iLogWarn = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogWarn",
			_category,
			_fmt,
		)
	}

	iLogError = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogError",
			_category,
			_fmt,
		)
	}

	iLogCritical = func(category int32, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogCritical",
			_category,
			_fmt,
		)
	}

	iLogMessage = func(category int32, priority LogPriority, fmt string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_priority := int32(priority)
		_fmt := internal.StringOnJSStack(fmt)
		js.Global().Get("Module").Call(
			"_SDL_LogMessage",
			_category,
			_priority,
			_fmt,
		)
	}

	iLogMessageV = func(category int32, priority LogPriority, fmt string, ap va_list) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_category := int32(category)
		_priority := int32(priority)
		_fmt := internal.StringOnJSStack(fmt)
		_ap := int32(ap)
		js.Global().Get("Module").Call(
			"_SDL_LogMessageV",
			_category,
			_priority,
			_fmt,
			_ap,
		)
	}

	/*iGetDefaultLogOutputFunction = func() LogOutputFunction {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDefaultLogOutputFunction",
		)

		return LogOutputFunction(ret.Int())
	}*/

	iGetLogOutputFunction = func(callback *LogOutputFunction, userdata *uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback, ok := internal.GetJSPointer(callback)
		if !ok {
			_callback = internal.StackAlloc(int(unsafe.Sizeof(*callback)))
		}
		_userdata, ok := internal.GetJSPointer(userdata)
		if !ok {
			_userdata = internal.StackAlloc(int(unsafe.Sizeof(*userdata)))
		}
		js.Global().Get("Module").Call(
			"_SDL_GetLogOutputFunction",
			_callback,
			_userdata,
		)
	}

	/*iSetLogOutputFunction = func(callback LogOutputFunction, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_SetLogOutputFunction",
			_callback,
			_userdata,
		)
	}*/

	iShowMessageBox = func(messageboxdata *messageBoxData, buttonid *int32) bool {
		panic("not implemented on js")
		/*internal.StackSave()
		defer internal.StackRestore()
		_messageboxdata := internal.StackAlloc(int(unsafe.Sizeof(*messageboxdata)))
		_buttonid := internal.StackAlloc(int(unsafe.Sizeof(*buttonid)))
		ret := js.Global().Get("Module").Call(
			"_SDL_ShowMessageBox",
			_messageboxdata,
			_buttonid,
		)

		return internal.GetBool(ret)*/
	}

	iShowSimpleMessageBox = func(flags MessageBoxFlags, title string, message string, window *Window) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_flags := int32(flags)
		_title := internal.StringOnJSStack(title)
		_message := internal.StringOnJSStack(message)
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ShowSimpleMessageBox",
			_flags,
			_title,
			_message,
			_window,
		)

		return internal.GetBool(ret)
	}

	iMetal_CreateView = func(window *Window) MetalView {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_Metal_CreateView",
			_window,
		)

		return MetalView(ret.Int())
	}

	iMetal_DestroyView = func(view MetalView) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_view := int32(view)
		js.Global().Get("Module").Call(
			"_SDL_Metal_DestroyView",
			_view,
		)
	}

	iMetal_GetLayer = func(view MetalView) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_view := int32(view)
		ret := js.Global().Get("Module").Call(
			"_SDL_Metal_GetLayer",
			_view,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iOpenURL = func(url string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_url := internal.StringOnJSStack(url)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenURL",
			_url,
		)

		return internal.GetBool(ret)
	}

	/*iCreateProcess = func(args *string, pipe_stdio bool) *Process {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_args, ok := internal.GetJSPointer(args)
		if !ok {
			_args = internal.StackAlloc()
		}
		_pipe_stdio := internal.NewBoolean(pipe_stdio)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateProcess",
			_args,
			_pipe_stdio,
		)

		_obj := &Process{}
		internal.StoreJSPointer(_obj, ret)
		return _obj
	}*/

	iCreateProcessWithProperties = func(props PropertiesID) *Process {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateProcessWithProperties",
			_props,
		)

		_obj := &Process{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetProcessProperties = func(process *Process) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetProcessProperties",
			_process,
		)

		return PropertiesID(ret.Int())
	}

	iReadProcess = func(process *Process, datasize *uintptr, exitcode *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		_datasize, ok := internal.GetJSPointer(datasize)
		if !ok {
			_datasize = internal.StackAlloc(int(unsafe.Sizeof(*datasize)))
		}
		_exitcode, ok := internal.GetJSPointer(exitcode)
		if !ok {
			_exitcode = internal.StackAlloc(int(unsafe.Sizeof(*exitcode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadProcess",
			_process,
			_datasize,
			_exitcode,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetProcessInput = func(process *Process) *IOStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetProcessInput",
			_process,
		)

		_obj := &IOStream{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iGetProcessOutput = func(process *Process) *IOStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetProcessOutput",
			_process,
		)

		_obj := &IOStream{}
		//internal.StoreJSPointer(_obj, ret)
		_ = ret
		return _obj
	}

	iKillProcess = func(process *Process, force bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		_force := internal.NewBoolean(force)
		ret := js.Global().Get("Module").Call(
			"_SDL_KillProcess",
			_process,
			_force,
		)

		return internal.GetBool(ret)
	}

	iWaitProcess = func(process *Process, block bool, exitcode *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		_block := internal.NewBoolean(block)
		_exitcode, ok := internal.GetJSPointer(exitcode)
		if !ok {
			_exitcode = internal.StackAlloc(int(unsafe.Sizeof(*exitcode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_WaitProcess",
			_process,
			_block,
			_exitcode,
		)

		return internal.GetBool(ret)
	}

	iDestroyProcess = func(process *Process) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_process, ok := internal.GetJSPointer(process)
		if !ok {
			_process = internal.StackAlloc(int(unsafe.Sizeof(*process)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyProcess",
			_process,
		)
	}

	iGetNumRenderDrivers = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumRenderDrivers",
		)

		return int32(ret.Int())
	}

	iGetRenderDriver = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderDriver",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iCreateWindowAndRenderer = func(title string, width int32, height int32, window_flags WindowFlags, window **Window, renderer **Renderer) bool {
		internal.StackSave()
		defer internal.StackRestore()

		_title := internal.StringOnJSStack(title)
		_width := int32(width)
		_height := int32(height)
		_window_flags := internal.NewBigInt(uint64(window_flags))
		_window := internal.StackAlloc(4)
		_renderer := internal.StackAlloc(4)

		ret := js.Global().Get("Module").Call(
			"_SDL_CreateWindowAndRenderer",
			_title,
			_width,
			_height,
			_window_flags,
			_window,
			_renderer,
		)
		windowPtr := internal.GetValue(_window, "*")
		rendererPtr := internal.GetValue(_renderer, "*")
		*window = internal.NewObject[Window](windowPtr)
		*renderer = internal.NewObject[Renderer](rendererPtr)

		return internal.GetBool(ret)
	}

	iCreateRenderer = func(window *Window, name string) *Renderer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateRenderer",
			_window,
			_name,
		)

		_obj := &Renderer{}
		_ = ret
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iCreateRendererWithProperties = func(props PropertiesID) *Renderer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateRendererWithProperties",
			_props,
		)
		_ = ret

		_obj := &Renderer{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iCreateSoftwareRenderer = func(surface *Surface) *Renderer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSoftwareRenderer",
			_surface,
		)
		_ = ret

		_obj := &Renderer{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iGetRenderer = func(window *Window) *Renderer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_window, ok := internal.GetJSPointer(window)
		if !ok {
			_window = internal.StackAlloc(int(unsafe.Sizeof(*window)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderer",
			_window,
		)
		_ = ret

		_obj := &Renderer{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iGetRenderWindow = func(renderer *Renderer) *Window {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderWindow",
			_renderer,
		)
		_ = ret

		_obj := &Window{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iGetRendererName = func(renderer *Renderer) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRendererName",
			_renderer,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetRendererProperties = func(renderer *Renderer) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRendererProperties",
			_renderer,
		)

		return PropertiesID(ret.Int())
	}

	iGetRenderOutputSize = func(renderer *Renderer, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderOutputSize",
			_renderer,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iGetCurrentRenderOutputSize = func(renderer *Renderer, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentRenderOutputSize",
			_renderer,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iCreateTexture = func(renderer *Renderer, format PixelFormat, access TextureAccess, w int32, h int32) *Texture {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_format := int32(format)
		_access := int32(access)
		_w := int32(w)
		_h := int32(h)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateTexture",
			_renderer,
			_format,
			_access,
			_w,
			_h,
		)

		_obj := internal.NewObject[Texture](ret)

		return _obj
	}

	iCreateTextureFromSurface = func(renderer *Renderer, surface *Surface) *Texture {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			panic("nil surface")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateTextureFromSurface",
			_renderer,
			_surface,
		)

		_obj := internal.NewObject[Texture](ret)

		return _obj
	}

	iCreateTextureWithProperties = func(renderer *Renderer, props PropertiesID) *Texture {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateTextureWithProperties",
			_renderer,
			_props,
		)
		_ = ret

		_obj := &Texture{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iGetTextureProperties = func(texture *Texture) PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureProperties",
			_texture,
		)

		return PropertiesID(ret.Int())
	}

	iGetRendererFromTexture = func(texture *Texture) *Renderer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRendererFromTexture",
			_texture,
		)
		_ = ret

		_obj := &Renderer{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iGetTextureSize = func(texture *Texture, w *float32, h *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureSize",
			_texture,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iSetTextureColorMod = func(texture *Texture, r uint8, g uint8, b uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextureColorMod",
			_texture,
			_r,
			_g,
			_b,
		)

		return internal.GetBool(ret)
	}

	iSetTextureColorModFloat = func(texture *Texture, r float32, g float32, b float32) bool {
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextureColorModFloat",
			_texture,
			r,
			g,
			b,
		)

		return internal.GetBool(ret)
	}

	iGetTextureColorMod = func(texture *Texture, r *uint8, g *uint8, b *uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureColorMod",
			_texture,
			_r,
			_g,
			_b,
		)

		return internal.GetBool(ret)
	}

	iGetTextureColorModFloat = func(texture *Texture, r *float32, g *float32, b *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureColorModFloat",
			_texture,
			_r,
			_g,
			_b,
		)

		return internal.GetBool(ret)
	}

	iSetTextureAlphaMod = func(texture *Texture, alpha uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_alpha := int32(alpha)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextureAlphaMod",
			_texture,
			_alpha,
		)

		return internal.GetBool(ret)
	}

	iSetTextureAlphaModFloat = func(texture *Texture, alpha float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_alpha := int32(alpha)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextureAlphaModFloat",
			_texture,
			_alpha,
		)

		return internal.GetBool(ret)
	}

	iGetTextureAlphaMod = func(texture *Texture, alpha *uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_alpha, ok := internal.GetJSPointer(alpha)
		if !ok {
			_alpha = internal.StackAlloc(int(unsafe.Sizeof(*alpha)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureAlphaMod",
			_texture,
			_alpha,
		)

		return internal.GetBool(ret)
	}

	iGetTextureAlphaModFloat = func(texture *Texture, alpha *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_alpha, ok := internal.GetJSPointer(alpha)
		if !ok {
			_alpha = internal.StackAlloc(int(unsafe.Sizeof(*alpha)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureAlphaModFloat",
			_texture,
			_alpha,
		)

		return internal.GetBool(ret)
	}

	iSetTextureBlendMode = func(texture *Texture, blendMode BlendMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_blendMode := int32(blendMode)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextureBlendMode",
			_texture,
			_blendMode,
		)

		return internal.GetBool(ret)
	}

	iGetTextureBlendMode = func(texture *Texture, blendMode *BlendMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_blendMode, ok := internal.GetJSPointer(blendMode)
		if !ok {
			_blendMode = internal.StackAlloc(int(unsafe.Sizeof(*blendMode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureBlendMode",
			_texture,
			_blendMode,
		)

		return internal.GetBool(ret)
	}

	iSetTextureScaleMode = func(texture *Texture, scaleMode ScaleMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_scaleMode := int32(scaleMode)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextureScaleMode",
			_texture,
			_scaleMode,
		)

		return internal.GetBool(ret)
	}

	iGetTextureScaleMode = func(texture *Texture, scaleMode *ScaleMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_scaleMode, ok := internal.GetJSPointer(scaleMode)
		if !ok {
			_scaleMode = internal.StackAlloc(int(unsafe.Sizeof(*scaleMode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextureScaleMode",
			_texture,
			_scaleMode,
		)

		return internal.GetBool(ret)
	}

	iUpdateTexture = func(texture *Texture, rect *Rect, pixels uintptr, pitch int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		_pixels, free := internal.ClonePtrArrayToJSHeap(*(**byte)(unsafe.Pointer(&pixels)), int(texture.H*pitch))
		defer free()
		_pitch := int32(pitch)
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateTexture",
			_texture,
			_rect,
			_pixels,
			_pitch,
		)

		internal.CopyJSToObject(texture, _texture)

		return internal.GetBool(ret)
	}

	iUpdateYUVTexture = func(texture *Texture, rect *Rect, Yplane *uint8, Ypitch int32, Uplane *uint8, Upitch int32, Vplane *uint8, Vpitch int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		_Yplane, ok := internal.GetJSPointer(Yplane)
		if !ok {
			_Yplane = internal.StackAlloc(int(unsafe.Sizeof(*Yplane)))
		}
		_Ypitch := int32(Ypitch)
		_Uplane, ok := internal.GetJSPointer(Uplane)
		if !ok {
			_Uplane = internal.StackAlloc(int(unsafe.Sizeof(*Uplane)))
		}
		_Upitch := int32(Upitch)
		_Vplane, ok := internal.GetJSPointer(Vplane)
		if !ok {
			_Vplane = internal.StackAlloc(int(unsafe.Sizeof(*Vplane)))
		}
		_Vpitch := int32(Vpitch)
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateYUVTexture",
			_texture,
			_rect,
			_Yplane,
			_Ypitch,
			_Uplane,
			_Upitch,
			_Vplane,
			_Vpitch,
		)

		return internal.GetBool(ret)
	}

	iUpdateNVTexture = func(texture *Texture, rect *Rect, Yplane *uint8, Ypitch int32, UVplane *uint8, UVpitch int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		_Yplane, ok := internal.GetJSPointer(Yplane)
		if !ok {
			_Yplane = internal.StackAlloc(int(unsafe.Sizeof(*Yplane)))
		}
		_Ypitch := int32(Ypitch)
		_UVplane, ok := internal.GetJSPointer(UVplane)
		if !ok {
			_UVplane = internal.StackAlloc(int(unsafe.Sizeof(*UVplane)))
		}
		_UVpitch := int32(UVpitch)
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateNVTexture",
			_texture,
			_rect,
			_Yplane,
			_Ypitch,
			_UVplane,
			_UVpitch,
		)

		return internal.GetBool(ret)
	}

	iLockTexture = func(texture *Texture, rect *Rect, pixels *uintptr, pitch *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		_pixels, ok := internal.GetJSPointer(pixels)
		if !ok {
			_pixels = internal.StackAlloc(int(unsafe.Sizeof(*pixels)))
		}
		_pitch, ok := internal.GetJSPointer(pitch)
		if !ok {
			_pitch = internal.StackAlloc(int(unsafe.Sizeof(*pitch)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_LockTexture",
			_texture,
			_rect,
			_pixels,
			_pitch,
		)

		return internal.GetBool(ret)
	}

	iLockTextureToSurface = func(texture *Texture, rect *Rect, surface **Surface) bool {
		internal.StackSave()
		defer internal.StackRestore()

		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		_surface := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_SDL_LockTextureToSurface",
			_texture,
			_rect,
			_surface,
		)
		surfacePtr := internal.GetValue(_surface, "*")
		*surface = internal.NewObject[Surface](surfacePtr)

		return internal.GetBool(ret)
	}

	iUnlockTexture = func(texture *Texture) {
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		js.Global().Get("Module").Call(
			"_SDL_UnlockTexture",
			_texture,
		)
	}

	iSetRenderTarget = func(renderer *Renderer, texture *Texture) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderTarget",
			_renderer,
			_texture,
		)

		return internal.GetBool(ret)
	}

	iGetRenderTarget = func(renderer *Renderer) *Texture {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderTarget",
			_renderer,
		)
		_ = ret

		_obj := &Texture{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iSetRenderLogicalPresentation = func(renderer *Renderer, w int32, h int32, mode RendererLogicalPresentation) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_w := int32(w)
		_h := int32(h)
		_mode := int32(mode)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderLogicalPresentation",
			_renderer,
			_w,
			_h,
			_mode,
		)

		return internal.GetBool(ret)
	}

	iGetRenderLogicalPresentation = func(renderer *Renderer, w *int32, h *int32, mode *RendererLogicalPresentation) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		_mode, ok := internal.GetJSPointer(mode)
		if !ok {
			_mode = internal.StackAlloc(int(unsafe.Sizeof(*mode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderLogicalPresentation",
			_renderer,
			_w,
			_h,
			_mode,
		)

		return internal.GetBool(ret)
	}

	iGetRenderLogicalPresentationRect = func(renderer *Renderer, rect *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderLogicalPresentationRect",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iRenderCoordinatesFromWindow = func(renderer *Renderer, window_x float32, window_y float32, x *float32, y *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_window_x := int32(window_x)
		_window_y := int32(window_y)
		_x, ok := internal.GetJSPointer(x)
		if !ok {
			_x = internal.StackAlloc(int(unsafe.Sizeof(*x)))
		}
		_y, ok := internal.GetJSPointer(y)
		if !ok {
			_y = internal.StackAlloc(int(unsafe.Sizeof(*y)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderCoordinatesFromWindow",
			_renderer,
			_window_x,
			_window_y,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iRenderCoordinatesToWindow = func(renderer *Renderer, x float32, y float32, window_x *float32, window_y *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_x := int32(x)
		_y := int32(y)
		_window_x, ok := internal.GetJSPointer(window_x)
		if !ok {
			_window_x = internal.StackAlloc(int(unsafe.Sizeof(*window_x)))
		}
		_window_y, ok := internal.GetJSPointer(window_y)
		if !ok {
			_window_y = internal.StackAlloc(int(unsafe.Sizeof(*window_y)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderCoordinatesToWindow",
			_renderer,
			_x,
			_y,
			_window_x,
			_window_y,
		)

		return internal.GetBool(ret)
	}

	iConvertEventToRenderCoordinates = func(renderer *Renderer, event *Event) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_event, ok := internal.GetJSPointer(event)
		if !ok {
			_event = internal.StackAlloc(int(unsafe.Sizeof(*event)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_ConvertEventToRenderCoordinates",
			_renderer,
			_event,
		)

		return internal.GetBool(ret)
	}

	iSetRenderViewport = func(renderer *Renderer, rect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderViewport",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iGetRenderViewport = func(renderer *Renderer, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderViewport",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iRenderViewportSet = func(renderer *Renderer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderViewportSet",
			_renderer,
		)

		return internal.GetBool(ret)
	}

	iGetRenderSafeArea = func(renderer *Renderer, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderSafeArea",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iSetRenderClipRect = func(renderer *Renderer, rect *Rect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderClipRect",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iGetRenderClipRect = func(renderer *Renderer, rect *Rect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_rect, ok := internal.GetJSPointer(rect)
		if !ok {
			_rect = internal.StackAlloc(int(unsafe.Sizeof(*rect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderClipRect",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iRenderClipEnabled = func(renderer *Renderer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderClipEnabled",
			_renderer,
		)

		return internal.GetBool(ret)
	}

	iSetRenderScale = func(renderer *Renderer, scaleX float32, scaleY float32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_scaleX := int32(scaleX)
		_scaleY := int32(scaleY)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderScale",
			_renderer,
			_scaleX,
			_scaleY,
		)

		return internal.GetBool(ret)
	}

	iGetRenderScale = func(renderer *Renderer, scaleX *float32, scaleY *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_scaleX, ok := internal.GetJSPointer(scaleX)
		if !ok {
			_scaleX = internal.StackAlloc(int(unsafe.Sizeof(*scaleX)))
		}
		_scaleY, ok := internal.GetJSPointer(scaleY)
		if !ok {
			_scaleY = internal.StackAlloc(int(unsafe.Sizeof(*scaleY)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderScale",
			_renderer,
			_scaleX,
			_scaleY,
		)

		return internal.GetBool(ret)
	}

	iSetRenderDrawColor = func(renderer *Renderer, r uint8, g uint8, b uint8, a uint8) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}

		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		_a := int32(a)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderDrawColor",
			_renderer,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iSetRenderDrawColorFloat = func(renderer *Renderer, r float32, g float32, b float32, a float32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderDrawColorFloat",
			_renderer,
			r,
			g,
			b,
			a,
		)

		return internal.GetBool(ret)
	}

	iGetRenderDrawColor = func(renderer *Renderer, r *uint8, g *uint8, b *uint8, a *uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderDrawColor",
			_renderer,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iGetRenderDrawColorFloat = func(renderer *Renderer, r *float32, g *float32, b *float32, a *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderDrawColorFloat",
			_renderer,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iSetRenderColorScale = func(renderer *Renderer, scale float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_scale := int32(scale)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderColorScale",
			_renderer,
			_scale,
		)

		return internal.GetBool(ret)
	}

	iGetRenderColorScale = func(renderer *Renderer, scale *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_scale, ok := internal.GetJSPointer(scale)
		if !ok {
			_scale = internal.StackAlloc(int(unsafe.Sizeof(*scale)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderColorScale",
			_renderer,
			_scale,
		)

		return internal.GetBool(ret)
	}

	iSetRenderDrawBlendMode = func(renderer *Renderer, blendMode BlendMode) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_blendMode := int32(blendMode)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderDrawBlendMode",
			_renderer,
			_blendMode,
		)

		return internal.GetBool(ret)
	}

	iGetRenderDrawBlendMode = func(renderer *Renderer, blendMode *BlendMode) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_blendMode, ok := internal.GetJSPointer(blendMode)
		if !ok {
			_blendMode = internal.StackAlloc(int(unsafe.Sizeof(*blendMode)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderDrawBlendMode",
			_renderer,
			_blendMode,
		)

		return internal.GetBool(ret)
	}

	iRenderClear = func(renderer *Renderer) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderClear",
			_renderer,
		)

		return internal.GetBool(ret)
	}

	iRenderPoint = func(renderer *Renderer, x float32, y float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_x := int32(x)
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderPoint",
			_renderer,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iRenderPoints = func(renderer *Renderer, points *FPoint, count int32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_points, free := internal.ClonePtrArrayToJSHeap(points, int(count))
		defer free()
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderPoints",
			_renderer,
			_points,
			_count,
		)

		return internal.GetBool(ret)
	}

	iRenderLine = func(renderer *Renderer, x1 float32, y1 float32, x2 float32, y2 float32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderLine",
			_renderer,
			x1,
			y1,
			x2,
			y2,
		)

		return internal.GetBool(ret)
	}

	iRenderLines = func(renderer *Renderer, points *FPoint, count int32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_points, free := internal.ClonePtrArrayToJSHeap(points, int(count))
		defer free()
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderLines",
			_renderer,
			_points,
			_count,
		)

		return internal.GetBool(ret)
	}

	iRenderRect = func(renderer *Renderer, rect *FRect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderRect",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iRenderRects = func(renderer *Renderer, rects *FRect, count int32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rects, free := internal.ClonePtrArrayToJSHeap(rects, int(count))
		defer free()
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderRects",
			_renderer,
			_rects,
			_count,
		)

		return internal.GetBool(ret)
	}

	iRenderFillRect = func(renderer *Renderer, rect *FRect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderFillRect",
			_renderer,
			_rect,
		)

		return internal.GetBool(ret)
	}

	iRenderFillRects = func(renderer *Renderer, rects *FRect, count int32) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rects, free := internal.ClonePtrArrayToJSHeap(rects, int(count))
		defer free()
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderFillRects",
			_renderer,
			_rects,
			_count,
		)

		return internal.GetBool(ret)
	}

	iRenderTexture = func(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderTexture",
			_renderer,
			_texture,
			_srcrect,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iRenderTextureRotated = func(renderer *Renderer, texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_dstrect := internal.CloneObjectToJSStack(dstrect)
		_angle := angle
		_center := internal.CloneObjectToJSStack(center)
		_flip := int32(flip)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderTextureRotated",
			_renderer,
			_texture,
			_srcrect,
			_dstrect,
			_angle,
			_center,
			_flip,
		)

		return internal.GetBool(ret)
	}

	iRenderTextureAffine = func(renderer *Renderer, texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		_srcrect := internal.CloneObjectToJSStack(srcrect)
		_origin := internal.CloneObjectToJSStack(origin)
		_right := internal.CloneObjectToJSStack(right)
		_down := internal.CloneObjectToJSStack(down)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderTextureAffine",
			_renderer,
			_texture,
			_srcrect,
			_origin,
			_right,
			_down,
		)

		return internal.GetBool(ret)
	}

	iRenderTextureTiled = func(renderer *Renderer, texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_srcrect, ok := internal.GetJSPointer(srcrect)
		if !ok {
			_srcrect = internal.StackAlloc(int(unsafe.Sizeof(*srcrect)))
		}
		_scale := int32(scale)
		_dstrect, ok := internal.GetJSPointer(dstrect)
		if !ok {
			_dstrect = internal.StackAlloc(int(unsafe.Sizeof(*dstrect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderTextureTiled",
			_renderer,
			_texture,
			_srcrect,
			_scale,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iRenderTexture9Grid = func(renderer *Renderer, texture *Texture, srcrect *FRect, left_width float32, right_width float32, top_height float32, bottom_height float32, scale float32, dstrect *FRect) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_srcrect, ok := internal.GetJSPointer(srcrect)
		if !ok {
			_srcrect = internal.StackAlloc(int(unsafe.Sizeof(*srcrect)))
		}
		_left_width := int32(left_width)
		_right_width := int32(right_width)
		_top_height := int32(top_height)
		_bottom_height := int32(bottom_height)
		_scale := int32(scale)
		_dstrect, ok := internal.GetJSPointer(dstrect)
		if !ok {
			_dstrect = internal.StackAlloc(int(unsafe.Sizeof(*dstrect)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderTexture9Grid",
			_renderer,
			_texture,
			_srcrect,
			_left_width,
			_right_width,
			_top_height,
			_bottom_height,
			_scale,
			_dstrect,
		)

		return internal.GetBool(ret)
	}

	iRenderGeometry = func(renderer *Renderer, texture *Texture, vertices *Vertex, num_vertices int32, indices *int32, num_indices int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_texture, _ := internal.GetJSPointer(texture)
		_vertices, freeVertices := internal.ClonePtrArrayToJSHeap(vertices, int(num_vertices))
		defer freeVertices()
		_num_vertices := int32(num_vertices)
		_indices, freeIndices := internal.ClonePtrArrayToJSHeap(indices, int(num_indices))
		defer freeIndices()
		_num_indices := int32(num_indices)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderGeometry",
			_renderer,
			_texture,
			_vertices,
			_num_vertices,
			_indices,
			_num_indices,
		)

		return internal.GetBool(ret)
	}

	iRenderGeometryRaw = func(renderer *Renderer, texture *Texture, xy *float32, xy_stride int32, color *FColor, color_stride int32, uv *float32, uv_stride int32, num_vertices int32, indices uintptr, num_indices int32, size_indices int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			_texture = internal.StackAlloc(int(unsafe.Sizeof(*texture)))
		}
		_xy, ok := internal.GetJSPointer(xy)
		if !ok {
			_xy = internal.StackAlloc(int(unsafe.Sizeof(*xy)))
		}
		_xy_stride := int32(xy_stride)
		_color, ok := internal.GetJSPointer(color)
		if !ok {
			_color = internal.StackAlloc(int(unsafe.Sizeof(*color)))
		}
		_color_stride := int32(color_stride)
		_uv, ok := internal.GetJSPointer(uv)
		if !ok {
			_uv = internal.StackAlloc(int(unsafe.Sizeof(*uv)))
		}
		_uv_stride := int32(uv_stride)
		_num_vertices := int32(num_vertices)
		_indices := internal.NewBigInt(indices)
		_num_indices := int32(num_indices)
		_size_indices := int32(size_indices)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderGeometryRaw",
			_renderer,
			_texture,
			_xy,
			_xy_stride,
			_color,
			_color_stride,
			_uv,
			_uv_stride,
			_num_vertices,
			_indices,
			_num_indices,
			_size_indices,
		)

		return internal.GetBool(ret)
	}

	iRenderReadPixels = func(renderer *Renderer, rect *Rect) *Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		_rect := internal.CloneObjectToJSStack(rect)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderReadPixels",
			_renderer,
			_rect,
		)

		_obj := internal.NewObject[Surface](ret)

		return _obj
	}

	iRenderPresent = func(renderer *Renderer) bool {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderPresent",
			_renderer,
		)

		return internal.GetBool(ret)
	}

	iDestroyTexture = func(texture *Texture) {
		_texture, ok := internal.GetJSPointer(texture)
		if !ok {
			panic("nil texture")
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyTexture",
			_texture,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(texture)))
	}

	iDestroyRenderer = func(renderer *Renderer) {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			return
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyRenderer",
			_renderer,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(renderer)))
	}

	iFlushRenderer = func(renderer *Renderer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_FlushRenderer",
			_renderer,
		)

		return internal.GetBool(ret)
	}

	iGetRenderMetalLayer = func(renderer *Renderer) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderMetalLayer",
			_renderer,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iGetRenderMetalCommandEncoder = func(renderer *Renderer) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderMetalCommandEncoder",
			_renderer,
		)

		return uintptr(internal.GetInt64(ret))
	}

	iAddVulkanRenderSemaphores = func(renderer *Renderer, wait_stage_mask uint32, wait_semaphore int64, signal_semaphore int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_wait_stage_mask := int32(wait_stage_mask)
		_wait_semaphore := internal.NewBigInt(wait_semaphore)
		_signal_semaphore := internal.NewBigInt(signal_semaphore)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddVulkanRenderSemaphores",
			_renderer,
			_wait_stage_mask,
			_wait_semaphore,
			_signal_semaphore,
		)

		return internal.GetBool(ret)
	}

	iSetRenderVSync = func(renderer *Renderer, vsync int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_vsync := int32(vsync)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetRenderVSync",
			_renderer,
			_vsync,
		)

		return internal.GetBool(ret)
	}

	iGetRenderVSync = func(renderer *Renderer, vsync *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_vsync, ok := internal.GetJSPointer(vsync)
		if !ok {
			_vsync = internal.StackAlloc(int(unsafe.Sizeof(*vsync)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetRenderVSync",
			_renderer,
			_vsync,
		)

		return internal.GetBool(ret)
	}

	iRenderDebugText = func(renderer *Renderer, x float32, y float32, str string) bool {
		internal.StackSave()
		defer internal.StackRestore()

		_renderer, _ := internal.GetJSPointer(renderer)
		// TODO: can we pass float instead of int32s?
		_x := int32(x)
		_y := int32(y)
		_str := internal.StringOnJSStack(str)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderDebugText",
			_renderer,
			_x,
			_y,
			_str,
		)

		return internal.GetBool(ret)
	}

	iRenderDebugTextFormat = func(renderer *Renderer, x float32, y float32, fmt string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_x := int32(x)
		_y := int32(y)
		_fmt := internal.StringOnJSStack(fmt)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderDebugTextFormat",
			_renderer,
			_x,
			_y,
			_fmt,
		)

		return internal.GetBool(ret)
	}

	iOpenTitleStorage = func(override string, props PropertiesID) *Storage {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_override := internal.StringOnJSStack(override)
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenTitleStorage",
			_override,
			_props,
		)
		_ = ret

		_obj := &Storage{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iOpenUserStorage = func(org string, app string, props PropertiesID) *Storage {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_org := internal.StringOnJSStack(org)
		_app := internal.StringOnJSStack(app)
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenUserStorage",
			_org,
			_app,
			_props,
		)
		_ = ret

		_obj := &Storage{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iOpenFileStorage = func(path string) *Storage {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenFileStorage",
			_path,
		)
		_ = ret

		_obj := &Storage{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iOpenStorage = func(iface *StorageInterface, userdata uintptr) *Storage {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_iface, ok := internal.GetJSPointer(iface)
		if !ok {
			_iface = internal.StackAlloc(int(unsafe.Sizeof(*iface)))
		}
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenStorage",
			_iface,
			_userdata,
		)
		_ = ret

		_obj := &Storage{}
		// internal.StoreJSPointer(_obj, ret)
		return _obj
	}

	iCloseStorage = func(storage *Storage) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CloseStorage",
			_storage,
		)

		return internal.GetBool(ret)
	}

	iStorageReady = func(storage *Storage) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_StorageReady",
			_storage,
		)

		return internal.GetBool(ret)
	}

	iGetStorageFileSize = func(storage *Storage, path string, length *uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		_length, ok := internal.GetJSPointer(length)
		if !ok {
			_length = internal.StackAlloc(int(unsafe.Sizeof(*length)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetStorageFileSize",
			_storage,
			_path,
			_length,
		)

		return internal.GetBool(ret)
	}

	iReadStorageFile = func(storage *Storage, path string, destination uintptr, length uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		_destination := internal.NewBigInt(destination)
		_length := internal.NewBigInt(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_ReadStorageFile",
			_storage,
			_path,
			_destination,
			_length,
		)

		return internal.GetBool(ret)
	}

	iWriteStorageFile = func(storage *Storage, path string, source uintptr, length uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		_source := internal.NewBigInt(source)
		_length := internal.NewBigInt(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_WriteStorageFile",
			_storage,
			_path,
			_source,
			_length,
		)

		return internal.GetBool(ret)
	}

	iCreateStorageDirectory = func(storage *Storage, path string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateStorageDirectory",
			_storage,
			_path,
		)

		return internal.GetBool(ret)
	}

	/*iEnumerateStorageDirectory = func(storage *Storage, path string, callback EnumerateDirectoryCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnStackGoToJS(path)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_EnumerateStorageDirectory",
			_storage,
			_path,
			_callback,
			_userdata,
		)

		return internal.GetBool(ret)
	}*/

	iRemoveStoragePath = func(storage *Storage, path string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_RemoveStoragePath",
			_storage,
			_path,
		)

		return internal.GetBool(ret)
	}

	iRenameStoragePath = func(storage *Storage, oldpath string, newpath string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_oldpath := internal.StringOnJSStack(oldpath)
		_newpath := internal.StringOnJSStack(newpath)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenameStoragePath",
			_storage,
			_oldpath,
			_newpath,
		)

		return internal.GetBool(ret)
	}

	iCopyStorageFile = func(storage *Storage, oldpath string, newpath string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_oldpath := internal.StringOnJSStack(oldpath)
		_newpath := internal.StringOnJSStack(newpath)
		ret := js.Global().Get("Module").Call(
			"_SDL_CopyStorageFile",
			_storage,
			_oldpath,
			_newpath,
		)

		return internal.GetBool(ret)
	}

	iGetStoragePathInfo = func(storage *Storage, path string, info *PathInfo) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		_info, ok := internal.GetJSPointer(info)
		if !ok {
			_info = internal.StackAlloc(int(unsafe.Sizeof(*info)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetStoragePathInfo",
			_storage,
			_path,
			_info,
		)

		return internal.GetBool(ret)
	}

	iGetStorageSpaceRemaining = func(storage *Storage) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetStorageSpaceRemaining",
			_storage,
		)

		return uint64(internal.GetInt64(ret))
	}

	iGlobStorageDirectory = func(storage *Storage, path string, pattern string, flags GlobFlags, count *int32) uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_storage, ok := internal.GetJSPointer(storage)
		if !ok {
			_storage = internal.StackAlloc(int(unsafe.Sizeof(*storage)))
		}
		_path := internal.StringOnJSStack(path)
		_pattern := internal.StringOnJSStack(pattern)
		_flags := int32(flags)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GlobStorageDirectory",
			_storage,
			_path,
			_pattern,
			_flags,
			_count,
		)

		return uintptr(internal.GetInt64(ret))
	}

	/*iSetX11EventHook = func(callback X11EventHook, userdata uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		js.Global().Get("Module").Call(
			"_SDL_SetX11EventHook",
			_callback,
			_userdata,
		)
	}*/

	iIsTablet = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_IsTablet",
		)

		return internal.GetBool(ret)
	}

	iIsTV = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_IsTV",
		)

		return internal.GetBool(ret)
	}

	iGetSandbox = func() Sandbox {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSandbox",
		)

		return Sandbox(ret.Int())
	}

	iOnApplicationWillTerminate = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_OnApplicationWillTerminate",
		)
	}

	iOnApplicationDidReceiveMemoryWarning = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_OnApplicationDidReceiveMemoryWarning",
		)
	}

	iOnApplicationWillEnterBackground = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_OnApplicationWillEnterBackground",
		)
	}

	iOnApplicationDidEnterBackground = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_OnApplicationDidEnterBackground",
		)
	}

	iOnApplicationWillEnterForeground = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_OnApplicationWillEnterForeground",
		)
	}

	iOnApplicationDidEnterForeground = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_OnApplicationDidEnterForeground",
		)
	}

	iGetDateTimeLocalePreferences = func(dateFormat *DateFormat, timeFormat *TimeFormat) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dateFormat, ok := internal.GetJSPointer(dateFormat)
		if !ok {
			_dateFormat = internal.StackAlloc(int(unsafe.Sizeof(*dateFormat)))
		}
		_timeFormat, ok := internal.GetJSPointer(timeFormat)
		if !ok {
			_timeFormat = internal.StackAlloc(int(unsafe.Sizeof(*timeFormat)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDateTimeLocalePreferences",
			_dateFormat,
			_timeFormat,
		)

		return internal.GetBool(ret)
	}

	iGetCurrentTime = func(ticks *Time) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ticks, ok := internal.GetJSPointer(ticks)
		if !ok {
			_ticks = internal.StackAlloc(int(unsafe.Sizeof(*ticks)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetCurrentTime",
			_ticks,
		)

		return internal.GetBool(ret)
	}

	iTimeToDateTime = func(ticks Time, dt *DateTime, localTime bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ticks := int32(ticks)
		_dt, ok := internal.GetJSPointer(dt)
		if !ok {
			_dt = internal.StackAlloc(int(unsafe.Sizeof(*dt)))
		}
		_localTime := internal.NewBoolean(localTime)
		ret := js.Global().Get("Module").Call(
			"_SDL_TimeToDateTime",
			_ticks,
			_dt,
			_localTime,
		)

		return internal.GetBool(ret)
	}

	iDateTimeToTime = func(dt *DateTime, ticks *Time) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dt, ok := internal.GetJSPointer(dt)
		if !ok {
			_dt = internal.StackAlloc(int(unsafe.Sizeof(*dt)))
		}
		_ticks, ok := internal.GetJSPointer(ticks)
		if !ok {
			_ticks = internal.StackAlloc(int(unsafe.Sizeof(*ticks)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_DateTimeToTime",
			_dt,
			_ticks,
		)

		return internal.GetBool(ret)
	}

	iTimeToWindows = func(ticks Time, dwLowDateTime *uint32, dwHighDateTime *uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ticks := int32(ticks)
		_dwLowDateTime, ok := internal.GetJSPointer(dwLowDateTime)
		if !ok {
			_dwLowDateTime = internal.StackAlloc(int(unsafe.Sizeof(*dwLowDateTime)))
		}
		_dwHighDateTime, ok := internal.GetJSPointer(dwHighDateTime)
		if !ok {
			_dwHighDateTime = internal.StackAlloc(int(unsafe.Sizeof(*dwHighDateTime)))
		}
		js.Global().Get("Module").Call(
			"_SDL_TimeToWindows",
			_ticks,
			_dwLowDateTime,
			_dwHighDateTime,
		)
	}

	iTimeFromWindows = func(dwLowDateTime uint32, dwHighDateTime uint32) Time {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_dwLowDateTime := int32(dwLowDateTime)
		_dwHighDateTime := int32(dwHighDateTime)
		ret := js.Global().Get("Module").Call(
			"_SDL_TimeFromWindows",
			_dwLowDateTime,
			_dwHighDateTime,
		)

		return Time(ret.Int())
	}

	iGetDaysInMonth = func(year int32, month int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_year := int32(year)
		_month := int32(month)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDaysInMonth",
			_year,
			_month,
		)

		return int32(ret.Int())
	}

	iGetDayOfYear = func(year int32, month int32, day int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_year := int32(year)
		_month := int32(month)
		_day := int32(day)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDayOfYear",
			_year,
			_month,
			_day,
		)

		return int32(ret.Int())
	}

	iGetDayOfWeek = func(year int32, month int32, day int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_year := int32(year)
		_month := int32(month)
		_day := int32(day)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetDayOfWeek",
			_year,
			_month,
			_day,
		)

		return int32(ret.Int())
	}

	iGetTicks = func() uint64 {
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTicks",
		)

		return uint64(internal.GetInt64(ret))
	}

	iGetTicksNS = func() uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTicksNS",
		)

		return uint64(internal.GetInt64(ret))
	}

	iGetPerformanceCounter = func() uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPerformanceCounter",
		)

		return uint64(internal.GetInt64(ret))
	}

	iGetPerformanceFrequency = func() uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPerformanceFrequency",
		)

		return uint64(internal.GetInt64(ret))
	}

	iDelay = func(ms uint32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ms := int32(ms)
		js.Global().Get("Module").Call(
			"_SDL_Delay",
			_ms,
		)
	}

	iDelayNS = func(ns uint64) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ns := internal.NewBigInt(ns)
		js.Global().Get("Module").Call(
			"_SDL_DelayNS",
			_ns,
		)
	}

	iDelayPrecise = func(ns uint64) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ns := internal.NewBigInt(ns)
		js.Global().Get("Module").Call(
			"_SDL_DelayPrecise",
			_ns,
		)
	}

	/*iAddTimer = func(interval uint32, callback TimerCallback, userdata uintptr) TimerID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_interval := int32(interval)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddTimer",
			_interval,
			_callback,
			_userdata,
		)

		return TimerID(ret.Int())
	}*/

	/*iAddTimerNS = func(interval uint64, callback NSTimerCallback, userdata uintptr) TimerID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_interval := internal.NewBigInt(interval)
		_callback := int32(callback)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_SDL_AddTimerNS",
			_interval,
			_callback,
			_userdata,
		)

		return TimerID(ret.Int())
	}*/

	iRemoveTimer = func(id TimerID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_id := int32(id)
		ret := js.Global().Get("Module").Call(
			"_SDL_RemoveTimer",
			_id,
		)

		return internal.GetBool(ret)
	}

	/*iRunApp = func(argc int32, argv *string, mainFunction main_func, reserved uintptr) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_argc := int32(argc)
		_argv, ok := internal.GetJSPointer(argv)
		if !ok {
			_argv = internal.StackAlloc()
		}
		_mainFunction := int32(mainFunction)
		_reserved := internal.NewBigInt(reserved)
		ret := js.Global().Get("Module").Call(
			"_SDL_RunApp",
			_argc,
			_argv,
			_mainFunction,
			_reserved,
		)

		return int32(ret.Int())
	}*/

	/*iEnterAppMainCallbacks = func(argc int32, argv *string, appinit AppInit_func, appiter AppIterate_func, appevent AppEvent_func, appquit AppQuit_func) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_argc := int32(argc)
		_argv, ok := internal.GetJSPointer(argv)
		if !ok {
			_argv = internal.StackAlloc()
		}
		_appinit := int32(appinit)
		_appiter := int32(appiter)
		_appevent := int32(appevent)
		_appquit := int32(appquit)
		ret := js.Global().Get("Module").Call(
			"_SDL_EnterAppMainCallbacks",
			_argc,
			_argv,
			_appinit,
			_appiter,
			_appevent,
			_appquit,
		)

		return int32(ret.Int())
	}*/

	iVulkan_LoadLibrary = func(path string) bool {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_LoadLibrary",
		)
		return false
	}

	iVulkan_GetVkGetInstanceProcAddr = func() FunctionPointer {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_GetVkGetInstanceProcAddr",
		)
		return FunctionPointer(0)
	}

	iVulkan_UnloadLibrary = func() {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_UnloadLibrary",
		)
	}

	iVulkan_GetInstanceExtensions = func(count *uint32) **byte {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_GetInstanceExtensions",
		)
		return nil
	}

	iVulkan_CreateSurface = func(window *Window, instance VkInstance, allocator *VkAllocationCallbacks, surface *VkSurfaceKHR) bool {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_CreateSurface",
		)
		return false
	}

	iVulkan_DestroySurface = func(instance VkInstance, surface VkSurfaceKHR, allocator *VkAllocationCallbacks) {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_DestroySurface",
		)
	}

	iVulkan_GetPresentationSupport = func(instance VkInstance, physicalDevice VkPhysicalDevice, queueFamilyIndex uint32) bool {
		panic("not implemented on js")
		js.Global().Get("Module").Call(
			"_SDL_Vulkan_GetPresentationSupport",
		)
		return false
	}
}
