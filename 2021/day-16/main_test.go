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
			desc:  "example",
			want:  16,
			input: util.PrepareExampleInput(`8A004A801A8002F478`),
		},
		{
			desc:  "example",
			want:  12,
			input: util.PrepareExampleInput(`620080001611562C8802118E34`),
		},
		{
			desc:  "example",
			want:  23,
			input: util.PrepareExampleInput(`C0015000016115A2E0802F182340`),
		},
		{
			desc:  "example",
			want:  31,
			input: util.PrepareExampleInput(`A0016C880162017C3686B18A3D4780`),
		},
		{
			desc:  "actual",
			want:  1007,
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
			desc:  "example",
			want:  3,
			input: util.PrepareExampleInput(`C200B40A82`),
		},
		{
			desc:  "example",
			want:  54,
			input: util.PrepareExampleInput(`04005AC33890`),
		},
		{
			desc:  "example",
			want:  7,
			input: util.PrepareExampleInput(`880086C3E88112`),
		},
		{
			desc:  "example",
			want:  9,
			input: util.PrepareExampleInput(`CE00C43D881120`),
		},
		{
			desc:  "example",
			want:  1,
			input: util.PrepareExampleInput(`D8005AC2A8F0`),
		},
		{
			desc:  "example",
			want:  0,
			input: util.PrepareExampleInput(`F600BC2D8F`),
		},
		{
			desc:  "example",
			want:  0,
			input: util.PrepareExampleInput(`9C005AC2F8F0`),
		},
		{
			desc:  "example",
			want:  1,
			input: util.PrepareExampleInput(`9C0141080250320F1802104A08`),
		},
		{
			desc:  "actual",
			want:  834151779165,
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

func Test_binToUint64(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "example",
			args: args{b: "011111100101"},
			want: 2021,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binToUint64(tt.args.b); got != tt.want {
				t.Errorf("binToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexToBin(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{h: "D2FE28"},
			want: "110100101111111000101000",
		},
		{
			name: "example",
			args: args{h: "38006F45291200"},
			want: "00111000000000000110111101000101001010010001001000000000",
		},
		{
			name: "example",
			args: args{h: "EE00D40C823060"},
			want: "11101110000000001101010000001100100000100011000001100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexToBin(tt.args.h); got != tt.want {
				t.Errorf("hexToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
