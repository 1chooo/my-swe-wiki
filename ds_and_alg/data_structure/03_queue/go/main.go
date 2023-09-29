package main

import "fmt"

type Queue struct {
	arr []int
}

func (q *Queue) initialize() {
	// 初始化一個空切片
	q.arr = []int{}
}

func (q *Queue) isEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) enqueue(data int) {
	q.arr = append(q.arr, data)
}

func (q *Queue) dequeue() int {
	if q.isEmpty() {
		fmt.Println("Queue Underflow")
		return -1
	}
	// 取得佇列前端元素
	frontElement := q.arr[0]
	// 刪除佇列前端元素
	q.arr = q.arr[1:]
	return frontElement
}

func (q *Queue) length() int {
	return len(q.arr)
}

func (q *Queue) clear() {
	// 清空佇列，即使在 Go 中不執行任何操作也足夠
	q.arr = []int{}
}

func main() {
	var queue Queue
	queue.initialize()

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	fmt.Println(queue.dequeue()) // Output: 1
	fmt.Println(queue.dequeue()) // Output: 2
	fmt.Println(queue.dequeue()) // Output: 3
	fmt.Println(queue.dequeue()) // Output: Queue Underflow
	
	fmt.Printf("Length of the queue: %d\n", queue.length()) // Output: 0
	
	queue.clear()
}
