package switch_test

import "testing"

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log(i, "Even")
		case 1, 3:
			t.Log(i, "Odd")
		default:
			t.Log(i, "it is not 0-3")
		}
	}

}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "Even")
		case i%2 == 1:
			t.Log(i, "Odd")
		default:
			t.Log(i, "unknown")
		}
	}
}
