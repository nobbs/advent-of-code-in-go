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
			want: 7 * 5,
			input: util.PrepareExampleInput(`16
			10
			15
			5
			1
			11
			7
			19
			6
			12
			4`),
		},
		{
			desc: "example",
			want: 22 * 10,
			input: util.PrepareExampleInput(`28
			33
			18
			42
			31
			14
			46
			20
			48
			47
			24
			23
			49
			45
			19
			38
			39
			11
			1
			32
			25
			35
			8
			17
			7
			9
			4
			2
			34
			10
			3`),
		},
		{
			desc:  "actual",
			want:  1914,
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
			input: util.PrepareExampleInput(`16
			10
			15
			5
			1
			11
			7
			19
			6
			12
			4`),
		},
		{
			desc: "example",
			want: 19208,
			input: util.PrepareExampleInput(`28
			33
			18
			42
			31
			14
			46
			20
			48
			47
			24
			23
			49
			45
			19
			38
			39
			11
			1
			32
			25
			35
			8
			17
			7
			9
			4
			2
			34
			10
			3`),
		},
		{
			desc:  "actual",
			want:  9256148959232,
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
