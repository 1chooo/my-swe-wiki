package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化 默認0
	s = "hello"
	t.Log(len(s))
	// s[1]='3' //string是不可變的byte slice
	s = "\xE6\xB2\xB9" //可以儲存任何二進制數據 s = "\xe5\x8a\xa0"
	// s="\xE4\xBA\xB5\xFF" //亂碼也可以
	t.Log(s) //長度為3

	s = "油"
	t.Log(len(s)) // 3 是byte數

	c := []rune(s)
	t.Log(len(c)) //1 個字
	// t.Log("rune size:",unsafe.Sizeof(c[0]))
	t.Logf("油 unicode %x", c[0]) //6cb9
	t.Logf("油 UTF8 %x", s)       //e6b2b9
	/* 字符 			"油" */
	/*	Unicode 		0x6CB9*/
	/*	UTF-8 			0xE6B2B9*/
	/*	string/[]byte 	[0xE6,0xB2,0xB9]*/
}

func TestStringtoRune(t *testing.T) {
	s := "加油各位"
	for _, c := range s {
		// c會自動轉rune 不是byte
		t.Logf("%[1]c %[1]d", c) // %[1]代表都是和第一個參數對應 （都和c對應）也就是第一個參數以%c跟%d格式化的意思
		// t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFn(t *testing.T) {
	//分割
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	//連接
	t.Log(strings.Join(parts, "-")) //A-B-C
}

func TestConv(t *testing.T) {
	// string跟int轉換
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
