package main

import (
	"fmt"
	"golang-data-structures/list"
	"golang-data-structures/stack"
)

func main() {
	intList := list.New()
	fmt.Println(intList)
	intList.Append(3)
	fmt.Println(intList)
	fmt.Println(intList.Tail)
	intList.Prepend(5)
	fmt.Println(intList)
	fmt.Println(intList.Head)
	fmt.Println(intList.Head.Next)

	intStack := stack.New()
	fmt.Println(intStack.Size)
	fmt.Println(intStack.IsEmpty())
	intStack.Push(4)
	intStack.Push(2)
	fmt.Println(intStack.Size)
	peek, _ := intStack.Peek()
	fmt.Println(peek)
	pop, _ := intStack.Pop()
	fmt.Println(pop)
	fmt.Println(intStack.Size)
}
