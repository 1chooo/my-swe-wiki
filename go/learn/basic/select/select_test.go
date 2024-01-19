package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	// time.Sleep(time.Millisecond * 50)
	return "Done service"
}

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

// 透過select多路選擇，實現超時機制
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
