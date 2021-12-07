package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solveProblem(input string, diffCalc func(a, b int) int) int {
	numbers := []int{}

	for _, token := range strings.Split(input, ",") {
		numbers = append(numbers, util.ParseInt(token))
	}

	sumAbsDiffs := func(num int) (result int) {
		for _, n := range numbers {
			result += diffCalc(n, num)
		}
		return
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	min := math.MaxInt
	for i, n := numbers[0], len(numbers)-1; i < numbers[n]; i++ {
		if diff := sumAbsDiffs(i); diff < min {
			min = diff
		}
	}

	return min
}

func partOne(lines []string) int {
	return solveProblem(lines[0], func(a, b int) int {
		return Abs(a - b)
	})
}

func partTwo(lines []string) int {
	return solveProblem(lines[0], func(a, b int) int {
		t := Abs(a - b)
		return (t*t + t) / 2
	})
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
