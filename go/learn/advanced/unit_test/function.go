package unittest

import (
	"bytes"
	"strings"
)

func Square(i int) int {
	return i * i
}

func ConcatStringByAdd(strings []string) string {
	ret := ""
	for _, elem := range strings {
		ret += elem
	}
	return ret
}

func ConcatStringByBytesBuffer(strings []string) string {
	var buf bytes.Buffer
	for _, elem := range strings {
		buf.WriteString(elem)
	}
	return buf.String()

}

func ConcatStringByStringsBuilder(string_slice []string) string {
	var str strings.Builder
	for _, elem := range string_slice {
		str.WriteString(elem)
	}
	return str.String()
}

/*

由於string是不可修改的，所以在使用“+”進行拼接字符串
，每次都會產生申請空間，拼接，複製等操作，數據量大的情況下非常消耗資源和性能。
而採用Buffer等方式，都是預先計算拼接字符串數組的總長度（如果可以知道長度），
申請空間，底層是slice數組，可以以append的形式向後進行追加。
最後在轉換為字符串。這申請了不斷申請空間的操作，也減少了空間的使用和拷貝的次數，自然性能也高不少。


如果需要拼接多次，應使用strings.Builder，最小化內存拷貝次數。

strings.Builder和bytes.Buffer底層都是使用[]byte實現的， 但是性能測試的結果顯示（https://gist.github.com/bwangelme/37facf96621fef19e2e70bce7a7b84），
 執行String()函數的時候，strings.Builder卻比bytes.Buffer快很多。

區別就在於 bytes.Buffer 是重新申請了一塊空間，存放生成的string變量，
Buffer的string是一種強轉，我們知道在強轉的時候是需要進行申請空間，並拷貝的。

而Builder只是指針的轉換，strings.Builder直接將底層的[]byte轉換成了string類型返回了回來，去掉了申請空間的操作。



在bytes.Buffer中也說明了，如果想更有效率地( efficiently )構建字符串，請使用strings.Builder類型
// String returns the contents of the unread portion of the buffer
// as a string. If the Buffer is a nil pointer, it returns "<nil>".
//
// To build strings more efficiently, see the strings.Builder type.
func (b *Buffer) String() string {
	if b == nil {
		// Special case, useful in debugging.
		return "<nil>"
	}
	return string(b.buf[b.off:])
}

*/
