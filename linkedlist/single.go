package linkedlist

type singleNode struct {
	ele  interface{}
	next *singleNode
}

type singleLinkedList struct {
	size int
	head *singleNode
}

// NewSingleLinkedList 初始化
func NewSingleLinkedList() ILinkedList {
	return &singleLinkedList{head: nil}
}

func (s *singleLinkedList) GetSize() int {
	return s.size
}

func (s *singleLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *singleLinkedList) AddHead(element interface{}) {
	s.Add(0, element)
}

func (s *singleLinkedList) AddTail(element interface{}) {
	s.Add(s.size, element)
}

// 添加节点
// 前  中  后  三种添加
// 链表为空或者中后添加都可以看作是元素追加
func (s *singleLinkedList) Add(index int, element interface{}) {
	rangeCheckForAdd(s.size, index)
	newNode := &singleNode{ele: element, next: s.head}
	if index == 0 {
		// 前插入  将新节点的next指向链表的头节点，然后再将链表的头节点更新为该新节点
		s.head = newNode
	} else {
		// 中间插入或者尾部插入
		prev := s.node(index - 1)
		newNode.next = prev.next
		prev.next = newNode
	}

	s.size++
}

// 删除节点
func (s *singleLinkedList) Remove(index int) (oldValue interface{}) {
	//  index < 0 || index > s.size  => out of range
	rangeCheck(s.size, index)
	node := s.head
	if index == 0 { //  删除头节点
		s.head = s.head.next
	} else { // 删除中间节点或者尾部节点
		prev := s.node(index - 1)
		node = prev.next
		prev.next = node.next
	}
	s.size--
	return node.ele
}

func (s *singleLinkedList) Clear() {
	s.size = 0
	s.head = nil
}

func (s *singleLinkedList) Set(index int, element interface{}) interface{} {
	rangeCheck(s.size, index)
	node := s.node(index)
	oldValue := node.ele
	node.ele = element
	return oldValue
}

func (s *singleLinkedList) Contains(element interface{}) bool {
	return s.IndexOf(element) != ElementNotFound
}

func (s *singleLinkedList) Get(index int) interface{} {
	rangeCheck(s.size, index)
	return s.node(index).ele
}

func (s *singleLinkedList) IndexOf(element interface{}) int {
	node := s.head
	i := 0
	for node != nil {
		if node.ele == element {
			return i
		}
		node = node.next
		i++
	}
	return ElementNotFound
}

// 根据节点位置查询节点
func (s *singleLinkedList) node(index int) *singleNode {
	node := s.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
