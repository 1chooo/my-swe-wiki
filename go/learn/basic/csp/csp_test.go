package csp_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done service"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

// 正常情況下，順序執行，耗时将是service()和otherTask()之和.
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// 為提高效能使用channel，來建構異步服務
/*
異步返回:利用Golang中的CSP並發機制實現FutureTask

熟悉Java的人對FutureTask有所了解.通俗來講就是當執行一個函數或一個task的時候,並不需要馬上知道它的結果
在需要結果時get結果.如果結果還沒有出來,就繼續阻塞在那裡;如果等到了結果就繼續向下執行.
這樣做的好處可以提高程序運行的效率,因為在執行task的時候,可以去執行其他任務
*/
func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("return result ")
		retCh <- ret //阻塞，等待取出，取出後才會往下跑
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 若沒取出，協成就不會跑到service exited，會一直在那裡等待
}

// 使用buffer channel解決
func AsyncServiceBufferChan() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result ")
		retCh <- ret //放完會繼續往下跑，不會阻塞了
		fmt.Println("service exited")
	}()
	return retCh
}
func TestAsyncServiceBufferChan(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

// 經典面試題：兩個線程輪流打印0到100？
// 建立兩個協程一個負責印奇數，一個印偶數
func goroutineOdd(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		<-ch
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
func goroutineEven(wg *sync.WaitGroup, ch chan struct{}) {
	defer func() {
		wg.Done()
		close(ch)
	}()

	for i := 0; i <= 100; i++ {
		ch <- struct{}{}
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
func TestCrossPrint0To100(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan struct{})
	// 這裡有一個技巧：為什麼使用struct 類型作為channel 的通知？
	// 很多開源代碼都是使用這種方式來作為信號通知機制，主要是因為空struct 在Go 中佔的內存是最少的。
	wg.Add(2)
	go goroutineOdd(&wg, ch)
	go goroutineEven(&wg, ch)
	wg.Wait()
}

// 題解：兩個協程都執行0到100次循環，但是不管哪個線程跑的快，在每次循環輸出時均會同步對齊， 每次循環時只輸出一個奇/偶值， 這樣也不用考慮兩個協程的啟動順序。
