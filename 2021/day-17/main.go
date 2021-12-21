package main

import (
	"fmt"
	"regexp"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type XY struct {
	x, y   int
	vx, vy int
}

type TargetArea struct {
	area [2]XY
}

func (t *TargetArea) isInArea(p *XY) bool {
	return t.area[0].x <= p.x && p.x <= t.area[1].x && t.area[0].y <= p.y && p.y <= t.area[1].y
}

func (t *TargetArea) hasOvershot(p *XY) bool {
	if p.x > t.area[1].x || p.y < t.area[0].y {
		return true
	}

	return false
}

func nextStep(p *XY) *XY {
	next := &XY{
		x:  p.x + p.vx,
		y:  p.y + p.vy,
		vy: p.vy - 1,
	}

	switch {
	case p.vx == 0:
		next.vx = 0
	case p.vx > 0:
		next.vx = p.vx - 1
	case p.vx < 0:
		next.vx = p.vx + 1
	}

	return next
}

func partOne(lines []string) int {
	r := regexp.MustCompile(`target area: x=([-\w]+)..([-\w]+), y=([-\w]+)..([-\w]+)`)
	groups := r.FindStringSubmatch(lines[0])

	target := &TargetArea{area: [2]XY{
		{
			x: util.ParseInt(groups[1]),
			y: util.ParseInt(groups[3]),
		},
		{
			x: util.ParseInt(groups[2]),
			y: util.ParseInt(groups[4]),
		},
	}}

	maxY := 0
	for vy := 1; vy <= util.MaxInt(util.AbsInt(target.area[0].y), util.AbsInt(target.area[1].y)); vy++ {
		for vx := 2 * target.area[1].x; vx > 0; vx-- {
			p := &XY{vx: vx, vy: vy}
			pY := 0
			for !target.hasOvershot(p) {
				p = nextStep(p)
				if p.y > pY {
					pY = p.y
				}

				if target.isInArea(p) {

					if pY > maxY {
						maxY = pY
					}

					break
				}
			}
		}
	}

	return maxY
}

func partTwo(lines []string) int {
	r := regexp.MustCompile(`target area: x=([-\w]+)..([-\w]+), y=([-\w]+)..([-\w]+)`)
	groups := r.FindStringSubmatch(lines[0])

	target := &TargetArea{area: [2]XY{
		{
			x: util.ParseInt(groups[1]),
			y: util.ParseInt(groups[3]),
		},
		{
			x: util.ParseInt(groups[2]),
			y: util.ParseInt(groups[4]),
		},
	}}

	count := 0
	for vy := util.MaxInt(util.AbsInt(target.area[0].y), util.AbsInt(target.area[1].y)); vy >= -util.MaxInt(util.AbsInt(target.area[0].y), util.AbsInt(target.area[1].y)); vy-- {
		for vx := 2 * target.area[1].x; vx > 0; vx-- {
			p := &XY{vx: vx, vy: vy}
			pY := 0
			for !target.hasOvershot(p) {
				p = nextStep(p)
				if p.y > pY {
					pY = p.y
				}

				if target.isInArea(p) {
					count++
					break
				}
			}
		}
	}

	return count
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
