package binarytreeinordertraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	v := []int{}
	if root == nil {
		return v
	}

	leftVal := inorderTraversal(root.Left)
	rightVal := inorderTraversal(root.Right)
	return append(append(leftVal, root.Val), rightVal...)
}
