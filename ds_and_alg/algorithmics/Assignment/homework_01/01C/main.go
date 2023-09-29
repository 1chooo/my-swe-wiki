package main

import (
	"fmt"
	"math"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)
	fmt.Scanln() // Consume the newline character

	for i := 0; i < runTimes; i++ {
		var inStr string
		fmt.Scanln(&inStr)

		columnNumber := NewColumnNumber(inStr)
		columnNumber.ShowAns()
	}
}

type ColumnNumber struct {
	inStr string
}

func NewColumnNumber(inStr string) *ColumnNumber {
	return &ColumnNumber{inStr: inStr}
}

func (cn *ColumnNumber) ShowAns() {
	len := len(cn.inStr)
	var outNum, sum int
	var out rune
	for i := 0; i < len; i++ {
		out = rune(cn.inStr[i])
		// Convert to ASCII code
		outNum = int(out - 'A' + 1)
		sum += outNum * int(math.Pow(26, float64(len-1-i)))
	}

	fmt.Println(sum)
}
