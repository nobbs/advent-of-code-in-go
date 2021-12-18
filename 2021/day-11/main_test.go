package main

import (
	"testing"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		desc  string
		want  int
		input []string
	}{
		{
			desc: "example",
			want: 1656,
			input: util.PrepareExampleInput(`5483143223
			2745854711
			5264556173
			6141336146
			6357385478
			4167524645
			2176841721
			6882881134
			4846848554
			5283751526`),
		},
		{
			desc:  "actual",
			want:  1637,
			input: util.ReadInputFile("./input.txt"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := partOne(tC.input)

			if got != tC.want {
				t.Errorf("partOne() = %v, want %v", got, tC.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		desc  string
		want  int
		input []string
	}{
		{
			desc: "example",
			want: 195,
			input: util.PrepareExampleInput(`5483143223
			2745854711
			5264556173
			6141336146
			6357385478
			4167524645
			2176841721
			6882881134
			4846848554
			5283751526`),
		},
		{
			desc:  "actual",
			want:  242,
			input: util.ReadInputFile("./input.txt"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := partTwo(tC.input)

			if got != tC.want {
				t.Errorf("partTwo() = %v, want %v", got, tC.want)
			}
		})
	}
}
