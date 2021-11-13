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
			want: 7,
			input: util.PrepareExampleInput(`..##.......
			#...#...#..
			.#....#..#.
			..#.#...#.#
			.#...##..#.
			..#.##.....
			.#.#.#....#
			.#........#
			#.##...#...
			#...##....#
			.#..#...#.#`),
		},
		{
			desc:  "actual",
			want:  148,
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
			want: 336,
			input: util.PrepareExampleInput(`..##.......
			#...#...#..
			.#....#..#.
			..#.#...#.#
			.#...##..#.
			..#.##.....
			.#.#.#....#
			.#........#
			#.##...#...
			#...##....#
			.#..#...#.#`),
		},
		{
			desc:  "actual",
			want:  727923200,
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
