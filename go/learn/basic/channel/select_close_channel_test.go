package channel

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel1(t *testing.T) {
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}

				time.Sleep(time.Millisecond * 5)

			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel_1(cancelChan)
	// 只會取消一個協程，因為只有一個阻塞被取出
	time.Sleep(time.Second * 1)
}

func TestCancel2(t *testing.T) {
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}

				time.Sleep(time.Millisecond * 5)

			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	// 若是透過close() ，用到同一個channel的所有的協程都會被關閉
	time.Sleep(time.Second * 1)
}
