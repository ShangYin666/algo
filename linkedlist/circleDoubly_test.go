package linkedlist

import "testing"

var cd = NewCircleDoublyLinkedList()

func Test_circleDoublyLinkedList_Add(t *testing.T) {
	cd.AddHead(5)
	cd.AddHead(4)
	//cd.AddHead(3)
	//cd.AddHead(2)
	traverseCircleDoublyLinkedList(cd.(*circleDoublyLinkedList))
}
