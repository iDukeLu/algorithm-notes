package linkedlist

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val}
}

func (n *Node[T]) Append(val T) *Node[T] {
	node := NewNode(val)
	n.Next = node
	return node
}

func (n *Node[T]) AppendNode(node *Node[T]) *Node[T] {
	n.Next = node
	return node
}

func (n *Node[T]) AppendNodes(nodes ...*Node[T]) *Node[T] {
	n.Next = nodes[0]
	return nodes[len(nodes)-1]
}

func (n *Node[T]) Foreach(f func(*Node[T])) {
	for n != nil {
		f(n)
		n = n.Next
	}
}
