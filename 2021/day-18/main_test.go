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
			want: 129,
			input: util.PrepareExampleInput(`[9,1]
			[1,9]`),
		},
		{
			desc: "example",
			want: 3488,
			input: util.PrepareExampleInput(`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
			[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
			[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
			[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
			[7,[5,[[3,8],[1,4]]]]
			[[2,[2,2]],[8,[8,1]]]
			[2,9]
			[1,[[[9,3],9],[[9,0],[0,7]]]]
			[[[5,[7,4]],7],1]
			[[[[4,2],2],6],[8,7]]`),
		},
		{
			desc: "example",
			want: 4140,
			input: util.PrepareExampleInput(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
			[[[5,[2,8]],4],[5,[[9,9],0]]]
			[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
			[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
			[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
			[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
			[[[[5,4],[7,7]],8],[[8,3],8]]
			[[9,3],[[9,9],[6,[4,9]]]]
			[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
			[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`),
		},
		{
			desc:  "actual",
			want:  4235,
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
			want: 3993,
			input: util.PrepareExampleInput(`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
			[[[5,[2,8]],4],[5,[[9,9],0]]]
			[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
			[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
			[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
			[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
			[[[[5,4],[7,7]],8],[[8,3],8]]
			[[9,3],[[9,9],[6,[4,9]]]]
			[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
			[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`),
		},
		{
			desc:  "actual",
			want:  4659,
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

func TestNode_preorderTraversal(t *testing.T) {
	testCases := []struct {
		desc  string
		input []string
	}{
		{
			desc:  "example",
			input: util.PrepareExampleInput(`[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]`),
		},
		{
			desc: "example",
			input: util.PrepareExampleInput(`[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
			`),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			root := parseNode(tC.input[0], nil)

			testFn := func(n *Node) {
				// check parents of children
				if n.left != nil {
					if n != n.left.parent {
						t.Error("Child-Parent nodes don't match!")
					}
				}
				if n.right != nil {
					if n != n.right.parent {
						t.Error("Child-Parent nodes don't match!")
					}
				}
			}

			root.inorderTraversal(testFn)
			root.preorderTraversal(testFn)
			root.postorderTraversal(testFn)
		})
	}
}

func TestNode_magnitude(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantMag int
	}{
		{
			input:   "[[1,2],[[3,4],5]]",
			wantMag: 143,
		},
		{
			input:   "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
			wantMag: 1384,
		},
		{
			input:   "[[[[1,1],[2,2]],[3,3]],[4,4]]",
			wantMag: 445,
		},
		{
			input:   "[[[[3,0],[5,3]],[4,4]],[5,5]]",
			wantMag: 791,
		},
		{
			input:   "[[[[5,0],[7,4]],[5,5]],[6,6]]",
			wantMag: 1137,
		},
		{
			input:   "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
			wantMag: 3488,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := parseNode(tt.input, nil)
			if gotMag := n.magnitude(); gotMag != tt.wantMag {
				t.Errorf("Node.magnitude() = %v, want %v", gotMag, tt.wantMag)
			}
		})
	}
}
