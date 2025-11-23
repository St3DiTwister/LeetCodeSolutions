package _112_Path_Sum

import (
	"LeetCodeSolutions/internal/structures"
	"testing"
)

func buildTree(vals []*int) *structures.TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	nodes := make([]*structures.TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &structures.TreeNode{Val: *v}
		}
	}

	for i := range nodes {
		if nodes[i] == nil {
			continue
		}
		l := 2*i + 1
		r := 2*i + 2
		if l < len(nodes) {
			nodes[i].Left = nodes[l]
		}
		if r < len(nodes) {
			nodes[i].Right = nodes[r]
		}
	}

	return nodes[0]
}

func pint(x int) *int { return &x }

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *structures.TreeNode
		targetSum int
		want      bool
	}{
		{
			name:      "empty tree",
			root:      nil,
			targetSum: 10,
			want:      false,
		},
		{
			name:      "single node equal",
			root:      &structures.TreeNode{Val: 5},
			targetSum: 5,
			want:      true,
		},
		{
			name:      "single node not equal",
			root:      &structures.TreeNode{Val: 5},
			targetSum: 4,
			want:      false,
		},
		{
			name: "example from LeetCode (path exists)",
			root: &structures.TreeNode{
				Val: 5,
				Left: &structures.TreeNode{
					Val: 4,
					Left: &structures.TreeNode{
						Val:   11,
						Left:  &structures.TreeNode{Val: 7},
						Right: &structures.TreeNode{Val: 2},
					},
				},
				Right: &structures.TreeNode{
					Val:  8,
					Left: &structures.TreeNode{Val: 13},
					Right: &structures.TreeNode{
						Val:   4,
						Right: &structures.TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			want:      true,
		},
		{
			name: "example from LeetCode (path does not exist)",
			root: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
				},
				Right: &structures.TreeNode{
					Val: 3,
				},
			},
			targetSum: 5,
			want:      false,
		},
		{
			name: "balanced with match",
			root: buildTree([]*int{
				pint(1),
				pint(2), pint(3),
				pint(4), pint(5), pint(6), pint(7),
			}),
			targetSum: 1 + 2 + 4,
			want:      true,
		},
		{
			name: "balanced no match",
			root: buildTree([]*int{
				pint(1),
				pint(2), pint(3),
				pint(4), pint(5), pint(6), pint(7),
			}),
			targetSum: 50,
			want:      false,
		},
		{
			name: "deep left path exists",
			root: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
					Left: &structures.TreeNode{
						Val: 3,
						Left: &structures.TreeNode{
							Val: 4,
						},
					},
				},
			},
			targetSum: 1 + 2 + 3 + 4,
			want:      true,
		},
		{
			name: "deep right path not exists",
			root: &structures.TreeNode{
				Val: 1,
				Right: &structures.TreeNode{
					Val: 2,
					Right: &structures.TreeNode{
						Val: 3,
						Right: &structures.TreeNode{
							Val: 4,
						},
					},
				},
			},
			targetSum: 100,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(tt.root, tt.targetSum)
			if got != tt.want {
				t.Fatalf("hasPathSum(...) = %v, want %v", got, tt.want)
			}
		})
	}
}
