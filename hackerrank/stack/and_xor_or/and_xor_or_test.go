package and_xor_or

import "testing"

func Test_Main(t *testing.T) {
	var n, expected uint32
	n = 5
	a := []uint32{9,6,3,5,2}
	expected = 15

	actual := max(n, a)

	if expected != actual {
		t.Errorf("expected %d != actual %d", expected, actual)
	}
}