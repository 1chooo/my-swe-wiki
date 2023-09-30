package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputStr string
	fmt.Scanln(&inputStr)

	inStr := strings.Fields(inputStr)

	len := len(inStr)

	sequence := make([]int, len)

	for i := 0; i < len; i++ {
		fmt.Sscanf(inStr[i], "%d", &sequence[i])
	}

	for i := 0; i < len; i++ {
		if i == len-1 {
			fmt.Println(sequence[i])
		} else {
			fmt.Printf("%d ", sequence[i])
		}
	}

	quickSort := NewQuickSort()
	quickSort.Sort(sequence, 0, len-1)
}

type QuickSort struct{}

func NewQuickSort() *QuickSort {
	return &QuickSort{}
}

func (qs *QuickSort) Sort(A []int, lb, rb int) {
	if lb >= rb {
		return
	}

	pivot := A[rb]
	l := lb
	r := rb - 1

	for {
		for A[l] < pivot {
			l++
		}
		for A[r] >= pivot && r > lb {
			r--
		}
		if l < r {
			A[l], A[r] = A[r], A[l]
			fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
		} else {
			break
		}
	}

	if A[rb] != A[l] {
		A[rb], A[l] = A[l], A[rb]
		fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
	}

	qs.Sort(A, lb, l-1)
	qs.Sort(A, l+1, rb)
}
