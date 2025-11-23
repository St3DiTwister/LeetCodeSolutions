package _136_Single_Number

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"example 1", []int{2, 2, 1}, 1},
		{"example 2", []int{4, 1, 2, 1, 2}, 4},
		{"negative", []int{-1, -1, -5}, -5},
		{"large", []int{10, 10, 100}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.in)
			if got != tt.want {
				t.Fatalf("singleNumber(%v) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
