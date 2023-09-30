package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)

	for i := 0; i < k; i++ {
		var n, m int
		fmt.Scan(&n, &m)

		str := make([]string, 101)
		differ := make([]int, 101)

		for j := 0; j < m; j++ {
			fmt.Scan(&str[j])
			differ[j] = 0

			for k := 0; k < n; k++ {
				for l := k + 1; l < n; l++ {
					if str[j][k] > str[j][l] {
						differ[j]++
					}
				}
			}
		}

		for j := 0; j < m-1; j++ {
			for k := 0; k < m-j-1; k++ {
				if differ[k+1] < differ[k] {
					tmp := differ[k]
					differ[k] = differ[k+1]
					differ[k+1] = tmp

					temp := str[k]
					str[k] = str[k+1]
					str[k+1] = temp
				}
			}
		}

		for j := 0; j < m; j++ {
			fmt.Println(str[j])
		}

		if i != k-1 {
			fmt.Println() // Print an empty line
			var blank string
			fmt.Scanln(&blank) // Read and discard the newline character
		}
	}
}
