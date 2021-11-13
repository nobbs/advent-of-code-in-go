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
			want: 37,
			input: util.PrepareExampleInput(`L.LL.LL.LL
			LLLLLLL.LL
			L.L.L..L..
			LLLL.LL.LL
			L.LL.LL.LL
			L.LLLLL.LL
			..L.L.....
			LLLLLLLLLL
			L.LLLLLL.L
			L.LLLLL.LL`),
		},
		{
			desc:  "actual",
			want:  2270,
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
			want: 26,
			input: util.PrepareExampleInput(`L.LL.LL.LL
			LLLLLLL.LL
			L.L.L..L..
			LLLL.LL.LL
			L.LL.LL.LL
			L.LLLLL.LL
			..L.L.....
			LLLLLLLLLL
			L.LLLLLL.L
			L.LLLLL.LL`),
		},
		{
			desc:  "actual",
			want:  2042,
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
