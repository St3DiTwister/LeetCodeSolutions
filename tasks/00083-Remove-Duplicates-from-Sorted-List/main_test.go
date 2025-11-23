package _0083_Remove_Duplicates_from_Sorted_List

import (
	"LeetCodeSolutions/internal/structures"
	"testing"
)

func buildList(vals []int) *structures.ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &structures.ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &structures.ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *structures.ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "empty list",
			in:   []int{},
			want: []int{},
		},
		{
			name: "single element",
			in:   []int{1},
			want: []int{1},
		},
		{
			name: "all duplicates",
			in:   []int{1, 1, 1, 1},
			want: []int{1},
		},
		{
			name: "simple duplicates",
			in:   []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "multiple duplicates",
			in:   []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "no duplicates",
			in:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "long mix",
			in:   []int{1, 1, 1, 2, 2, 3, 3, 3, 4, 4},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.in)
			got := deleteDuplicates(head)
			gotSlice := listToSlice(got)

			if len(gotSlice) != len(tt.want) {
				t.Fatalf("length mismatch: got %v, want %v", gotSlice, tt.want)
			}

			for i := range gotSlice {
				if gotSlice[i] != tt.want[i] {
					t.Fatalf("deleteDuplicates(%v) = %v, want %v", tt.in, gotSlice, tt.want)
				}
			}
		})
	}
}
