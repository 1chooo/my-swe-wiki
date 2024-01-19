package err_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

// os.Exit退出時不會調用defer指定的函數，也不會輸出調用訊息
func TestOSExit(t *testing.T) {
	defer func() {
		fmt.Println("defer-> Finally")
	}()
	fmt.Println("Start os.Exit()")
	os.Exit(-1)

}

// panic用於不能恢復的錯誤，退出前會執行defer指定的內容
func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("defer-> Finally")
	}()
	fmt.Println("Start panic")
	panic(errors.New("Something wrong!"))

}

// Recover

// In java
// try{
// 	...
// }catch(Throwable t){

// }
// Go 可以在panic之前調用defer來實現recover
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start panic")
	panic(errors.New("Something wrong!"))

}

// 常見的"錯誤"使用panic
// 什麼也沒做只是寫個log，可能會導致殭屍服務進程，導致health check失效
// 不如採用let it crash來重起
// defer func() {
// 	if err := recover(); err != nil {
// 		log.Error("recovered panic", err)
// 	}
// }()
