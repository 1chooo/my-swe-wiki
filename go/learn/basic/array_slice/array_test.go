package array_slice

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	// 宣告方式
	var arr [3]int
	t.Log(arr[1], arr[2])
	arr1 := [4]int{1, 2, 3, 4}
	t.Log(arr1)
	arr2 := [...]int{1, 3, 5, 7, 9}
	t.Log(arr2)
	arr1[1] = 99
	t.Log(arr1)

	arr4 := [2][2]int{{1, 2}, {3, 4}}
	t.Log(arr4)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 5, 7, 9}
	// 1. normal
	for i := 0; i < len(arr); i++ {
		t.Log("index:", i, " value:", arr[i])
	}
	// 2. use range
	for index, value := range arr {
		t.Log("index:", index, " value:", value)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr_sec := arr[3:] //[4 5]
	t.Log(arr_sec)
}
