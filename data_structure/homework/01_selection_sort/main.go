package main

import (
	"fmt"
)

func main() {
	var studentNum int

	fmt.Scanln(&studentNum)

	id := make([]int, studentNum)
	score := make([]int, studentNum)

	for i := 0; i < studentNum; i++ {
		fmt.Scanln(&id[i], &score[i])
	}

	selectionSort(id, score)

	for i := 0; i < studentNum; i++ {
		fmt.Println(id[i])
	}
}

func selectionSort(idList, scoreList []int) {
	amount := len(idList)
	var max int

	for i := 0; i < amount-1; i++ {
		max = i
		for j := i + 1; j < amount; j++ {
			if scoreList[j] > scoreList[max] {
				max = j
			}
		}

		keyId := idList[max]
		keyScore := scoreList[max]
		for max > i {
			idList[max] = idList[max-1]
			scoreList[max] = scoreList[max-1]
			max--
		}

		idList[i] = keyId
		scoreList[i] = keyScore
	}
}
