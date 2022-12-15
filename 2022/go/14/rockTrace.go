package day14

type rockTrace struct {
	coords []coord
}

func isOnLine(x, y int, c1, c2 coord) bool {
	minX := c1.x
	maxX := c2.x

	minY := c1.y
	maxY := c2.y

	if minX > maxX {
		minX, maxX = maxX, minX
	}
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	if minX == maxX {
		return x == minX && y >= minY && y <= maxY
	} else {
		return y == minY && x >= minX && x <= maxX
	}
}

func (rockTrace *rockTrace) isOn(x, y int) bool {
	for i := 0; i < len(rockTrace.coords)-1; i++ {
		if isOnLine(x, y, rockTrace.coords[i], rockTrace.coords[i+1]) {
			return true
		}
	}

	return false
}

func (rockTrace *rockTrace) addCoord(c coord) {
	if rockTrace.coords == nil {
		rockTrace.coords = []coord{}
	}

	rockTrace.coords = append(rockTrace.coords, c)
}

func (rockTrace *rockTrace) isBelow(y int) bool {
	for _, c := range rockTrace.coords {
		if y <= c.y {
			return false
		}
	}

	return true
}
