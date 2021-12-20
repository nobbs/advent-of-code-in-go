package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Rules map[string]string

func parseInput(lines []string) (map[string]int, Rules) {
	pairs := map[string]int{}
	template := lines[0]
	for i := 0; i < len(template)-1; i++ {
		p := template[i : i+2]
		pairs[p]++
	}

	rules := Rules{}
	r := regexp.MustCompile(`^(\w\w) -> (\w)$`)
	for _, line := range lines[2:] {
		groups := r.FindStringSubmatch(line)
		rules[groups[1]] = groups[2]
	}

	return pairs, rules
}

func doStep(pairs map[string]int, r Rules) map[string]int {
	next := map[string]int{}

	for k, v := range pairs {
		// check if a rule exists for this pair, if so, add new pairs based on this rule
		if l, ok := r[k]; ok {
			next[string(k[0])+l] += v
			next[l+string(k[1])] += v
		}
	}

	return next
}

func solve(iterations int, lines []string) int {
	pairs, rules := parseInput(lines)

	// perform iterations
	for i := 0; i < iterations; i++ {
		pairs = doStep(pairs, rules)
	}

	// count characters, all of them are twice in there
	// (first and second char. in a pair)
	counts := map[string]int{}
	for k, v := range pairs {
		counts[string(k[0])] += v
		counts[string(k[1])] += v
	}

	max, min := 0, math.MaxInt64
	for _, count := range counts {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	// return the final result
	return (max - min + 1) / 2
}

func partOne(lines []string) int {
	return solve(10, lines)
}

func partTwo(lines []string) int {
	return solve(40, lines)
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
