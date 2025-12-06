package main

import (
	"testing"
)

func TestParseOperations(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name           string
		args           args
		wantNumbers    []int
		wantOperations []string
	}{
		{
			name: "Test 0",
			args: args{
				line: "*   +   *   +  ",
			},
			wantNumbers:    []int{4, 4, 4, 3},
			wantOperations: []string{"*", "+", "*", "+"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOperations, gotNumbers := parseOperations(tt.args.line)
			if !equalIntSlices(gotNumbers, tt.wantNumbers) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, gotNumbers, tt.wantNumbers)
			}
			if !equalStringSlices(gotOperations, tt.wantOperations) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, gotOperations, tt.wantOperations)
			}
		})
	}
}

func TestParseNumbersLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name           string
		args           args
		wantNumbers    []int
		wantOperations []string
	}{
		{
			name: "Test 0",
			args: args{
				line: "123 328  51 64 ",
			},
			wantNumbers:    []int{123, 328, 51, 64},
			wantOperations: []string{},
		},
		{
			name: "Test 1",
			args: args{
				line: " 45 64  387 23 ",
			},
			wantNumbers:    []int{45, 64, 387, 23},
			wantOperations: []string{},
		},
		{
			name: "Test 2",
			args: args{
				line: "  6 98  215 314",
			},
			wantNumbers:    []int{6, 98, 215, 314},
			wantOperations: []string{},
		},
		{
			name: "Test 3",
			args: args{
				line: "*   +   *   +  ",
			},
			wantNumbers:    []int{},
			wantOperations: []string{"*", "+", "*", "+"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumbers, gotOperations := parseLine(tt.args.line)
			if !equalIntSlices(gotNumbers, tt.wantNumbers) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, gotNumbers, tt.wantNumbers)
			}
			if !equalStringSlices(gotOperations, tt.wantOperations) {
				t.Errorf("%s parseLine(%v) = %v, want %v", tt.name, tt.args.line, gotOperations, tt.wantOperations)
			}
		})
	}
}

var textGrid = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

// 123 328  51 64
//  45 64  387 23
//   6 98  215 314
// *   +   *   +

func TestParseLines(t *testing.T) {
	wantNumbers := [][]int{
		{1, 24, 356},
		{369, 248, 8},
		{32, 581, 175},
		{623, 431, 4},
	}
	wantOperations := []string{
		"*", "+", "*", "+",
	}

	lines := splitLines(textGrid)
	gotNumbers, gotOperations := parseLines(lines)

	if len(gotNumbers) != len(wantNumbers) {
		t.Fatalf("parseLines() numbers length = %v, want %v", len(gotNumbers), len(wantNumbers))
	}

	for i := range wantNumbers {
		if !equalIntSlices(gotNumbers[i], wantNumbers[i]) {
			t.Errorf("parseLines() line %d = %v, want %v", i, gotNumbers[i], wantNumbers[i])
		}
	}

	for i := range wantOperations {
		if gotOperations[i] != wantOperations[i] {
			t.Errorf("parseLines() operation %d = %v, want %v", i, gotOperations[i], wantOperations[i])
		}
	}
}

func TestReverseIntMatrices(t *testing.T) {
	numbers := [][]int{
		{123, 328, 51, 64},
		{45, 64, 387, 23},
		{6, 98, 215, 314},
	}
	wantNumbers := [][]int{
		{123, 45, 6},
		{328, 64, 98},
		{51, 387, 215},
		{64, 23, 314},
	}

	gotNumbers := reverseIntMatrices(numbers)

	if len(gotNumbers) != len(wantNumbers) {
		t.Fatalf("reverseIntMatrices() numbers length = %v, want %v", len(gotNumbers), len(wantNumbers))
	}

	for i := range wantNumbers {
		if !equalIntSlices(gotNumbers[i], wantNumbers[i]) {
			t.Errorf("parseLines() line %d = %v, want %v", i, gotNumbers[i], wantNumbers[i])
		}
	}
}

func equalIntSlices(a, b []int) bool {
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

func equalStringSlices(a, b []string) bool {
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

func TestProcessLines(t *testing.T) {
	want := 3263827

	lines := splitLines(textGrid)
	numbers, operations := parseLines(lines)
	got := processLines(numbers, operations)
	if got != want {
		t.Errorf("processLines() = %v, want %v", got, want)
	}
}
