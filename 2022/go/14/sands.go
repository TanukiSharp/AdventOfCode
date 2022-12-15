package day14

type sands struct {
	grains []*sand
}

var _false bool = false
var _true bool = false
var _pfalse *bool = &_false
var _ptrue *bool = &_true

func (sands *sands) isOnMoving(x, y int) bool {
	return sands.isOn(x, y, _pfalse)
}

func (sands *sands) isOnFixed(x, y int) bool {
	return sands.isOn(x, y, _ptrue)
}

func (sands *sands) isOnAny(x, y int) bool {
	return sands.isOn(x, y, nil)
}

func (sands *sands) isOn(x, y int, onlyMoving *bool) bool {
	for _, grain := range sands.grains {
		if grain.pos.x == x && grain.pos.y == y {
			if onlyMoving == nil || (*onlyMoving == grain.isMoving) {
				return true
			}
		}
	}

	return false
}

func (sands *sands) emit(rocks rockTraces) bool {
	x := 500
	y := 0

	if rocks.isOn(x, y) || sands.isOnFixed(x, y) {
		return false
	}

	sands.grains = append(sands.grains, &sand{
		pos:      coord{x: x, y: y},
		isMoving: true,
	})

	return true
}

func (sands *sands) update(rocks rockTraces) bool {
	result := false

	for _, grain := range sands.grains {
		if grain.isMoving == false {
			continue
		}

		if grain.update(sands, rocks) {
			result = true
		}
	}

	return result
}

type sand struct {
	pos      coord
	isMoving bool
}

func (grain *sand) canMoveTo(c coord, sands *sands, rocks rockTraces) bool {
	return sands.isOnFixed(c.x, c.y) == false && rocks.isOn(c.x, c.y) == false
}

func (grain *sand) update(sands *sands, rocks rockTraces) bool {
	coordsToMove := []coord{
		/* down  */ {x: grain.pos.x + 0, y: grain.pos.y + 1},
		/* left  */ {x: grain.pos.x - 1, y: grain.pos.y + 1},
		/* right */ {x: grain.pos.x + 1, y: grain.pos.y + 1},
	}

	for _, c := range coordsToMove {
		if grain.canMoveTo(c, sands, rocks) {
			grain.pos.x = c.x
			grain.pos.y = c.y
			return true
		}
	}

	grain.isMoving = false

	return false
}
