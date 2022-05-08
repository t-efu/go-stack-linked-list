package stack

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(v interface{}) {
	s.top = &Node{
		value: v,
		next:  s.top,
	}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	popped := s.top.value
	s.top = s.top.next
	s.size--
	return popped
}

func (s *Stack) Print() {
	temp := s.top
	fmt.Print("[")
	for {
		if temp != nil {
			fmt.Printf(" %v", temp.value)
		}
		temp = temp.next
		if temp != nil {
			fmt.Printf(",")
		}
		if temp == nil {
			fmt.Println(" ]")
			break
		}
	}
}

func (s *Stack) Len() int {
	return s.size
}
