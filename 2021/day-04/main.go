package main

import (
	"fmt"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Cell struct {
	val    int
	marked bool
}

type Board struct {
	b          [][]Cell
	rows, cols int
}

func NewBoard() *Board {
	b := &Board{}
	b.rows, b.cols = 5, 5
	b.b = [][]Cell{}

	for r := 0; r < b.rows; r++ {
		b.b = append(b.b, make([]Cell, b.cols))
	}

	return b
}

func parseBoards(lines []string) []Board {
	boards := []Board{}
	b := NewBoard()
	r := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			boards = append(boards, *b)
			b = NewBoard()
			r = 0
		} else {
			for c, v := range strings.Fields(line) {
				b.b[r][c] = Cell{val: util.ParseInt(v), marked: false}
			}
			r++
		}
	}

	return boards
}

func (b *Board) hasWon() bool {
	checkRow := func(r int) bool {
		for c := 0; c < b.cols; c++ {
			if !b.b[r][c].marked {
				return false
			}
		}
		return true
	}

	checkColumn := func(c int) bool {
		for r := 0; r < b.rows; r++ {
			if !b.b[r][c].marked {
				return false
			}
		}
		return true
	}

	for r := 0; r < b.rows; r++ {
		if checkRow(r) {
			return true
		}
	}

	for c := 0; c < b.cols; c++ {
		if checkColumn(c) {
			return true
		}
	}

	return false
}

func (b *Board) markNumber(num int) bool {
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.cols; c++ {
			if b.b[r][c].val == num {
				b.b[r][c].marked = true
				return true
			}
		}
	}
	return false
}

func (b *Board) score(called int) int {
	sum := 0

	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.cols; c++ {
			if !b.b[r][c].marked {
				sum += b.b[r][c].val
			}
		}
	}

	return sum * called
}

func partOne(lines []string) int {
	// parse boards
	boards := parseBoards(append(lines[2:], ""))

	// parse called numbers and play them
	for _, n := range strings.Split(lines[0], ",") {
		num := util.ParseInt(n)

		for _, board := range boards {
			if board.markNumber(num) {
				if board.hasWon() {
					return board.score(num)
				}
			}
		}
	}

	return -1
}

func partTwo(lines []string) int {
	// parse boards
	boards := parseBoards(append(lines[2:], ""))
	boardsDone := 0

	// parse called numbers
	for _, n := range strings.Split(lines[0], ",") {
		num := util.ParseInt(n)
		for _, board := range boards {
			if !board.hasWon() && board.markNumber(num) {
				if board.hasWon() {
					boardsDone++
					if boardsDone == len(boards) {
						return board.score(num)
					}
				}
			}
		}
	}

	return -1
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
