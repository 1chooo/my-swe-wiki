package main

import (
	"fmt"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 1; i <= runTimes; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		oddSum := NewOddSum(a, b)
		oddSum.ShowAns(i)
	}
}

type OddSum struct {
	a, b, sum int
}

func NewOddSum(a, b int) *OddSum {
	return &OddSum{a: a, b: b, sum: 0}
}

func (os *OddSum) CountOddSum() {
	for i := os.a; i <= os.b; i++ {
		if i%2 == 1 {
			os.sum += i
		}
	}
}

func (os *OddSum) ShowAns(num int) {
	os.CountOddSum()
	fmt.Printf("Case %d: %d\n", num, os.sum)
}
