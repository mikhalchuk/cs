package simple_text_editor

import (
	"stack"
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

var st = &stack.StackList{}
var s = make([]string, 0)

func main() {

	var q, i uint32

	reader := bufio.NewReader(os.Stdin)
	q = readNextInt(reader)

	for i = 0; i < q; i++ {

		line := readNextString(reader)
		val, pr := step(line, false)
		if pr {
			fmt.Println(val)
		}
	}
}

func readNextInt(reader *bufio.Reader) uint32 {
	line := readNextString(reader)
	v, _ := strconv.ParseInt(line, 10, 32)
	return uint32(v)
}

func readNextString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.Trim(s, "\n")
}

func step(line string, undo bool) (string, bool) {

	cmd := strings.Split(line, " ")

	switch cmd[0] {
	case "1":
		k := string(cmd[1])
		for i := 0; i < len(k); i++ {
			s = append(s, string(k[i]))
		}
		if !undo {
			st.Push(strings.Join([]string{"2", strconv.Itoa(len(k))}, " "))
		}
		return "", false
	case "2":
		k, _ := strconv.ParseUint(string(cmd[1]), 10, 32)

		var i uint64
		var toPush = make([]string, 0)
		for i = uint64(len(s)) - k; i < uint64(len(s)); i++ {
			toPush = append(toPush, s[i])
		}

		s = append(s[:0], s[:uint64(len(s)) - k]...)
		if !undo {
			st.Push(strings.Join([]string{"1", strings.Join(toPush, "")}, " "))
		}
		return "", false
	case "3":
		k, _ := strconv.ParseUint(string(cmd[1]), 10, 32)
		k -= 1
		return string(s[k]), true
	case "4":
		l, _ := st.Pop()
		return step(l.(string), true)
	}

	return  "", false
}