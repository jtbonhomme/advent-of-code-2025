package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 0",
			args: args{
				line: "1122",
			},
			want: []int{1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseLine(tt.args.line)
			if !equalSlices(got, tt.want) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetMaxJoltage(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				line: "987654321111111",
			},
			want: 987654321111,
		},
		{
			name: "Test 2",
			args: args{
				line: "811111111111119",
			},
			want: 811111111119,
		},
		{
			name: "Test 3",
			args: args{
				line: "234234234234278",
			},
			want: 434234234278,
		},
		{
			name: "Test 4",
			args: args{
				line: "818181911112111",
			},
			want: 888911112111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank := parseLine(tt.args.line)
			got := getMaxJoltage(bank)
			if got != tt.want {
				t.Errorf("%s getMaxJoltage(%v) = %v, want %v", tt.name, bank, got, tt.want)
			}
		})
	}
}
