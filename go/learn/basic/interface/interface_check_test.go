package interface_test

import (
	"fmt"
	"testing"
)

/*
	常見問題 var _ I = (*T)(nil) 的作用？
	作用：用簡單的語法，檢查這個結構是否實現了我這個接口
	理解：可以把=左右兩邊分開
	左邊：var _ I 等價於我們平時用的var variable type
	右邊：(* T)(nil)等價於var variable *T nil
*/
type I1 interface{}

type I2 interface {
	say()
}

type TestStruct struct{}

func TestStructImpInterface(t *testing.T) {
	// 看TestStruct有沒有實現I1這個interface
	var _ I1 = (*TestStruct)(nil) // 將nil轉為*TestStruct後判斷是否實現I1接口
	// var _ I2 = (*TestStruct)(nil) // 編譯會失敗,因為TestStruct沒有實現say方法,除非外面新增func (*TestStruct) say() {}

	// 繁瑣的寫法
	var testStruct *TestStruct = nil
	var i I1 = testStruct // Verify that *T implements I.
	fmt.Println(i)

	/*這些寫法也都可以 下面再詳細講區別*/
	var _ I1 = TestStruct{}
	var _ I1 = &TestStruct{}
	var _ I1 = new(TestStruct)

	var a T2
	a.Sing()
}

type I interface {
	Sing()
}

type T struct {
}

func (t T) Sing() {
}

type T2 struct {
}

func (t *T2) Sing() {
}

// 編譯通過
var _ I = T{}

// 編譯通過
var _ I = &T{}

// 編譯失敗
// var _ I = T2{}

// 編譯通過
var _ I = &T2{}

/*
T實現了 Sing 方法，*T2實現了 Sing 方法。

我們都知道，Go 語言中是按值傳遞的。

那對於 T2 來說，調用 Sing 方法時，
copy 一個副本，然後取地址，通過這個地址是找不到原始調用的那個結構體的，
但是receiver 是個指針，表示此次調用是需要改變調用者內部變量的，
很明顯，以 T2 類型調用無法完達到這個目的，所以這裡是需要報錯的。
而以 &T2 調用 Sing 方法，則可以，因此不報錯。

而對於 T 來說，不管是否有指針調用，都不會報錯，實際上，Go 語言會自動實現 *T 的 Sing 方法。
*/
