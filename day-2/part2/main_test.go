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
			want: 33,
		},
		{
			name: "Test 1",
			args: args{
				low:  1188511880,
				high: 1188511890,
			},
			want: 1188511885,
		},
		{
			name: "Test 2",
			args: args{
				low:  112112111,
				high: 112112113,
			},
			want: 112112112,
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
