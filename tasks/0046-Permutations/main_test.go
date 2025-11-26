package _046_Permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{"example 1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"example 2", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"example 3", []int{1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
