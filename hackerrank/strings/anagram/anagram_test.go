package main

import "testing"

func Test_Calculate(t *testing.T) {
	input := []string{
		"xxyx",
		"aaabbb",
		"ab",
		"abc",
		"mnop",
		"xyyx",
		"xaxbbbxx",
	}

	results := []int16{
		1,
		3,
		1,
		-1,
		2,
		0,
		1,
	}

	for i, s := range input {
		result := calculate(s)
		if result != results[i] {
			t.Errorf("Should match %s (%d != %d)", s, results[i], result)
		}
	}
}