package main

import (
	"fmt"
)

func main() {
	var runTimes int
	fmt.Scan(&runTimes)

	for i := 0; i < runTimes; i++ {
		var page int
		fmt.Scan(&page)

		addThePage := NewAddThePage()
		addThePage.GetTotalPages(page)
	}
}

type AddThePage struct{}

func NewAddThePage() *AddThePage {
	return &AddThePage{}
}

func (atp *AddThePage) GetTotalPages(page int) {
	i := 0
	total := 0
	rp := 0

	for {
		i++
		total += i

		if total > page {
			rp = total
			break
		}
	}

	fmt.Printf("%d %d\n", (rp - page), i)
}
