package constant

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday)
}

const (
	Readable   = 1 << iota //0001
	Writable               //0010
	Executable             //0100
)

func TestConst2(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestBitClear(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	//true true true
	a = a &^ Readable //按位清零
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	//false true true
}
