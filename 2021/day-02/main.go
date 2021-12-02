package main

import (
	"fmt"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Position struct {
	x int
	y int
}

type PositionAndAim struct {
	Position
	aim int
}

func partOne(lines []string) int {
	curPosition := Position{x: 0, y: 0}

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		move, change := tokens[0], util.ParseInt(tokens[1])

		switch move {
		case "forward":
			curPosition.x += change
		case "down":
			curPosition.y += change
		case "up":
			curPosition.y -= change
		}
	}

	return curPosition.x * curPosition.y
}

func partTwo(lines []string) int {
	curPosition := PositionAndAim{Position: Position{0, 0}, aim: 0}

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		move, change := tokens[0], util.ParseInt(tokens[1])

		switch move {
		case "forward":
			curPosition.x += change
			curPosition.y += change * curPosition.aim
		case "down":
			curPosition.aim += change
		case "up":
			curPosition.aim -= change
		}
	}

	return curPosition.x * curPosition.y
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
