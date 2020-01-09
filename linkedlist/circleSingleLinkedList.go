package linkedlist

type circleSingleLinkedList struct {
	size       int
	head, last *singleNode
}

func newCircleSingleLinkedList() ILinkedList {
	return &circleSingleLinkedList{head: nil, last: nil}
}

func (cs *circleSingleLinkedList) GetSize() int {
	panic("implement me")
}

func (cs *circleSingleLinkedList) IsEmpty() bool {
	panic("implement me")
}

func (cs *circleSingleLinkedList) AddHead(element interface{}) {
	panic("implement me")
}

func (cs *circleSingleLinkedList) AddTail(element interface{}) {
	panic("implement me")
}

func (cs *circleSingleLinkedList) Add(index int, element interface{}) {
	panic("implement me")
}

func (cs *circleSingleLinkedList) Remove(index int) interface{} {
	panic("implement me")
}

func (cs *circleSingleLinkedList) Clear() {
	panic("implement me")
}

func (cs *circleSingleLinkedList) Set(index int, element interface{}) interface{} {
	panic("implement me")
}

func (cs *circleSingleLinkedList) Contains(element interface{}) bool {
	panic("implement me")
}

func (cs *circleSingleLinkedList) Get(index int) interface{} {
	panic("implement me")
}

func (cs *circleSingleLinkedList) IndexOf(element interface{}) int {
	panic("implement me")
}
