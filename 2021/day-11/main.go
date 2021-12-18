package main

import (
	"fmt"

	"github.com/juju/collections/deque"
	"github.com/nobbs/advent-of-code-in-go/util"
)

type Point struct {
	r, c int
}

var (
	D = [8]Point{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
)

type Grid [][]int

func NewGrid(height, width int) (grid Grid) {
	for i := 0; i < height; i++ {
		grid = append(grid, make([]int, width))
	}
	return
}

func parseInput(lines []string) (height, width int, grid Grid) {
	// parsing the octopusmap
	width, height = len(lines[0]), len(lines)
	grid = NewGrid(height, width)

	for r, line := range lines {
		for c, v := range line {
			grid[r][c] = util.ParseInt(string(v))
		}
	}

	return
}

func (g *Grid) step() (Grid, int) {
	width, height := len((*g)[0]), len(*g)
	next := NewGrid(height, width)

	flashes := 0
	deque := deque.New()
	toReset := []Point{}

	// first, increase all octopi by 1
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			next[r][c] = ((*g)[r][c] + 1) % 10
			if next[r][c] == 0 {
				deque.PushFront(Point{r, c})
				toReset = append(toReset, Point{r, c})
			}
		}
	}

	// now flash all that are 'ready'
	for np, ok := deque.PopFront(); ok; np, ok = deque.PopFront() {
		p := np.(Point)
		flashes++

		for _, d := range D {
			nr, nc := p.r+d.r, p.c+d.c

			if 0 <= nr && nr < height && 0 <= nc && nc < width {
				next[nr][nc] = (next[nr][nc] + 1) % 10
				if next[nr][nc] == 0 {
					deque.PushBack(Point{nr, nc})
					toReset = append(toReset, Point{nr, nc})
				}
			}
		}
	}

	// now reset all, that have flashed
	for _, p := range toReset {
		next[p.r][p.c] = 0
	}

	return next, flashes
}

func (g *Grid) print() {
	for _, r := range *g {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func partOne(lines []string) int {
	_, _, grid := parseInput(lines)

	totalFlashes := 0
	for s := 0; s < 100; s++ {
		g, n := grid.step()
		totalFlashes += n
		grid = g
	}

	return totalFlashes
}

func partTwo(lines []string) int {
	height, width, grid := parseInput(lines)

	for s := 0; true; s++ {
		g, n := grid.step()
		grid = g

		if n == height*width {
			return s + 1
		}
	}

	return -1
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
