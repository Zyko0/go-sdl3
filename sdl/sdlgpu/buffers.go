package sdlgpu

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/sdl"
)

// GPUTransferBuffer

type TypedTransferBuffer[T any] struct {
	buffer *sdl.GPUTransferBuffer
	size   uint32
	length uint32
}

// Raw returns the underlying GPUTransferBuffer.
func (b *TypedTransferBuffer[T]) Raw() *sdl.GPUTransferBuffer {
	return b.buffer
}

// Length returns the number of elements in the buffer.
func (b *TypedTransferBuffer[T]) Length() uint32 {
	return b.length
}

// Size returns the total size of the buffer in bytes.
func (b *TypedTransferBuffer[T]) Size() uint32 {
	return b.size * b.length
}

func (b *TypedTransferBuffer[T]) Location(index uint32) *sdl.GPUTransferBufferLocation {
	return &sdl.GPUTransferBufferLocation{
		TransferBuffer: b.buffer,
		Offset:         index * b.size,
	}
}

func (b *TypedTransferBuffer[T]) Map(device *sdl.GPUDevice, cycle bool) ([]T, error) {
	ptr, err := device.MapTransferBuffer(b.buffer, cycle)
	if err != nil {
		return nil, err
	}

	return unsafe.Slice((*T)(unsafe.Pointer(ptr)), b.length), nil
}

func CreateTypedTransferBuffer[T any](device *sdl.GPUDevice, usage sdl.GPUTransferBufferUsage, length uint32, props sdl.PropertiesID) (*TypedTransferBuffer[T], error) {
	var t T

	size := uint32(unsafe.Sizeof(t))
	buffer, err := device.CreateTransferBuffer(&sdl.GPUTransferBufferCreateInfo{
		Usage: usage,
		Size:  size * length,
		Props: props,
	})
	if err != nil {
		return nil, err
	}

	return &TypedTransferBuffer[T]{
		buffer: buffer,
		size:   size,
		length: length,
	}, nil
}

// GPUBuffer

type TypedBuffer[T any] struct {
	buffer *sdl.GPUBuffer
	size   uint32
	length uint32
}

// Raw returns the underlying GPUBuffer.
func (b *TypedBuffer[T]) Raw() *sdl.GPUBuffer {
	return b.buffer
}

// Length returns the number of elements in the buffer.
func (b *TypedBuffer[T]) Length() uint32 {
	return b.length
}

// Size returns the total size of the buffer in bytes.
func (b *TypedBuffer[T]) Size() uint32 {
	return b.size * b.length
}

func (b *TypedBuffer[T]) Location(index uint32) *sdl.GPUBufferLocation {
	return &sdl.GPUBufferLocation{
		Buffer: b.buffer,
		Offset: index * b.size,
	}
}

func (b *TypedBuffer[T]) Binding(index uint32) *sdl.GPUBufferBinding {
	return &sdl.GPUBufferBinding{
		Buffer: b.buffer,
		Offset: index * b.size,
	}
}

func (b *TypedBuffer[T]) Region(start, end uint32) *sdl.GPUBufferRegion {
	return &sdl.GPUBufferRegion{
		Buffer: b.buffer,
		Offset: start * b.size,
		Size:   (end - start) * b.size,
	}
}

func CreateTypedBuffer[T any](device *sdl.GPUDevice, usage sdl.GPUBufferUsageFlags, length uint32, props sdl.PropertiesID) (*TypedBuffer[T], error) {
	var t T

	size := uint32(unsafe.Sizeof(t))
	buffer, err := device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: usage,
		Size:  size * length,
		Props: props,
	})
	if err != nil {
		return nil, err
	}

	return &TypedBuffer[T]{
		buffer: buffer,
		size:   size,
		length: length,
	}, nil
}
