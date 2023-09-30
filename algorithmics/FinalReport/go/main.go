package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Graph 定義了一個圖形，包含節點和邊的資訊
type Graph struct {
	d [11][11]int // 節點之間的距離資訊
}

// NewGraph 創建一個新的圖形
func NewGraph() *Graph {
	return &Graph{}
}

// FindShortestPath 找到從節點 A 到目標節點的最短路徑
func (g *Graph) FindShortestPath(stage int) int {
	switch stage {
	case 1:
		return g.Stage1()
	case 2:
		return g.Stage2()
	case 3:
		return g.Stage3()
	case 4, 5:
		return g.Stage4Or5()
	default:
		return math.MaxInt32
	}
}

// Stage1 實現了 stage1 的邏輯
func (g *Graph) Stage1() int {
	small := math.MaxInt32
	for i := 1; i <= 3; i++ {
		if small > g.d[0][i] {
			small = g.d[0][i]
		}
	}
	return small
}

// Stage2 實現了 stage2 的邏輯
func (g *Graph) Stage2() int {
	small := math.MaxInt32
	for j := 0; j < 3; j++ {
		temp := g.d[0][j+1] + g.min(g.d[j+1][4], g.d[j+1][5], g.d[j+1][6])
		if small > temp {
			small = temp
		}
	}
	return small
}

// Stage3 實現了 stage3 的邏輯
func (g *Graph) Stage3() int {
	small := math.MaxInt32
	for n := 1; n <= 3; n++ {
		for j := 4; j <= 6; j++ {
			temp := g.d[0][n] + g.d[n][j] + g.min(g.d[j][7], g.d[j][8], g.d[j][9])
			if small > temp {
				small = temp
			}
		}
	}
	return small
}

// Stage4Or5 實現了 stage4 和 stage5 的邏輯
func (g *Graph) Stage4Or5() int {
	small := math.MaxInt32
	for n := 1; n <= 3; n++ {
		for j := 4; j <= 6; j++ {
			for k := 7; k <= 9; k++ {
				temp := g.d[0][n] + g.d[n][j] + g.d[j][k] + g.d[k][10]
				if small > temp {
					small = temp
				}
			}
		}
	}
	return small
}

// min 返回三個整數的最小值
func (g *Graph) min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var times, stage int
	fmt.Scan(&times)
	var g Graph

	for b := 0; b < times; b++ {
		for c := 1; c <= 3; c++ {
			fmt.Scan(&g.d[0][c])
		}
		for c := 1; c <= 3; c++ {
			for e := 4; e <= 6; e++ {
				fmt.Scan(&g.d[c][e])
			}
		}
		for c := 4; c <= 6; c++ {
			for e := 7; e <= 9; e++ {
				fmt.Scan(&g.d[c][e])
			}
		}
		for c := 7; c <= 9; c++ {
			fmt.Scan(&g.d[c][10])
		}
		fmt.Scan(&stage)

		shortestPath := g.FindShortestPath(stage)
		fmt.Println(shortestPath)
	}
}
