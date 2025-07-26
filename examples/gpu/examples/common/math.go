package common

func ClampInt[T int32 | uint32](x, a, b T) T {
	if x < a {
		return a
	} else if x > b {
		return b
	}
	return x
}
