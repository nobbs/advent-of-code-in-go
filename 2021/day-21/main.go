package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Dice struct {
	sides int
	rolls int
}

type Deterministic struct {
	positions [2]int
	scores    [2]int
	dice      Dice
}

type Dirac struct {
	positions     [2]int
	scores        [2]int
	currentPlayer int
}

func (d *Dice) roll() int {
	v := (d.rolls)%d.sides + 1
	d.rolls++
	return v
}

func (g *Deterministic) move(player int, moves int) {
	g.positions[player] = (g.positions[player]+moves-1)%10 + 1
	g.scores[player] += g.positions[player]
}

func (g *Deterministic) play() int {
	currentPlayer := 0

	for g.scores[0] < 1000 && g.scores[1] < 1000 {
		a := g.dice.roll()
		b := g.dice.roll()
		c := g.dice.roll()

		g.move(int(currentPlayer), a+b+c)

		currentPlayer = currentPlayer ^ 1
	}

	return currentPlayer ^ 1
}

func partOne(lines []string) int {
	positionA := util.ParseInt(lines[0][len(lines[0])-1:])
	positionB := util.ParseInt(lines[1][len(lines[1])-1:])

	g := Deterministic{positions: [2]int{positionA, positionB}}
	g.dice = Dice{sides: 100}
	winner := g.play()

	return g.scores[winner^1] * g.dice.rolls
}

func partTwo(lines []string) int {
	positionA := util.ParseInt(lines[0][len(lines[0])-1:])
	positionB := util.ParseInt(lines[1][len(lines[1])-1:])

	distribution := map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
	cache := map[Dirac]int{}

	g := Dirac{positions: [2]int{positionA, positionB}}

	var play func(g *Dirac) int
	play = func(g *Dirac) int {
		// check if game is already cached
		if v, ok := cache[*g]; ok {
			return v
		}

		// if not, play all possible games starting from this one
		sum := 0
		for k, v := range distribution {
			next := Dirac{
				currentPlayer: g.currentPlayer ^ 1,
				positions:     g.positions,
				scores:        g.scores,
			}
			next.positions[g.currentPlayer] = (g.positions[g.currentPlayer]+k-1)%10 + 1
			next.scores[g.currentPlayer] += next.positions[g.currentPlayer]

			// check if this game is won
			if next.scores[g.currentPlayer] >= 21 {
				// if player A has won, add the number of possibilities
				if g.currentPlayer == 0 {
					sum += v
				}
			} else {
				// play again
				sum += v * play(&next)
			}
		}

		// cache the results
		cache[*g] = sum
		return sum
	}

	return play(&g)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
