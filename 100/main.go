package main

// [1,2,3]
var testCase1 *TreeNode = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

// [1,2]
var testCase2 *TreeNode = &TreeNode{1, &TreeNode{2, nil, nil}, nil}

// [1,2,1]
var testCase3 *TreeNode = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// If one of the nodes is nil, then the trees are not the same
	if p == nil || q == nil {
		return false
	}
	// If the values are not the same, then the trees are not the same
	if p.Val != q.Val {
		return false
	}
	// Recursively check the left and right nodes
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
