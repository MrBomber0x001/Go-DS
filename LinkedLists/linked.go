package main

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	myList := linkedlist{}
	node1 := &node{data: 47}
	node2 := &node{data: 47}
	node3 := &node{data: 47}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
}
