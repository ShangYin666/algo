package linkedlist

import (
	"fmt"
	"strings"
)

func rangeCheck(target int, index int) {
	if index < 0 || index >= target {
		panic("Out Of Range")
	}
}

func rangeCheckForAdd(target int, index int) {
	if index < 0 || index > target {
		panic("Out Of Range")
	}
}
func traverseSingleLinkedList(s *singleLinkedList) {
	node := s.head
	str := ""
	fmt.Println(strings.Repeat("-", 120))
	fmt.Printf("%v\t\t%v\t\t%v\t\t\t%v\n", "节点下标", "节点元素", "节点地址", "该节点的next")
	i := 0
	for node != nil {
		fmt.Printf("%v\t\t\t%v\t\t\t%p\t%p\n", i, node.ele, node, node.next)
		str += fmt.Sprintf("%v->", node.ele)
		node = node.next
		i++
	}
	fmt.Println(strings.Repeat("-", 120))
	fmt.Printf("[%d] ： %v\n", s.size, strings.Trim(str, "->"))
}

func traverseDoublyLinkedList(s *doublyLinkedList) {
	node := s.head
	str := ""
	fmt.Println(strings.Repeat("-", 120))
	fmt.Printf("%v\t\t%v\t\t%v\t\t\t%v\t\t\t%v\n", "节点下标", "节点元素", "节点地址", "该节点的prev", "该节点的next")
	i := 0
	for node != nil {
		fmt.Printf("%v\t\t\t%v\t\t\t%p\t%p\t%p\n", i, node.ele, node, node.prev, node.next)
		str += fmt.Sprintf("%v->", node.ele)
		node = node.next
		i++
	}
	fmt.Println(strings.Repeat("-", 120))
	if s.head != nil {
		fmt.Printf("head:%v %p\n", s.head.ele, s.head)
	}
	if s.last != nil {
		fmt.Printf("last:%v %p\n", s.last.ele, s.last)
	}
	fmt.Printf("%d ： [ %v ]\n", s.size, strings.Trim(str, "->"))
}

func traverseCircleSingleLinkedList(s *circleSingleLinkedList) {
	node := s.head
	str := ""
	fmt.Println(strings.Repeat("-", 120))
	fmt.Printf("%v\t\t%v\t\t%v\t\t\t%v\n", "节点下标", "节点元素", "节点地址", "该节点的next")
	for i := 0; i < s.size; i++ {
		fmt.Printf("%v\t\t\t%v\t\t\t%p\t%p\n", i, node.ele, node, node.next)
		str += fmt.Sprintf("%v->", node.ele)
		node = node.next
	}
	fmt.Println(strings.Repeat("-", 120))
	if s.head != nil {
		fmt.Printf("head:%v %p %v\n", s.head.ele, s.head, s.head.next.ele)
	}

	fmt.Printf("[%d] ： %v\n", s.size, strings.Trim(str, "->"))
}

func traverseCircleDoublyLinkedList(s *circleDoublyLinkedList) {
	node := s.head
	str := ""
	fmt.Println(strings.Repeat("-", 120))
	fmt.Printf("%v\t\t%v\t\t%v\t\t\t%v\t\t\t%v\n", "节点下标", "节点元素", "节点地址", "该节点的prev", "该节点的next")
	for i := 0; i < s.size; i++ {
		fmt.Printf("%v\t\t\t%v\t\t\t%p\t%p\t%p\n", i, node.ele, node, node.prev, node.next)
		str += fmt.Sprintf("%v->", node.ele)
		node = node.next
	}
	fmt.Println(strings.Repeat("-", 120))
	if s.head != nil {
		fmt.Printf("head:%v %p %v\n", s.head.ele, s.head, s.head.next.ele)
	}
	if s.last != nil {
		fmt.Printf("last:%v %p\n", s.last.ele, s.last)
	}
	fmt.Printf("[%d] ： %v\n", s.size, strings.Trim(str, "->"))
}
