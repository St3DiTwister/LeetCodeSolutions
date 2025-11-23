package _100_Same_Tree

import (
	"LeetCodeSolutions/internal/structures"
	"testing"
)

func buildTreeFromArray(vals []*int) *structures.TreeNode {
	// vals — это массив указателей:
	// nil → нет узла
	// *int → значение узла
	//
	// Дерево строим как complete binary tree по индексам:
	// left = 2*i+1, right = 2*i+2
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
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < len(nodes) {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < len(nodes) {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}

func pint(v int) *int { return &v }

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *structures.TreeNode
		q    *structures.TreeNode
		want bool
	}{
		{
			name: "both nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "one nil one non-nil",
			p:    &structures.TreeNode{Val: 1},
			q:    nil,
			want: false,
		},
		{
			name: "single node equal",
			p:    &structures.TreeNode{Val: 1},
			q:    &structures.TreeNode{Val: 1},
			want: true,
		},
		{
			name: "single node not equal",
			p:    &structures.TreeNode{Val: 1},
			q:    &structures.TreeNode{Val: 2},
			want: false,
		},
		{
			name: "same structure and values",
			p: buildTreeFromArray([]*int{
				pint(1), pint(2), pint(3),
			}),
			q: buildTreeFromArray([]*int{
				pint(1), pint(2), pint(3),
			}),
			want: true,
		},
		{
			name: "same structure different value",
			p: buildTreeFromArray([]*int{
				pint(1), pint(2), pint(3),
			}),
			q: buildTreeFromArray([]*int{
				pint(1), pint(2), pint(4),
			}),
			want: false,
		},
		{
			name: "different structure",
			p: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
				},
			},
			q: &structures.TreeNode{
				Val: 1,
				Right: &structures.TreeNode{
					Val: 2,
				},
			},
			want: false,
		},
		{
			name: "bigger equal tree",
			p: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
					Left: &structures.TreeNode{
						Val: 4,
					},
				},
				Right: &structures.TreeNode{
					Val: 3,
					Right: &structures.TreeNode{
						Val: 5,
					},
				},
			},
			q: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
					Left: &structures.TreeNode{
						Val: 4,
					},
				},
				Right: &structures.TreeNode{
					Val: 3,
					Right: &structures.TreeNode{
						Val: 5,
					},
				},
			},
			want: true,
		},
		{
			name: "bigger tree differ in leaf",
			p: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
					Left: &structures.TreeNode{
						Val: 4,
					},
				},
				Right: &structures.TreeNode{
					Val: 3,
					Right: &structures.TreeNode{
						Val: 5,
					},
				},
			},
			q: &structures.TreeNode{
				Val: 1,
				Left: &structures.TreeNode{
					Val: 2,
					Left: &structures.TreeNode{
						Val: 4,
					},
				},
				Right: &structures.TreeNode{
					Val: 3,
					Right: &structures.TreeNode{
						Val: 6, // отличается
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSameTree(tt.p, tt.q)
			if got != tt.want {
				t.Fatalf("isSameTree(...) = %v, want %v", got, tt.want)
			}
		})
	}
}
