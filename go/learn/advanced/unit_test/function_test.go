package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := Square(inputs[i])
		assert.Equal(t, expected[i], ret)
	}
}

func TestConcatStringByAdd(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	excepted := "abcd"
	ret := ConcatStringByAdd(input)
	assert.Equal(t, ret, excepted)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	excepted := "abcd"
	ret := ConcatStringByBytesBuffer(input)
	assert.Equal(t, ret, excepted)
}

func TestConcatStringByStringsBuilder(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	excepted := "abcd"
	ret := ConcatStringByStringsBuilder(input)
	assert.Equal(t, ret, excepted)
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	input := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStringByAdd(input)
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	input := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStringByBytesBuffer(input)
	}
	b.StopTimer()
}

func BenchmarkConcatStringByStringsBuilder(b *testing.B) {
	input := []string{"a", "b", "c", "d"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatStringByStringsBuilder(input)
	}
	b.StopTimer()
}

/*
goos: linux
goarch: amd64
pkg: advanced/unit_test
cpu: Intel(R) Xeon(R) W-1370P @ 3.60GHz
BenchmarkConcatStringByAdd
					16代表CPU核數(GOMAXPROCS)	  執行次數              單次耗費時間		allocs/op 代表每次執行都需要搭配一個記憶體空間，而一個記憶體空間為 X B(Bytes)/op
BenchmarkConcatStringByAdd-16               	17524599	        66.08 ns/op	      12 B/op	       3 allocs/op （每次迭代觸發3次內存分配，每次迭代內存使用量為12Bytes）
BenchmarkConcatStringByBytesBuffer
BenchmarkConcatStringByBytesBuffer-16       	13806450	       130.1 ns/op	      68 B/op	       2 allocs/op
BenchmarkConcatStringByStringsBuilder
BenchmarkConcatStringByStringsBuilder-16    	36199176	        30.09 ns/op	       8 B/op	       1 allocs/op
PASS
coverage: 92.3% of statements
ok  	advanced/unit_test	4.243s
*/
