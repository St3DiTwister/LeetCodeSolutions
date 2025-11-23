package _069_Sqrt

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"x=0", 0, 0},
		{"x=1", 1, 1},

		{"x=2", 2, 1},
		{"x=3", 3, 1},
		{"x=4", 4, 2},
		{"x=8", 8, 2},
		{"x=9", 9, 3},

		{"x=15", 15, 3},
		{"x=16", 16, 4},
		{"x=17", 17, 4},
		{"x=24", 24, 4},
		{"x=25", 25, 5},
		{"x=26", 26, 5},

		{"x=100", 100, 10},
		{"x=101", 101, 10},
		{"x=2147395599", 2147395599, 46339},
		{"x=2147483647", 2147483647, 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.x)
			if got != tt.want {
				t.Fatalf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
