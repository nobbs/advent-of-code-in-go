package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) (r int) {
	var prev int = util.ParseInt(lines[0])
	for _, line := range lines[1:] {
		curr := util.ParseInt(line)
		if curr > prev {
			r++
		}
		prev = curr
	}

	return r
}

func partTwo(lines []string) (r int) {
	slidingWindow := func(index int) int {
		return util.ParseInt(lines[index]) + util.ParseInt(lines[index+1]) + util.ParseInt(lines[index+2])
	}

	var prev int = slidingWindow(0)
	for i := 1; i < len(lines)-2; i++ {
		curr := slidingWindow(i)
		if curr > prev {
			r++
		}
		prev = curr
	}

	return r
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
