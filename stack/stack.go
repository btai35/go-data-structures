package stack

import (
	"errors"
	"golang-data-structures/node"
)

type Stack struct {
	Head *node.Node
	Size int
}

func New() *Stack {
	return &Stack{Head: nil, Size: 0}
}

func (s *Stack) Pop() (int, error) {
	if s.Size < 1 {
		return -1, errors.New("error: stack is empty")
	}

	num := s.Head.Data

	s.Head = s.Head.Next
	s.Size = s.Size - 1

	return num, nil
}

func (s *Stack) Push(num int) {
	new := node.New(num, s.Head)

	s.Head = new
	s.Size = s.Size + 1
}

func (s *Stack) Peek() (int, error) {
	if s.Size < 1 {
		return -1, errors.New("error: stack is empty")
	}

	return s.Head.Data, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}
