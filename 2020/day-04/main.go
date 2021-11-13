package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Passport struct {
	string
}

func partOne(lines []string) int {
	passports := parsePassports(lines)

	checkValidity := func(passport Passport) bool {
		metadata := []string{
			"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
		}

		for _, data := range metadata {
			if !strings.Contains(passport.string, data) {
				return false
			}
		}

		return true
	}

	valid := 0
	for _, p := range passports {
		if checkValidity(p) {
			valid++
		}
	}

	return valid
}

func partTwo(lines []string) int {
	passports := parsePassports(lines)

	isYearValid := func(p *Passport, field string, from, to int) bool {
		r := regexp.MustCompile(field + `:(\d{4})`)
		matches := r.FindStringSubmatch(p.string)

		if len(matches) != 2 {
			return false
		}

		year := util.ParseInt(matches[1])

		return from <= year && year <= to
	}

	isHeightValid := func(p *Passport) bool {
		rMetric := regexp.MustCompile(`hgt:(\d{3})cm`).FindStringSubmatch(p.string)
		rImperial := regexp.MustCompile(`hgt:(\d{2})in`).FindStringSubmatch(p.string)

		switch {
		case len(rMetric) == 2:
			height := util.ParseInt(rMetric[1])
			return 150 <= height && height <= 193
		case len(rImperial) == 2:
			height := util.ParseInt(rImperial[1])
			return 59 <= height && height <= 76
		default:
			return false
		}
	}

	isHairColorValid := func(p *Passport) bool {
		color := regexp.MustCompile(`hcl:(#[0-9a-f]{6})\b`).FindStringSubmatch(p.string)
		return len(color) == 2
	}

	isEyeColorValid := func(p *Passport) bool {
		color := regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)\b`).FindStringSubmatch(p.string)
		return len(color) == 2
	}

	isPIDValid := func(p *Passport) bool {
		id := regexp.MustCompile(`pid:(\d{9})\b`).FindStringSubmatch(p.string)
		return len(id) == 2
	}

	checkValidity := func(p *Passport) bool {
		metadata := []string{
			"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid",
		}

		for _, data := range metadata {
			if !strings.Contains(p.string, data) {
				return false
			}
		}

		switch {
		case !isYearValid(p, "byr", 1920, 2002):
			return false
		case !isYearValid(p, "iyr", 2010, 2020):
			return false
		case !isYearValid(p, "eyr", 2020, 2030):
			return false
		case !isHeightValid(p):
			return false
		case !isHairColorValid(p):
			return false
		case !isEyeColorValid(p):
			return false
		case !isPIDValid(p):
			return false
		default:
			return true
		}
	}

	valid := 0
	for _, p := range passports {
		if checkValidity(&p) {
			valid++
		}
	}

	return valid
}

func parsePassports(lines []string) []Passport {
	passports := []Passport{}

	var currentPassport string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			passports = append(passports, Passport{string: currentPassport})
			currentPassport = ""
			continue
		} else {
			currentPassport = strings.Join([]string{currentPassport, line}, " ")
		}
	}

	return passports
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
