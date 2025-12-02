package _278_First_Bad_Version

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"example 1", 5, 4},
		{"example 2", 1, 1},
		{"example 3", 10, 3},
		{"example 4", 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isBad := func(version int) bool {
				return version >= tt.want
			}

			if got := firstBadVersion(tt.num, isBad); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
