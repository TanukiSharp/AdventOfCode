package day15

import "aoc/2022/shared"

type zone struct {
	start int
	end   int
}

func (z *zone) isEmpty() bool {
	return z.size() <= 0
}

func (z *zone) size() int {
	return z.end - z.start + 1
}

func (z *zone) isFullyContaining(other *zone) bool {
	return other.start >= z.start && other.end <= z.end
}

func (z *zone) isOverlapping(other *zone) bool {
	start := shared.Max(z.start, other.start)
	end := shared.Min(z.end, other.end)

	return start <= end
}

func (z *zone) canMerge(other *zone) bool {
	return other.isEmpty() || z.isFullyContaining(other) || z.isOverlapping(other)
}

func (z *zone) merge(other *zone) *zone {
	if other.isEmpty() || z.isFullyContaining(other) {
		return z
	}

	z.start = shared.Min(z.start, other.start)
	z.end = shared.Max(z.end, other.end)

	return z
}
