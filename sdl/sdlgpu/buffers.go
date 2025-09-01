package sdlgpu

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Test

type bufferItem struct {
	ptr    reflect.Value
	typ    reflect.Type
	offset uint32
	length uint32
}

type BaseLayout struct {
	items []bufferItem
	size  uint32
}

func (l *BaseLayout) TotalSize() uint32 {
	return l.size
}

func (l *BaseLayout) Offset(slot, index uint32) uint32 {
	addr := l.items[slot].offset
	addr += index * uint32(l.items[slot].typ.Size())
	return addr
}

func (l *BaseLayout) Size(slot uint32) uint32 {
	return max(l.items[slot].length, 1) * uint32(l.items[slot].typ.Size())
}

func (l *BaseLayout) Map(ptr uintptr) {
	for _, item := range l.items {
		addr := unsafe.Pointer(ptr + uintptr(item.offset))
		if item.length > 0 {
			item.ptr.Elem().Set(
				reflect.SliceAt(item.typ, addr, int(item.length)),
			)
		} else {
			item.ptr.Elem().Set(
				reflect.NewAt(item.typ, addr),
			)
		}
	}
}

func Register(layout any) {
	t := reflect.TypeOf(layout)
	if t.Kind() != reflect.Pointer {
		panic("layout must be a pointer to a struct")
	}

	t = t.Elem()
	if t.Kind() != reflect.Struct {
		panic("layout must be a pointer to a struct")
	}

	if t.NumField() > 0 && t.Field(0).Type != reflect.TypeOf(BaseLayout{}) {
		panic("first field must be sdlgpu.BufferLayout")
	}
	if t.NumField() < 2 {
		panic("no buffer type to register")
	}

	v := reflect.ValueOf(layout).Elem()
	info := (*BaseLayout)(v.Field(0).Addr().UnsafePointer())
	info.size = 0
	info.items = make([]bufferItem, t.NumField()-1)

	var slotsUsed bool
	var offset uint32
	for i := 1; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		tag := fieldType.Tag.Get("gpu")
		if tag == "" {
			continue
		}

		// Tag parsing
		fmt.Println("tag:", tag)
		var slot, length = -1, -1
		for _, decl := range strings.Split(tag, ",") {
			parts := strings.Split(decl, ":")
			switch parts[0] {
			case "slot":
				if len(parts) != 2 {
					panic("invalid struct tag: " + decl)
				}
				n, err := strconv.Atoi(parts[1])
				if err != nil {
					panic("invalid slot number: " + parts[1])
				}
				if n < 0 {
					panic("slot number must be >= 0: " + parts[1])
				}
				if n >= t.NumField()-1 {
					panic("slot number out of range: " + parts[1])
				}
				slot = n
			case "len":
				if len(parts) != 2 {
					panic("invalid struct tag: " + decl)
				}
				n, err := strconv.Atoi(parts[1])
				if err != nil {
					panic("invalid length: " + parts[1])
				}
				if n <= 0 {
					panic("length must be positive: " + parts[1])
				}
				length = n
			case "-":
			default:
				panic("unknown struct tag: " + parts[0])
			}
		}
		// Set slot mode if first item
		if i == 1 {
			slotsUsed = slot >= 0
		}
		// Check slot mode consistency
		if slotsUsed != (slot >= 0) {
			panic("inconsistent slot mode in struct tags")
		}
		// Assign slot by struct field order if not specified explicitely
		if !slotsUsed {
			slot = i - 1
		}
		// Check if field is of either pointer of slice type
		switch fieldValue.Kind() {
		case reflect.Pointer:
			length = 0
		case reflect.Slice:
			if length <= 0 {
				panic("length must be specified for slices")
			}
		default:
			panic("field must be a pointer or a slice: " + fieldType.Name)
		}

		info.items[slot] = bufferItem{
			ptr:    fieldValue.Addr(),
			typ:    fieldType.Type.Elem(),
			offset: offset,
			length: uint32(length),
		}
		// Increase offset by field size (or slice element size times length)
		size := uint32(fieldType.Type.Elem().Size()) * uint32(max(length, 1))
		info.size += size
		offset += size
	}
}
