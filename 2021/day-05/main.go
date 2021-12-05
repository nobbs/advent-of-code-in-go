package main

import (
	"fmt"
	"regexp"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Point struct {
	x, y int
}

type Grid map[Point]uint

var (
	origin = Point{0, 0}
)

func (p *Point) DirectionVector(q *Point) *Point {
	d := func(a, b int) int {
		if a == b {
			return 0
		} else if a < b {
			return 1
		}
		return -1
	}

	return &Point{
		x: d(p.x, q.x),
		y: d(p.y, q.y),
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Norm(p, q *Point) int {
	return Abs(p.x-q.x) + Abs(p.y-q.y)
}

func (p *Point) isNonDiagonalLine(q *Point) bool {
	d := p.DirectionVector(q)
	return Norm(d, &origin) != 2
}

func pointsOnLine(p, q *Point) func() (*Point, bool) {
	i := -1
	d := p.DirectionVector(q)

	return func() (*Point, bool) {
		if p.x+i*d.x == q.x && p.y+i*d.y == q.y {
			return nil, false
		}
		i++
		return &Point{
			x: p.x + i*d.x,
			y: p.y + i*d.y,
		}, true
	}
}

func partOne(lines []string) int {
	r := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	grid := Grid{}
	for _, line := range lines {
		g := r.FindStringSubmatch(line)
		p := &Point{x: util.ParseInt(g[1]), y: util.ParseInt(g[2])}
		q := &Point{x: util.ParseInt(g[3]), y: util.ParseInt(g[4])}

		if p.isNonDiagonalLine(q) {
			iter := pointsOnLine(p, q)
			for s, ok := iter(); ok; s, ok = iter() {
				grid[*s]++
			}
		}
	}

	result := 0
	for _, v := range grid {
		if v > 1 {
			result++
		}
	}

	return result
}

func partTwo(lines []string) int {
	r := regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

	grid := Grid{}
	for _, line := range lines {
		g := r.FindStringSubmatch(line)
		p := &Point{x: util.ParseInt(g[1]), y: util.ParseInt(g[2])}
		q := &Point{x: util.ParseInt(g[3]), y: util.ParseInt(g[4])}

		iter := pointsOnLine(p, q)
		for s, ok := iter(); ok; s, ok = iter() {
			grid[*s]++
		}
	}

	result := 0
	for _, v := range grid {
		if v > 1 {
			result++
		}
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
