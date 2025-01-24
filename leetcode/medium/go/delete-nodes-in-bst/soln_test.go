package deletenodesinbst

import (
	"fmt"
	"testing"
)

func TestSoln(t *testing.T) {
	tree := makeTree([]int{1, 2})
	bfsTraverse(tree)
	fmt.Println()
	r := deleteNode(tree, 1)
	bfsTraverse(r)
	fmt.Println()
}

func TestSoln2(t *testing.T) {
	tree := makeTree([]int{2})
	bfsTraverse(tree)
	fmt.Println()
	r := deleteNode(tree, 2)
	bfsTraverse(r)
}

func makeTree(n []int) *TreeNode {
	var root *TreeNode
	for _, v := range n {
		root = insertNode(root, v)
	}
	return root
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertNode(root.Left, val)
	} else {
		root.Right = insertNode(root.Right, val)
	}
	return root
}

func bfsTraverse(root *TreeNode) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
