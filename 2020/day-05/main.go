package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Seat struct {
	row, column int
}

func (s *Seat) id() int {
	return 8*s.row + s.column
}

func partitionSeat(side string, from, to int) int {
	if len(side) == 0 {
		return from
	}

	half := (to-from)/2 + 1

	if side[0] == 'F' || side[0] == 'L' {
		return partitionSeat(side[1:], from, to-half)
	} else {
		return partitionSeat(side[1:], from+half, to)
	}
}

func partOne(lines []string) int {
	results := []int{}

	for _, line := range lines {
		row := partitionSeat(line[:7], 0, 127)
		column := partitionSeat(line[7:], 0, 7)
		seat := Seat{row: row, column: column}
		results = append(results, seat.id())
	}

	var max int
	for i, e := range results {
		if i == 0 || e > max {
			max = e
		}
	}

	return int(max)
}

func partTwo(lines []string) int {
	seats := []Seat{}

	for _, line := range lines {
		row := partitionSeat(line[:7], 0, 127)
		column := partitionSeat(line[7:], 0, 7)
		seat := Seat{row: row, column: column}
		seats = append(seats, seat)
	}

	seatGrid := [128][8]bool{}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			seatGrid[i][j] = false
		}
	}

	for _, seat := range seats {
		seatGrid[seat.row][seat.column] = true
	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			if (j == 0 || seatGrid[i][j-1]) && !seatGrid[i][j] && (j == 7 || seatGrid[i][j+1]) {
				return i*8 + j
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
