package main

import (
	"fmt"
	"golang-data-structures/list"
)

func main() {
	linkedList := list.New(3)
	fmt.Println(linkedList)
	linkedList.Append(3)
	fmt.Println(linkedList)
	fmt.Println(linkedList.Tail)
	linkedList.Prepend(5)
	fmt.Println(linkedList)
	fmt.Println(linkedList.Head)
	fmt.Println(linkedList.Head.Next)
}
