package main

import (
	"fmt"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	uglyNumber := NewUglyNumber()

	for i := 0; i < runTimes; i++ {
		var inNum int
		fmt.Scan(&inNum)
		result := uglyNumber.FindFlag(inNum)
		fmt.Println(result)
	}
}

type UglyNumber struct{}

func NewUglyNumber() *UglyNumber {
	return &UglyNumber{}
}

func (ug *UglyNumber) FindFlag(a int) int {
	// Set "count" equals to 1 because 1 also is ugly number.
	count := 1

	// Then start counting from 2
	num := 2

	for {
		b := num

		// The domain of "Ugly Number" is that only divided into 2, 3, 5.
		for b%2 == 0 {
			b = b / 2
		}

		for b%3 == 0 {
			b = b / 3
		}

		for b%5 == 0 {
			b = b / 5
		}

		if b == 1 {
			count++
		}

		if count == a {
			return num
		} else {
			num++
		}
	}
}
