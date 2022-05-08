package main

import "github.com/t-efu/go-stack-linked-list"

func main() {
	s := stack.Stack{}
	s.Pop()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	s.Pop()
	s.Print()
	s.Push(4)
	s.Print()
}
