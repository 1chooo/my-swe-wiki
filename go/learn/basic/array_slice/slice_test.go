package array_slice

import "testing"

// slice struct 含 ptr(指向存儲空間的指標) len（數組長度） cap（數組最大容量）
func TestSliceInit(t *testing.T) {
	//宣告 1
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))
	//宣告 2
	s1 := []int{1, 2, 3, 4}
	t.Log(s0, len(s1), cap(s1))
	//宣告 3
	s2 := make([]int, 3, 5)     //make(類型,len,cap)
	t.Log(s2, len(s2), cap(s2)) //[0 0 0] 3 5
	s2 = append(s2, 7)
	t.Log(s2, len(s2), cap(s2)) //[0 0 0 7] 4 5

}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i) // 分配記憶體+複製一份達成append（有消耗）
		t.Log(s, len(s), cap(s))
	}
	//  [0] 1 1
	//  [0 1] 2 2
	//  [0 1 2] 3 4
	//  [0 1 2 3] 4 4
	//  [0 1 2 3 4] 5 8
	//  [0 1 2 3 4 5] 6 8
	//  [0 1 2 3 4 5 6] 7 8
	//  [0 1 2 3 4 5 6 7] 8 8
	//  [0 1 2 3 4 5 6 7 8] 9 16
	//  [0 1 2 3 4 5 6 7 8 9] 10 16
	// 當append的超出cap的數量時 cap會乘二（cap = cap*2）
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) //[Apr May Jun] 3 9
	// cap=9 是因為"Apr"～"Dec" 有9位

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) //[Jun Jul Aug] 3 7

	summer[0] = "Unknown" //Jun改Unknown
	t.Log(summer)         // [Unknown Jul Aug]
	t.Log(Q2)             // Q2也會跟著變[Apr May Unknown]
	summer[0] = "Jun"     //先恢復正常

	t.Log("--APPEND 的效果--")
	Q2 = append(Q2, "APPEND")
	t.Log(Q2)     // Q2成功append [Apr May Jun APPEND]
	t.Log(summer) //但是summer變成[Jun APPEND Aug] 造成Jul被改成APPEND
	t.Log(year)   //[Jan Feb Mar Apr May Jun APPEND Aug Sep Oct Nov Dec]
}
