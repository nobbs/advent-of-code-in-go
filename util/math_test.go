package util

import (
	"testing"
)

func TestExtendedGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name:  "example",
			args:  args{1432, 123211},
			want:  1,
			want1: -22973,
			want2: 267,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ExtendedGCD(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("ExtendedGCD() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtendedGCD() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ExtendedGCD() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestInverseModulo(t *testing.T) {
	type args struct {
		a int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{123, 4567},
			want: 854,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InverseModulo(tt.args.a, tt.args.n); got != tt.want {
				t.Errorf("InverseModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChineseRemainderTheorem(t *testing.T) {
	type args struct {
		n []int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{n: []int{3, 4, 5}, a: []int{0, 3, 4}},
			want: 39,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChineseRemainderTheorem(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("ChineseRemainderTheorem() = %v, want %v", got, tt.want)
			}
		})
	}
}
