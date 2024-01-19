package series

import "fmt"

func GetFibonacci(n int) ([]int, error) {

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 小寫 不對外公開 只能內部使用
// func square(n int) int {
// 	return n * n
// }

// init方法
// 1.	在main被執行前，所有依賴的package的init方法都會被執行
// 2.	不同包的init函數按照包導入的依賴關係決定執行順序 （go會自己處理這個依賴順序）
// 3.	每個包可以有多個init函數
// 4.	包的每個源文件也可以有多個init函數，這點比較特殊

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
