package interface_test

import (
	"fmt"
	"strconv"
	"testing"
)

// 接口為非入侵性，實現不依賴於接口定義
// 所以接口的定義可以包含在接口使用者包內，不會有循環依賴
// 不像java打包的時候要做一個只有interface定義的pkg的

// 接口定義
type Programmer interface {
	WriteHelloWorld() string
}

// 接口實現
type GoProgrammer struct{}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	p := new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

/* 1. Go 透過 interface 實現了 duck-typing (鴨子類型)

“當看到一隻鳥走起來像鴨子、游泳起來像鴨子、叫起來也像鴨子，那麼這隻鳥就可以被稱為鴨子。”
意思就是: 一個東西究竟是不是鴨子，取決於它能不能滿足鴨子的工作。
duck typing 多見於動態語言,例如PHP,Python等.

在鴨子類型中，關注的不是對象的類型本身，而是它是如何使用的。
例如，在不使用鴨子類型的語言中，我們可以編寫一個函數，它接受一個類型為鴨的對象，並調用它的走和叫方法。
在使用鴨子類型的語言中，這樣的一個函數可以接受一個任意類型的對象，並調用它的走和叫方法。
對於動態類型的語言來說， 如果這些需要被調用的方法不存在，那麼將引發一個運行時錯誤。
而對於golang來說， 如果這些被調用的方法不存在，編譯時就會報錯。

*/
/* python code
def say_quack(duck)
    duck.quack()

class RealDcuk:
    def quack(self):
        print("quack quack")

class ToyDuck:
    def quack(self):
        print("squee squee")

duck = RealDuck()
say_quack(duck)

toyDuck = ToyDuck()
say_quack(duck)

可以看出動態語言的duck typing非常靈活方便，類型的檢測和使用不依賴於編譯器的靜態檢測，而是依賴文檔、清晰的代碼和測試來確保正確使用。
這樣其實是犧牲了安全性來換取靈活性。
假設你沒有認真看文檔，不知道say_quack方法的duck參數是需要quack方法， 你編寫了一個Dog類，它只有一個run方法， 你把它的對象當成參數給say_quack編譯時也是不會報錯的。
只有在運行時才會報錯， 這樣就存在很大的安全隱患。有沒有一種折中（tradeoff）， 兼顧這種duck typing的靈活性和靜態檢測的安全性呢？
*/

// go語言接口的隱式實現

//定義一個鴨子接口
//Go 接口是一組方法的集合，可以理解為抽象的類型。它提供了一種非侵入式的接口。任何類型，只要實現了該接口中方法集，那麼就屬於這個類型。
type Duck interface {
	Gaga()
}

// 任何擁有Gaga方法的類型， 都隱式地（implicitly）實現了Duck接口， 並能當做Duck接口使用

// 假設現在有一個真鴨子類型
type RealDuck struct{}

// 真鴨子聲明方法-滿足鴨子會嘎嘎叫的特性
func (rd RealDuck) Gaga() {
	fmt.Println("gaga gaga")
}

// 假設現在有一個玩具鴨類型
type ToyDuck struct{}

// 玩具鴨聲明方法-滿足鴨子會嘎嘎叫的特性
func (td ToyDuck) Gaga() {
	fmt.Println("kaka kaka")
}

//要調用的函數 - 負責執行鴨子能做的事情,注意這裡的參數,有類型限制為Duck接口
func DuckSay(d Duck) {
	d.Gaga()
}

func TestQuack(t *testing.T) {
	fmt.Println("duck typing")
	// 實例化對象
	var rd RealDuck //真鴨子
	var td ToyDuck  //玩具鴨

	//調用方法
	rd.Gaga()
	td.Gaga()
	DuckSay(rd) // 因為真鴨子實現了所有鴨子的函數，所以可以這麼用
	DuckSay(td) // 因為假鴨子實現了所有鴨子的函數，所以可以這麼用
}

/*
如果你有一個Dog類型， 它沒有Gaga方法， 當你用它做DuckSay參數時， 編譯時就會報錯。
*/
type Dog struct{}

func TestDog(t *testing.T) {
	// dog := Dog{}
	// DuckSay(dog)

	//cannot use dog (variable of type Dog) as type Duck in argument to sayQuack:
	// Dog does not implement Duck (missing quack method)
}

/*
另外來說， 如果接口使用者定義了一個新的接口也擁有Gaga方法， 那上面的RealDuck和ToyDuck也可以當做新的接口來使用。
這樣就達到了一個靈活性和安全性的平衡。因為go對接口的實現是隱式的， 所以它的接口類型在使用之前是不固定的，
它可以靈活的變成各種接口類型，只要它滿足使用者的對接口的要求。
又因為使用者使用接口時在編譯時就對接口實現者有沒有滿足接口需求進行了檢測，所以又兼顧了安全性。
*/

/* 2.	---空interface---
空 interface(interface{})不包含任何的 method，正因為如此，所有的型別都實現了空 interface。
空 interface 對於描述起不到任何的作用(因為它不包含任何的 method），但是空 interface 在我們需要儲存任意型別的數值的時候相當有用，
因為它可以儲存任意型別的數值。有點類似於 C 語言的 void*型別。
*/

func TestEmptyInterface(t *testing.T) {
	var a interface{}
	var i int = 5
	s := "Hello world"
	// a 可以儲存任意型別的數值
	a = i
	t.Logf("%d", a)
	a = s
	t.Logf("%s", a)
}

// 一個函式把 interface{} 作為參數，那麼他可以接受任意型別的值作為參數，如果一個函式回傳 interface{}，那麼也就可以回傳任意型別的值。

/* 3.	---interface變數儲存的型別---
我們知道 interface 的變數裡面可以儲存任意型別的數值(該型別實現了 interface)。那麼我們怎麼反向知道這個變數裡面實際儲存了的是哪個型別的物件呢？
目前常用的有兩種方法：
3.1. Comma-ok 斷言
Go 語言裡面有一個語法，可以直接判斷是否是該型別的變數： value, ok = element.(T)，
這裡 value 就是變數的值，ok 是一個 bool 型別，element 是 interface 變數，T 是斷言的型別。
如果 element 裡面確實儲存了 T 型別的數值，那麼 ok 回傳 true，否則回傳 false。
*/

type Element interface{}
type List []Element

type Person struct {
	Name string
	Age  int
}

/*
fmt.Println 是我們常用的一個函式，但是你是否注意到它可以接受任意型別的資料。
開啟 fmt 的原始碼檔案，你會看到這樣一個定義:
type Stringer interface {
     String() string
}
也就是說，任何實現了 String 方法的型別都能作為參數被 fmt.Println 呼叫
*/

//定義了 String 方法，實現了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.Name + " - age: " + strconv.Itoa(p.Age) + " years)"
}

func TestCommaOk(t *testing.T) {
	list := make(List, 3)
	list[0] = 1       // int
	list[1] = "Hello" // string
	list[2] = Person{Name: "Raymond", Age: 25}

	for index, element := range list {
		if value, ok := element.(int); ok {
			t.Logf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			t.Logf("list[%d] is an string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			t.Logf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			t.Logf("list[%d] is of a different type\n", index)
		}
	}
}

// 3.2	switch 測試
func TestTypeAssertionsSwitch(t *testing.T) {
	list := make(List, 3)
	list[0] = 1       // int
	list[1] = "Hello" // string
	list[2] = Person{Name: "Raymond", Age: 25}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			t.Logf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			t.Logf("list[%d] is an string and its value is %s\n", index, value)
		case Person:
			t.Logf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			t.Logf("list[%d] is of a different type\n", index)
		}
	}
}

// 3.3 傳值還是傳指針？

type notifier1 interface {
	notifyPointer()
}
type notifier2 interface {
	notifyValue()
}

func sendNotificationPointer(n notifier1) {
	n.notifyPointer()
}

func sendNotificationValue(n notifier2) {
	n.notifyValue()
}

type user struct {
	name  string
	email string
}

func (u *user) notifyPointer() { // 這裡的接收者是指針
	fmt.Println(u.name, u.email)
}
func (u user) notifyValue() { // 這裡的接收者是值
	fmt.Println(u.name, u.email)
}
func TestNotify(t *testing.T) {
	u := user{"Raymond", "raymond@gmail.com"}
	// 接收者是指針時
	// sendNotificationPointer(u)  // 傳遞的是值 會編譯失敗
	sendNotificationPointer(&u) // 傳遞的是指针，會成功

	// 接收者是值時，無論傳遞值還是指針都會通過
	sendNotificationValue(u)
	sendNotificationValue(&u)

}

// 3.4 什麼時候接收者用值or指針？ 先單看Struct結構體指針類型方法

type myStruct struct {
	Name string
}

//定義這個結構的改名方法
func (m myStruct) ChangeName(newName string) {
	m.Name = newName
}

func TestChangeName(t *testing.T) {
	//創建這個結構體變量
	mystruct := myStruct{
		Name: "raymond",
	}

	//調用改名函數
	mystruct.ChangeName("RAYMOND")

	//不會變
	fmt.Println(mystruct.Name)
}

/*這樣的方法不會改掉結構體變量內的字段值。
就算是結構體方法，如果不使用指針，方法內還是傳遞結構體的值。*/

// 當要改變結構體內的值時，要使用指針定義結構
func (m *myStruct) ChangeName2(newName string) {
	m.Name = newName
}
func TestChangeName2(t *testing.T) {
	//創建這個結構體變量
	mystruct := myStruct{
		Name: "raymond",
	}

	//調用改名函數
	/*
		***當使用指針類型定義方法後，結構體類型的變量調用方法時，
		   會自動取得該結構體的指針類型並傳入方法。***
	*/
	mystruct.ChangeName2("RAYMOND")

	//變了
	fmt.Println(mystruct.Name)
}

// 3.5 再看看指針類型的接口實現

// 定義一個接口
type myInterfaceI interface {
	ChangeNameI(string)
	SayMyNameI()
}

type myStructI struct {
	Name string
}

// 定義接收指針的改名方法
func (m *myStructI) ChangeNameI(newName string) {
	m.Name = newName
}

// 定義接收變量的方法
func (m myStructI) SayMyNameI() {
	fmt.Println(m.Name)
}

// 一個使用接口作為參數的函數
func SetName(s myInterfaceI, name string) {
	s.ChangeNameI(name)
}

func TestChangeNameI(t *testing.T) {
	// 檢核myStructI有沒有實現myInterfaceI （interface_check_test.go再詳細講）
	var _ myInterfaceI = (*myStructI)(nil)
	//創建這個結構體變量
	mystruct := myStructI{
		Name: "raymond",
	}
	// 調用函數，無法編譯通過
	// SetName(mystruct, "RAYMOND")

	/*
		cannot use mystruct (type myStructI) as type myInterfaceI in argument to SetNameI:
		myStructI does not implement myInterfaceI (ChangeNameI method has pointer receiver)
	*/
	// myStructI類型沒有實現接口方法ChangeNameI，
	// 也就是說func (m *myStructI) ChangeName(newName string) 並不算實現了接口，
	// 因為它是*myStructI類型實現的，而不是myStructI。

	// 直接用結構體卻是可以調用的
	mystruct.ChangeNameI("RAYMOND")
	mystruct.SayMyNameI()

	//在調用SetName時，得用&mystruct 替代mystruct
	SetName(&mystruct, "RAYMOND!")
	mystruct.SayMyNameI()

}

/*
為什麼結構體類型實現的接口該結構體的指針類型也算實現了，而指針類型實現的接口，不算是該結構體實現了接口呢？

**
   原因是，結構體類型定義的方法可以被該結構體的指針類型調用；
   而結構體類型調用該指針類型的方法時是被轉換成指針，不是直接調用。
**

所以，&mystruct直接實現了接口定義的ChangeNameI和SayMyNameI兩個方法，
而mystruct只能實現了SayMyNameI，mystruct調用ChangeNameI方法其實轉換成指針類型後調用的，不算實現了接口。
*/
