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
			want: 24000,
			input: util.PrepareExampleInput(`1000
			2000
			3000
			
			4000
			
			5000
			6000
			
			7000
			8000
			9000
			
			10000`),
		},
		{
			desc:  "actual",
			want:  70374,
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
			want: 45000,
			input: util.PrepareExampleInput(`1000
			2000
			3000
			
			4000
			
			5000
			6000
			
			7000
			8000
			9000
			
			10000`),
		},
		{
			desc:  "actual",
			want:  204610,
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
