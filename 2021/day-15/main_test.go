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
			want: 40,
			input: util.PrepareExampleInput(`1163751742
				1381373672
				2136511328
				3694931569
				7463417111
				1319128137
				1359912421
				3125421639
				1293138521
				2311944581`),
		},
		{
			desc:  "actual",
			want:  685,
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
			want: 315,
			input: util.PrepareExampleInput(`1163751742
				1381373672
				2136511328
				3694931569
				7463417111
				1319128137
				1359912421
				3125421639
				1293138521
				2311944581`),
		},
		{
			desc:  "actual",
			want:  2995,
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
