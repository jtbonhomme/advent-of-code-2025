package main

import (
	"testing"
)

func TestCountInvalidIDsInRange(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 0",
			args: args{
				low:  11,
				high: 22,
			},
			want: 2,
		},
		{
			name: "Test 1",
			args: args{
				low:  95,
				high: 115,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				low:  998,
				high: 1012,
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				low:  998,
				high: 1112,
			},
			want: 2,
		},
		{
			name: "Test 4",
			args: args{
				low:  998,
				high: 2021,
			},
			want: 11,
		},
		{
			name: "Test 5",
			args: args{
				low:  111,
				high: 999,
			},
			want: 0,
		},
		{
			name: "Test 6",
			args: args{
				low:  123122,
				high: 123124,
			},
			want: 1,
		},
		{
			name: "Test 7",
			args: args{
				low:  38593856,
				high: 38593862,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countInvalidIDsInRange(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("%s countInvalidIDsInRange(%v, %v) got %v, want %v", tt.name, tt.args.low, tt.args.high, got, tt.want)
			}
		})
	}

}
