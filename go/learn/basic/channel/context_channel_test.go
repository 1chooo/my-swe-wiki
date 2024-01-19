package channel_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

/*
1. Context 作用

每個Context 都會從最頂層Goroutine 一層一層傳遞到最下層，context.Context 可以在上層Goroutine 執行出現錯誤，會將信號及時同步到下一次層，
這樣，上層因為某些原因失敗時， 下層就可以停掉無用的工作，以減少資源損耗。
實際應用： RPC 超時時間設置

context 中一般意義context.WithValue 能從父上下文中創建一個子上下文，傳值的子上下文使用context.valueCtx 類型。
WithValue 是一對kv 類型，可用來傳值，實際應用：傳遞全局唯一的調用鏈

*/

/*
2. Context 接口
Context 是Go 語言再1.7 版本引入的標準庫接口，有以下需要實現的方法

Deadlime 返回context.Context 被取消的時間，也就是完成工作的截止時間
Done 返回一個Channel ，這個Channel 會在當前工作完成或者被取消後關閉，多次調用Done 方法會返回同一個Channle
Err 返回context.Context 結束的原因，只會在Done 方法對應的Channel 關閉時返回非空的值
Value 從context.Context 中獲取對應的值，對同一個上下文來說，多次調用Value 並傳入相同的Key 會返回相同的結果。改方法可以用來傳遞特定的請求。

如果context.Context 被取消，會返回Canceled 錯誤；
如果context.Context 超時，會返回DeadlineExceeded 錯誤；
*/

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// 3. 使用context 同步信號
/*
創建一個過期時間為1s 的上下文， 並向上下文傳入handle 函數，該方法會使用500ms 的時間處理傳入的請求。
*/

func TestContext(t *testing.T) {
	// 創建一個過期時間為1s 的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// ctx丟到goroutine
	go handle(ctx, 500*time.Millisecond)

	// 等到ctx過期
	for {
		select {
		case <-ctx.Done():
			fmt.Println("main", ctx.Err())
			return
		default:
			fmt.Println("wait for context deadline...")
			time.Sleep(time.Millisecond * 1500)
		}
	}

}

func handle(ctx context.Context, duration time.Duration) {

	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	// 500ms到走下面出去
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

//有時候任務是關聯的，像是起一個任務A他有兩個子任務B,C，當取消掉A的時候 BC也要跟著取消
/*
根 context: 通過context.Background()建立
子 context: context.WithCancel(parentContext)建立
	ctx,cancel := context.WithCancel(context.Background())
	返回cancel方法，跟ctx可以繼續往下傳
當前context被取消時，基於他的子context都會被取消
接收取消通知: <-ctx.Done()
*/

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// 透過context來關閉channel
func TestCancelByContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}

				time.Sleep(time.Millisecond * 5)

			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}

/*
調用context的WithCancel方法會返回一個可被取消的ctx和CancelFunc，需要取消ctx時，調用cancel函數即可。
而context有個Done方法，這個方法返回一個channel，當Context被取消時，這個channel會被關閉。
消費中的協程通過select監聽這個channel，收到關閉信號後一個return就能結束消費。

CancelFunc主要用途是預防系統做不必要的工作。
比如用戶請求A接口時，A接口內部需要請求A database、B cache 、C System獲取各種數據，把這些數據經過計算後組裝到一起返回給調用方。

但如果用戶在訪問網站時覺得沒意思，去其他網站了。此時若你的服務收到用戶請求後繼續去訪問其他C system、B database就是浪費資源。
比較符合直覺的做法是：當業務請求取消時，你的系統也應該停止請求下游系統。
當用戶取消訪問時，只要context監聽取消事件並在用戶取消時發送取消事件，就可以取消請求了。

除了用戶中途取消請求的情況，還有一種情況也可以用到cancelFunc：
服務A的返回數據依賴服務B和服務C的相關接口，若服務B或者服務C掛了，此次請求就算失敗了，沒必要再訪問另一個服務，此時也可以用CancelFunc

*/

func getUserInfoBySystemA(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	// 模擬請求出錯
	return errors.New("failed")
}

func getOrderInfoBySystemB(ctx context.Context) {
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("process finished")
	case <-ctx.Done():
		fmt.Println("process cancelled")
	}
}

func TestServiceFailed(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	//併發從兩個服務獲取相關數據
	go func() {
		err := getUserInfoBySystemA(ctx)
		if err != nil {
			// 發生錯誤，調用cancelFunc
			cancel()
		}
	}()

	getOrderInfoBySystemB(ctx)
}
