package linkedlist

type circleDoublyLinkedList struct {
	size       int
	head, last *doublyNode
}

func newCircleDoublyLinkedList() ILinkedList {
	return &circleDoublyLinkedList{head: nil, last: nil}
}

func (cd *circleDoublyLinkedList) GetSize() int {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) IsEmpty() bool {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) AddHead(element interface{}) {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) AddTail(element interface{}) {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Add(index int, element interface{}) {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Remove(index int) interface{} {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Clear() {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Set(index int, element interface{}) interface{} {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Contains(element interface{}) bool {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Get(index int) interface{} {
	panic("implement me")
}

func (cd *circleDoublyLinkedList) IndexOf(element interface{}) int {
	panic("implement me")
}
