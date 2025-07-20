package sdl

import (
	"errors"
	"reflect"
	"unsafe"
)

// NewProperties is a convenience function to create a new PropertiesID from the given map.
func NewProperties(properties map[string]any) (PropertiesID, error) {
	props, err := CreateProperties()
	if err != nil {
		return props, err
	}

	for name, anyValue := range properties {
		var err error
		switch value := anyValue.(type) {
		case uintptr:
			ok := props.SetPointerProperty(name, (*byte)(unsafe.Pointer(value)))
			if !ok {
				err = errors.New("failed to set pointer property")
			}
		case unsafe.Pointer:
			ok := props.SetPointerProperty(name, (*byte)(unsafe.Pointer(value)))
			if !ok {
				err = errors.New("failed to set pointer property")
			}
		case *byte:
			ok := props.SetPointerProperty(name, value)
			if !ok {
				err = errors.New("failed to set pointer property")
			}
		case string:
			err = props.SetStringProperty(name, value)
		case int:
			err = props.SetNumberProperty(name, int64(value))
		case int32:
			err = props.SetNumberProperty(name, int64(value))
		case int64:
			err = props.SetNumberProperty(name, value)
		case uint:
			err = props.SetNumberProperty(name, int64(value))
		case uint32:
			err = props.SetNumberProperty(name, int64(value))
		case uint64:
			err = props.SetNumberProperty(name, int64(value)) // precision loss
		case float32:
			err = props.SetFloatProperty(name, value)
		case float64:
			err = props.SetFloatProperty(name, float32(value)) // precision loss
		case bool:
			err = props.SetBooleanProperty(name, value)
		default:
			return props, errors.New("unknown type: " + reflect.TypeOf(anyValue).String())
		}
		if err != nil {
			return props, err
		}
	}

	return props, nil
}

// IOFromBytes creates a stream from a byte slice.
func IOFromBytes(b []byte) (*IOStream, error) {
	stream := IOFromDynamicMem()
	if _, err := stream.Write(b); err != nil {
		return nil, err
	}
	if _, err := stream.Seek(0, IO_SEEK_SET); err != nil {
		return nil, err
	}
	return stream, nil
}
