package main

import (
	"fmt"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	gamma, epsilon := 0, 0
	I, J := len(lines[0]), len(lines)

	for i := 0; i < I; i++ {
		ones := 0
		for j := 0; j < J; j++ {
			if lines[j][i] == '1' {
				ones++
			}
		}

		switch {
		case ones > J/2:
			gamma += 1 << (I - (i + 1))
		default:
			epsilon += 1 << (I - (i + 1))
		}
	}

	return gamma * epsilon
}

func partTwo(lines []string) int {
	filterOxygen, filterCO2 := map[int]bool{}, map[int]bool{}
	numChars, numLines := len(lines[0]), len(lines)

	for j := 0; j < numLines; j++ {
		filterOxygen[j] = true
		filterCO2[j] = true
	}

	checkFiltered := func(filter, ones, zeros map[int]bool, i int) {
		for j := 0; j < numLines; j++ {
			if _, ok := filter[j]; ok {
				switch {
				case lines[j][i] == '1':
					ones[j] = true
				case lines[j][i] == '0':
					zeros[j] = true
				}
			}
		}
	}

	removeFromFilter := func(filter, remove map[int]bool) {
		for k := range remove {
			delete(filter, k)
		}
	}

	getFinalNumber := func(lines []string, filter map[int]bool) int {
		bitstring := ""
		for k := range filter {
			bitstring = lines[k]
		}

		result := 0
		for i, c := range bitstring {
			if c == '1' {
				result += 1 << (len(bitstring) - (i + 1))
			}
		}

		return result
	}

	oxygen := 0
	for i := 0; i < numChars; i++ {
		ones, zeros := map[int]bool{}, map[int]bool{}
		checkFiltered(filterOxygen, ones, zeros, i)

		if len(ones) >= len(zeros) {
			removeFromFilter(filterOxygen, zeros)
		} else {
			removeFromFilter(filterOxygen, ones)
		}

		if len(filterOxygen) == 1 {
			oxygen = getFinalNumber(lines, filterOxygen)
		}
	}

	co2 := 0
	for i := 0; i < numChars; i++ {
		ones, zeros := map[int]bool{}, map[int]bool{}
		checkFiltered(filterCO2, ones, zeros, i)

		if len(ones) >= len(zeros) {
			removeFromFilter(filterCO2, ones)
		} else {
			removeFromFilter(filterCO2, zeros)
		}

		if len(filterCO2) == 1 {
			co2 = getFinalNumber(lines, filterCO2)
		}
	}

	return oxygen * co2
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
