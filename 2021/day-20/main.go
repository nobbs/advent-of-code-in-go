package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type XY struct {
	x, y int
}

type Grid struct {
	g [][]byte

	w, h, offset int
}

func dotsToInt(s string) int {
	s = strings.ReplaceAll(s, "#", "1")
	s = strings.ReplaceAll(s, ".", "0")

	i, _ := strconv.ParseInt(s, 2, 64)

	return int(i)
}

func parseInput(lines []string) *Grid {
	grid := &Grid{}
	grid.offset = 0
	grid.w = len(lines[0])
	grid.h = len(lines)

	for i := 0; i < grid.h; i++ {
		grid.g = append(grid.g, make([]byte, grid.w))
	}

	for y, line := range lines {
		for x, cell := range line {
			grid.g[y][x] = byte(cell)
		}
	}

	return grid
}

func (grid *Grid) String() string {
	builder := strings.Builder{}

	for _, line := range grid.g {
		builder.WriteString(fmt.Sprintln(string(line)))
	}

	return builder.String()
}

func (g *Grid) neighbours(x, y int, border byte) string {
	D := [9]XY{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{0, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	builder := strings.Builder{}

	for i := 0; i < 9; i++ {
		nx := x + D[i].x
		ny := y + D[i].y

		if 0 <= nx && nx < g.w && 0 <= ny && ny < g.h {
			builder.WriteByte(g.g[ny][nx])
		} else {
			builder.WriteByte(border)
		}
	}

	return builder.String()
}

func (g *Grid) enhanceImage(iem string) *Grid {
	next := &Grid{}
	next.offset = g.offset + 1
	next.w = g.w + 2
	next.h = g.h + 2

	var border byte
	if iem[0] == '.' {
		border = '.'
	} else {
		if next.offset%2 == 1 {
			border = '.'
		} else {
			border = '#'
		}
	}

	for i := 0; i < next.h; i++ {
		next.g = append(next.g, make([]byte, next.w))
	}

	for y := 0; y < next.h; y++ {
		for x := 0; x < next.w; x++ {
			iemIndex := dotsToInt(g.neighbours(x-1, y-1, border))
			next.g[y][x] = iem[iemIndex]
		}
	}

	return next
}

func partOne(lines []string) int {
	iem := lines[0]
	grid := parseInput(lines[2:])

	grid = grid.enhanceImage(iem)
	grid = grid.enhanceImage(iem)

	count := 0
	for _, line := range grid.g {
		for _, cell := range line {
			if cell == '#' {
				count++
			}
		}
	}

	return count
}

func partTwo(lines []string) int {
	iem := lines[0]
	grid := parseInput(lines[2:])

	for i := 0; i < 50; i++ {
		grid = grid.enhanceImage(iem)
	}

	count := 0
	for _, line := range grid.g {
		for _, cell := range line {
			if cell == '#' {
				count++
			}
		}
	}

	return count
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
