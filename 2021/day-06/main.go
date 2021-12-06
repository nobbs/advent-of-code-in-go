package main

import (
	"fmt"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	fishes := []int{}
	for _, n := range strings.Split(lines[0], ",") {
		fishes = append(fishes, util.ParseInt(n))
	}

	// play 80 rounds
	for i := 0; i < 80; i++ {
		newFishes := []int{}

		for i, n := range fishes {
			if n == 0 {
				newFishes = append(newFishes, 8)
				fishes[i] = 6
			} else {
				fishes[i] -= 1
			}
		}

		fishes = append(fishes, newFishes...)
	}

	return len(fishes)
}

func partTwo(lines []string) int {
	fishes := map[int]int{}
	for _, n := range strings.Split(lines[0], ",") {
		fishes[util.ParseInt(n)]++
	}

	// play 256 rounds
	for i := 0; i < 256; i++ {
		nextRound := map[int]int{}
		for k := 0; k <= 8; k++ {
			nextRound[k] += fishes[(k+1)%9]
		}
		nextRound[6] += fishes[0]
		fishes = nextRound
	}

	result := 0
	for _, v := range fishes {
		result += v
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
