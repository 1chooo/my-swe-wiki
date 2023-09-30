package main

import (
	"fmt"
	"sort"
)

func main() {
	for {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		myArray := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&myArray[i])
		}

		ageSort := NewAgeSort(myArray)
		ageSort.Sort()
	}
}

type AgeSort struct {
	myArray []int
}

func NewAgeSort(myArray []int) *AgeSort {
	return &AgeSort{
		myArray: myArray,
	}
}

func (as *AgeSort) Sort() {
	sort.Ints(as.myArray)

	for i, age := range as.myArray {
		if i == len(as.myArray)-1 {
			fmt.Println(age)
		} else {
			fmt.Printf("%d ", age)
		}
	}
}
