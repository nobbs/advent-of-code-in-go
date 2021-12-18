package main

import (
	"fmt"
	"sort"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Point struct {
	r, c int
}

var (
	D = [4]Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func parseInput(lines []string) (height, width int, grid [][]int) {
	// parsing the heightmap
	width, height = len(lines[0]), len(lines)
	for i := 0; i < height; i++ {
		grid = append(grid, make([]int, width))
	}

	for i, line := range lines {
		for j, v := range line {
			grid[i][j] = util.ParseInt(string(v))
		}
	}

	return
}

func findLowPoints(height, width int, grid [][]int) []Point {
	// and now looking for low points
	isLowPoint := func(r, c int) bool {
		for _, d := range D {
			nr, nc := r+d.r, c+d.c
			if 0 <= nr && nr < height && 0 <= nc && nc < width {
				if grid[nr][nc] <= grid[r][c] {
					return false
				}
			}
		}
		return true
	}

	// compute the solution
	lowPoints := []Point{}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if isLowPoint(r, c) {
				lowPoints = append(lowPoints, Point{r, c})
			}
		}
	}

	return lowPoints
}

func partOne(lines []string) (result int) {
	height, width, grid := parseInput(lines)

	for _, p := range findLowPoints(height, width, grid) {
		result += 1 + grid[p.r][p.c]
	}
	return
}

func partTwo(lines []string) int {
	height, width, grid := parseInput(lines)
	lowPoints := findLowPoints(height, width, grid)

	var basinFill func(r, c int, visited [][]bool) int
	basinFill = func(r, c int, visited [][]bool) int {
		visited[r][c] = true
		size := 1

		for _, d := range D {
			nr, nc := r+d.r, c+d.c
			if 0 <= nr && nr < height && 0 <= nc && nc < width {
				if !visited[nr][nc] && grid[nr][nc] != 9 && grid[nr][nc] > grid[r][c] {
					size += basinFill(nr, nc, visited)
				}
			}
		}

		return size
	}

	// for every low point, do a basin fill to figure out their size
	basinSizes := []int{}
	for _, p := range lowPoints {
		visited := [][]bool{}
		for i := 0; i < height; i++ {
			visited = append(visited, make([]bool, width))
		}
		basinSizes = append(basinSizes, basinFill(p.r, p.c, visited))
	}

	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
