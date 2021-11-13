package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type PasswordSet struct {
	from, to int
	char     string
	input    string
}

func partOne(lines []string) int {
	passwords := parsePasswords(lines)

	isValid := func(p PasswordSet) bool {
		n := strings.Count(p.input, p.char)
		return p.from <= n && n <= p.to
	}

	validPasswords := 0
	for _, problem := range passwords {
		if isValid(problem) {
			validPasswords++
		}
	}

	return validPasswords
}

func partTwo(lines []string) int {
	passwords := parsePasswords(lines)

	isValid := func(p PasswordSet) bool {
		first := string(p.input[p.from-1])
		second := string(p.input[p.to-1])

		if (first == p.char && second != p.char) || (first != p.char && second == p.char) {
			return true
		}
		return false
	}

	validPasswords := 0
	for _, problem := range passwords {
		if isValid(problem) {
			validPasswords++
		}
	}

	return validPasswords
}

func parsePasswords(lines []string) []PasswordSet {
	passwords := []PasswordSet{}

	for _, line := range lines {
		password := parseInputLine(line)
		passwords = append(passwords, password)
	}

	return passwords
}

func parseInputLine(line string) PasswordSet {
	r := regexp.MustCompile(`^(\d+)-(\d+)\s(\w):\s(\w*)$`)
	groups := r.FindStringSubmatch(line)

	if len(groups) != 5 {
		panic("input line does not match required layout")
	}

	from := util.ParseInt(groups[1])
	to := util.ParseInt(groups[2])

	return PasswordSet{from: from, to: to, char: groups[3], input: groups[4]}
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
