package sharemem_test

import (
	"sync"
	"testing"
	"time"
)

// 線程不安全的程序（這邊是協程）
func TestCounterGoroutineUnsafe(t *testing.T) {

	counter := 0
	// 開5000個協程來執行+1
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
	//counter = 4824 ，併發導致輸出不到5000
}

// 像i++ 這樣的操作就不是原子性的，在系統調用執行時，這個操作被分為了三步
// 首先是cpu 從內存中讀取到i 的值，讀到cpu 裡面的寄存器（寄存器是存在於cpu 中，空間比內存小，訪問速度比內存快）裡，
// 然后在寄存器裡執行了類似於ADD 這樣的指令然後就完成了+1 的操作，操作結果依舊放在寄存器裡，
// 最后cpu 裡的寄存器把值寫回到內存中，這樣才完成了一個i++ 的操作。
// 這樣的三步就相當於讀取，修改，寫入的過程。

// 由於線程是搶占式執行的，這就導致了，一個線程在自增執行了一半的時候，另一個線程也來讀取，然後自增。
// 假如有兩個線程分別在兩個cpu 中同時從執行i++ 的操作，那麼它們將同時讀到i 的值(此時讀取到的值相同)，
// 然後分別執行加1，再寫回到內存中，此時就會發現i 的值就只是加了一次。因為操作是非原子性，導致了出現了這樣的錯誤，從而導致發生線程不安全
// 如果runtime.GOMAXPROCS(1)，限制只用一個cpu，那結果就會是5000

func TestCounterGoroutineSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	// 開5000個協程來執行+1
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}() // 最後自動釋放鎖
			mut.Lock() //加鎖
			counter++
		}()
	}
	time.Sleep(1 * time.Second) //一定要sleep 不然協程還沒執行完 主程序就完了
	t.Logf("counter = %d", counter)
	//counter = 5000
}

// 正常來講不會透過sleep，應該要透過waitgroup來控制協程完成
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	// 開5000個協程來執行+1
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}() // 最後自動釋放鎖
			mut.Lock() //加鎖
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
	//counter = 5000
}
