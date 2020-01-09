package linkedlist

type circleSingleLinkedList struct {
	size       int
	head, last *singleNode
}

func newCircleSingleLinkedList() ILinkedList {
	return &circleSingleLinkedList{head: nil, last: nil}
}

func (cs *circleSingleLinkedList) GetSize() int {
	return cs.size
}

func (cs *circleSingleLinkedList) IsEmpty() bool {
	return cs.size == 0
}

func (cs *circleSingleLinkedList) AddHead(element interface{}) {
	cs.Add(0, element)
}

func (cs *circleSingleLinkedList) AddTail(element interface{}) {
	cs.Add(cs.size, element)
}

func (cs *circleSingleLinkedList) Add(index int, element interface{}) {
	rangeCheckForAdd(cs.size, index)
	panic("implement me")
}

func (cs *circleSingleLinkedList) Remove(index int) interface{} {
	rangeCheck(cs.size, index)
	panic("implement me")
}

func (cs *circleSingleLinkedList) Clear() {
	cs.size = 0
	cs.head = nil
	cs.last = nil
}

func (cs *circleSingleLinkedList) Set(index int, element interface{}) interface{} {
	rangeCheck(cs.size, index)
	node := cs.getNode(index)
	oldValue := node.ele
	node.ele = element
	return oldValue
}

func (cs *circleSingleLinkedList) Contains(element interface{}) bool {
	return cs.IndexOf(element) != ElementNotFound
}

func (cs *circleSingleLinkedList) Get(index int) interface{} {
	rangeCheck(cs.size, index)
	node := cs.getNode(index)
	return node.ele
}

func (cs *circleSingleLinkedList) IndexOf(element interface{}) int {
	node := cs.head
	for i := 0; i < cs.size; i++ {
		if node.ele == element {
			return i
		}
		node = node.next
	}
	return ElementNotFound
}
func (cs *circleSingleLinkedList) getNode(index int) *singleNode {
	node := cs.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
