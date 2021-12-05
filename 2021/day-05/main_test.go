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
			want: 5,
			input: util.PrepareExampleInput(`0,9 -> 5,9
			8,0 -> 0,8
			9,4 -> 3,4
			2,2 -> 2,1
			7,0 -> 7,4
			6,4 -> 2,0
			0,9 -> 2,9
			3,4 -> 1,4
			0,0 -> 8,8
			5,5 -> 8,2`),
		},
		{
			desc:  "actual",
			want:  5280,
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
			want: 12,
			input: util.PrepareExampleInput(`0,9 -> 5,9
			8,0 -> 0,8
			9,4 -> 3,4
			2,2 -> 2,1
			7,0 -> 7,4
			6,4 -> 2,0
			0,9 -> 2,9
			3,4 -> 1,4
			0,0 -> 8,8
			5,5 -> 8,2`)},
		{
			desc:  "actual",
			want:  16716,
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
