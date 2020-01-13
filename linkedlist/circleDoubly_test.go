package linkedlist

import (
	"fmt"
	"testing"
)

var cd = NewCircleDoublyLinkedList()

func Test_circleDoublyLinkedList_Add(t *testing.T) {
	cd.Add(0, 111)
	cd.Add(0, 222)
	cd.Add(0, 333)
	for i := 5; i > 0; i-- {
		cd.AddHead(i)
	}

	cd.Add(1, 444)
	cd.Add(2, 555)
	for i := 6; i < 10; i++ {
		cd.AddTail(i)
	}

	fmt.Println(cd.GetSize(), cd.IsEmpty(), cd.Remove(2))
	traverseCircleDoublyLinkedList(cd.(*circleDoublyLinkedList))
}
