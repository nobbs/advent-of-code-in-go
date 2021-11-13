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
			want: 4,
			input: util.PrepareExampleInput(`light red bags contain 1 bright white bag, 2 muted yellow bags.
				dark orange bags contain 3 bright white bags, 4 muted yellow bags.
				bright white bags contain 1 shiny gold bag.
				muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
				shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
				dark olive bags contain 3 faded blue bags, 4 dotted black bags.
				vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
				faded blue bags contain no other bags.
				dotted black bags contain no other bags.`),
		},
		{
			desc:  "actual",
			want:  287,
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
			want: 32,
			input: util.PrepareExampleInput(`light red bags contain 1 bright white bag, 2 muted yellow bags.
				dark orange bags contain 3 bright white bags, 4 muted yellow bags.
				bright white bags contain 1 shiny gold bag.
				muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
				shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
				dark olive bags contain 3 faded blue bags, 4 dotted black bags.
				vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
				faded blue bags contain no other bags.
				dotted black bags contain no other bags.`),
		},
		{
			desc: "example",
			want: 126,
			input: util.PrepareExampleInput(`shiny gold bags contain 2 dark red bags.
				dark red bags contain 2 dark orange bags.
				dark orange bags contain 2 dark yellow bags.
				dark yellow bags contain 2 dark green bags.
				dark green bags contain 2 dark blue bags.
				dark blue bags contain 2 dark violet bags.
				dark violet bags contain no other bags.`),
		},
		{
			desc:  "actual",
			want:  48160,
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
