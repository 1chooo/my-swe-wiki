package main

import "fmt"

type matrix struct {
	row   int
	col   int
	value int
}

func COMPARE(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func transMatrix(a []matrix, b []matrix) {
	rowTerms := make([]int, 100000)
	startingPos := make([]int, 100000)

	if a[0].value > 0 {
		for i := 0; i < a[0].col; i++ {
			rowTerms[i] = 0
		}
		for i := 1; i <= a[0].value; i++ {
			rowTerms[a[i].col]++
		}
		startingPos[0] = 1
		for i := 1; i <= a[0].value; i++ {
			startingPos[i] = startingPos[i-1] + rowTerms[i-1]
		}
		for i := 1; i <= a[0].value; i++ {
			j := startingPos[a[i].col]
			b[j].row = a[i].col
			b[j].col = a[i].row
			b[j].value = a[i].value
			startingPos[a[i].col]++
		}
	}
}

func multMatrix(a []matrix, b []matrix, newB []matrix, d []matrix) {
	totald := 0
	rowBegin := 1
	row := a[1].row
	sum := 0

	rowsA := a[0].row
	colsB := b[0].col
	totala := a[0].value
	totalb := b[0].value

	transMatrix(b, newB)

	a[totala+1].row = a[0].row
	newB[totalb+1].row = b[0].col
	newB[totalb+1].col = 0

	for i := 1; i <= totala; {
		for j := 1; j <= totalb+1; {
			if a[i].row != row {
				storesum(d, &totald, row, newB[j].row, &sum)
				i = rowBegin
				for newB[j].row == newB[j-1].row {
					j++
				}
			} else if newB[j].row != newB[j-1].row {
				storesum(d, &totald, row, newB[j].row, &sum)
				i = rowBegin
			} else {
				switch COMPARE(a[i].col, newB[j].col) {
				case 1:
					j++
				case -1:
					i++
				case 0:
					sum += (a[i].value * newB[j].value)
					i++
					j++
				}
			}
		}
		for a[i].row == row {
			i++
		}
		rowBegin = i
		row = a[i].row
	}
	d[0].row = rowsA
	d[0].col = colsB
	d[0].value = totald
}


func storesum(d []matrix, totald *int, row, column int, sum *int) {
	if *sum != 0 {
		*totald++
		d[*totald].row = row
		d[*totald].col = column
		d[*totald].value = *sum
		*sum = 0
	}
}

func printMatrix(matrix []matrix) {
	for i := 0; i <= matrix[0].value; i++ {
		fmt.Printf("%d %d %d\n", matrix[i].row, matrix[i].col, matrix[i].value)
	}
}

func main() {
	var aRow, aCol, bRow, bCol, countA, countB, tmp int
	countA, countB = 1, 1

	a := make([]matrix, 100000)
	b := make([]matrix, 100000)
	newB := make([]matrix, 100000)
	result := make([]matrix, 100000)

	fmt.Scanf("%d %d", &aRow, &aCol)
	for i := 1; i <= aRow; i++ {
		for j := 1; j <= aCol; j++ {
			fmt.Scanf("%d", &tmp)
			if tmp != 0 {
				a[countA].row = i - 1
				a[countA].col = j - 1
				a[countA].value = tmp
				countA++
			}
		}
	}
	a[countA].col = 0
	a[0].row = aRow
	a[0].col = aCol
	a[0].value = countA - 1
	a[countA].row = aRow

	fmt.Scanf("%d %d", &bRow, &bCol)
	for i := 1; i <= bRow; i++ {
		for j := 1; j <= bCol; j++ {
			fmt.Scanf("%d", &tmp)
			if tmp != 0 {
				b[countB].row = i - 1
				b[countB].col = j - 1
				b[countB].value = tmp
				countB++
			}
		}
	}
	b[countB].col = 0
	b[0].row = bRow
	b[0].col = bCol
	b[0].value = countB - 1
	b[countB].row = b[0].row

	multMatrix(a, b, newB, result)
	printMatrix(result)
}
