package _070_Climbing_Stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"n=0", 0, 1},
		{"n=1", 1, 1},
		{"n=2", 2, 2},

		{"n=3", 3, 3},
		{"n=4", 4, 5},
		{"n=5", 5, 8},
		{"n=6", 6, 13},
		{"n=7", 7, 21},
		{"n=8", 8, 34},
		{"n=9", 9, 55},
		{"n=10", 10, 89},

		{"n=20", 20, 10946},
		{"n=30", 30, 1346269},

		{"n=45", 45, 1836311903},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Fatalf("climbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
