package main

import (
	"fmt"
	"sort"

	"github.com/juju/collections/deque"
	"github.com/nobbs/advent-of-code-in-go/util"
)

var (
	pairs = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
		'>': '<',
	}
)

func partOne(lines []string) int {
	result := 0
	scoring := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, line := range lines {
		d := deque.New()
		isCorrupted := false

		for _, token := range line {
			if isCorrupted {
				break
			}

			switch token {
			case '[', '(', '<', '{':
				d.PushFront(token)
			case ']', ')', '>', '}':
				last, ok := d.Front()
				if !ok || last.(rune) != pairs[token] {
					result += scoring[token]
					isCorrupted = true
				} else {
					d.PopFront()
				}
			}
		}
	}

	return result
}

func partTwo(lines []string) int {
	scores := []int{}
	scoring := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	for _, line := range lines {
		d := deque.New()
		isCorrupted := false

		for _, token := range line {
			if isCorrupted {
				break
			}

			switch token {
			case '[', '(', '<', '{':
				d.PushFront(token)
			case ']', ')', '>', '}':
				last, ok := d.Front()
				if !ok || last.(rune) != pairs[token] {
					isCorrupted = true
				} else {
					d.PopFront()
				}
			}
		}

		// if there are still elements in the deque, then the line was incomplete
		if !isCorrupted && d.Len() != 0 {
			score := 0
			for n, ok := d.PopFront(); ok; n, ok = d.PopFront() {
				score = score*5 + scoring[n.(rune)]
			}
			scores = append(scores, score)
		}

	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	return scores[(len(scores)-1)/2]
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
