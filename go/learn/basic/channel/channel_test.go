package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

/*上面是數據量固定的情況，也就是Receiver跟Producer知道彼此有多少數據要傳輸跟接受
  但是正常來說不會知道，而且也有可能存在多個Producer對多個Receiver的情況
  所以要將通道關閉
*/

func dataProducer2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) //數據發送完 ，直接關閉channel，就不用通知所有的receiver數據發送完畢
		// ch <- 11 //往關閉的通道上發，會發生panic
		wg.Done()
	}()
}
func dataReceiver2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// for i := 0; i < 10; i++ {
		for { // 改成不知道數量的形式
			// data := <-ch
			// 這個情境，要判斷channel有沒有被關閉，要改寫成這樣
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}

func TestCloseChannel2(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer2(ch, &wg)
	wg.Add(1)
	dataReceiver2(ch, &wg)
	wg.Wait()
}

// 多個receiver 也可以正常執行，只是因為多協程 順序會不固定
func TestCloseChannel3(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer2(ch, &wg)

	wg.Add(1)
	dataReceiver2(ch, &wg)

	wg.Add(1)
	dataReceiver2(ch, &wg)

	wg.Wait()
}
