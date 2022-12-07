package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
	root *directory
	cwd  *directory
}

func (*Day7) Day() int {
	return 7
}

func (*Day7) IsTest() bool {
	return false
}

func (puzzle *Day7) Run(input string) {
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			puzzle.parseLine(line)
		}
	}

	//puzzle.print()
	puzzle.part1()
	puzzle.part2()
}

const printIndentSize = 2

func (puzzle *Day7) part1() {
	total := 0

	puzzle.part1Core(puzzle.root, &total, 100000)

	fmt.Printf("Part1: %d\n", total)
}

func (puzzle *Day7) part1Core(dir *directory, total *int, threshold int) {
	if dir.totalSize <= threshold {
		*total += dir.totalSize
	}

	for _, sub := range dir.directories {
		puzzle.part1Core(sub, total, threshold)
	}
}

func (puzzle *Day7) part2() {
	const maxSize = 70_000_000
	const neededForUpdate = 30_000_000

	availableSize := maxSize - puzzle.root.totalSize
	needToFree := neededForUpdate - availableSize

	bestDirectory := puzzle.root
	puzzle.findDirectory(&bestDirectory, puzzle.root, needToFree)

	fmt.Printf("Part2: %d\n", bestDirectory.totalSize)
}

func (puzzle *Day7) findDirectory(best **directory, current *directory, size int) {
	for _, dir := range current.directories {
		if dir.totalSize >= size && dir.totalSize < (*best).totalSize {
			*best = dir
		}
		puzzle.findDirectory(best, dir, size)
	}
}

func (puzzle *Day7) print() {
	puzzle.root.print(0)
}

func (puzzle *Day7) parseCommand(line string) {
	if strings.HasPrefix(line, "cd ") {
		dir := line[3:]
		if strings.HasPrefix(dir, "/") {
			puzzle.root = puzzle.newDirectory("/", nil)
			puzzle.cwd = puzzle.root
		} else if dir == ".." {
			puzzle.cwd = puzzle.cwd.parent
		} else {
			puzzle.cwd = puzzle.cwd.ensureSubDirectory(puzzle, dir)
		}
	}

	// Nothing to do on 'ls', just wait for comming lines.
}

func (puzzle *Day7) parseFsEntry(line string) {
	if strings.HasPrefix(line, "dir ") {
		name := line[4:]
		puzzle.cwd.ensureSubDirectory(puzzle, name)
	} else {
		parts := strings.Split(line, " ")
		size, _ := strconv.Atoi(parts[0])
		name := parts[1]
		puzzle.cwd.addFile(puzzle, name, size)
	}
}

func (puzzle *Day7) parseLine(line string) {
	if strings.HasPrefix(line, "$ ") {
		puzzle.parseCommand(line[2:])
	} else {
		puzzle.parseFsEntry(line)
	}
}

type file struct {
	name      string
	size      int
	directory *directory

	_ struct{}
}

func (f *file) print(level int) {
	format := fmt.Sprintf("%%%ds- %%s (file, size=%%d)\n", level*printIndentSize)
	fmt.Printf(format, "", f.name, f.size)
}

type directory struct {
	name        string
	totalSize   int
	parent      *directory
	directories []*directory
	files       []*file

	_ struct{}
}

func (*Day7) newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:        name,
		parent:      parent,
		directories: []*directory{},
		files:       []*file{},
	}
}

func (parent *directory) ensureSubDirectory(puzzle *Day7, name string) *directory {
	for _, existing := range parent.directories {
		if existing.name == name {
			return existing
		}
	}

	new := puzzle.newDirectory(name, parent)
	parent.directories = append(parent.directories, new)

	return new
}

func (dir *directory) updateTotalSize(deltaSize int) {
	dir.totalSize += deltaSize
	if dir.parent != nil {
		dir.parent.updateTotalSize(deltaSize)
	}
}

func (dir *directory) addFile(puzzle *Day7, name string, size int) {
	file := &file{
		name:      name,
		size:      size,
		directory: dir,
	}

	dir.files = append(dir.files, file)
	dir.updateTotalSize(size)
}

func (dir *directory) print(level int) {
	format := fmt.Sprintf("%%%ds- %%s (dir, totalSize=%%d)\n", level*printIndentSize)
	fmt.Printf(format, "", dir.name, dir.totalSize)

	for _, childDir := range dir.directories {
		childDir.print(level + 1)
	}

	for _, file := range dir.files {
		file.print(level + 1)
	}
}
