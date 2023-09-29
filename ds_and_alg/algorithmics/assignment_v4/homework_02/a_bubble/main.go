package main

import (
	"fmt"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 0; i < runTimes; i++ {
		var length int
		fmt.Scan(&length)

		myArray := make([]int, length)
		for j := 0; j < length; j++ {
			fmt.Scan(&myArray[j])
		}

		bubbleSort := NewBubbleSort(myArray)
		bubbleSort.ShowAns()
	}
}

type BubbleSort struct {
	myArray    []int
	swapTimes  int
}

func NewBubbleSort(myArray []int) *BubbleSort {
	return &BubbleSort{
		myArray:   myArray,
		swapTimes: 0,
	}
}

func (bs *BubbleSort) Sort() {
	length := len(bs.myArray)
	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			left := bs.myArray[j]
			right := bs.myArray[j+1]
			if left > right {
				bs.myArray[j], bs.myArray[j+1] = right, left
				bs.swapTimes++
			}
		}
	}
}

func (bs *BubbleSort) ShowAns() {
	bs.Sort()
	fmt.Printf("Optimal swapping takes %d swaps.\n", bs.swapTimes)
}
