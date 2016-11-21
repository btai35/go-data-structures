package stack

import "golang-data-structures/node"

type Stack struct {
	Head *node.Node
  Size int
}

type New() *Stack {
  return &Stack{Head: nil, Size: 0}
}

func (s *Stack) Pop() (int, error) {
  if s.Size < 1 {
    return 0, errors.New("error: stack is empty")
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
