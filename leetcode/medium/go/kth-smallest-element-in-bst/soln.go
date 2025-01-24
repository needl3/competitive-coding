package kthsmallestelementinbst

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

func kthSmallest(root *TreeNode, k int) int {
	q := []*TreeNode{}
	cptr := root
	for {
		if cptr == nil {
			for len(q) != 0 {
				cptr = q[len(q)-1]
				q = q[:len(q)-1]
				k--
				if k == 0 {
					return cptr.Val
				}
				if cptr.Right != nil {
					cptr = cptr.Right
					break
				}
			}
		}

		q = append(q, cptr)
		cptr = cptr.Left
	}
}
