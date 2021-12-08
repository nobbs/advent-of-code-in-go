package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

var (
	numbers = []string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdfg",
	}
	chars = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
)

const (
	offset = 97
)

func partOne(lines []string) int {
	result := 0

	for _, line := range lines {
		secondPart := strings.Split(line, " | ")[1]
		for _, token := range strings.Fields(secondPart) {
			switch len(token) {
			case 2, 3, 4, 7:
				result++
			}
		}
	}
	return result
}

func partTwo(lines []string) int {
	result := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		inputMap := [10][7]bool{}
		alreadyFound := map[byte]int{}

		// convert the 'learning' input to a bitmap
		for i, token := range strings.Fields(parts[0]) {
			for _, c := range token {
				inputMap[i][c-offset] = true
			}
		}

		// custom rowSum function counts the remaining unknown characters in a row of the bitmap
		rowSum := func(i int) int {
			l := 0
			for _, c := range inputMap[i] {
				if c {
					l++
				}
			}
			return l
		}

		// counts how often the character (column) is remaining in all the rows of the bitmap
		columnSum := func(c int) int {
			l := 0
			for r := range inputMap {
				if inputMap[r][c] {
					l++
				}
			}
			return l
		}

		// remove already mapped characters from the bitmap
		eliminate := func() {
			for _, v := range alreadyFound {
				for i := range inputMap {
					inputMap[i][v] = false
				}
			}
		}

		// eliminate the easy ones, 'b', 'e' and 'f'
		for c := range inputMap[0] {
			switch columnSum(c) {
			case 4:
				alreadyFound['e'-offset] = c
			case 6:
				alreadyFound['b'-offset] = c
			case 9:
				alreadyFound['f'-offset] = c
			}
		}

		// try to find the 'c'
		// it's the one with 8 occurences that also appears in a 1-length token
		eliminate()
		for r := range inputMap {
			if rowSum(r) == 1 {
				for c, v := range inputMap[r] {
					if v && columnSum(c) == 8 {
						alreadyFound['c'-offset] = c
						break
					}
				}
				break
			}
		}

		// try to find the d
		// it's the one with 8 occurences that also appears in a 1-length token
		eliminate()
		for r := range inputMap {
			if rowSum(r) == 1 {
				for c, v := range inputMap[r] {
					if v && columnSum(c) == 7 {
						alreadyFound['d'-offset] = c
						break
					}
				}
			}
		}

		// 'g' and 'a' last
		eliminate()
		for i := range inputMap[0] {
			switch columnSum(i) {
			case 7:
				alreadyFound['g'-offset] = i
			case 8:
				alreadyFound['a'-offset] = i
			}
		}

		// now it's time to decode the other values
		reverseMapping := map[byte]byte{}
		for i, c := range alreadyFound {
			reverseMapping[byte(c+offset)] = i + offset
		}

		decode := func(token string) int {
			newToken := []byte{}
			for _, c := range token {
				newToken = append(newToken, reverseMapping[byte(c)])
			}
			sort.Slice(newToken, func(i, j int) bool { return newToken[i] < newToken[j] })

			for i := range numbers {
				if string(newToken) == numbers[i] {
					return i
				}
			}
			return -1
		}

		number := 0
		for _, token := range strings.Fields(parts[1]) {
			number = 10*number + decode(token)
		}
		result += number
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
