package linkedlist

type circleSingleLinkedList struct {
	size       int
	head  *singleNode
}

// NewCircleSingleLinkedList 初始化
func NewCircleSingleLinkedList() ILinkedList {
	return &circleSingleLinkedList{head: nil}
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

// 环形单链表节点插入操作
// 前 中 后三个位置插入，需要考虑链表为空的情况
func (cs *circleSingleLinkedList) Add(index int, element interface{}) {
	rangeCheckForAdd(cs.size, index)

	newNode := &singleNode{ele: element}
	if index == 0 { // 头插入
		// 获取最后一个节点
		lastNode := cs.getNode(cs.size -1 )
		if lastNode == nil {
			lastNode = newNode
		}
		// 更新头节点
		newNode.next = cs.head
		cs.head = newNode
		// 将尾节点的next指针指向头节点
		lastNode.next = cs.head
	} else {
		prevNode := cs.getNode(index - 1)
		next := prevNode.next
		prevNode.next = newNode
		newNode.next = next
	}

	cs.size++
}

// 删除节点
func (cs *circleSingleLinkedList) Remove(index int) interface{} {
	rangeCheck(cs.size, index)
	node := cs.head
	if index == 0 { // 删除头节点
		if cs.size == 1 {  // 只有一个元素
			cs.head = nil
		}else{
			node := cs.head
			cs.head = node.next
		}
	} else {  // 删除中间节点 或者 是尾节点
		prevNode := cs.getNode(index - 1)
		node = prevNode.next
		prevNode.next = node.next
	}
	cs.size--
	return node.ele
}

func (cs *circleSingleLinkedList) Clear() {
	cs.size = 0
	cs.head = nil
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
