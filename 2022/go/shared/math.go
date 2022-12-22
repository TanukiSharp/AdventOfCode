package shared

type OrderedNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T OrderedNumber](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T OrderedNumber](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Sign[T OrderedNumber](value T) T {
	if value > 0 {
		return +1
	}
	if value < 0 {
		return -1
	}
	return 0
}

func Abs[T OrderedNumber](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
