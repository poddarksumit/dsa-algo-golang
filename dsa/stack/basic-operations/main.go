package main

import (
	"fmt"

	stack_linked_list "github.com/poddarksumit/dsa-algo-golang/dsa/stack/stack-linked-list"
)

func main() {

	fmt.Println("Welcome to Basic Operations of Stack")

	// Stack = []
	stack := stack_linked_list.Stack[int]{}
	fmt.Println("Stack Empty? ", stack.IsEmpty())
	data, err := stack.Peek()
	if err != nil {
		fmt.Println("Stack Peek threw error:: ", err)
	} else {
		fmt.Println("Stack Peek gave data:: ", data)
	}

	stack.Push(1)

	fmt.Println("Stack Empty? ", stack.IsEmpty())

	// Stack = [1]
	data, err = stack.Peek()
	if err != nil {
		fmt.Println("Stack Peek threw error:: ", err)
	} else {
		fmt.Println("Stack Peek gave data:: ", data)
	}

	// Stack = [1]
	data, err = stack.Pop()
	fmt.Println("Stack Empty? ", stack.IsEmpty())
	if err != nil {
		fmt.Println("Stack Pop threw error:: ", err)
	} else {
		fmt.Println("Stack Pop gave data:: ", data)
	}

	// Stack = []
	data, err = stack.Pop()
	if err != nil {
		fmt.Println("Stack Pop threw error:: ", err)
	} else {
		fmt.Println("Stack Pop gave data:: ", data)
	}

	stack.Push(10)
	stack.Push(11)
	stack.Push(9)
	stack.Push(67)
	stack.Push(34)
	stack.Push(91)
	stack.Push(82)
	stack.Push(61)

	// stack = [61,82,91,34,67,9,11,10]
	fmt.Println("Stack Empty? ", stack.IsEmpty())

	data, err = stack.Pop()
	if err != nil {
		fmt.Println("Stack Pop threw error:: ", err)
	} else {
		fmt.Println("Stack Pop gave data:: ", data)
	}

	data, err = stack.Peek()
	if err != nil {
		fmt.Println("Stack Peek threw error:: ", err)
	} else {
		fmt.Println("Stack Peek gave data:: ", data)
	}

	data, err = stack.Pop()
	if err != nil {
		fmt.Println("Stack Pop threw error:: ", err)
	} else {
		fmt.Println("Stack Pop gave data:: ", data)
	}

	fmt.Println("Bye")
}
