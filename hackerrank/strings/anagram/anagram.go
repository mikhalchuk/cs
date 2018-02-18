package main

import (
	"bufio"
	"strconv"
	"strings"
	"os"
	"fmt"
)

func main() {
	var i uint8

	reader := bufio.NewReader(os.Stdin)
	n := readNextInt(reader)

	for i =0; i < n; i++ {
		line := readNextString(reader)
		result := calculate(line)
		fmt.Println(result)
	}
}

func readNextInt(reader *bufio.Reader) uint8 {
	line := readNextString(reader)
	v, _ := strconv.ParseInt(line, 10, 8)
	return uint8(v)
}

func readNextString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.Trim(s, "\n")
}

func calculate(s string) int16 {
	len := len(s)
	if len % 2 != 0 {
		return -1
	}

	var replace int16
	half := len / 2
	dict := make(map[string]uint16, half)

	for i := 0; i < half; i++ {
		dict[string(s[i])] += 1
	}

	for i := half; i < len; i++ {
		if _, ok := dict[string(s[i])]; ok {
			if dict[string(s[i])] > 0 {
				dict[string(s[i])] -= 1
			} else {
				replace += 1
			}
		} else {
			replace += 1
		}
	}

	return replace
}