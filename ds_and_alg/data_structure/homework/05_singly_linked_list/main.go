package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	data int64
	next *Node
}

func getsize(head *Node) int {
	size := 0

	for head != nil {
		size++
		head = head.next
	}

	return size
}

func addBack(num int64, head **Node) {
	temp := &Node{data: num, next: nil}

	if *head == nil {
		*head = temp
	} else {
		before := *head
		for before.next != nil {
			before = before.next
		}
		before.next = temp
	}
}

func addFront(num int64, head **Node) {
	temp := &Node{data: num, next: *head}
	*head = temp
}

func addIndex(index, num int64, head **Node) {
	before := *head

	for i := int64(0); i <= index; i++ {
		if int64(getsize(*head)) < index {
			break
		}

		if index == 0 {
			addFront(num, head)
			break
		} else if index == i {
			temp := &Node{data: num, next: before.next}
			before.next = temp
		} else if i == int64(getsize(*head)-1) {
			addBack(num, head)
			break
		} else if before.next == nil {
			break
		} else if index-i > 1 {
			before = before.next
		}
	}
}

func deleteIndex(index int64, head **Node) {
	before := *head

	for i := int64(0); i <= index; i++ {
		if index == 0 {
			if *head != nil {
				*head = (*head).next
			}
			break
		} else if before.next == nil {
			break
		} else if index-i == 1 {
			continue
		} else if i == index {
			before.next = before.next.next
		} else {
			before = before.next
		}
	}
}

func split(input string, head **Node) {
	commands := strings.Fields(input)

	for i := 0; i < len(commands); i++ {
		command := commands[i]
		switch command {
		case "addBack":
			num, _ := strconv.ParseInt(commands[i+1], 10, 64)
			addBack(num, head)
			i++
		case "addFront":
			num, _ := strconv.ParseInt(commands[i+1], 10, 64)
			addFront(num, head)
			i++
		case "addIndex":
			index, _ := strconv.ParseInt(commands[i+1], 10, 64)
			num, _ := strconv.ParseInt(commands[i+2], 10, 64)
			addIndex(index, num, head)
			i += 2
		case "deleteIndex":
			index, _ := strconv.ParseInt(commands[i+1], 10, 64)
			deleteIndex(index, head)
			i++
		case "exit":
			return
		default:
			// Malicious input
			return
		}
	}
}

func print(head *Node) {
	for head != nil {
		fmt.Print(head.data, "-->")
		head = head.next
	}

	if head != nil {
		fmt.Println(head.data, "-->null")
	} else {
		fmt.Println("null")
	}
}

func main() {
	var input string
	head := (*Node)(nil)

	fmt.Scan(&input)

	split(input, &head)

	print(head)
}
