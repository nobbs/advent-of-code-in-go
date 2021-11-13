package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

const (
	inactive rune = '.'
	active   rune = '#'
)

type Point interface {
	Add(Point) Point
}

type Point3D struct {
	x, y, z int
}

type Point4D struct {
	x, y, z, w int
}

func (p *Point3D) Add(q *Point3D) Point3D {
	return Point3D{p.x + q.x, p.y + q.y, p.z + q.z}
}

func (p *Point4D) Add(q *Point4D) Point4D {
	return Point4D{p.x + q.x, p.y + q.y, p.z + q.z, p.w + q.w}
}

type Grid interface {
	print()
	fill([]string)
	countNeighbors(Point, []Point) int
	cycle() Grid
}

type Grid3D map[Point3D]bool
type Grid4D map[Point4D]bool

func (g *Grid3D) print() {
	for k, p := range *g {
		if p {
			fmt.Println(k, p)
		}
	}
}

func (g *Grid4D) print() {
	for k, p := range *g {
		if p {
			fmt.Println(k, p)
		}
	}
}

func (g *Grid3D) fill(lines []string) {
	for y, line := range lines {
		for x, cell := range line {
			if cell == active {
				(*g)[Point3D{x, y, 0}] = true
			}
		}
	}
}

func (g *Grid4D) fill(lines []string) {
	for y, line := range lines {
		for x, cell := range line {
			if cell == active {
				(*g)[Point4D{x, y, 0, 0}] = true
			}
		}
	}
}

func (g *Grid3D) countNeighbors(directions []Point3D, p Point3D) (sum int) {
	for _, d := range directions {
		np := p.Add(&d)
		if v, ok := (*g)[np]; ok && v {
			sum++
		}
	}
	return
}

func (g *Grid4D) countNeighbors(directions []Point4D, p Point4D) (sum int) {
	for _, d := range directions {
		np := p.Add(&d)
		if v, ok := (*g)[np]; ok && v {
			sum++
		}
	}
	return
}

func generateDirections3D() (p []Point3D) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == y && y == z && z == 0 {
					continue
				}
				p = append(p, Point3D{x, y, z})
			}
		}
	}
	return
}

func generateDirections4D() (p []Point4D) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x == y && y == z && z == w && w == 0 {
						continue
					}
					p = append(p, Point4D{x, y, z, w})
				}
			}
		}
	}
	return
}

func (g *Grid3D) cycle() Grid3D {
	ng := Grid3D{}
	directions := generateDirections3D()

	for k := range *g {
		for _, d := range append(directions, Point3D{0, 0, 0}) {
			cp := k.Add(&d)
			if _, ok := ng[cp]; !ok {
				sum := g.countNeighbors(directions, cp)
				gv := (*g)[cp]

				switch {
				case gv && (sum == 2 || sum == 3):
					ng[cp] = true
				case !gv && (sum == 3):
					ng[cp] = true
				default:
					ng[cp] = false
				}
			}

		}
	}

	return ng
}

func (g *Grid4D) cycle() Grid4D {
	ng := Grid4D{}
	directions := generateDirections4D()

	for k := range *g {
		for _, d := range append(directions, Point4D{0, 0, 0, 0}) {
			cp := k.Add(&d)
			if _, ok := ng[cp]; !ok {
				sum := g.countNeighbors(directions, cp)
				gv := (*g)[cp]

				switch {
				case gv && (sum == 2 || sum == 3):
					ng[cp] = true
				case !gv && (sum == 3):
					ng[cp] = true
				default:
					ng[cp] = false
				}
			}

		}
	}

	return ng
}

func partOne(lines []string) int {
	grid := Grid3D{}
	grid.fill(lines)

	for i := 0; i < 6; i++ {
		grid = grid.cycle()
	}

	result := 0
	for _, v := range grid {
		if v {
			result++
		}
	}

	return result
}

func partTwo(lines []string) int {
	grid := Grid4D{}
	grid.fill(lines)

	for i := 0; i < 6; i++ {
		grid = grid.cycle()
	}

	result := 0
	for _, v := range grid {
		if v {
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
