package main

import (
	"fmt"

	stack_linked_list "github.com/poddarksumit/dsa-algo-golang/dsa/stack/stack-linked-list"
)

func main() {

	fmt.Println("Welcome to Check Valid Parentheses using Stack")

	isValid := CheckValidParentheses("{abc}")

	fmt.Println("Is valid string:: ", isValid)
}

func CheckValidParentheses(text string) bool {

	if text == "" {
		return false
	}

	stack := stack_linked_list.Stack[rune]{}
	parenthesesCloseMap := map[rune]rune{'}': '{', ']': '[', ')': '('}
	validParenthesesMap := map[rune]bool{'(': true, ')': true, '{': true, '}': true, '[': true, ']': true}

	for _, textChar := range text {

		if validParenthesesMap[textChar] {

			parenthesesCloseMapValue := parenthesesCloseMap[textChar]

			if parenthesesCloseMapValue == 0 {
				stack.Push(textChar)
			} else {
				data, err := stack.Peek()

				if err == nil && data == parenthesesCloseMapValue {
					stack.Pop()
				} else {
					return false
				}
			}
		}
	}

	return stack.IsEmpty()
}
