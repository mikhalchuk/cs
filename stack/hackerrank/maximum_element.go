package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"stack"
)

func main() {
	stack := &stack.StackList{}

	reader := bufio.NewReader(os.Stdin)

	var n, i int64
	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		el, _ := reader.ReadString('\n')
		el = strings.Trim(el, "\n")
		s := strings.Split(el, " ")

		c, _ := strconv.ParseInt(s[0], 10, 64)
		if c == 1 {
			var p int64
			t, _ := strconv.ParseInt(s[1], 10, 64)
			in, err := stack.Peek()
			if err != nil {
				p = t
			} else {
				f := in.(int64)
				if f > t {
					p = f
				} else {
					p = t
				}
			}
			stack.Push(p)
		} else if c == 2 {
			stack.Pop()
		} else {
			max, _ := stack.Peek()
			fmt.Println(max)
		}
	}
}