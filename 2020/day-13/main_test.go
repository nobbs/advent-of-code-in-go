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
			want: 295,
			input: util.PrepareExampleInput(`939
			7,13,x,x,59,x,31,19`),
		},
		{
			desc:  "actual",
			want:  2305,
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
			want: 1068781,
			input: util.PrepareExampleInput(`939
			7,13,x,x,59,x,31,19`),
		},
		{
			desc: "example",
			want: 3417,
			input: util.PrepareExampleInput(`939
			17,x,13,19`),
		},
		{
			desc: "example",
			want: 754018,
			input: util.PrepareExampleInput(`939
			67,7,59,61`),
		},
		{
			desc: "example",
			want: 779210,
			input: util.PrepareExampleInput(`939
			67,x,7,59,61`),
		},
		{
			desc: "example",
			want: 1261476,
			input: util.PrepareExampleInput(`939
			67,7,x,59,61`),
		},
		{
			desc: "example",
			want: 1202161486,
			input: util.PrepareExampleInput(`939
			1789,37,47,1889`),
		},
		{
			desc:  "actual",
			want:  552612234243498,
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
