package _100_Same_Tree

import (
	"LeetCodeSolutions/internal/structures"
)

func isSameTree(p *structures.TreeNode, q *structures.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
