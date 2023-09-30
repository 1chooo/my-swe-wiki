package main

import (
	"fmt"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 0; i < runTimes; i++ {
		var inStr1, inStr2 string
		fmt.Scan(&inStr1)
		fmt.Scan(&inStr2)

		allInAll := NewAllInAll(inStr1, inStr2)
		allInAll.ShowAns()
	}
}

type AllInAll struct {
	inStr1 string
	inStr2 string
}

func NewAllInAll(inStr1, inStr2 string) *AllInAll {
	return &AllInAll{
		inStr1: inStr1,
		inStr2: inStr2,
	}
}

func (aia *AllInAll) Judge() bool {
	len1 := len(aia.inStr1)
	len2 := len(aia.inStr2)

	for j := len1 - 1; j >= 0; j-- {
		for k := len2 - 1; k >= 0; k-- {
			if aia.inStr1[j] == aia.inStr2[k] {
				len2-- // this will prevent from counting repeatedly.
				break
			}
		}
	}

	swapTimes := 0

	// If the order of the array isn't descending,
	// it represents that "inStr1" is not in "inStr2".
	// Using the concept of the "Bubble Sort".

	for i := 0; i < len(aia.inStr2)-1; i++ {
		first := aia.inStr2[i]
		last := aia.inStr2[i+1]
		if first < last {
			swapTimes++
		}
	}

	return len(aia.inStr2) == len1 && swapTimes == 0
}

func (aia *AllInAll) ShowAns() {
	if aia.Judge() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
