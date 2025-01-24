package insertintobinarysearchtree

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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	cptr := root
	node := &TreeNode{
		Val: val,
	}

	if cptr == nil {
		return node
	}

	for cptr != nil {
		if val < cptr.Val {
			if cptr.Left != nil {
				cptr = cptr.Left
			} else {
				cptr.Left = node
				break
			}
		} else if val > cptr.Val {
			if cptr.Right != nil {
				cptr = cptr.Right
			} else {
				cptr.Right = node
				break
			}
		}
	}

	return root
}
