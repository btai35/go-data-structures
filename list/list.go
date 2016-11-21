package list

import "golang-data-structures/node"

type List struct {
	Head *node.Node
	Tail *node.Node
	Size int
}

func New() *List {
	return &List{Head: nil, Tail: nil, Size: 0}
}

func (l *List) Append(num int) {
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

func (l *List) Prepend(num int) {
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
