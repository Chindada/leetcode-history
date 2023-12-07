package main

// [1,null,2,3]
var testCase1 *TreeNode = &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

// []
var testCase2 *TreeNode = nil

// [1]
var testCase3 *TreeNode = &TreeNode{1, nil, nil}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of a binary tree, return the preorder traversal of its nodes' values.

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}
