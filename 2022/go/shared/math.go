package shared

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Sign(value int) int {
	if value > 0 {
		return +1
	}
	if value < 0 {
		return -1
	}
	return 0
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
