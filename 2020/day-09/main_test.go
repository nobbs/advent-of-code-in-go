package main

import (
	"testing"

	"github.com/nobbs/advent-of-code-in-go/util"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		desc           string
		want           int
		input          []string
		preambleLength int
	}{
		{
			desc: "example",
			want: 65,
			input: util.PrepareExampleInput(`20
			1
			2
			3
			4
			5
			6
			7
			8
			9
			10
			11
			12
			13
			14
			15
			16
			17
			18
			19
			21
			22
			23
			24
			25
			45
			65
			64`),
			preambleLength: 25,
		},
		{
			desc: "example",
			want: 127,
			input: util.PrepareExampleInput(`35
			20
			15
			25
			47
			40
			62
			55
			65
			95
			102
			117
			150
			182
			127
			219
			299
			277
			309
			576`),
			preambleLength: 5,
		},
		{
			desc:           "actual",
			want:           14144619,
			input:          util.ReadInputFile("./input.txt"),
			preambleLength: 25,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := partOne(tC.input, tC.preambleLength)

			if got != tC.want {
				t.Errorf("partOne() = %v, want %v", got, tC.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		desc           string
		want           int
		input          []string
		preambleLength int
	}{
		{
			desc: "example",
			want: 62,
			input: util.PrepareExampleInput(`35
			20
			15
			25
			47
			40
			62
			55
			65
			95
			102
			117
			150
			182
			127
			219
			299
			277
			309
			576`),
			preambleLength: 5,
		},
		{
			desc:           "actual",
			want:           1766397,
			input:          util.ReadInputFile("./input.txt"),
			preambleLength: 25,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := partTwo(tC.input, tC.preambleLength)

			if got != tC.want {
				t.Errorf("partTwo() = %v, want %v", got, tC.want)
			}
		})
	}
}
