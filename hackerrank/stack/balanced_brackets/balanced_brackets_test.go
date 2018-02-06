package balanced_brackets

import "testing"

const n  = 4

func Test_Balanced(t *testing.T) {

	input := [n]string{
		"",
		"{[()]}",
		"{[(])}",
		"{{[[(())]]}}",
	}

	expected := [n]bool{
		true,
		true,
		false,
		true,
	}

	for i := 0; i < n; i++ {
		actual := balanced(input[i])

		if actual != expected[i] {
			t.Errorf("They should match: [%d]", i)
		}
	}
}