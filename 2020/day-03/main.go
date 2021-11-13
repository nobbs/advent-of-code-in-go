package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	result := 0

	// walk in a (3, 1) stride from top-left to bottom and count #
	Y := len(lines)
	X := len(lines[0])

	for y, x := 0, 0; y < Y; y, x = y+1, (x+3)%X {
		if lines[y][x] == '#' {
			result++
		}
	}

	return result
}

func partTwo(lines []string) int {
	Y := len(lines)
	X := len(lines[0])

	checkSlope := func(dy, dx int) int {
		result := 0

		for y, x := 0, 0; y < Y; y, x = y+dy, (x+dx)%X {
			if lines[y][x] == '#' {
				result++
			}
		}

		return result
	}

	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, slope := range slopes {
		product *= checkSlope(slope[1], slope[0])
	}

	return product
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
