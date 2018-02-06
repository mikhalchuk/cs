package maximum_element

import (
	"testing"
)

func Test_main(t *testing.T)  {

	input := [10][2]int64{}
	input[0][0] = 1
	input[0][1] = 97
	input[1][0] = 2
	input[2][0] = 1
	input[2][1] = 20
	input[3][0] = 2
	input[4][0] = 1
	input[4][1] = 26
	input[5][0] = 1
	input[5][1] = 20
	input[6][0] = 2
	input[7][0] = 3
	input[8][0] = 1
	input[8][1] = 91
	input[9][0] = 3


	cnt := 0
	res := [2]int64{26, 91}

	for _, val := range input {
		max, _ := step(val)
		if val[0] == 3 {
			if max != res[cnt] {
				t.Errorf("%d != %d step is wrong", res[cnt], max)
			}
			cnt++
		}
	}
}