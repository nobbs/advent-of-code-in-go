package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string, preambleLength int) int {
	numbers := []int{}

	for _, line := range lines {
		n := util.ParseInt(line)
		numbers = append(numbers, n)
	}

	for i := preambleLength; i < len(numbers); i++ {
		current := numbers[i]

		found := false
		for j := i - preambleLength; j < i; j++ {

			for k := j + 1; k < i; k++ {
				if current == numbers[j]+numbers[k] {
					found = true
					break
				}
			}

		}
		if !found {
			return current
		}
	}

	return -1
}

func partTwo(lines []string, preambleLength int) int {
	goal := partOne(lines, preambleLength)

	numbers := []int{}

	for _, line := range lines {
		n := util.ParseInt(line)
		numbers = append(numbers, n)
	}

	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		min, max := sum, sum

		for j := i + 1; j < len(numbers) && sum < goal; j++ {
			next := numbers[j]
			sum += next

			if min > next {
				min = next
			}
			if max < next {
				max = next
			}
		}

		if sum == goal {
			return min + max
		}
	}

	return -1
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines, 25)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines, 25)
	fmt.Println("Solution for part 2:", solutionTwo)
}
