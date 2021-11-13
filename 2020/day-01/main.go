package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	numbers := parseNumbers(lines)
	length := len(numbers)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i]+numbers[j] == 2020 {
				return numbers[i] * numbers[j]
			}
		}
	}

	return 0
}

func partTwo(lines []string) int {
	numbers := parseNumbers(lines)
	length := len(numbers)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				u, v, w := numbers[i], numbers[j], numbers[k]
				if u+v+w == 2020 {
					return u * v * w
				}
			}
		}
	}

	return 0
}

func parseNumbers(lines []string) (numbers []int) {
	for _, text := range lines {
		n := util.ParseInt(text)
		numbers = append(numbers, n)
	}

	return numbers
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
