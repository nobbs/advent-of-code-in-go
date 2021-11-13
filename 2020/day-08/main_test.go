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
			input: util.PrepareExampleInput(`nop +0
			acc +1
			jmp +4
			acc +3
			jmp -3
			acc -99
			acc +1
			jmp -4
			acc +6`),
		},
		{
			desc:  "actual",
			want:  1782,
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
			want: 8,
			input: util.PrepareExampleInput(`nop +0
			acc +1
			jmp +4
			acc +3
			jmp -3
			acc -99
			acc +1
			jmp -4
			acc +6`),
		},
		{
			desc:  "actual",
			want:  797,
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
