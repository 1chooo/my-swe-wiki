package goroutine_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i) //值傳遞會複製一份 丟進協程
	}
	time.Sleep(time.Millisecond * 50)
}

// 錯誤的寫法
func TestGoroutineWrongUse(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i) //直接訪問到i，i被共享了
		}()
	}
	time.Sleep(time.Millisecond * 50)
}

func memConsumed() uint64 {
	runtime.GC() //GC，排除物件影響
	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)
	return memStat.Sys
}

//在Go語言中，對於多執行緒是相當友善好用的，相對其他語言所需要的資源與行數都少很多。
//以Java 8為例，執行一個Thread 預設需要分配1MB 記憶體，而Golang只需要幾kB 。
//goroutine 所佔用的記憶體，均在stack中進行管理
//goroutine 所佔用的棧空間大小，由 runtime 按需進行分配
func TestGetGoroutineMemConsume(t *testing.T) {
	var c chan int
	var wg sync.WaitGroup
	const goroutineNum = 1e4 // 1 * 10^4

	noop := func() {
		wg.Done()
		<-c //防止goroutine退出，記憶體被釋放
	}

	wg.Add(goroutineNum)
	before := memConsumed() //獲取建立goroutine前記憶體
	for i := 0; i < goroutineNum; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed() //獲取建立goroutine後記憶體
	//計算單個Goroutine記憶體佔用大小（2~4kb）
	fmt.Printf("====>%.3f KB\n", float64(after-before)/goroutineNum/1024)
}

func TestGoroutinePROCS(t *testing.T) {
	runtime.GOMAXPROCS(1) // 設置進程綁定的邏輯處理器
	//對於邏輯處理器的個數，不是越多越好，要根據電腦的實際物理核數，如果不是多核的，設置再多的邏輯處理器個數也沒用，
	//如果需要設置的話，一般我們採用如下代碼設置。
	// runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			fmt.Println("A:", i)
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			fmt.Println("B:", i)
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Wait()
}

// 展示主執行緒執行結束後，會將子執行緒release
func TestGoroutineRelease(t *testing.T) {
	//     執行子執行序
	go func() {
		time.Sleep(100000000)
		fmt.Println("Goroutine Done!")
	}()
	fmt.Println("Done!")
}

// 以上執行的結果為"Done！"，原因是在未執行完Goroutine的時候就自動的被釋放掉了，導致不會印出Goroutine Done！。

/*
一般來說使用多執行緒中，最常會遇到會5個問題如下:
	1.多執行緒相互溝通
	2.等待一執行緒結束後再接續工作
	3.多執行緒共用同一個變數
	4.不同執行緒產出影響後續邏輯
	5.兄弟執行緒間不求同生只求同死
	根據上述問題，基本上都可以透過channel, context, sync.WaitGroup, Select, sync.Mutex等方式解決，下面詳細解析如何解決:
*/

// 1. 多執行緒相互溝通
// 執行緒間的存取有兩種方式:
// 1-共用透過記憶體 => 在這邊都是以記憶體的方式進行存取
// 2-透過Socket的方式

// Goroutine的溝通主要可以透過channel、全域變數進行操作。Channel有點類似Linux C語言中pipe的方式，主要分成分為寫入端與讀取端。而全域變數的方式就是單純變數。
// 首先Channel的部份，宣告的方式是透過chan關鍵字宣告，搭配make 關鍵字令出空間，語法為: make(chan 型別 容量)

// 範例: channel控制執行緒，收集兩個執行序的資料 1、2
func TestGoroutineByChannel(t *testing.T) {
	// 宣告channel make(chan 型態 <容量>)
	val := make(chan int)
	// 執行第一個執行緒
	go func() {
		fmt.Println("intput val 1")
		val <- 1 //注入資料1
	}()
	// 執行第二個執行緒
	go func() {
		fmt.Println("intput val 2")
		val <- 2 //注入資料2
		time.Sleep(time.Millisecond * 100)
	}()
	ans := []int{}
	for {
		ans = append(ans, <-val) //取出資料
		fmt.Println(ans)
		if len(ans) == 2 {
			break
		}
	}
}

// 另一個方式就是比較傳統的方式進行存取，直接使用變數進行存取如下:
// 範例: 共用變數
func TestGoroutineByValue(t *testing.T) {
	val := 1
	// 執行第一個執行緒
	go func() {
		fmt.Println("first", val)
	}()
	// 執行第二個執行緒
	go func() {
		fmt.Println("sec ", val)
	}()
	time.Sleep(time.Millisecond * 500)
}

// 2. 等待一執行緒結束後再接續工作
// Java可以聯想到Join的概念，而在Golang中要做到等待的這件事情有兩個方法，一個是sync.WaitGroup、另一個是channel。
// 首先Sync.WaitGroup 像是一個計數器，啟動一條Goroutine 計數器 +1; 反之結束一條 -1。若計數器為複數代表Error。

//範例: 等待一執行緒結束後再接續工作(使用WaitGroup)
func TestGoroutineWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	// 執行執行緒
	go func() {
		defer wg.Done() //defer表示最後執行，因此該行為最後執行wg.Done()將計數器-1
		defer log.Println("goroutine drop out")
		log.Println("start a go routine")
		time.Sleep(time.Second) //休息一秒鐘
	}()
	wg.Add(1)                         //計數器+1
	time.Sleep(time.Millisecond * 30) //休息30 ms
	log.Println("wait a goroutine")
	wg.Wait() //等待計數器歸0
}

// Channel 的作法是利用等待提取、等待可注入會lock住的特性，達到Sync.WaitGroup 的功能。
// 範例:不同執行緒產出影響後續邏輯，使用多路復用
func TestGoroutineByChannel2(t *testing.T) {
	forever := make(chan int) //宣告一個channel
	//執行執行序
	go func() {
		defer log.Println("goroutine drop out")
		log.Println("start a go routine")
		time.Sleep(time.Second) //等待1秒鐘
		forever <- 1            //注入1進入forever channel
	}()
	time.Sleep(time.Millisecond * 30) //等待30 ms
	log.Println("wait a goroutine")
	<-forever // 取出forever channel 的資料
}

// 3. 多執行緒共用同一個變數
// 在多執行緒的世界，只是讀取一個共用變數是不會有問題的，但若是要進行修改可能會因為多個執行緒正在存取造成concurrent 錯誤。
// 若要解決這種情況，必須在存取時先將資源lock住，就可以避免這種問題。

// example 5: 多執行緒共用同一個變數

//範例: 多個執行序讀寫同一個變數
func TestGoroutineUseLock(t *testing.T) {
	var lock sync.Mutex   // 宣告Lock 用以資源佔有與解鎖
	var wg sync.WaitGroup // 宣告WaitGroup 用以等待執行序
	val := 0
	// 執行 執行緒: 將變數val+1
	go func() {
		defer wg.Done() //wg 計數器-1
		//使用for迴圈將val+1
		for i := 0; i < 10; i++ {
			lock.Lock() //佔有資源
			val++
			fmt.Printf("First gorutine val++ and val = %d\n", val)
			lock.Unlock() //釋放資源
			time.Sleep(3000)
		}
	}()
	// 執行 執行緒: 將變數val+1
	go func() {
		defer wg.Done() //wg 計數器-1
		//使用for迴圈將val+1
		for i := 0; i < 10; i++ {
			lock.Lock() //佔有資源
			val++
			fmt.Printf("Sec gorutine val++ and val = %d\n", val)
			lock.Unlock() // 釋放資源
			time.Sleep(1000)
		}
	}()
	wg.Add(2) //記數器+2
	wg.Wait() //等待計數器歸零
}

// sync.Mutex: 宣告資源鎖
// Lock: 在存取時需要將資源鎖住
// Unlock: 存取結束後需要釋放出來給需要的執行序使用

// 4. 不同執行緒產出影響後續邏輯
// 執行多執行緒控制時，可能會多個執行緒產生出的結果都不一樣，但每個結果都會影響下一步的動作。
// 例如: 在做error控制時，只要某一個Goroutine 錯誤時，就做相對應的處置，這樣的需求中，需要提不同錯誤不同的對應處置。
// 此時在這種情況下，就需要select多路複用的方式解:

// example 6: 不同執行緒產出影響後續邏輯

//範例:不同執行緒產出影響後續邏輯，使用多路復用。
func TestGoroutineUseSelect(t *testing.T) {
	firstRoutine := make(chan string) //宣告給第1個執行序的channel
	secRoutine := make(chan string)   //宣告給第2個執行序的channel
	rand.Seed(time.Now().UnixNano())

	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r)) //隨機等待 0~100 ms
		firstRoutine <- "first goroutine"
	}()
	go func() {
		r := rand.Intn(100)
		time.Sleep(time.Microsecond * time.Duration(r)) //隨機等待 0~100 ms
		secRoutine <- "Sec goroutine"
	}()
	select {
	case f := <-firstRoutine: //第1個執行序先執行後所要做的動作
		fmt.Println(f)
		return
	case s := <-secRoutine: //第2個執行序先執行後所要做的動作
		fmt.Println(s)
		return
	}
}

// 上面程式碼的例子，當其中一條Goroutine先結束時，主程式就會自動結束。
// 而Select的用法就是去聽哪一個channel已經先被注入資料，而做相對應的動作，若同時則是隨機採用對應的方案。

// 5. 兄弟執行緒間不求同生只求同死
// 在Goroutine主要的基本用法與應用，在上述都可以做到。
// 在這一章節主要是介紹一些進階用法" Context"。這種用法主要是在go 1.7之後才正式被收入官方套件中，使得更方便的控制Goroutine的生命週期。

// 主要提供以下幾種方法:
// WithCancel: 當parent呼叫cancel方法之後，所有相依的Goroutine 都會透過context接收parent要所有子執行序結束的訊息。
// WithDeadline: 當所設定的時間到時所有相依的Goroutine 都會透過context接收parent要所有子執行序結束的訊息。
// WithTimeout: 當所設定的日期到時所有相依的Goroutine 都會透過context接收parent要所有子執行序結束的訊息。
// WithValue: parent可透過訊息的方式與所有相依的Goroutine進行溝通。

// 以WithTimeout作為例子，下面例子是透過context的方式設定當超過10 ms沒結束Goroutine的執行，
// 則會發起"context deadline exceed"的錯誤訊息，或者成功執行就發出overslept的訊息

// 範例: 兄弟執行緒間不求同生只求同死，使用context​

const shortDuration = 1001 * time.Millisecond

var wg sync.WaitGroup //宣告計數器

func aRoutine(ctx context.Context) {
	defer wg.Done() //當該執行緒執行到最後計數器-1
	select {
	case <-time.After(1 * time.Second): // 1秒之後繼續執行,改成2就會先觸發到Deadline，就會走下面那條
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // context deadline exceeded
	}
}

func TestGoroutineUseContext(t *testing.T) {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d) //宣告一個context.WithDeadline並注入1.001秒之類為執行完的執行緒將發產出ctx.Err
	defer cancel()                                               // 程式最後執行WithDeadline失效
	go aRoutine(ctx)                                             // 啟動aRoutine執行序
	wg.Add(1)                                                    // 計數器+1
	wg.Wait()                                                    //等待計數器歸零
}

// Tips: context.Background(): 取得Context的實體
// context.WithDeadline(Context實體, 時間): 使用WithDeadline並設定好時間 Cancel 則是在程式結束前需要被使用，否則會有memory leak的錯誤訊息

// 總結
// 在Golang多執行緒的世界中，最常用的就是共用變數、channel、 Select、sync.WaitGroup、sync.Lock等方式，比較進階的用法是Context。
// Context主要就是官方提供一個interface使得大家更方便的去操作，若使用者不想使用也是可以透過channel自行實作。
