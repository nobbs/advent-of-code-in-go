package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func UintToBinary(x uint64) string {
	return fmt.Sprintf("%036s", strconv.FormatUint(x, 2))
}

func parseMask(line string) (uint64, uint64, []uint) {
	r := regexp.MustCompile(`^mask = (\w+)$`)
	mask := r.FindStringSubmatch(line)[1]

	var orMask uint64 = 0
	var andMask uint64 = 0
	xMask := []uint{}
	for i := 0; i < len(mask); i++ {
		c := mask[i]

		switch c {
		case '1':
			orMask += 1 << (len(mask) - i - 1)
		case '0':
			andMask += 1 << (len(mask) - i - 1)
		case 'X':
			xMask = append(xMask, uint(len(mask)-i-1))
		}
	}

	return orMask, andMask, xMask
}

func parseMemory(line string) (uint64, uint64) {
	r := regexp.MustCompile(`^mem\[(\d+)\] = (\w+)$`)
	groups := r.FindStringSubmatch(line)

	index, _ := strconv.ParseUint(groups[1], 10, 64)
	value, _ := strconv.ParseUint(groups[2], 10, 64)

	return index, value
}

func partOne(lines []string) int {
	var allOnes uint64 = 1<<36 - 1
	var maskO uint64
	var maskA uint64

	memory := map[uint64]uint64{}

	for _, line := range lines {
		if strings.Contains(line, "mask = ") {
			maskO, maskA, _ = parseMask(line)
		} else if strings.Contains(line, "mem[") {
			index, value := parseMemory(line)
			memory[index] = maskO | (value & (allOnes ^ maskA))
		}
	}

	var sum uint64 = 0
	for _, v := range memory {
		sum += v
	}

	return int(sum)
}

func generatePossibilities(n uint64, mask []uint) []uint64 {
	generate := func(m uint64, i int) (uint64, uint64) {
		a := m & ((1<<36 - 1) - (1 << mask[i]))
		b := m | (1 << mask[i])
		return a, b
	}

	result := []uint64{n}
	for i := range mask {
		nextResult := []uint64{}
		for _, x := range result {
			a, b := generate(x, i)
			nextResult = append(nextResult, a, b)
		}
		result = nextResult
	}

	return result
}

func partTwo(lines []string) int {
	var maskO uint64
	var maskX []uint

	memory := map[uint64]uint64{}

	for _, line := range lines {
		if strings.Contains(line, "mask = ") {
			maskO, _, maskX = parseMask(line)
		} else if strings.Contains(line, "mem[") {
			index, value := parseMemory(line)
			for _, address := range generatePossibilities(index|maskO, maskX) {
				memory[address] = value
			}
		}
	}

	var sum uint64 = 0
	for _, v := range memory {
		sum += v
	}

	return int(sum)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
