package linkedlist

import (
	"fmt"
	"testing"
)

var d = NewDoublyLinkedList()

func Test_doublyLinkedList_Add(t *testing.T) {
	for i := 8; i > 0; i-- {
		d.AddHead(i)
	}

	//for i := 5; i < 10; i++ {
	//	d.AddTail(i)
	//}
	//d.Set(7, 99999)
	//d.Add(d.GetSize()-1, 000)

	traverseDoublyLinkedList(d.(*doublyLinkedList))
	//log.Println(d.GetSize(), d.Get(3), d.Contains(8), d.IndexOf(1))

}

func TestDoublyLinkedList_Remove(t *testing.T) {
	for i := 0; i < 10; i++ {
		d.AddTail(i)
	}

	traverseDoublyLinkedList(d.(*doublyLinkedList))
	fmt.Println(d.GetSize())
	fmt.Println("删除头节点：", d.Remove(0))
	fmt.Println("删除尾节点：", d.Remove(d.GetSize()-1))
	fmt.Println("删除节点：", d.Remove(1))
	//
	traverseDoublyLinkedList(d.(*doublyLinkedList))
	fmt.Println(d.IsEmpty())
}

func BenchmarkDoublyLinkedList_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d.AddHead(i)
	}
}
