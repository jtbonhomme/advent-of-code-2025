package main

import (
	"testing"
)

func TestComputeNewDial(t *testing.T) {
	type args struct {
		currentDial int
		i           int
		dist        int
	}
	type result struct {
		newDial int
		zero    int
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "Test 0",
			args: args{
				currentDial: 0,
				i:           -1,
				dist:        252,
			},
			want: result{newDial: 48, zero: 2},
		},
		{
			name: "Test 1",
			args: args{
				currentDial: 50,
				i:           1,
				dist:        30,
			},
			want: result{newDial: 80, zero: 0},
		},
		{
			name: "Test 2",
			args: args{
				currentDial: 90,
				i:           1,
				dist:        20,
			},
			want: result{newDial: 10, zero: 1},
		},
		{
			name: "Test 3",
			args: args{
				currentDial: 10,
				i:           -1,
				dist:        20,
			},
			want: result{newDial: 90, zero: 1},
		},
		{
			name: "Test 4",
			args: args{
				currentDial: 99,
				i:           1,
				dist:        1,
			},
			want: result{newDial: 0, zero: 1},
		},
		{
			name: "Test 5",
			args: args{
				currentDial: 5,
				i:           -1,
				dist:        10,
			},
			want: result{newDial: 95, zero: 1},
		},
		{
			name: "Test 6",
			args: args{
				currentDial: 95,
				i:           1,
				dist:        5,
			},
			want: result{newDial: 0, zero: 1},
		},
		{
			name: "Test 7",
			args: args{
				currentDial: 95,
				i:           1,
				dist:        100,
			},
			want: result{newDial: 95, zero: 1},
		},
		{
			name: "Test 8",
			args: args{
				currentDial: 95,
				i:           -1,
				dist:        100,
			},
			want: result{newDial: 95, zero: 1},
		},
		{
			name: "Test 9",
			args: args{
				currentDial: 95,
				i:           1,
				dist:        200,
			},
			want: result{newDial: 95, zero: 2},
		},
		{
			name: "Test 10",
			args: args{
				currentDial: 95,
				i:           -1,
				dist:        200,
			},
			want: result{newDial: 95, zero: 2},
		},
		{
			name: "Test 11",
			args: args{
				currentDial: 50,
				i:           1,
				dist:        1000,
			},
			want: result{newDial: 50, zero: 10},
		},
		{
			name: "Test 12",
			args: args{
				currentDial: 84,
				i:           1,
				dist:        16,
			},
			want: result{newDial: 0, zero: 1},
		},
		{
			name: "Test 13",
			args: args{
				currentDial: 84,
				i:           -1,
				dist:        84,
			},
			want: result{newDial: 0, zero: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if dial, zero := computeNewDial(tt.args.currentDial, tt.args.i, tt.args.dist); dial != tt.want.newDial || zero != tt.want.zero {
				t.Errorf("%s computeNewDial(%v, %v, %v) got (%v, %v), want (%v, %v)", tt.name, tt.args.currentDial, tt.args.i, tt.args.dist, dial, zero, tt.want.newDial, tt.want.zero)
			}
		})
	}
}
