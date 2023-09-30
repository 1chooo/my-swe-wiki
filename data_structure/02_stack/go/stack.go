package main

import (
	"fmt"
)

type Stack struct {
	arr []int
}

func (s *Stack) initialize() {
	// Initialize an empty slice
	s.arr = []int{}
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) push(data int) {
	s.arr = append(s.arr, data)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("Stack Underflow")
		return -1
	}
	// Get the top element of the stack
	topElement := s.arr[len(s.arr)-1]
	// Remove the top element from the stack
	s.arr = s.arr[:len(s.arr)-1]
	return topElement
}

func (s *Stack) peek() int {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.arr[len(s.arr)-1]
}

func (s *Stack) clear() {
	// Clear the stack, simply by creating a new empty slice
	s.arr = []int{}
}

func main() {
	var stack Stack
	stack.initialize()

	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack.pop())  // Output: 3
	fmt.Println(stack.pop())  // Output: 2
	fmt.Println(stack.peek()) // Output: 1 (peek without removing)
	stack.clear()             // Clear the stack
	fmt.Println(stack.pop())  // Output: Stack Underflow
}
