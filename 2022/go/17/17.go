package day17

import (
	"aoc/2022/shared"
	"fmt"
	"math"
	"strings"
)

type Puzzle struct{}

func (*Puzzle) Day() int     { return 17 }
func (*Puzzle) IsTest() bool { return true }

func (*Puzzle) Run(input string) {
	shapes := []*Shape{
		NewShape([]Coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}),         // Horizontal line.
		NewShape([]Coord{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}), // Plus.
		NewShape([]Coord{{2, 0}, {2, 1}, {0, 2}, {1, 2}, {2, 2}}), // Reversed L.
		NewShape([]Coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}}),         // Vertical line.
		NewShape([]Coord{{0, 0}, {0, 1}, {1, 0}, {1, 1}}),         // Square.
	}

	blows := parseInput(input)
	chamber := NewChamber(7, shapes)

	part1(chamber, blows)
	part2(chamber, blows)
}

func part1(chamber *Chamber, blows []int64) {
	blowIndex := 0

	for chamber.fixedBlocksCount < 2022 {
		chamber.RunTurn(blows[blowIndex], false)
		blowIndex = (blowIndex + 1) % len(blows)
	}

	fmt.Printf("Part1: %d\n", chamber.highestStart)
}

func part2(chamber *Chamber, blows []int64) {
	blowIndex := 0

	for chamber.fixedBlocksCount < 1_000_000_000_000 {
		chamber.RunTurn(blows[blowIndex], false)
		blowIndex = (blowIndex + 1) % len(blows)
	}

	fmt.Printf("Part2: %d\n", chamber.highestStart)
}

type Coord struct {
	x, y int64
}

type Shape struct {
	width  int64
	height int64
	coords []Coord
}

func NewShape(coords []Coord) *Shape {
	var width int64 = math.MinInt64
	var height int64 = math.MinInt64

	for _, c := range coords {
		width = shared.Max(width, c.x)
		height = shared.Max(height, c.y)
	}

	return &Shape{
		width:  width + 1,
		height: height + 1,
		coords: coords,
	}
}

func (s *Shape) CreateBlock() *Block {
	return &Block{
		isMoving: true,
		dx:       2,
		dy:       0,
		shape:    s,
	}
}

func (s *Shape) IsCollidingAt(x, y int64) bool {
	for _, c := range s.coords {
		if c.x == x && c.y == y {
			return true
		}
	}
	return false
}

func (s *Shape) PrintAt(dx int64) {
	for y := (int64)(0); y < s.height; y++ {
		for i := (int64)(0); i < dx; i++ {
			fmt.Print(" ")
		}

		for x := (int64)(0); x < s.width; x++ {
			if s.IsCollidingAt(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Block struct {
	isMoving bool
	dx       int64
	dy       int64
	shape    *Shape
}

func NewBlock(x, y int64, shape *Shape) *Block {
	return &Block{
		isMoving: true,
		dx:       x,
		dy:       y,
		shape:    shape,
	}
}

func (block *Block) IsCollidingAt(x, y int64) bool {
	at := Coord{x, y}
	for _, coord := range block.shape.coords {
		coord = Coord{coord.x + block.dx, block.dy - coord.y}
		if coord == at {
			return true
		}
	}
	return false
}

func (lhs *Block) IsAABBColliding(rhs *Block) bool {
	if lhs.dy > rhs.dy {
		if (lhs.dy - rhs.dy) >= lhs.shape.height {
			return false
		}
	} else if rhs.dy > lhs.dy {
		if (rhs.dy - lhs.dy) >= rhs.shape.height {
			return false
		}
	}

	if lhs.dx > rhs.dx {
		if (lhs.dx - rhs.dx) >= rhs.shape.width {
			return false
		}
	} else if rhs.dx > lhs.dx {
		if (rhs.dx - lhs.dx) >= lhs.shape.width {
			return false
		}
	}

	return true
}

func (lhs *Block) IsColliding(rhs *Block) bool {
	if lhs.IsAABBColliding(rhs) == false {
		return false
	}

	for _, coord1 := range lhs.shape.coords {
		coord1 = Coord{coord1.x + lhs.dx, lhs.dy - coord1.y}
		for _, coord2 := range rhs.shape.coords {
			coord2 = Coord{coord2.x + rhs.dx, rhs.dy - coord2.y}
			if coord1 == coord2 {
				return true
			}
		}
	}
	return false
}

func (b *Block) Print() {
	for y := (int64)(0); y < b.shape.height; y++ {
		for i := (int64)(0); i < b.dx; i++ {
			fmt.Print(" ")
		}

		for x := (int64)(0); x < b.shape.width; x++ {
			if b.shape.IsCollidingAt(x, y) {
				if b.isMoving {
					fmt.Print("@")
				} else {
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Chamber struct {
	width             int64
	highestStart      int64
	shapes            []*Shape
	currentShapeIndex int
	blocks            []*Block
	fixedBlocksCount  int64
}

func NewChamber(width int64, shapes []*Shape) *Chamber {
	return &Chamber{
		width:             width,
		highestStart:      0,
		shapes:            shapes,
		currentShapeIndex: 0,
		blocks:            []*Block{},
	}
}

func (c *Chamber) RunTurn(blow int64, print bool) bool {
	blockCount := len(c.blocks)

	var block *Block
	shape := c.shapes[c.currentShapeIndex]

	if blockCount == 0 || c.blocks[blockCount-1].isMoving == false {
		block = NewBlock(2, c.highestStart+shape.height+3, shape)
		c.blocks = append(c.blocks, block)
		c.UpdateStart()
		if print {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			c.Print()
			fmt.Printf("%d units tall.", c.highestStart)
			fmt.Print("") // For breakpoint.
		}
	} else {
		block = c.blocks[blockCount-1]
	}

	block.dx += blow

	if c.IsColliding(block) {
		block.dx -= blow
	}

	if print {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		c.Print()
		fmt.Printf("%d units tall.", c.highestStart)
		fmt.Print("") // For breakpoint.
	}

	block.dy--

	result := true

	if c.IsColliding(block) {
		block.dy++
		block.isMoving = false
		c.fixedBlocksCount++
		c.currentShapeIndex = (c.currentShapeIndex + 1) % len(c.shapes)
		result = false
	}

	c.UpdateStart()

	if print {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		c.Print()
		fmt.Printf("%d units tall.", c.highestStart)
		fmt.Print("") // For breakpoint.
	}

	return result
}

func (c *Chamber) UpdateStart() {
	c.highestStart = 0
	for _, block := range c.blocks {
		c.highestStart = shared.Max(c.highestStart, block.dy)
	}
}

func (c *Chamber) Print() {
	for y := c.highestStart; y >= 1; y-- {
		fmt.Print("|")

		for x := (int64)(0); x < c.width; x++ {
			block := c.GetCollidingBlockAt(x, y)
			if block == nil {
				fmt.Print(".")
			} else if block.isMoving {
				fmt.Print("@")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println("|")
	}

	fmt.Print("+")
	for x := (int64)(0); x < c.width; x++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (c *Chamber) GetCollidingBlockAt(x, y int64) *Block {
	for _, block := range c.blocks {
		if block.IsCollidingAt(x, y) {
			return block
		}
	}
	return nil
}

func (c *Chamber) IsColliding(block *Block) bool {
	if block.dy <= 0 || block.dx < 0 || (block.dx+block.shape.width) > c.width {
		return true
	}

	for i := len(c.blocks) - 1; i >= 0; i-- {
		otherBlock := c.blocks[i]

		if otherBlock.isMoving {
			continue
		}

		if block.IsColliding(otherBlock) {
			return true
		}
	}

	return false
}

func parseInput(input string) []int64 {
	result := []int64{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			break
		}

		for _, c := range line {
			if c == '<' {
				result = append(result, -1)
			} else if c == '>' {
				result = append(result, +1)
			}
		}
	}

	return result
}
