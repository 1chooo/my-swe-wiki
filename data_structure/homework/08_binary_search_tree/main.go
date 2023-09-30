package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

func insert(root *node, value int) *node {
	if root == nil {
		root = &node{
			value: value,
			left:  nil,
			right: nil,
		}
	} else if value < root.value {
		root.left = insert(root.left, value)
	} else if value > root.value {
		root.right = insert(root.right, value)
	}
	return root
}

func depthFirstSearch(root *node, sum, target int) bool {
	if root == nil {
		return false
	}

	if (root.value + sum) == target && root.left == nil && root.right == nil {
		return true
	}

	if depthFirstSearch(root.left, sum+root.value, target) || depthFirstSearch(root.right, sum+root.value, target) {
		return true
	}

	return false
}

func main() {
	var amount, target int
	fmt.Scanf("%d", &amount)

	var BST *node

	for i := 0; i < amount; i++ {
		var value int
		fmt.Scanf("%d", &value)
		BST = insert(BST, value)
	}

	fmt.Scanf("%d", &target)
	if depthFirstSearch(BST, 0, target) {
		fmt.Println("There exists at least one path in the binary search tree.")
	} else {
		fmt.Println("There is no path in the binary search tree.")
	}
}
