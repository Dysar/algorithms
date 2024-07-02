package main

import "github.com/golang-collections/collections/stack"

func isValid(s string) bool {
	// it has to be even num of characters
	if len(s)%2 != 0 {
		return false
	}

	st := stack.New()
	for _, ch := range s {
		char := string(ch)

		if char == "(" || char == "[" || char == "{" {
			st.Push(char)
		} else if char == ")" && st.Peek() == "(" {
			st.Pop()
			continue
		} else if char == "}" && st.Peek() == "{" {
			st.Pop()
		} else if char == "]" && st.Peek() == "[" {
			st.Pop()
		} else {
			return false //when cases like }}, )), ([}}]) this will return false
		}

	}

	return st.Len() == 0
}
