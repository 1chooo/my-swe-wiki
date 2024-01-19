package channel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 利用channel時要注意goroutine阻塞洩漏問題
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

// 遇到只要處理第一筆資料的情境
func FirstResponseUnBufferCh() string {
	num0fRunner := 10
	ch := make(chan string)
	for i := 0; i < num0fRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret //除了第一筆 ch會阻塞
		}(i)
	}
	return <-ch
}
func TestFirstResponseUnBufferCh(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) //算goroutine數量
	t.Log(FirstResponseUnBufferCh())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
	// 發現goroutine阻塞與洩漏

}

// 改用buffer channel來防止阻塞
func FirstResponseBufferCh() string {
	num0fRunner := 10
	ch := make(chan string, num0fRunner)
	for i := 0; i < num0fRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}
func TestFirstResponseBufferCh(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponseBufferCh())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}

// 遇到要等到所有goroutune都執行完的情境，除了用sync.WaitGroup以外，也可以利用channel的這種方式解決
func AllResponseBufferCh() string {
	num0fRunner := 10
	ch := make(chan string, num0fRunner)
	for i := 0; i < num0fRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < num0fRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponseBufferCh(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponseBufferCh())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
