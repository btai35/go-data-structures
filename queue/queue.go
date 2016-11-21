package queue

import "golang-data-structures/node"

type Queue struct {
	Head *node.Node
	Tail *node.Node
	Size int
}

func New() *Queue {
	return &Queue{Head: nil, Tail: nil, Size: 0}
}
