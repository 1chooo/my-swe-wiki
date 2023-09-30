package main

import (
	"fmt"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 0; i < runTimes; i++ {
		var inNum int
		fmt.Scan(&inNum)
		fibonacci := NewFibonacci(inNum)
		fibonacci.RecursiveFibonacci(inNum)
		fibonacci.ShowAns()
	}
}

type Fibonacci struct {
	n int
}

func NewFibonacci(n int) *Fibonacci {
	return &Fibonacci{n: n}
}

func (f *Fibonacci) RecursiveFibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	return f.RecursiveFibonacci(n-1) + f.RecursiveFibonacci(n-2)
}

func (f *Fibonacci) ShowAns() {
	fmt.Println(f.RecursiveFibonacci(f.n))
}
