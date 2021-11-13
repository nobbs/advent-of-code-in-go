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
			want:  436,
			input: util.PrepareExampleInput(`0,3,6`),
		},
		{
			desc:  "example",
			want:  1,
			input: util.PrepareExampleInput(`1,3,2`),
		},
		{
			desc:  "example",
			want:  10,
			input: util.PrepareExampleInput(`2,1,3`),
		},
		{
			desc:  "example",
			want:  27,
			input: util.PrepareExampleInput(`1,2,3`),
		},
		{
			desc:  "example",
			want:  78,
			input: util.PrepareExampleInput(`2,3,1`),
		},
		{
			desc:  "example",
			want:  438,
			input: util.PrepareExampleInput(`3,2,1`),
		},
		{
			desc:  "example",
			want:  1836,
			input: util.PrepareExampleInput(`3,1,2`),
		},
		{
			desc:  "actual",
			want:  639,
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
			desc:  "example",
			want:  175594,
			input: util.PrepareExampleInput(`0,3,6`),
		},
		{
			desc:  "example",
			want:  2578,
			input: util.PrepareExampleInput(`1,3,2`),
		},
		{
			desc:  "example",
			want:  3544142,
			input: util.PrepareExampleInput(`2,1,3`),
		},
		{
			desc:  "example",
			want:  261214,
			input: util.PrepareExampleInput(`1,2,3`),
		},
		{
			desc:  "example",
			want:  6895259,
			input: util.PrepareExampleInput(`2,3,1`),
		},
		{
			desc:  "example",
			want:  18,
			input: util.PrepareExampleInput(`3,2,1`),
		},
		{
			desc:  "example",
			want:  362,
			input: util.PrepareExampleInput(`3,1,2`),
		},
		{
			desc:  "actual",
			want:  266,
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
