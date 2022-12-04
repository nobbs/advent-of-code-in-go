package main

import (
	"fmt"
	"sort"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	var result int

	calories := 0
	for _, line := range lines {
		if line == "" {
			result = util.MaxInt(result, calories)
			calories = 0
		} else {
			calories += util.ParseInt(line)
		}
	}

	return result
}

func partTwo(lines []string) int {
	allCalories := []int{}

	calories := 0
	for _, line := range lines {
		if line == "" {
			allCalories = append(allCalories, calories)
			calories = 0
		} else {
			calories += util.ParseInt(line)
		}
	}
	allCalories = append(allCalories, calories)

	sort.Slice(allCalories, func(i, j int) bool {
		return allCalories[i] > allCalories[j]
	})

	return allCalories[0] + allCalories[1] + allCalories[2]
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
