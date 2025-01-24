package constructbtreefrompreorderandinordertraversal

import (
	"fmt"
	"testing"
)

func TestSoln(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{1, 2}
	tree := buildTree(preorder, inorder)
	bfsTraverse(tree)
}

func bfsTraverse(root *TreeNode) {
	queue := []*TreeNode{root}
	node := queue[0]
	for len(queue) > 0 && node != nil {
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
