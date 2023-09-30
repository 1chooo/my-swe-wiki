package main

import (
	"fmt"
)

type element struct {
	row int
	col int
	dir int
}

type offset struct {
	vert int
	horiz int
}

var stack [100]element
var top = -1

var move = []offset{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func push(item element) {
	top++
	stack[top] = item
}

func pop() {
	if top == -1 {
		stack[top] = element{}
	} else {
		top--
	}
}

func main() {
	var length, width, nextRow, nextCol, dir, col, row int
	found := false
	var position element

	length, width = 0, 0 // Initialize to zero
	fmt.Scanf("%d %d", &length, &width)

	maze := make([][]int, length+2)
	mark := make([][]int, length+2)

	for i := range maze {
		maze[i] = make([]int, width+2)
		mark[i] = make([]int, width+2)
	}

	for i := 1; i <= length; i++ {
		for j := 1; j <= width; j++ {
			fmt.Scanf("%d", &maze[i][j])
		}
	}

	for i := 0; i <= width+1; i++ {
		maze[0][i] = 1
		maze[length+1][i] = 1
	}

	for i := 0; i <= length+1; i++ {
		maze[i][0] = 1
		maze[i][width+1] = 1
	}

	for i := 0; i < length+2; i++ {
		for j := 0; j < width+2; j++ {
			mark[i][j] = 0
		}
	}

	mark[1][1] = 1
	top = 0
	stack[0] = element{row: 1, col: 1, dir: 1}

	for top > -1 && !found {
		position = stack[top]
		row, col, dir = position.row, position.col, position.dir
		pop()

		for dir < 4 && !found {
			if dir > 3 && top == -1 {
				break
			}
			nextRow = row + move[dir].vert
			nextCol = col + move[dir].horiz

			if nextRow == length && nextCol == width {
				found = true
			} else if maze[nextRow][nextCol] == 0 && mark[nextRow][nextCol] == 0 {
				mark[nextRow][nextCol] = 1
				position.dir = dir + 1
				position.row, position.col = row, col
				push(position)
				row, col, dir = nextRow, nextCol, 0
			} else {
				dir++
			}
		}
	}

	if found {
		for i := 0; i <= top; i++ {
			fmt.Printf("(%d,%d) ", stack[i].row-1, stack[i].col-1)
		}
		fmt.Printf("(%d,%d) ", row-1, col-1)
		fmt.Printf("(%d,%d)", nextRow-1, nextCol-1)
	} else if dir > 3 {
		fmt.Println("Can't reach the exit!")
	}
}
