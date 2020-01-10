package linkedlist

type circleDoublyLinkedList struct {
	size       int
	head, last *doublyNode
}

// NewCircleDoublyLinkedList 初始化
func NewCircleDoublyLinkedList() ILinkedList {
	return &circleDoublyLinkedList{head: nil, last: nil}
}

func (cd *circleDoublyLinkedList) GetSize() int {
	return cd.size
}

func (cd *circleDoublyLinkedList) IsEmpty() bool {
	return cd.size == 0
}

func (cd *circleDoublyLinkedList) AddHead(element interface{}) {
	cd.Add(0, element)
}

func (cd *circleDoublyLinkedList) AddTail(element interface{}) {
	cd.Add(cd.size, element)
}

func (cd *circleDoublyLinkedList) Add(index int, element interface{}) {
	rangeCheckForAdd(cd.size, index)
}

func (cd *circleDoublyLinkedList) Remove(index int) interface{} {
	rangeCheck(cd.size, index)
	panic("implement me")
}

func (cd *circleDoublyLinkedList) Clear() {
	cd.size = 0
	cd.head = nil
	cd.last = nil
}

func (cd *circleDoublyLinkedList) Set(index int, element interface{}) interface{} {
	rangeCheck(cd.size, index)
	node := cd.getNode(index)
	oldValue := node.ele
	node.ele = element
	return oldValue
}

func (cd *circleDoublyLinkedList) Contains(element interface{}) bool {
	return cd.IndexOf(element) != ElementNotFound
}

func (cd *circleDoublyLinkedList) Get(index int) interface{} {
	rangeCheck(cd.size, index)
	return cd.getNode(index).ele
}

func (cd *circleDoublyLinkedList) IndexOf(element interface{}) int {
	node := cd.head
	for i := 0; i < cd.size; i++ {
		if node.ele == element {
			return i
		}
		node = node.next
	}
	return ElementNotFound
}

func (cd *circleDoublyLinkedList) getNode(index int) *doublyNode {
	if index <= cd.size>>1 {
		node := cd.head
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	}
	node := cd.last
	for i := cd.size - 1; i > index; i-- {
		node = node.prev
	}
	return node
}
