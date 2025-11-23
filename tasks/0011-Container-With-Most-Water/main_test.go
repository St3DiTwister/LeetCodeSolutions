package _011_Container_With_Most_Water

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "leetcode example",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "two elements equal",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "two elements different",
			height: []int{1, 3},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Fatalf("maxArea(%v), want %d", got, tt.want)
			}
		})
	}
}
