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
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				currentDial: 50,
				i:           1,
				dist:        30,
			},
			want: 80,
		},
		{
			name: "Test 2",
			args: args{
				currentDial: 90,
				i:           1,
				dist:        20,
			},
			want: 10,
		},
		{
			name: "Test 3",
			args: args{
				currentDial: 10,
				i:           -1,
				dist:        20,
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeNewDial(tt.args.currentDial, tt.args.i, tt.args.dist); got != tt.want {
				t.Errorf("computeNewDial() = %v, want %v", got, tt.want)
			}
		})
	}
}
