package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Interval struct {
	from, to int
}

func (i *Interval) Contains(n int) bool {
	if i.from <= n && n <= i.to {
		return true
	}

	return false
}

func parseInput(lines []string) ([]Interval, []int, [][]int) {
	intervals := []Interval{}

	// parse intervals
	i := 0
	r := regexp.MustCompile(`.*: (\d+)-(\d+) or (\d+)-(\d+)$`)
	for ; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "" {
			break
		}

		groups := r.FindStringSubmatch(lines[i])
		intervals = append(intervals, Interval{util.ParseInt(groups[1]), util.ParseInt(groups[2])})
		intervals = append(intervals, Interval{util.ParseInt(groups[3]), util.ParseInt(groups[4])})
	}

	// parse your ticket
	i += 2
	ownNumbers := []int{}
	for _, n := range strings.Split(lines[i], ",") {
		ownNumbers = append(ownNumbers, util.ParseInt(n))
	}
	i += 3

	otherNumbers := [][]int{}
	// parse nearby tickets
	for ; i < len(lines); i++ {
		numbers := []int{}
		for _, n := range strings.Split(lines[i], ",") {
			numbers = append(numbers, util.ParseInt(n))
		}
		otherNumbers = append(otherNumbers, numbers)
	}

	return intervals, ownNumbers, otherNumbers
}

func validateTickets(intervals []Interval, ownNumbers []int, otherNumbers [][]int) (int, [][]int) {
	validNumbers := [][]int{}
	result := 0

	// and now check the other tickets numbers against your intervals
	for _, numbers := range otherNumbers {
		ticketIsValid := true
		for _, n := range numbers {
			isValid := false
			for _, interval := range intervals {
				if interval.Contains(n) {
					isValid = true
				}
			}

			if !isValid {
				result += n
				ticketIsValid = false
			}
		}

		if ticketIsValid {
			validNumbers = append(validNumbers, numbers)
		}
	}

	return result, validNumbers
}

func partOne(lines []string) int {
	intervals, ownNumbers, otherNumbers := parseInput(lines)
	result, _ := validateTickets(intervals, ownNumbers, otherNumbers)
	return result
}

// my solution for the second part is really not that great, so I'll try to
// explain it a bit with comments
func partTwo(lines []string) int {
	// same as part one, but validNumbers now contains only the valid rows
	intervals, ownNumbers, otherNumbers := parseInput(lines)
	_, validNumbers := validateTickets(intervals, ownNumbers, otherNumbers)

	// nBuckets is the number of different fields (each field has got two
	// intervals). buckets is a 2d-matrix that will be used to count whether
	// the how many of the valid tickets column `col` matches the field `row`.
	nBuckets := len(intervals) / 2
	buckets := [][]int{}
	for i := 0; i < nBuckets; i++ {
		buckets = append(buckets, make([]int, nBuckets))
	}

	// max is used to find the max value in the buckets section.
	max := 0
	for row := 0; row < len(validNumbers); row++ {
		for col := 0; col < nBuckets; col++ {
			for bucketIdx := 0; bucketIdx < nBuckets; bucketIdx++ {
				curr := validNumbers[row][col]
				if intervals[2*bucketIdx].Contains(curr) || intervals[2*bucketIdx+1].Contains(curr) {
					buckets[bucketIdx][col]++

					if buckets[bucketIdx][col] > max {
						max = buckets[bucketIdx][col]
					}
				}
			}
		}
	}

	for r, row := range buckets {
		for c := range row {
			if buckets[r][c] == max {
				buckets[r][c] = 1
			} else {
				buckets[r][c] = 0
			}
		}
	}

	countOnes := func(col int) (sum int) {
		for r := 0; r < nBuckets; r++ {
			sum += buckets[r][col]
		}
		return
	}

	findOne := func(col int) int {
		for r := 0; r < nBuckets; r++ {
			if buckets[r][col] == 1 {
				return r
			}
		}
		return -1
	}

	clearColAndRow := func(row, col int) {
		for r := 0; r < nBuckets; r++ {
			buckets[r][col] = 0
		}
		for c := 0; c < nBuckets; c++ {
			buckets[row][c] = 0
		}
	}

	order := make([]int, nBuckets)
	for left := nBuckets; left > 0; {
		for col := 0; col < nBuckets; col++ {
			if ones := countOnes(col); ones == 1 {
				idx := findOne(col)
				order[col] = idx
				clearColAndRow(idx, col)
				left--
			}
		}
	}

	result := 1
	for i, c := range order {
		if c < 6 {
			result *= ownNumbers[i]
		}
	}
	fmt.Println(order)

	return result
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
