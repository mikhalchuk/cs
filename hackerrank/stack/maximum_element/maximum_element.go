package maximum_element

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"stack"
)

var st = &stack.StackList{}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var n, i int64
	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		s, _ := reader.ReadString('\n')

		cmd, input := parseCommand(s)

		max, peek := step([2]int64{cmd,input})

		if peek {
			fmt.Println(max)
		}
	}
}

func step(input [2]int64) (int64, bool) {
	cmd := input[0]
	switch cmd {
	case 1:
		newVal := input[1]
		var toPush int64
		head, err := st.Peek()
		if err != nil {
			toPush = newVal
		} else {
			if head.(int64) > newVal {
				toPush = head.(int64)
			} else {
				toPush = newVal
			}
		}
		st.Push(toPush)
	case 2:
		st.Pop()
	case 3:
		val, _ := st.Peek()
		return val.(int64), true
	}

	return 0, false
}

func parseCommand(s string) (cmd, val int64) {
	s = strings.Trim(s, "\n")
	args := strings.Split(s, " ")
	cmd = parseInt(args[0])
	if len(args) == 1 {
		return
	}

	val = parseInt(args[1])
	return
}

func parseInt(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}