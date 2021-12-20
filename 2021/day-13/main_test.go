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
			want: 17,
			input: util.PrepareExampleInput(`6,10
			0,14
			9,10
			0,3
			10,4
			4,11
			6,0
			6,12
			4,1
			0,13
			10,12
			3,4
			3,0
			8,4
			1,10
			2,14
			8,10
			9,0
			
			fold along y=7
			fold along x=5`),
		},
		{
			desc:  "actual",
			want:  807,
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
		want  string
		input []string
	}{
		{
			desc: "example",
			want: `#####
#...#
#...#
#...#
#####
.....
.....
`,
			input: util.PrepareExampleInput(`6,10
			0,14
			9,10
			0,3
			10,4
			4,11
			6,0
			6,12
			4,1
			0,13
			10,12
			3,4
			3,0
			8,4
			1,10
			2,14
			8,10
			9,0
			
			fold along y=7
			fold along x=5`),
		},
		{
			desc: "actual",
			want: `#.....##..#..#.####..##..#..#.####...##.
#....#..#.#..#.#....#..#.#..#.#.......#.
#....#....####.###..#....#..#.###.....#.
#....#.##.#..#.#....#.##.#..#.#.......#.
#....#..#.#..#.#....#..#.#..#.#....#..#.
####..###.#..#.####..###..##..####..##..
`,
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
