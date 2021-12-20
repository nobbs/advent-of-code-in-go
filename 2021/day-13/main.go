package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type XY struct {
	x, y int
}

type Grid struct {
	g map[XY]bool
	X int
	Y int
}

type Fold struct {
	n int
	u bool
}

func parseInput(lines []string) (Grid, []Fold) {
	folds := []Fold{}
	grid := Grid{g: map[XY]bool{}}

	r1 := regexp.MustCompile(`^(\d+),(\d+)$`)
	r2 := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if !strings.Contains(line, "fold along") {
			groups := r1.FindStringSubmatch(line)
			x, y := util.ParseInt(groups[1]), util.ParseInt(groups[2])
			grid.g[XY{x, y}] = true

			if x+1 > grid.X {
				grid.X = x + 1
			}
			if y+1 > grid.Y {
				grid.Y = y + 1
			}
		} else {
			groups := r2.FindStringSubmatch(line)
			n, u := util.ParseInt(groups[2]), groups[1] == "y"
			folds = append(folds, Fold{n, u})
		}
	}

	return grid, folds
}

func (g *Grid) fold(f Fold) *Grid {
	newGrid := &Grid{g: map[XY]bool{}}

	if f.u {
		// upwards
		for x := 0; x < g.X; x++ {
			for dy := 0; dy < g.Y; dy++ {
				_, ok1 := g.g[XY{x, f.n - dy}]
				_, ok2 := g.g[XY{x, f.n + dy}]

				if ok1 || ok2 {
					newGrid.g[XY{x, f.n - dy}] = true
				}
			}
		}
		newGrid.X = g.X
		newGrid.Y = g.Y / 2
	} else {
		// to the left
		for y := 0; y < g.Y; y++ {
			for dx := 0; dx < g.X; dx++ {
				_, ok1 := g.g[XY{f.n - dx, y}]
				_, ok2 := g.g[XY{f.n + dx, y}]

				if ok1 || ok2 {
					newGrid.g[XY{f.n - dx, y}] = true
				}
			}
		}
		newGrid.X = g.X / 2
		newGrid.Y = g.Y
	}

	return newGrid
}

func (g *Grid) print() string {
	builder := strings.Builder{}

	for y := 0; y < g.Y; y++ {
		for x := 0; x < g.X; x++ {
			if _, ok := g.g[XY{x, y}]; ok {
				builder.WriteString("#")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func partOne(lines []string) int {
	grid, folds := parseInput(lines)

	// only perform the first fold
	grid = *grid.fold(folds[0])

	return len(grid.g)
}

func partTwo(lines []string) string {
	grid, folds := parseInput(lines)

	// perform all folds and print out the result
	for _, fold := range folds {
		grid = *grid.fold(fold)
	}

	return grid.print()
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:\n" + solutionTwo)
}
