package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	var runTimes int
	fmt.Scanln(&runTimes)
	ansList := make([]int, 0, runTimes)

	for i := 0; i < runTimes; i++ {
		down, right := readInput()
		total := down + right
		temp1 := big.NewInt(1)
		temp2 := big.NewInt(1)
		temp3 := big.NewInt(1)

		if right == 0 && down == 0 {
			ansList = append(ansList, 0)
		} else {
			for i := int64(total); i > 1; i-- {
				temp1.Mul(temp1, big.NewInt(i))
			}
			for i := int64(right); i > 1; i-- {
				temp2.Mul(temp2, big.NewInt(i))
			}
			for i := int64(down); i > 1; i-- {
				temp3.Mul(temp3, big.NewInt(i))
			}

			ans := new(big.Int).Div(temp1, new(big.Int).Mul(temp2, temp3))
			ansList = append(ansList, int(ans.Int64()))
		}
	}

	for _, ans := range ansList {
		fmt.Println(ans)
	}
}

func readInput() (int64, int64) {
	var temp string
	fmt.Scanln(&temp)
	size := strings.Split(temp, " ")
	down, _ := strconv.ParseInt(size[0], 10, 64)
	right, _ := strconv.ParseInt(size[1], 10, 64)
	return down - 1, right - 1
}
