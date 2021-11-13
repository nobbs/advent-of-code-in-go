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
			want: 165,
			input: util.PrepareExampleInput(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
			mem[8] = 11
			mem[7] = 101
			mem[8] = 0`),
		},
		{
			desc:  "actual",
			want:  4886706177792,
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
			want: 208,
			input: util.PrepareExampleInput(`mask = 000000000000000000000000000000X1001X
			mem[42] = 100
			mask = 00000000000000000000000000000000X0XX
			mem[26] = 1`),
		},
		{
			desc:  "actual",
			want:  3348493585827,
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
