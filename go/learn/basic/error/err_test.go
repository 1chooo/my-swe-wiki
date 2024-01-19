package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThenTwoError = errors.New("n should be not less than 2")
var LargeThenHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThenTwoError
	}
	if n > 100 {
		return nil, LargeThenHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(5); err != nil {
		if err == LargeThenHundredError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

// 這樣寫錯誤一多 嵌套會越來越多 （要避免，及早失敗返回最好）
func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

// 應該要反方向寫
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err == nil {

		fmt.Println("Error", err)
		return

	}
	fmt.Println(list)
}
