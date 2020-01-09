package linkedlist

import "fmt"

type doublyNode struct {
	ele        interface{}
	prev, next *doublyNode
}

type doublyLinkedList struct {
	size int
	head *doublyNode
	last *doublyNode
}

func newDoublyLinkedList() ILinkedList {
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
	rangeCheckForAdd(d.size, index)
	newNode := &doublyNode{ele: element}
	if d.head == nil {
		d.head = newNode
		d.last = newNode
	} else if index == 0 {
		d.head.prev = newNode
		newNode.next = d.head
		d.head = newNode
	} else {
		prevNode := d.getNode(index - 1) // 前驱节点
		nextNode := prevNode.next        // 后继节点
		if nextNode != nil {
			nextNode.prev = newNode
		}
		prevNode.next = newNode

		newNode.prev = prevNode
		newNode.next = nextNode
		if d.size == index {
			d.last = newNode
		}
	}
	d.size++
}

// 删除节点
// 前 中 后 删除
func (d *doublyLinkedList) Remove(index int) interface{} {
	rangeCheck(d.size, index)
	node := d.getNode(index)
	preNode := node.prev
	nextNode := node.next
	if preNode != nil {
		preNode.next = preNode.next.next
	} else { // 头节点
		fmt.Println("000-0000")
		d.head = node.next
	}
	if nextNode != nil {
		nextNode.prev = nextNode.prev.prev
	} else { // 尾节点
		fmt.Println("1111-111")
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
	node := d.getNode(index)
	oldValue := node.ele
	node.ele = element
	return oldValue
}

func (d *doublyLinkedList) Contains(element interface{}) bool {
	return d.IndexOf(element) != ElementNotFound
}

func (d *doublyLinkedList) Get(index int) interface{} {
	return d.getNode(index)
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

func (d *doublyLinkedList) getNode(index int) *doublyNode {
	if index <= d.size>>1 { // 向后查找节点  d.size==7   则 d.size>>1 == 3
		node := d.head
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	} else { // 向前查找元素
		node := d.last
		for i := d.size - 1; i > index; i-- {
			node = node.prev
		}
		return node
	}
}
