package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

func inorderT(root *node) {
	if root == nil {
		return
	}
	inorderT(root.left)
	fmt.Printf("%d ", root.data)
	inorderT(root.right)
}

func createNode(value int) *node {
	newNode := &node{
		data:  value,
		left:  nil,
		right: nil,
	}
	return newNode
}

func insertLevelOrder(arr []int, root *node, index, max int) *node {
	if index < max {
		temp := createNode(arr[index])
		root = temp

		root.left = insertLevelOrder(arr, root.left, 2*index+1, max)
		root.right = insertLevelOrder(arr, root.right, 2*index+2, max)
	}
	return root
}

func main() {
	const MAX_SIZE = 1000
	arr := make([]int, MAX_SIZE)
	i := 0
	var temp int

	for {
		_, err := fmt.Scanf("%d", &temp)
		if err != nil {
			break
		}
		arr[i] = temp
		i++
	}

	var root *node
	root = insertLevelOrder(arr, root, 0, i)
	inorderT(root)
}
