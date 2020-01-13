package linkedlist

type doublyNode struct {
	ele        interface{}
	prev, next *doublyNode
}

type doublyLinkedList struct {
	size int
	head *doublyNode
	last *doublyNode
}

// NewDoublyLinkedList 初始化
func NewDoublyLinkedList() ILinkedList {
	return &doublyLinkedList{head: nil}
}

func (d *doublyLinkedList) GetSize() int {
	return d.size
}

func (d *doublyLinkedList) IsEmpty() bool {
	return d.size == 0
}

func (d *doublyLinkedList) AddHead(element interface{}) {
	d.Add(0, element)
}

func (d *doublyLinkedList) AddTail(element interface{}) {
	d.Add(d.size, element)
}

// 添加节点
// 链表为空 和 前插 中插 后插 三种情况
// 中插和后插可以看作是往当前节点的前一个节点的追加操作
func (d *doublyLinkedList) Add(index int, element interface{}) {
	//  index < 0 || index > s.size  => out of range
	rangeCheckForAdd(d.size, index)
	if d.size == index { // 往后面添加节点
		oldLast := d.last
		node := &doublyNode{ele: element, prev: oldLast, next: nil}
		if oldLast == nil {
			d.head = node
			d.last = node
		} else {
			d.last = node
			oldLast.next = node
		}
	} else {
		next := d.node(index) // 后一个节点
		prev := next.prev     // 前一个节点
		// 创建新节点, prev指向原链表的尾节点, next指向首节点
		node := &doublyNode{ele: element, prev: prev, next: next}
		next.prev = node
		if prev != nil { // 非头节点插入
			prev.next = node
		} else {
			d.head = node
		}
	}
	d.size++
}

// 删除节点
// 前 中 后 删除
func (d *doublyLinkedList) Remove(index int) interface{} {
	rangeCheck(d.size, index)

	node := d.node(index)
	preNode := node.prev
	next := node.next
	if preNode != nil {
		preNode.next = preNode.next.next
	} else { // 头节点
		d.head = node.next
	}
	if next != nil {
		next.prev = next.prev.prev
	} else { // 尾节点
		d.last = preNode
	}
	d.size--
	return node.ele
}

func (d *doublyLinkedList) Clear() {
	d.size = 0
	d.head = nil
}

func (d *doublyLinkedList) Set(index int, element interface{}) interface{} {
	rangeCheck(d.size, index)
	node := d.node(index)
	oldValue := node.ele
	node.ele = element
	return oldValue
}

func (d *doublyLinkedList) Contains(element interface{}) bool {
	return d.IndexOf(element) != ElementNotFound
}

func (d *doublyLinkedList) Get(index int) interface{} {
	return d.node(index)
}

func (d *doublyLinkedList) IndexOf(element interface{}) int {
	node := d.head
	for i := 0; i < d.size; i++ {
		if node.ele == element {
			return i
		}
		node = node.next
	}
	return ElementNotFound
}

func (d *doublyLinkedList) node(index int) *doublyNode {
	if index < d.size>>1 { // 向后查找节点  d.size==7   则 d.size>>1 == 3
		node := d.head
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	}
	// 向前查找元素
	node := d.last
	for i := d.size - 1; i > index; i-- {
		node = node.prev
	}
	return node
}
