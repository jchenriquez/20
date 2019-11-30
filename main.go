package main

import "fmt"

var PAREN_MATCH = map[rune]rune {
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, char := range s {
		opening, isClosing := PAREN_MATCH[char]

		if isClosing {
			if len(stack) == 0 {
				return false
			}

			if opening != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Printf("isValid %v\n", isValid("{[]}"))
}
