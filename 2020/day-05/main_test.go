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
			desc:  "example",
			want:  357,
			input: util.PrepareExampleInput(`FBFBBFFRLR`),
		},
		{
			desc:  "example",
			want:  567,
			input: util.PrepareExampleInput(`BFFFBBFRRR`),
		},
		{
			desc:  "example",
			want:  119,
			input: util.PrepareExampleInput(`FFFBBBFRRR`),
		},
		{
			desc:  "example",
			want:  820,
			input: util.PrepareExampleInput(`BBFFBBFRLL`),
		},
		{
			desc:  "actual",
			want:  935,
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
			desc:  "actual",
			want:  743,
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
