package main

import (
	"fmt"
)

type Node struct {
	priority int
	data     string
}

type MaxHeap struct {
	counter    int
	maxAmount  int
	root       []*Node
}

func generateMaxHeap(total int) *MaxHeap {
	return &MaxHeap{
		counter:   0,
		maxAmount: total,
		root:      make([]*Node, total),
	}
}

func insert(maxHeap *MaxHeap, priority int, mission string) {
	if maxHeap.counter >= maxHeap.maxAmount {
		fmt.Println("超出上限")
		return
	}

	i := maxHeap.counter
	node := &Node{
		priority: priority,
		data:     mission,
	}
	maxHeap.root[i] = node

	for i > 0 && priority > maxHeap.root[(i-1)/2].priority {
		maxHeap.root[i], maxHeap.root[(i-1)/2] = maxHeap.root[(i-1)/2], maxHeap.root[i]
		i = (i - 1) / 2
	}

	maxHeap.counter++
}

func extractMaxHeap(maxHeap *MaxHeap) {
	if maxHeap.counter < 1 {
		fmt.Println("Queue Empty")
		return
	}

	i := maxHeap.counter
	fmt.Println(maxHeap.root[0].data)
	maxHeap.root[0], maxHeap.root[i] = maxHeap.root[i], maxHeap.root[0]

	maxHeap.counter--

	maxHeapify(maxHeap, 0)
}

func maxHeapify(maxHeap *MaxHeap, index int) {
	left := 2*index + 1
	right := left + 1
	max := index

	if right < maxHeap.counter && maxHeap.root[right].priority > maxHeap.root[left].priority {
		max = right
	} else if left < maxHeap.counter && maxHeap.root[left].priority > maxHeap.root[right].priority {
		max = left
	}

	if max != index {
		maxHeap.root[index], maxHeap.root[max] = maxHeap.root[max], maxHeap.root[index]
		maxHeapify(maxHeap, max)
	}
}

func main() {
	var total int
	fmt.Scanf("%d", &total)
	maxHeap := generateMaxHeap(total)

	for i := 0; i < total; i++ {
		var priority int
		var buffer string

		fmt.Scanf("%s %d", &buffer, &priority)
		insert(maxHeap, priority, buffer)
	}

	fmt.Println("First three things to do:")
	for i := 0; i < 3; i++ {
		extractMaxHeap(maxHeap)
	}
}
