package list

type List struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Next *Node
	Data int // TODO: Make generic
}

func New(num int) *List {
	head := &Node{Next: nil, Data: num}
	return &List{Head: head, Tail: head, Size: 1}
}

func (l *List) Append(num int) {
	tail := l.Tail
	new := Node{Next: nil, Data: num}
	tail.Next = &new

	l.Tail = &new
	l.Size = l.Size + 1
}

func (l *List) Prepend(num int) {
	head := l.Head
	new := Node{Next: head, Data: num}

	l.Head = &new
	l.Size = l.Size + 1
}
