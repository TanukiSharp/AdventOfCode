package day15

import (
	"aoc/2022/shared"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (*Puzzle) Day() int     { return 15 }
func (*Puzzle) IsTest() bool { return false }

func (puzzle *Puzzle) Run(input string) {
	data := parseInput(input)

	part1(puzzle.IsTest(), data)
	part2(puzzle.IsTest(), data)
}

func part1(isTest bool, data *data) {
	if isTest {
		fmt.Printf("Part1: %d\n", findScannedDistanceAt(data, true, 10))
	} else {
		fmt.Printf("Part1: %d\n", findScannedDistanceAt(data, true, 2_000_000))
	}
}

func part2(isTest bool, data *data) {
	const xMul = 4_000_000

	min := 0
	max := 20
	if isTest == false {
		min = 0
		max = xMul
	}

	if isTest {
		fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvv")
		for y := min; y <= max; y++ {
			printZones(y, min, max-min, data, getMergedZonesAt(data, y))
		}
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^")
		fmt.Println()
		fmt.Println()
	}

	// Could be optimized by sorting sensors by Y coordinate and
	// avoid checkiong passed sensors out of range from current Y.

	for y := min; y <= max; y++ {
		zones := getMergedZonesAt(data, y)

		x := findHole(min, max, zones)

		if x != nil {
			result := (*x)*xMul + y
			fmt.Printf("Part2: %d\n", result)
			break
		}
	}
}

func findScannedDistanceAt(data *data, includeBeacons bool, y int) int {
	zones := getMergedZonesAt(data, y)

	total := 0

	for _, zone := range zones {
		total += zone.size()
	}

	if includeBeacons {
		for beacon := range data.beacons {
			if beacon.y == y {
				total--
			}
		}
	}

	return total
}

func getMergedZonesAt(data *data, y int) []*zone {
	zones := []*zone{}

	for _, sensor := range data.sensors {
		zone := sensor.getZoneAt(y)

		if zone.isEmpty() {
			continue
		}

		zones = append(zones, zone)
	}

	for {
	label:

		for i := 0; i < len(zones); i++ {
			for j := 0; j < len(zones); j++ {
				if i == j {
					continue
				}

				if zones[i].canMerge(zones[j]) {
					zones[i].merge(zones[j])
					zones = append(zones[:j], zones[j+1:]...)
					goto label
				}
			}
		}

		break
	}

	return zones
}

func findHole(minX, maxX int, zones []*zone) *int {
	sort.Slice(zones, func(i, j int) bool { return zones[i].start < zones[j].start })

	if minX < zones[0].start {
		return &minX
	}

	if maxX > zones[len(zones)-1].end {
		return &maxX
	}

	for i := 0; i < len(zones)-1; i++ {
		if zones[i+1].start-zones[i].end == 2 {
			res := zones[i].end + 1
			return &res
		}
	}

	return nil
}

func printZones(mapY, mapMinX, mapWidth int, data *data, zones []*zone) {
	line := make([]rune, mapWidth)

	for i := 0; i < mapWidth; i++ {
		line[i] = '.'

		mapX := i + mapMinX

		if hasBeaconAt(data, mapX, mapY) {
			line[i] = 'B'
			continue
		}

		if hasSensorAt(data, mapX, mapY) {
			line[i] = 'S'
			continue
		}

		for _, zone := range zones {
			if zone.isEmpty() {
				continue
			}

			if zone.start <= mapX && mapX <= zone.end {
				line[i] = '#'
				break
			}
		}
	}

	fmt.Printf("%-2d - %s\n", mapY, (string)(line))
}

func parseInput(input string) *data {
	sensors := []sensor{}
	beacons := shared.NewHashSet[beacon]()

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		regex, _ := regexp.Compile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
		matches := regex.FindSubmatch(([]byte)(line))

		sx, _ := strconv.Atoi((string)(matches[1]))
		sy, _ := strconv.Atoi((string)(matches[2]))
		bx, _ := strconv.Atoi((string)(matches[3]))
		by, _ := strconv.Atoi((string)(matches[4]))

		beacon := beacon{x: bx, y: by}
		beacons.Add(beacon)

		sensors = append(sensors, sensor{
			x:      sx,
			y:      sy,
			beacon: beacon,
			radius: shared.Abs(sx-bx) + shared.Abs(sy-by),
		})
	}

	return &data{
		sensors: sensors,
		beacons: beacons,
	}
}

func hasBeaconAt(data *data, x, y int) bool {
	for beacon := range data.beacons {
		if beacon.x == x && beacon.y == y {
			return true
		}
	}
	return false
}

func hasSensorAt(data *data, x, y int) bool {
	for _, sensor := range data.sensors {
		if sensor.x == x && sensor.y == y {
			return true
		}
	}
	return false
}

type data struct {
	sensors []sensor
	beacons shared.HashSet[beacon]
}

type beacon struct {
	x int
	y int
}

type sensor struct {
	x      int
	y      int
	beacon beacon
	radius int
}

func (sensor *sensor) getZoneAt(y int) *zone {
	distanceToSensor := shared.Abs(y - sensor.y)
	delta := sensor.radius - distanceToSensor
	return &zone{start: sensor.x - delta, end: sensor.x + delta}
}
