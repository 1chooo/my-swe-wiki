package channel

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

/*多用於自定義的連接池 ex: db,網路*/
type ReusableObj struct{}

type ObjPool struct {
	bufChan chan *ReusableObj //用於緩存可重用對象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := new(ObjPool)
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超時控制
		return nil, errors.New("Time Out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default: // 如果有人誤放，為了防止放滿，寫default返回異常
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	// 創建數量obj數量＝10的pool
	pool := NewObjPool(10)
	// 把10個obj都撈出來
	for i := 0; i < 10; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T", v)
		}
	}
	// 撈11個obj出來，最後一個因為阻塞，會timeout
	// if v, err := pool.GetObj(time.Second * 1); err != nil {
	// 	t.Error(err)
	// } else {
	// 	fmt.Printf("%T", v)
	// }

	// 把10個對象存回去
	for i := 0; i < 10; i++ {
		if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
			t.Error(err)
		}
	}

	// 多存一個對象回去的話，也會因為阻塞返回異常
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }

	fmt.Println("Done")
}

/* 為什麼不用sync包pool (sync.Pool（）)?

與sync包pool的不同 一句話總結：保存和復用臨時對象，減少內存分配，降低GC壓力。*/

/*
Pool 的目的是緩存已分配但未使用的項目以供以後重用，減輕垃圾收集器的壓力。
也就是說，它使構建高效、線程安全的空閒列表變得容易。但是，它並不適用於所有空閒列表。
一個 Pool 可以安全地同時被多個 goroutine 使用。


sync.Pool 是一個並發安全的緩存池，能夠並發且安全地存儲、獲取元素/對象。
常用於對象實例創建會佔用較多資源的場景。但是它不具有嚴格的緩存作用，因為Pool 中的元素/對象的釋放時機是隨機的。
作為緩存的一種姿勢，sync.Pool 能夠避免元素/對象的申請內存操作和初始化過程，以提高性能。
當然，這裡有個trick，釋放元素/對象的操作是直接將元素/對象放回池子，從而免去了真正釋放的操作。

另外，不考慮內存浪費和初始化消耗的情況下，“使用sync.Pool 管理多個對象”和“直接New 多個對象”兩者的區別在於：
後者會創建出更多的對象，並發高時會給GC 帶來非常大的負擔，進而影響整體程序的性能。
因為Go 申請內存是程序員觸發的，而回收卻是Go 內部runtime GC 回收器來執行的。即，使用sync.Pool 還可以減少GC 次數。

一句話總結：保存和復用臨時對象，減少內存分配，降低GC 壓力。

*/
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

/*
json 的反序列化在文本解析和網絡通信過程中非常常見，
當程序並發度非常高的情況下，短時間內需要創建大量的臨時對象。
而這些對像是都是分配在heap上的，會給GC 造成很大壓力，嚴重影響程序的性能。
*/
// func unmarsh() {
// 	stu := &Student{}
// 	json.Unmarshal(buf, stu)
// }

var studentPool = sync.Pool{
	//對像池中沒有對象時，將會調用New 函數創建。
	New: func() interface{} {
		return new(Student)
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		//Get()用於從對像池中獲取對象，因為返回值是interface{}，因此需要類型轉換。
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		//Put()則是在對象使用完畢後，返回對像池。
		studentPool.Put(stu)
	}
}

/*
在這個例子中，因為Student 結構體內存佔用較小，內存分配幾乎不耗時間。
而標準庫json 反序列化時利用了反射，效率是比較低的，佔據了大部分時間，
因此兩種方式最終的執行時間幾乎沒什麼變化。但是內存佔用差了一個數量級，
使用了 sync.Pool 後，內存佔用僅為未使用的234/5096 = 1/22，對GC 的影響就很大了。
*/

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

/*
這個例子創建了一個 bytes.Buffer 對像池，而且每次只執行一個簡單的 Write 操作，存粹的內存搬運工，耗時幾乎可以忽略。
而內存分配和回收的耗時佔比較多，因此對程序整體的性能影響更大。
*/

/*
也正因為，我們不能對sync.Pool 中保存的元素做任何假設，以下事情是都可以發生的：
	1.Pool 池裡的元素隨時可能釋放掉，釋放策略完全由runtime 內部管理；
	2.Get 獲取到的元素對象可能是剛創建的，也可能是之前創建好cache 住的。使用者無法區分；
	3.Pool 池裡面的元素個數你無法知道；
所以，只有的你的場景滿足以上的假定，才能正確的使用Pool 。
sync.Pool 本質用途是增加臨時對象的重用率，減少GC 負擔。重點：臨時對象。
所以說，像socket 這種帶狀態的，長期有效的資源是不適合Pool 的。
*/
