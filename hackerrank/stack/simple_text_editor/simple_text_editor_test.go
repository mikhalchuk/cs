package simple_text_editor

import (
	"testing"
)


func Test_Steps(t *testing.T) {

	input := [8]string{
		"1 abc",
		"3 3",
		"2 3",
		"1 xy",
		"3 2",
		"4",
		"4",
		"3 1",
	}

	res := [3]string{
		"c",
		"y",
		"a",
	}

	cnt := 0
	for _, s := range input{
		val, pr := step(s, false)
		if pr {
			if val != res[cnt] {
				t.Errorf("They should match %d", cnt)
			}
			cnt++
		}
	}
}