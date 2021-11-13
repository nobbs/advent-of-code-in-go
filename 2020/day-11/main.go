package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Grid [][]rune

const (
	nothing  rune = '.'
	empty    rune = 'L'
	occupied rune = '#'
)

var (
	dirY = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	dirX = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
)

func (g Grid) print() {
	for _, row := range g {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func NewGrid(width, height int) Grid {
	g := Grid{}
	for y := 0; y < height; y++ {
		g = append(g, make([]rune, width))
	}
	return g
}

func ImportGrid(lines []string) Grid {
	height, width := len(lines), len(lines[0])
	g := NewGrid(width, height)

	for y, line := range lines {
		for x, cell := range line {
			g[y][x] = cell
		}
	}

	return g
}

func countNeighbors(g Grid, y, x int) int {
	sum := 0
	height, width := len(g), len(g[0])

	for i := 0; i < 8; i++ {
		dy, dx := dirY[i], dirX[i]
		ny, nx := y+dy, x+dx
		if 0 <= ny && ny < height && 0 <= nx && nx < width {
			if g[ny][nx] == occupied {
				sum++
			}
		}
	}

	return sum
}

func countNeighborsInLine(g Grid, y, x int) int {
	sum := 0
	height, width := len(g), len(g[0])

	for i := 0; i < 8; i++ {
		dy, dx := dirY[i], dirX[i]

	innerloop:
		for m := 1; m < width && m < height; m++ {
			ny, nx := y+m*dy, x+m*dx

			if 0 <= ny && ny < height && 0 <= nx && nx < width {
				switch g[ny][nx] {
				case occupied:
					sum++
					break innerloop
				case empty:
					break innerloop
				}
			}
		}
	}

	return sum
}

func iterateModel(g Grid, counterFn func(Grid, int, int) int, limit int) (Grid, bool) {
	hasChanged := false
	height, width := len(g), len(g[0])

	newGrid := NewGrid(width, height)

	for y, row := range g {
		for x, cell := range row {
			switch cell {
			case nothing:
				newGrid[y][x] = nothing
			case empty:
				if counterFn(g, y, x) == 0 {
					newGrid[y][x] = occupied
					hasChanged = true
				} else {
					newGrid[y][x] = empty
				}
			case occupied:
				if counterFn(g, y, x) >= limit {
					newGrid[y][x] = empty
					hasChanged = true
				} else {
					newGrid[y][x] = occupied
				}
			}
		}
	}

	return newGrid, hasChanged
}

func partOne(lines []string) int {
	grid := ImportGrid(lines)

	// iterate model til nothing changes anymore
	for i := true; i; grid, i = iterateModel(grid, countNeighbors, 4) {
		continue
	}

	// count occupied seats
	result := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == occupied {
				result++
			}
		}
	}

	return result
}

func partTwo(lines []string) int {
	grid := ImportGrid(lines)

	// iterate model til nothing changes anymore
	for i := true; i; grid, i = iterateModel(grid, countNeighborsInLine, 5) {

		// grid.print()
		// fmt.Println()
		continue
	}

	// count occupied seats
	result := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == occupied {
				result++
			}
		}
	}

	return result
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
