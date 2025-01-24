package binarytree

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

func searchBST(root *TreeNode, val int) *TreeNode {
	cptr := root
	for cptr != nil && cptr.Val != val {
		if val < cptr.Val {
			cptr = cptr.Left
		} else if val > cptr.Val {
			cptr = cptr.Right
		}
	}
	return cptr
}
