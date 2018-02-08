package and_xor_or

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"os"
	"stack"
)

var st = &stack.StackList{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNextInt(reader)
	line := readNextString(reader)
	input := strings.Split(line, " ")

	var a []uint32
	for _, x := range input {
		number, _ := strconv.ParseInt(x, 10, 32)
		a = append(a, uint32(number))
	}

	fmt.Println(max(n, a))
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

func max(n uint32, a []uint32) uint32 {
	var i, max uint32
	max = 0

	for i = 0; i < n; i++ {
		for !st.IsEmpty() {
			peek, _ := st.Peek()
			top := peek.(uint32)

			newMax := f(a[i], top)
			if newMax > max {
				max = newMax
			}

			if a[i] < top {
				st.Pop()
			} else {
				break
			}
		}
		st.Push(a[i])
	}
	return max
}

func f(x uint32, y uint32) uint32 {
	return x ^ y
}
