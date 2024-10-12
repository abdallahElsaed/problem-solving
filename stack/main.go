package main

import "fmt"

func main() {
	//! 155. Min Stack
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin()) // Should print -3
	stack.Pop()
	fmt.Println(stack.Top())    // Should print 0
	fmt.Println(stack.GetMin()) // Should print -2	
}