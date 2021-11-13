package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func partOne(lines []string) int {
	arrival := util.ParseInt(lines[0])
	buses := []int{}

	for _, x := range strings.Split(lines[1], ",") {
		n, err := strconv.Atoi(x)

		if err == nil {
			buses = append(buses, n)
		}
	}

	bestBus := 0
	bestBusDeparture := 1 << 32
	for _, n := range buses {
		for v := n; v <= arrival+n*2; v += n {
			if arrival <= v && v <= bestBusDeparture {
				bestBus = n
				bestBusDeparture = v
				break
			}
		}
	}

	return (bestBusDeparture - arrival) * bestBus
}

type Bus struct {
	id, offset int
}

func partTwo(lines []string) int {
	buses := []Bus{}

	offset := 0
	for _, x := range strings.Split(lines[1], ",") {
		n, err := strconv.Atoi(x)

		if err == nil {
			buses = append(buses, Bus{id: n, offset: offset})
		}
		offset++
	}

	num, rem := make([]int, len(buses)), make([]int, len(buses))
	for i := range buses {
		num[i] = buses[i].id
		rem[i] = (buses[0].offset - buses[i].offset + buses[i].id) % buses[i].id
	}

	return util.ChineseRemainderTheorem(num, rem)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
