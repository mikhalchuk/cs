package balanced_brackets

import (
	"stack"
	"bufio"
	"os"
	"fmt"
	"strings"
)

var brackets = map[string]string{"{":"}","(":")","[":"]",}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var i, n int
	fmt.Scanf("%d", &n)

	for i = 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.Trim(line, "\n")

		balanced := balanced(line)

		if balanced {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

// balanced check is the string balanced or not
func balanced(s string) bool {

	st := &stack.StackList{}

	len := len(s)
	if len == 0 {
		return true
	}

	for i := 0; i < len; i++ {
		ch := string(s[i])
		if isLeftBracket(ch) {
			st.Push(ch)
		} else {
			el, err := st.Pop()
			if err != nil {
				return false
			}

			bracket := findCorrespondingBracket(el.(string))
			if ch != bracket {
				return false
			}
		}
	}

	return st.IsEmpty()
}

// isLeftBracket check is the character left bracket or not
func isLeftBracket(s string) bool {
	_, exist := brackets[s]
	return exist
}

// findCorrespondingBracket find corresponding  bracket
func findCorrespondingBracket(s string) string {
	return brackets[s]
}