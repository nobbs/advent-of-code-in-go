package main

import (
	"fmt"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type mem struct {
	p, pp int
}

func play(input string, totalTurns int) int {
	startTurns := 0
	memory := map[int]mem{}
	prev := 0

	// parse starting input
	for i, n := range strings.Split(input, ",") {
		num := util.ParseInt(n)
		memory[num] = mem{i, -1}
		prev = num
		startTurns++
	}

	for i := startTurns; i < totalTurns; i++ {
		num := memory[prev]

		next := 0
		if num.pp >= 0 {
			next = num.p - num.pp
		}

		cur, ok := memory[next]
		if ok {
			memory[next] = mem{i, cur.p}
		} else {
			memory[next] = mem{i, -1}
		}
		prev = next
	}

	return prev

}

func partOne(lines []string) int {
	return play(lines[0], 2020)
}

func partTwo(lines []string) int {
	return play(lines[0], 30000000)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
