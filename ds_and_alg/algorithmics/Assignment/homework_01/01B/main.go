package main

import (
	"fmt"
	"math"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 0; i < runTimes; i++ {
		var inNum int
		fmt.Scan(&inNum)
		myArray := []int{inNum}

		happyNumber := NewHappyNumber(inNum, myArray)
		happyNumber.JudgeHappyNumber()
	}
}

type HappyNumber struct {
	inNum   int
	myArray []int
}

func NewHappyNumber(inNum int, myArray []int) *HappyNumber {
	return &HappyNumber{inNum: inNum, myArray: myArray}
}

func (hn *HappyNumber) JudgeHappyNumber() {
	for {
		if hn.inNum == 4 {
			fmt.Println("Not Happy")
			break
		}

		sum := 0
		tempNum := hn.inNum
		for tempNum != 0 {
			r := tempNum % 10
			sum += int(math.Pow(float64(r), 2))
			tempNum /= 10
		}
		hn.myArray = append(hn.myArray, sum)

		last := len(hn.myArray) - 1
		if hn.myArray[last] == 4 {
			fmt.Println("Not Happy")
			break
		} else if hn.myArray[last] == 1 {
			fmt.Println("Happy")
			break
		}

		hn.inNum = sum

		len1 := len(hn.myArray)
		for j := 0; j < len(hn.myArray); j++ {
			for k := 0; k < len(hn.myArray); k++ {
				if j != k && hn.myArray[j] == hn.myArray[k] {
					hn.myArray = append(hn.myArray[:k], hn.myArray[k+1:]...)
				}
			}
		}
		len2 := len(hn.myArray)
		if len1 < len2 {
			fmt.Println("Not Happy")
			break
		}
	}
}
