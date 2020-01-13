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
	// 如果 index == size, 说明添加的索引是最后位置
	if cd.size == index {
		// 创建新节点, prev指向原链表的尾节点, next指向首节点
		newNode := &doublyNode{next: cd.head, prev: cd.last, ele: element}
		if cd.size == 0 { // 空链表
			cd.head = newNode
			cd.last = newNode
			newNode.prev = cd.last
			newNode.next = cd.last
		} else {
			// 原链表尾节点next指向newNode
			cd.last.next = newNode
			// 原链表头结点prev指向newNode
			cd.head.prev = newNode
			// last指向新的尾节点
			cd.last = newNode
		}
	} else {
		// 中间节点或者首节点
		next := cd.node(index)
		prev := next.prev
		newNode := &doublyNode{element, prev, next}
		prev.next = newNode
		next.prev = newNode
		if index == 0 { //  如果是首节点添加元素
			cd.head = newNode
		}
	}
	cd.size++
}

func (cd *circleDoublyLinkedList) Remove(index int) interface{} {
	rangeCheck(cd.size, index)
	node := cd.node(index)
	if cd.size == 1 {
		cd.head = nil
		cd.last = nil
	} else {
		prev := node.prev
		next := node.next
		next.prev = prev
		prev.next = next
		// 如果node == cd.head, 说明删除的是第一个节点
		if node == cd.head {
			cd.head = next
		}
		// 如果next == cd.last, 说明删除的是最后一个节点
		if next == cd.last {
			cd.last = prev
		}
	}
	cd.size--
	return node.ele
}

func (cd *circleDoublyLinkedList) Clear() {
	cd.size = 0
	cd.head = nil
	cd.last = nil
}

func (cd *circleDoublyLinkedList) Set(index int, element interface{}) interface{} {
	rangeCheck(cd.size, index)
	node := cd.node(index)
	oldValue := node.ele
	node.ele = element
	return oldValue
}

func (cd *circleDoublyLinkedList) Contains(element interface{}) bool {
	return cd.IndexOf(element) != ElementNotFound
}

func (cd *circleDoublyLinkedList) Get(index int) interface{} {
	rangeCheck(cd.size, index)
	return cd.node(index).ele
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

func (cd *circleDoublyLinkedList) node(index int) *doublyNode {
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
