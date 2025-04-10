package main

import (
	"fmt"

	"strings"

	stack_linked_list "github.com/poddarksumit/dsa-algo-golang/dsa/stack/stack-linked-list"
)

// Leetcode 2000. Reverse Prefix of Word (https://leetcode.com/problems/reverse-prefix-of-word/description/)
func main() {

	fmt.Println("Welcome to Reverse Prefix of Word")

	reversedWord := reversePrefixUsingStack("abcdefd", 'd')

	fmt.Println("Reversed word 'reversePrefixUsingStack' :: ", reversedWord)

	reversedWord = reversePrefixUsingString("abcdefd", 'd')

	fmt.Println("Reversed word 'reversePrefixUsingString' :: ", reversedWord)

	reversedWord = reversePrefix("ab", 'b')

	fmt.Println("Reversed word 'reversePrefix' :: ", reversedWord)
}

func reversePrefix(word string, ch byte) string {

	if word == "" {
		return word
	}

	var finalReversedWordBuilder strings.Builder
	tempReversedWord := ""
	var wordSoFarBuilder strings.Builder
	chIndex := -1

	for index := 0; index < len(word); index++ {

		wordSoFarBuilder.WriteByte(word[index])

		if chIndex == -1 {
			tempReversedWord = string(word[index]) + tempReversedWord

			if word[index] == ch {
				chIndex = index
				finalReversedWordBuilder.WriteString(tempReversedWord)
				tempReversedWord = ""
				wordSoFarBuilder = strings.Builder{}
			}
		}
	}

	if wordSoFarBuilder.Len() != 0 {
		finalReversedWordBuilder.WriteString(wordSoFarBuilder.String())
	}

	return finalReversedWordBuilder.String()
}

func reversePrefixUsingString(word string, ch rune) string {

	if word == "" {
		return word
	}

	var finalReversedWordBuilder strings.Builder
	var wordSoFarBuilder strings.Builder
	var tempReversedWord = ""

	for _, wordCh := range word {

		tempReversedWord = string(wordCh) + tempReversedWord
		wordSoFarBuilder.WriteRune(wordCh)

		if wordCh == ch {
			finalReversedWordBuilder.WriteString(tempReversedWord)
			tempReversedWord = ""
			wordSoFarBuilder = strings.Builder{}
		}
	}

	if wordSoFarBuilder.Len() != 0 {
		finalReversedWordBuilder.WriteString(wordSoFarBuilder.String())
	}

	return finalReversedWordBuilder.String()
}

func reversePrefixUsingStack(word string, ch rune) string {

	if word == "" {
		return word
	}

	stack := stack_linked_list.Stack[rune]{}
	var reverseWord = ""
	var wordsSoFar = ""

	for _, wordCh := range word {
		wordsSoFar += string(wordCh)
		stack.Push(wordCh)
		reverseWord, stack = ReverseWordFromStack(reverseWord, stack, ch)

		if stack.IsEmpty() {
			wordsSoFar = ""
		}

	}

	reverseWord, stack = ReverseWordFromStack(reverseWord, stack, ch)

	if reverseWord == "" {
		return word
	} else {
		return reverseWord + wordsSoFar
	}
}

func ReverseWordFromStack(reverseWord string, stack stack_linked_list.Stack[rune], ch rune) (string, stack_linked_list.Stack[rune]) {
	if stack.IsEmpty() {
		return reverseWord, stack
	}

	var stackCh, _ = stack.Peek()
	if stackCh != ch {
		return reverseWord, stack
	}

	for !stack.IsEmpty() {
		stackCh, _ = stack.Pop()
		reverseWord += string(stackCh)
	}

	return reverseWord, stack
}
