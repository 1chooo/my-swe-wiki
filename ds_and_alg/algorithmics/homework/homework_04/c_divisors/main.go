package main

import (
	"fmt"
	"math"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 0; i < runTimes; i++ {
		var l, u int
		fmt.Scan(&l, &u)

		divisors := NewDivisors()
		divisors.BiggestDivisors(l, u)
	}
}

type Divisors struct{}

func NewDivisors() *Divisors {
	return &Divisors{}
}

func (d *Divisors) BiggestDivisors(a, b int) {
	max := 0
	divisor := 0
	num := 0

	for i := a; i <= b; i++ {
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				divisor += 2
			}
		}

		if divisor > max {
			max = divisor
			num = i
		}
		divisor = 0
	}

	fmt.Printf("Between %d and %d, %d has a maximum of %d divisors.\n", a, b, num, max)
}
