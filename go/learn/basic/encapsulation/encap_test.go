package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
Is Go an object-oriented language?
Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy.
The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general.
There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing.
Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data,
even built-in types such as plain, “unboxed” integers. They are not restricted to structs (classes).

Also, the lack of a type hierarchy makes “objects” in Go feel much more lightweight than in languages such as C++ or Java.
*/
// Go官方說
// Go 是面向對象的語言嗎？
// 是和不是。儘管 Go 有類型和方法並允許面向對象的編程風格，但沒有類型層次結構。Go 中的“interface”概念提供了一種不同的方法，
// 我們認為這種方法易於使用，並且在某些方面更通用。還有一些方法可以將類型嵌入到其他類型中，以提供與子類化類似但不完全相同的東西。
// 此外，Go 中的方法比 C++ 或 Java 中的方法更通用：它們可以為任何類型的數據定義，甚至是內置類型，例如普通的“unboxed”整數。它們不限於structs（classes）。

// 此外，沒有類型層次結構使得 Go 中的“objects”感覺比 C++ 或 Java 等語言更輕量。

// go不支持繼承 接口利用duck type
// e2 := new(Employee)//注意這裡返回的引用/指針相當於e:=&Employee{}
//與其他語言的差異：通過實例的指針訪問成員不需要使用->

type Employee struct {
	ID   string //不像java要逗號，直接換行就好
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e1 := Employee{"0", "Bob", 20}
	e2 := Employee{Name: "Mike", Age: 30}
	e3 := new(Employee) //回傳指針，注意這裡返回的引用(指針)相當於e3:=&Employee{}
	e3.ID = "3"
	e3.Name = "Raymond"
	e3.Age = 25
	t.Log(e1)
	t.Log(e2)
	t.Log(e2.ID)
	t.Log(e3)
	t.Logf("e2 is %T", e2)
	t.Logf("e3 is %T", e3)
}

func TestStructOperations(t *testing.T) {
	e1 := Employee{"0", "Bob", 20}
	fmt.Printf(" Origin Address is %x \n", unsafe.Pointer(&e1.Name))
	t.Log(e1.ToString())

	e2 := &Employee{Name: "Mike", Age: 30} //指向實例的指針
	fmt.Printf(" Origin Address is %x \n", unsafe.Pointer(&e2.Name))
	t.Log(e2.ToString()) //指針也可以直接調用，不像java需要使用->符號 可以直接用點. 他會自己判斷，雖然是指針但是一樣可以訪問在結構上定義的方法
}

//第一種定義方法在實例對應方法被調用時，實例成員會進行值的複製，可以發現到ToString() Address變了
func (e Employee) ToString() string {
	fmt.Printf(" Address in ToString() is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

//通常情況下為了避免內存拷貝，我們使用第二種定義方法

type Employee2 struct {
	ID   string
	Name string
	Age  int
}

//第二種定義方法 接收指向Employee2的指針
func (e *Employee2) ToString() string {
	fmt.Printf(" Address in ToString() is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func TestStructOperations2(t *testing.T) {
	e1 := Employee2{"0", "Bob", 20}
	fmt.Printf(" Origin Address is %x \n", unsafe.Pointer(&e1.Name))
	t.Log(e1.ToString()) //一樣結構體可以調用到 指向結構體的指針 的方法

	e2 := &Employee2{Name: "Mike", Age: 30}
	fmt.Printf(" Origin Address is %x \n", unsafe.Pointer(&e2.Name))
	t.Log(e2.ToString())
}
