package main

import (
	"fmt"
	"regexp"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Point struct {
	x, y int
}

var (
	origin = Point{0, 0}
)

func (p *Point) manhattenDistance(q *Point) int {
	return Abs(p.x-q.x) + Abs(p.y-q.y)
}

func partOne(lines []string) int {
	facing := Point{1, 0}
	ship := Point{0, 0}

	r := regexp.MustCompile(`^(\w)(\d+)$`)

	for _, line := range lines {
		groups := r.FindStringSubmatch(line)
		op := rune(groups[1][0])
		change := util.ParseInt(groups[2])

		switch op {
		case 'N':
			ship.y += change
		case 'S':
			ship.y -= change
		case 'E':
			ship.x += change
		case 'W':
			ship.x -= change
		case 'R':
			for n := 0; n < change/90; n++ {
				facing.x, facing.y = facing.y, -facing.x
			}
		case 'L':
			for n := 0; n < change/90; n++ {
				facing.x, facing.y = -facing.y, facing.x
			}
		case 'F':
			ship.x += change * facing.x
			ship.y += change * facing.y
		}
	}

	return ship.manhattenDistance(&origin)
}

func partTwo(lines []string) int {
	ship := Point{0, 0}
	waypoint := Point{10, 1}

	r := regexp.MustCompile(`^(\w)(\d+)$`)

	for _, line := range lines {
		groups := r.FindStringSubmatch(line)
		op := rune(groups[1][0])
		change := util.ParseInt(groups[2])

		switch op {
		case 'N':
			waypoint.y += change
		case 'S':
			waypoint.y -= change
		case 'E':
			waypoint.x += change
		case 'W':
			waypoint.x -= change
		case 'R':
			for n := 0; n < change/90; n++ {
				waypoint.x, waypoint.y = waypoint.y, -waypoint.x
			}
		case 'L':
			for n := 0; n < change/90; n++ {
				waypoint.x, waypoint.y = -waypoint.y, waypoint.x
			}
		case 'F':
			ship.x += change * waypoint.x
			ship.y += change * waypoint.y
		}
	}

	return ship.manhattenDistance(&origin)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
