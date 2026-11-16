//go:build js

package internal

import (
	"unsafe"
)

// ClonePtrString returns a newly allocated string from a uintptr
func ClonePtrString(ptr uintptr) string {
	return "" + PtrToString(ptr)
}

// PtrToString returns a string pointing to the provided ptr char data
func PtrToString(ptr uintptr) string {
	if addr, ok := lookupHeap(ptr); ok {
		return heapString(addr)
	}
	p := *(**uint8)(unsafe.Pointer(&ptr))
	if p == nil {
		return ""
	}
	if *p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := unsafe.Pointer(p); *(*byte)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}
	return string(unsafe.Slice(p, n))
}

// StringToNullablePtr returns a uintptr pointing to the provided string data
func StringToNullablePtr(s string) *byte {
	if len(s) == 0 {
		return nil
	}
	return StringToPtr(s)
}

// StringToPtr returns a NUL-terminated copy of s whose heap mirror lives
// as long as the copy.
func StringToPtr(s string) *byte {
	b := pinnable[byte](len(s) + 1)
	copy(b, s)
	pin(b)
	return &b[0]
}
