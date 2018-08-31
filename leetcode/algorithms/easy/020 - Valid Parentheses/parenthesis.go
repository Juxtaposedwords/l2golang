package parenthesis

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.
*/
import (
	"fmt"
	"strings"
)

type stack struct {
	inner []string
}

func (s *stack) push(i string) {
	s.inner = append(s.inner, i)
}
func (s *stack) pop() error {
	if len(s.inner) < 1 {
		return fmt.Errorf("There are no item in this stack")
	}
	s.inner = s.inner[:len(s.inner)-1]
	return nil
}
func (s *stack) top() string {
	if len(s.inner) < 1 {
		return "0"
	}
	return s.inner[len(s.inner)-1]
}

func isValid(s string) bool {
	// We'll want to keep track of every character we found
	var c stack
	// Go over every character
	closer := map[string]string{"}": "{", "]": "[", ")": "("}
	l := strings.Split(s, "")
	for _, e := range l {
		val, ok := closer[e]
		switch {
		case !ok:
			c.push(e)
		case val == c.top():
			if err := c.pop(); err != nil {
				return false
			}
		default:
			return false
		}

	}
	if len(c.inner) > 0 {
		return false
	} else {
		return true
	}
}
