package _0083_Remove_Duplicates_from_Sorted_List

import (
	"LeetCodeSolutions/internal/structures"
)

func deleteDuplicates(head *structures.ListNode) *structures.ListNode {
	start := head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return start
}
