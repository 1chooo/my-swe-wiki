package client_test

import (
	"basic/package_test/series"
	"testing"
)

func TestPkgFib(t *testing.T) {
	series.GetFibonacci(10)
	// series.square(10) 小寫函式 無法訪問
}
