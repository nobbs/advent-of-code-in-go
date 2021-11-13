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
			want: 6,
			input: util.PrepareExampleInput(`abcx
			abcy
			abcz`),
		},
		{
			desc: "example",
			want: 11,
			input: util.PrepareExampleInput(`abc

			a
			b
			c
			
			ab
			ac
			
			a
			a
			a
			a
			
			b`),
		},
		{
			desc:  "actual",
			want:  6534,
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
			want: 6,
			input: util.PrepareExampleInput(`abc

			a
			b
			c
			
			ab
			ac
			
			a
			a
			a
			a
			
			b`),
		},
		{
			desc:  "actual",
			want:  3402,
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
