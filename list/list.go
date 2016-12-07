package list

import "golang-data-structures/node"

type ListNode struct {
	Head *node.Node
	Tail *node.Node
	Size int
}

func New() *ListNode {
	return &ListNode{Head: nil, Tail: nil, Size: 0}
}

func (l *ListNode) Append(num int) {
	new := node.New(num, nil)

	if l.Size == 0 {
		l.Head = new
		l.Tail = new
		l.Size = 1

		return
	}

	l.Tail.Next = new
	l.Tail = new
	l.Size = l.Size + 1
}

func (l *ListNode) Prepend(num int) {
	if l.Size == 0 {
		new := node.New(num, nil)
		l.Head = new
		l.Tail = new
		l.Size = 1
		return
	}

	head := l.Head
	new := node.New(num, head)
	l.Head = new
	l.Size = l.Size + 1
}

func reverse(l *ListNode) *ListNode {
	if l == nil {
		return l
	}

	var prev *ListNode
	curr := l
	for hasNext(curr) {
		prev = &ListNode{Val: curr.Val, Next: prev}
		curr = curr.Next
	}

	curr.Next = prev

	return curr
}
