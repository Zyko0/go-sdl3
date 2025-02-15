package internal

var (
	fnLastErr func() error
	fnFree    func(uintptr)
)

func SetSDLLastErrFunc(fn func() error) {
	fnLastErr = fn
}

func SetSDLFreeFunc(fn func(uintptr)) {
	fnFree = fn
}

// Impl

func LastErr() error {
	return fnLastErr()
}

func Free(ptr uintptr) {
	fnFree(ptr)
}
