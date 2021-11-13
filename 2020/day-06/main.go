package main

import (
	"fmt"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	lines = append(lines, "")

	var sum int
	var currentBlock string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			// check block
			chars := [26]int{}

			for _, b := range []byte(currentBlock) {
				chars[int(b)-int('a')]++
			}

			for _, c := range chars {
				if c > 0 {
					sum++
				}
			}

			// and start anew
			currentBlock = ""
		} else {
			currentBlock = strings.Join([]string{currentBlock, line}, "")
		}
	}

	return sum
}

func partTwo(lines []string) int {
	lines = append(lines, "")

	var sum, n int
	var currentBlock string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			// check block
			chars := [26]int{}
			for _, b := range []byte(currentBlock) {
				chars[int(b)-int('a')]++
			}

			for _, c := range chars {
				if c == n {
					sum++
				}
			}

			// and start anew
			currentBlock = ""
			n = 0
		} else {
			currentBlock = strings.Join([]string{currentBlock, line}, "")
			n++
		}
	}

	return sum
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
