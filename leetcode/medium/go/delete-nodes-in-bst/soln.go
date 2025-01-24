package deletenodesinbst

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

func deleteNode(root *TreeNode, key int) *TreeNode {
	cptr := root
	var parent *TreeNode
	dir := ""
	for cptr != nil {
		if cptr.Val == key {
			newRoot := deleteTheNode(cptr, parent, dir)
			if parent == nil {
				return newRoot
			}
			break
		}
		parent = cptr
		if key < cptr.Val {
			dir = "l"
			cptr = cptr.Left
		} else {
			dir = "r"
			cptr = cptr.Right
		}
	}
	return root
}

func deleteTheNode(cptr, parent *TreeNode, direction string) *TreeNode {
	if ok, newRoot := checkAndDeleteLeafNode(cptr, parent, direction); ok {
		return newRoot
	} else if ok, newRoot := checkAndDeleteSingleChildNode(cptr, parent, direction); ok {
		return newRoot
	} else {
		return deleteBinaryNode(cptr, parent, direction)
	}
}

func deleteBinaryNode(cptr, parent *TreeNode, direction string) *TreeNode {
	if direction == "l" {
		parent.Left = cptr.Left
	} else if direction == "r" {
		parent.Right = cptr.Left
	}
	attachNode(cptr.Left, cptr.Right)
	return cptr.Left
}

func checkAndDeleteLeafNode(cptr, parent *TreeNode, direction string) (bool, *TreeNode) {
	if cptr.Right != nil || cptr.Left != nil {
		return false, nil
	}
	if direction == "l" {
		parent.Left = nil
	} else if direction == "r" {
		parent.Right = nil
	}
	return true, nil
}

func checkAndDeleteSingleChildNode(cptr, parent *TreeNode, direction string) (bool, *TreeNode) {
	if cptr.Right == nil {
		if cptr.Left != nil {
			if direction == "l" {
				parent.Left = cptr.Left
			} else if direction == "r" {
				parent.Right = cptr.Left
			}
			return true, cptr.Left
		}
		return false, nil
	} else if cptr.Left == nil {
		if direction == "l" {
			parent.Left = cptr.Right
		} else if direction == "r" {
			parent.Right = cptr.Right
		}
		return true, cptr.Right
	}
	return false, nil
}

func attachNode(root, node *TreeNode) {
	cptr := root
	for {
		if node.Val < cptr.Val {
			if cptr.Left == nil {
				cptr.Left = node
				break
			} else {
				cptr = cptr.Left
			}
		} else {
			if cptr.Right == nil {
				cptr.Right = node
				break
			} else {
				cptr = cptr.Right
			}
		}
	}
}
