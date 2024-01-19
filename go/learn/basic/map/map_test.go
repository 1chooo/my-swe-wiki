package map_test

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"unsafe"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	t.Log(m1[2])
	t.Logf("len m1: %d", len(m1))

	m2 := map[int]string{}
	m2[4] = "GGGG"
	t.Log(m2[4])
	t.Log(m2)
	t.Logf("len m2: %d", len(m2))

	// 用make先定義好capacity的話，就不用每次自增長就得要記憶體然後複製一份，進而提高性能
	m3 := make(map[int]string, 10)
	t.Logf("len m3: %d", len(m3))

}

func TestInitMap2(t *testing.T) {
	// 初始化方式1：直接声明
	// var m1 map[string]int
	// m1["a"] = 1
	// t.Log(m1, unsafe.Sizeof(m1))
	// panic: assignment to entry in nil map
	// 向 map 写入要非常小心，因为向未初始化的 map（值为 nil）写入会引发 panic，所以向 map 写入时需先进行判空操作
	// 初始化方式2：使用字面量
	m2 := map[string]int{}
	m2["a"] = 2
	t.Log(m2, unsafe.Sizeof(m2)) // map[a:2] 8
	// 初始化方式3：使用make创建
	m3 := make(map[string]int)
	m3["a"] = 3
	t.Log(m3, unsafe.Sizeof(m3)) // map[a:3] 8
}

func TestAccessNotExistKey(t *testing.T) {

	m1 := map[int]int{}
	t.Log(m1[1]) // 不存在的key 但return 0 不是nil
	m1[2] = 0
	t.Log(m1[2]) // 會跟存在 但值為0的狀況搞混
	// 在Java會有空指針異常，但go不會有問題

	//所以大多要這樣寫
	if v, ok := m1[3]; ok {
		t.Log("key 3 exist", v, ok)
	} else {
		t.Log("key 3 not exist", v, ok)
	}
}

// map travel
func TestTravelMap(t *testing.T) {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	// 但是map是無序的
	for key, value := range m1 {
		t.Log(key, value)
	}
}

// 2. 為什麼遍歷map是無序的？
// 3. 如何實現有序遍歷map？
/*遍歷順序隨機*/
// map 在沒有被修改的情況下，使用range 多次遍歷map 時輸出的key 和value 的順序可能不同。
// 這是Go 語言的設計者們有意為之，在每次range 時的順序被隨機化，旨在提示開發者們，Go 底層實現並不保證map 遍歷順序穩定，請大家不要依賴range 遍歷結果順序。
//map 本身是無序的(問題9 下面會講)，且遍歷時順序還會被隨機化，如果想順序遍歷map，需要對map key 先排序，再按照key 的順序遍歷map
func TestMapRange(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	t.Log("first range:")
	// 默認無序
	for i, v := range m {
		t.Logf("m[%v]=%v ", i, v)
	}
	t.Log("\nsecond range:")
	for i, v := range m {
		t.Logf("m[%v]=%v ", i, v)
	}

	// 實現有序排列
	var sl []int
	// 把key取出放到slice
	for k := range m {
		sl = append(sl, k)
	}
	// 排序slice
	sort.Ints(sl)
	// 以slice中的 key 顺序遍歷 map 就是有序的了
	for _, k := range sl {
		t.Log(k, m[k])
	}
}

/*共享存儲空間*/
func TestMapShareMemory(t *testing.T) {
	//map底層數據結構是通過指針指向實際的元素存儲空間，這種情況下，對其中一個map的更改，會影響到其他map
	m1 := map[string]int{}
	m2 := m1
	m1["a"] = 1
	t.Log(m1, len(m1))
	// map[a:1] 1
	t.Log(m2, len(m2))
	// map[a:1]
}

// 4. 為什麼Go map是非線程安全的？
/*非線程安全*/
// map默認是並發不安全的，原因如下：
// Go 官方在經過了長時間的討論後，認為Go map 更應適配典型使用場景（不需要從多個goroutine 中進行安全訪問），
// 而不是為了小部分情況（並發訪問），導致大部分程序付出加鎖代價（性能），決定了不支持。

// 場景: 2個協程同時讀和寫，以下程序會出現致命錯誤：fatal error: concurrent map writes
func TestConcurrentMapFatal(t *testing.T) {
	m := make(map[int]int)
	go func() {
		//開一個協程寫map
		for i := 0; i < 10000; i++ {

			m[i] = i
		}
	}()

	go func() {
		//開一個協程讀map
		for i := 0; i < 10000; i++ {

			fmt.Println(m[i])
		}
	}()

	//time.Sleep(time.Second * 20)
	for {

	}
}

// 5. 線程安全的map如何實現?
// 如果想實現map線程安全，有兩種方式：
// 方式一：使用讀寫鎖 map+ sync.RWMutex
func BenchmarkMapConcurrencySafeByMutex(b *testing.B) {
	var lock sync.Mutex //互斥锁
	m := make(map[int]int, 0)
	var wg sync.WaitGroup
	//b.N 是 go 語言內建提供的循環，根據一秒鐘的時間計算
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			m[i] = i
		}(i)
	}
	wg.Wait()
	b.Log(len(m), b.N)
}

// 方式二：使用golang提供的sync.Map
// sync.map是用讀寫分離實現的，其思想是空間換時間。和map+RWLock的實現方式相比，它做了一些優化：
// 可以無鎖訪問read map，而且會優先操作read map，倘若只操作read map就可以滿足要求(增刪改查遍歷)，
// 那就不用去操作write map(它的讀寫都要加鎖)，所以在某些特定場景中它發生鎖競爭的頻率會遠遠小於map+RWLock的實現方式。
// sync.map 類型，針對以下場景進行了性能優化：
// 當一個給定的鍵的條目只被寫入一次但被多次讀取時。例如在僅會增長的緩存中，就會有這種業務場景。
// 當多個 goroutines 讀取、寫入和覆蓋不相干的鍵集合的條目時。
func BenchmarkMapConcurrencySafeBySyncMap(b *testing.B) {
	var m sync.Map
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}
	wg.Wait()
	b.Log(b.N)
}

// 6. Go sync.map 和原生map 誰的性能好，為什麼？
// sync.map分析：https://www.readfog.com/a/1636088860119240704
/*
在寫入元素上，最慢的是 sync.map 類型，其次是原生 map + 互斥鎖（Mutex），最快的是原生 map + 讀寫鎖（RwMutex）
在查找元素上，最慢的是原生 map + 互斥鎖，其次是原生 map + 讀寫鎖。最快的是 sync.map 類型。
在刪除元素上，最慢的是原生 map + 讀寫鎖，其次是原生 map + 互斥鎖，最快的是 sync.map 類型

根據上述的壓測結果，我們可以得出 sync.Map 類型：

	在讀和刪場景上的性能是最佳的，領先一倍有多。

	在寫入場景上的性能非常差，落後原生 map + 鎖整整有一倍之多。

因此在實際的業務場景中。假設是讀多寫少的場景，會更建議使用 sync.Map 類型。

但若是那種寫多的場景，例如多 goroutine 批量的循環寫入，那就建議另闢途徑了，性能不忍直視（無性能要求另當別論）。
*/

/*
一些深入問題
// https://juejin.cn/post/7056290831182856205#heading-8
1. --map的底層實現原理--
Go中的map是一個指針，佔用8個字節，指向hmap結構體
每個map的底層結構是hmap，hmap包含很多個結構為bmap的bucket數組。每個bucket底層都採用linked list結構。

// A header for a Go map.
type hmap struct {
    count     int
    // 代表哈希表中的元素个数，调用len(map)时，返回的就是该字段值。
    flags     uint8
    // 状态标志，下文常量中会解释四种状态位含义。
    B         uint8
    // buckets（桶）的对数log_2
    // 如果B=5，则buckets数组的长度 = 2^5=32，意味着有32个桶
    noverflow uint16
    // 溢出桶的大概数量
    hash0     uint32
    // 哈希种子

    buckets    unsafe.Pointer
    // 指向buckets数组的指针，数组大小为2^B，如果元素个数为0，它为nil。
    oldbuckets unsafe.Pointer
    // 如果发生扩容，oldbuckets是指向老的buckets数组的指针，老的buckets数组大小是新的buckets的1/2;非扩容状态下，它为nil。
    nevacuate  uintptr
    // 表示扩容进度，小于此地址的buckets代表已搬迁完成。

    extra *mapextra
    // 这个字段是为了优化GC扫描而设计的。当key和value均不包含指针，并且都可以inline时使用。extra是指向mapextra类型的指针。
 }

bmap就是我們常說的“桶”，一個桶裡面會最多裝8 個key，這些key 之所以會落入同一個桶，是因為它們經過哈希計算後，
哈希結果是“一類”的，關於key的定位我們在map的查詢和插入中詳細說明。
在桶內，又會根據key 計算出來的hash 值的高8 位來決定key 到底落入桶內的哪個位置（一個桶內最多有8個位置)。

// A bucket for a Go map.
type bmap struct {
    tophash [bucketCnt]uint8
    // len为8的数组
    // 用来快速定位key是否在这个bmap中
    // 桶的槽位数组，一个桶最多8个槽位，如果key所在的槽位在tophash中，则代表该key在这个桶中
}
//底层定义的常量
const (
    bucketCntBits = 3
    bucketCnt     = 1 << bucketCntBits
    // 一个桶最多8个位置
）

但这只是表面(src/runtime/hashmap.go)的结构，编译期间会给它加料，动态地创建一个新的结构：

type bmap struct {
  topbits  [8]uint8
  keys     [8]keytype
  values   [8]valuetype
  pad      uintptr
  overflow uintptr
  // 溢出桶
}





7. --為什麼Go map 的負載因子是6.5？--
8. --map擴容策略是什麼?--
map 擴容的時機：在向map 插入新key 的時候，會進行條件檢測，符合下面這2 個條件，就會觸發擴容：
1、裝載因子超過閾值
源碼裡定義的閾值是6.5 (loadFactorNum/loadFactorDen)，是經過測試後取出的一個比較合理的因子
我們知道，每個bucket 有8 個空位，在沒有溢出，且所有的桶都裝滿了的情況下，裝載因子算出來的結果是8。
因此當裝載因子超過6.5 時，表明很多bucket 都快要裝滿了，查找效率和插入效率都變低了。在這個時候進行擴容是有必要的。
對於條件1，元素太多，而bucket 數量太少，很簡單：將B 加1，bucket 最大數量( 2^B)直接變成原來bucket 數量的2 倍。
於是，就有新老bucket 了。注意，這時候元素都在老bucket 裡，還沒遷移到新的bucket 來。新bucket 只是最大數量變為原來最大數量的2 倍( 2^B * 2) 。
2、overflow 的bucket 數量過多
在裝載因子比較小的情況下，這時候map 的查找和插入效率也很低，而第1 點識別不出來這種情況。
表面現象就是計算裝載因子的分子比較小，即map 裡元素總數少，但是bucket 數量多（真實分配的bucket 數量多，包括大量的overflow bucket）
不難想像造成這種情況的原因：不停地插入、刪除元素。先插入很多元素，導致創建了很多bucket，但是裝載因子達不到第1 點的臨界值，未觸發擴容來緩解這種情況。
之後，刪除元素降低元素總數量，再插入很多元素，導致創建很多的overflow bucket，但就是不會觸發第1 點的規定，你能拿我怎麼辦？
overflow bucket 數量太多，導致key 會很分散，查找插入效率低得嚇人，因此出台第2 點規定。
這就像是一座空城，房子很多，但是住戶很少，都分散了，找起人來很困難
對於條件2，其實元素沒那麼多，但是overflow bucket 數特別多，說明很多bucket 都沒裝滿。
解決辦法就是開闢一個新bucket 空間，將老bucket 中的元素移動到新bucket，使得同一個bucket 中的key 排列地更緊密。
這樣，原來，在overflow bucket 中的key 可以移動到bucket 中來。結果是節省空間，提高bucket 利用率，map 的查找和插入效率自然就會提升。

由於map 擴容需要將原有的key/value 重新搬遷到新的內存地址，如果有大量的key/value 需要搬遷，會非常影響性能。
因此Go map 的擴容採取了一種稱為“漸進式”的方式，原有的key 並不會一次性搬遷完畢，每次最多只會搬遷2 個bucket。
上面說的 hashGrow() 函數實際上並沒有真正地“搬遷”，它只是分配好了新的buckets，並將老的buckets 掛到了oldbuckets 字段上。
真正搬遷buckets 的動作在 growWork() 函數中，而調用 growWork() 函數的動作是在mapassign 和mapdelete 函數中。
也就是插入或修改、刪除key 的時候，都會嘗試進行搬遷buckets 的工作。先檢查oldbuckets 是否搬遷完畢，具體來說就是檢查oldbuckets 是否為nil。

如果未遷移完畢，賦值/刪除的時候，擴容完畢後（預分配內存），不會馬上就進行遷移。而是採取增量擴容的方式，當有訪問到具體bukcet 時，才會逐漸的進行遷移（將oldbucket 遷移到bucket）
*/

/*
9.--為什麼map 是無序的？--
遍歷的過程，就是按順序遍歷bucket，同時按順序遍歷bucket 中的key。
上面說過，map遍歷是無序的，如果想實現有序遍歷，可以先對key進行排序

如果發生過遷移，key 的位置發生了重大的變化，有些key 飛上高枝，有些key 則原地不動。這樣，遍歷map 的結果就不可能按原來的順序了。
如果就一個寫死的map，不會向map 進行插入刪除的操作，按理說每次遍歷這樣的map 都會返回一個固定順序的key/value 序列吧。
但是Go 杜絕了這種做法（加入隨機數在range），因為這樣會給新手程序員帶來誤解，以為這是一定會發生的事情，在某些情況下，可能會釀成大錯。
Go 做得更絕，當我們在遍歷map 時，並不是固定地從0 號bucket 開始遍歷，每次都是從一個**隨機值序號的bucket開始遍歷，
並且是從這個bucket 的一個隨機序號的cell **開始遍歷。這樣，即使你是一個寫死的map，僅僅只是遍歷它，也不太可能會返回一個固定序列的key/value 對了。


*/

/* --總結--
map是引用類型
map遍歷是無序的
map是非線程安全的
map的哈希衝突解決方式是鍊錶法
map的擴容不是一定會新增空間，也有可能是只是做了內存整理
map的遷移是逐步進行的，在每次賦值時，會做至少一次遷移工作
map中刪除key，有可能導致出現很多空的kv，這會導致遷移操作，如果可以避免，盡量避免
*/
