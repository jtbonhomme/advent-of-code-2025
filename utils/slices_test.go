package utils

import (
	"testing"
)

func TestEqualSlices(t *testing.T) {
	type args struct {
		a []any
		b []any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equal int slices",
			args: args{
				a: []any{1, 2, 3},
				b: []any{1, 2, 3},
			},
			want: true,
		},
		{
			name: "Different lengths",
			args: args{
				a: []any{1, 2, 3},
				b: []any{1, 2},
			},
			want: false,
		},
		{
			name: "Different contents",
			args: args{
				a: []any{1, 2, 3},
				b: []any{1, 2, 4},
			},
			want: false,
		},
		{
			name: "Equal string slices",
			args: args{
				a: []any{"a", "b", "c"},
				b: []any{"a", "b", "c"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualSlices(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EqualSlices(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
