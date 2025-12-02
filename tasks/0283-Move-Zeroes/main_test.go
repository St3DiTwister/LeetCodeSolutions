package _283_Move_Zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"example 2", []int{0}, []int{0}},
		{"example 3", []int{0, 0, 1}, []int{1, 0, 0}},
		{"example 4", []int{1, 0, 1}, []int{1, 1, 0}},
		{"example 5", []int{1, 0, 0, 2}, []int{1, 2, 0, 0}},
		{"example 5", []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveZeroes(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
