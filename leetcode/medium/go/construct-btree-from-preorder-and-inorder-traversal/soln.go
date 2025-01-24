package constructbtreefrompreorderandinordertraversal

import (
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
1.26
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeRecursively(preorder, inorder, 0)
}

func buildTreeRecursively(preorder, inorder []int, preorderRootIndex int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[preorderRootIndex],
	}

	inorderRootIndex := slices.Index(inorder, node.Val)
	if inorderRootIndex == -1 {
		return nil
	}

	if len(preorder) >= 2*preorderRootIndex+1 {
		node.Left = buildTree(preorder[2*preorderRootIndex+1:], inorder[:inorderRootIndex])
	}
	if len(preorder) >= 2*preorderRootIndex+2 {
		node.Right = buildTree(preorder[2*preorderRootIndex+2:], inorder[inorderRootIndex+1:])
	}
	return node
}
