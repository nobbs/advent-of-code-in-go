package main

import (
	"fmt"
	"regexp"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func parseLine(line string) (string, int, int) {
	r := regexp.MustCompile(`^(nop|acc|jmp)\s(\+|-)(\d*)$`)
	groups := r.FindStringSubmatch(line)

	n := util.ParseInt(groups[2] + groups[3])
	switch groups[1] {
	case "nop":
		return "nop", 0, n
	case "acc":
		return "acc", n, 1
	case "jmp":
		return "jmp", 0, n
	}

	return "", 0, 0
}

type Steps struct {
	op           string
	change, step int
	visited      bool
}

func checkLoops(steps []Steps) (int, bool) {
	acc, num := 0, 0

	for i := 0; num < len(steps); {
		if !steps[i].visited {
			steps[i].visited = true
			acc += steps[i].change

			if steps[i].op == "nop" || steps[i].op == "acc" {
				i++
			} else {
				i += steps[i].step
			}

			if i == len(steps) {
				return acc, false
			}

			num++
		} else {
			return acc, true
		}
	}

	return acc, false
}

func partOne(lines []string) int {
	size := len(lines)
	steps := make([]Steps, size)

	for i, line := range lines {
		op, change, step := parseLine(line)
		steps[i] = Steps{op: op, change: change, step: step, visited: false}
	}

	acc, _ := checkLoops(steps)
	return acc
}

func deepCopy(in []Steps) []Steps {
	out := make([]Steps, len(in))

	for i := range in {
		out[i] = in[i]
	}

	return out
}

func partTwo(lines []string) int {
	size := len(lines)
	steps := make([]Steps, size)

	for i, line := range lines {
		op, change, step := parseLine(line)
		steps[i] = Steps{op: op, change: change, step: step}
	}

	for i, elem := range steps {
		if elem.op == "nop" || elem.op == "jmp" {
			stepsCopy := deepCopy(steps)

			switch elem.op {
			case "nop":
				stepsCopy[i].op = "jmp"
			case "jmp":
				stepsCopy[i].op = "nop"
			}

			if acc, hasLooped := checkLoops(stepsCopy); !hasLooped {
				return acc
			}
		}
	}

	return 0
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
