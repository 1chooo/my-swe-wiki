package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
函數是一等公民
與其他語言差異
1.可以有多個返回值
2.所有參數都是傳值：slice,map,channel會有傳引用的錯覺（其實是傳值，只是是傳結構，但是結構裡包含的指針依然指向同一個記憶體位置）
3.函數可以作為變量的值
4.函數可以作為參數和返回值
*/

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 計算函數運行時間 有點類似functional programing or 裝飾子
func timeSpent(inner func(op int) int) func(n int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestTimeSpent(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(1000))
}

// 可變參數 透過...
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

// defer 類似java try finally
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err") //defer一樣會執行
	// fmt.Println("不會被執行")
}
func Clear() {
	fmt.Println("Clear resource")
}
