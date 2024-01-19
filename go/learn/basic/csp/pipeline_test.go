package csp_test

import (
	"fmt"
	"testing"
)

/*實現一個功能，給定一個切片，然後求它的子項的平方和。

例如，[1, 2, 3] -> 1^2 + 2^2 + 3^2 = 14。

正常的邏輯，遍歷切片，然後求平方累加。使用pipeline 模式，可以把求和和求平方拆分出來並行計算。*/
/*
* Pipeline 模式
 */

func generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

func TestPipeline(t *testing.T) {
	// [1, 2, 3]
	fmt.Println(<-sum(power(generator(3))))
}
