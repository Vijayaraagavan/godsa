package algorithm

import "fmt"

type bstNode struct {
	value *int
	left  *bstNode
	right *bstNode
}

func CreateBST() {
	var root bstNode
	var input = []int{50, 30, 20, 40, 70, 60, 80}
	for _, v := range input {
		insertBstNode(&root, v)
	}
	fmt.Printf("bst traverse: ")
	bstTraverse(root)
}
func insertBstNode(n *bstNode, key int) {
	if n.value == nil {
		n.value = &key
		return
	}
	if *n.value > key {
		if n.left == nil {
			n.left = &bstNode{value: &key}
		} else {
			insertBstNode(n.left, key)
		}
		// n.left = &bstNode{value: key}
		return
	} else {
		if n.right == nil {
			n.right = &bstNode{value: &key}
		} else {
			insertBstNode(n.right, key)
		}
		// n.right = &bstNode{value: key}
		return
	}
}

func bstTraverse(n bstNode) {
	fmt.Printf("%d -> ", *n.value)
	if n.left != nil {
		bstTraverse(*n.left)
	}
	if n.right != nil {
		bstTraverse(*n.right)
	}
}
