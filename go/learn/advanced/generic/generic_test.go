package generic

import (
	"fmt"
	"testing"
)

// 對map裡面進行sum操作(非泛型操作)
func TestSumWithoutGeneric(t *testing.T) {
	// 1.初始化包含int64元素的map
	ints := map[string]int64{
		"first":  50,
		"second": 50,
	}
	// 2.初始化包含float64元素的map
	floats := map[string]float64{
		"first":  45.44,
		"second": 54.55,
	}
	// 要定義int用和float用的func
	fmt.Printf("非泛型方式求和: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

}
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// 透過Interface來統一接口的方式
func TestSumInterface(t *testing.T) {
	// 1.初始化包含int64元素的map
	ints := map[string]int64{
		"first":  50,
		"second": 50,
	}
	// 2.初始化包含float64元素的map
	floats := map[string]float64{
		"first":  45.44,
		"second": 54.55,
	}

	fmt.Printf("使用Interface方式求和: %v and %v\n",
		SumInterface(ints),
		SumInterface(floats))

}

func SumInterface(m interface{}) (ret interface{}) {
	switch mm := m.(type) {
	case map[string]int64:
		var s int64
		for _, v := range mm {
			s += v
		}
		ret = s
		return ret

	case map[string]float64:
		var s float64
		for _, v := range mm {
			s += v
		}
		ret = s
		return ret
	}
	return nil

}

// 使用泛型來操作
func TestSumGeneric(t *testing.T) {
	// 1.初始化包含int64元素的map
	ints := map[string]int64{
		"first":  50,
		"second": 50,
	}
	// 2.初始化包含float64元素的map
	floats := map[string]float64{
		"first":  45.44,
		"second": 54.55,
	}

	fmt.Printf("使用Generic泛型的方式求和: %v and %v\n",
		SumGeneric(ints),
		SumGeneric(floats))
}

func SumGeneric[K comparable, v int64 | float64](m map[K]v) v {
	var s v
	for _, v := range m {
		s += v
	}
	return s
}

// 使用泛型加上接口約束來實現
type Number interface {
	int64 | float64 //聲明聯合類型
}

func SumGeneric2[K comparable, v Number](m map[K]v) v {
	var s v
	for _, v := range m {
		s += v
	}
	return s
}
func TestSumGeneric2(t *testing.T) {
	// 1.初始化包含int64元素的map
	ints := map[string]int64{
		"first":  50,
		"second": 50,
	}
	// 2.初始化包含float64元素的map
	floats := map[string]float64{
		"first":  45.44,
		"second": 54.55,
	}

	fmt.Printf("使用Generic泛型的方式求和: %v and %v\n",
		SumGeneric2(ints),
		SumGeneric2(floats))
}

// 對多參數的調用
type age interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

// 正常接口約束調用
func newGenrics[num age](s1 num) {
	val := float64(s1) + 1
	fmt.Println(val)
}

// 支援兩個參數進去，但這種寫法要注意 s1 ,s2 類型要一樣（都是int64 or 都是float64）
func total[num age](s1, s2 num) {
	val := float64(s1) + float64(s2)
	fmt.Println(val)
}

// 如果要接收不同類型的兩個參數，要這樣寫，宣告兩個age型態之後再轉給s1,s2
func summary[num1, num2 age](s1 num1, s2 num2) {
	val := float64(s1) + float64(s2)
	fmt.Println(val)
}

func TestMutiInput(t *testing.T) {
	var sum1 int64 = 10
	var sum2 float64 = 11.1
	newGenrics(sum1)
	newGenrics(sum2)

	// var sum3 int64 = 10 //會錯不行
	var sum3 float64 = 10
	var sum4 float64 = 11.1
	total(sum3, sum4)

	var sum5 int64 = 10
	var sum6 float64 = 11.1
	summary(sum5, sum6) // 可以成功相加
}

// 將bubblesort改為泛型
type NumberForSort interface {
	int | int32 | int64 | float32 | float64
}

func bubbleSort[n NumberForSort](array []n) []n {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func TestBubbleSort(t *testing.T) {
	n1 := []int{12, 21, 1, 3, 99, 5}
	n2 := []float64{12.1, 21.1, 1.1, 3.1, 99.1, 5.1}
	fmt.Println(bubbleSort(n1))
	fmt.Println(bubbleSort(n2))
}

/* Generic 的一些限制 */
// 1. 具有泛型的函數中，不支持 [類型聲明]
func GenericRestrict[K comparable, v Number](m map[K]v) v {
	/*
		type A struct {
			aa string
		}
		a := new(A)
		fmt.Println(a.aa)
	*/
	// 上面這樣聲明類型會報錯

	var s v
	for _, v := range m {
		s += v
	}
	return s
}

func TestGenericRestrict(t *testing.T) {
	ints := map[string]int64{
		"first":  50,
		"second": 50,
	}
	GenericRestrict(ints)
	// type declarations inside generic functions are not currently supported
}

// 2. 雖然允許泛型調用共同實現的接口方法，但是不允許調用結構體字段
// 接口約束
type Haha interface {
	A | B  //嵌入式的一個聯合的結構體類型聲明
	Haha() //方法聲明
}
type A struct {
	Aa string
	Cc string
}

type B struct {
	Bb string
	Cc string
}

// A和B 同時實現接口Haha的方法
func (a A) Haha() {
	fmt.Println(a.Aa)
}
func (b B) Haha() {
	fmt.Println(b.Bb)
}

func SayHaha[T Haha](t T) {
	t.Haha()
}

// 即使A和B 都有包含字段Cc，也是不允許這樣訪問的
// func SayHaha2[T A | B](t T) {
// 	t.Cc
// }

func TestSayHaha(t *testing.T) {
	a := A{"aa", "cc"}
	b := B{"bb", "cc"}
	SayHaha(a)
	SayHaha(b)
}
