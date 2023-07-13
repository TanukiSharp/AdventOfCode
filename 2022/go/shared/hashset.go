package shared

type HashSet[T comparable] map[T]struct{}

func NewHashSet[T comparable]() HashSet[T] {
	return HashSet[T]{}
}

func (h HashSet[T]) Add(value T) bool {
	_, ok := h[value]

	if ok {
		return false
	}

	h[value] = struct{}{}

	return true
}

func (h HashSet[T]) Contains(value T) bool {
	_, ok := h[value]
	return ok
}

func (h HashSet[T]) Size() int {
	return len(h)
}

func (h HashSet[T]) Remove(value T) {
	delete(h, value)
}
