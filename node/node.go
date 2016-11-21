package node

type Node struct {
	Next *Node
	Data int // TODO: Make generic
}

func New(num int, next *Node) *Node {
	return &Node{Data: num, Next: next}
}
