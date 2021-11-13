package main

import (
	"fmt"
	"sort"

	"github.com/nobbs/advent-of-code-in-go/util"
)

// global memoization to prevent dfs from searching the same path multiple times
var memo = map[uint]uint{}

func dfs(adjacency [][]uint, start, goal uint) uint {
	x, ok := memo[start]
	if ok {
		return x
	}

	if start == goal {
		memo[start] = 1
		return 1
	}

	var sum uint = 0
	for _, v := range adjacency[start] {
		if v > 0 {
			sum += dfs(adjacency, v, goal)
		}
	}

	memo[start] = sum
	return sum
}

func partOne(lines []string) int {
	numbers := []uint{0}
	for _, line := range lines {
		n := util.ParseInt(line)
		numbers = append(numbers, uint(n))
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	x, y := 0, 1
	for i := 0; i < len(numbers)-1; i++ {
		d := numbers[i+1] - numbers[i]
		if d == 1 {
			x++
		} else if d == 3 {
			y++
		}
	}

	return x * y
}

func partTwo(lines []string) int {
	// reset memoization
	memo = make(map[uint]uint)

	numbers := []uint{0}
	for _, line := range lines {
		n := util.ParseInt(line)
		numbers = append(numbers, uint(n))
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	total := len(numbers)
	adjacencyList := make([][]uint, numbers[len(numbers)-1])

	for i := 0; i < total; i++ {
		for j := i + 1; j < total && numbers[j]-numbers[i] <= 3; j++ {
			I, J := numbers[i], numbers[j]
			adjacencyList[I] = append(adjacencyList[I], J)
		}
	}

	x := dfs(adjacencyList, 0, numbers[len(numbers)-1])
	return int(x)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
