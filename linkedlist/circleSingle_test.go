package linkedlist

import (
	"testing"
)

var cs = NewCircleSingleLinkedList()

func Test_circleSingleLinkedList_Add(t *testing.T) {
	//
	for i := 5; i >= 0; i-- {
		cs.AddHead(i)
	}
	cs.Add(5, 111)
	for i := 6; i < 10; i++ {
		cs.AddTail(i)
	}
	t.Log(cs.Get(8))

	traverseCircleSingleLinkedList(cs.(*circleSingleLinkedList))
}
func TestCircleSingleLinkedList_Remove(t *testing.T) {
	for i := 5; i > 0; i-- {
		cs.AddHead(i)
	}
	cs.Add(5, 111)
	for i := 6; i < 10; i++ {
		cs.AddTail(i)
	}
	cs.Remove(0)
	cs.Remove(4)
	cs.Remove(cs.GetSize() - 1)

	traverseCircleSingleLinkedList(cs.(*circleSingleLinkedList))

	cs.Clear()

	cs.AddHead(1)
	cs.Remove(0)
	traverseCircleSingleLinkedList(cs.(*circleSingleLinkedList))

}

func TestCircleSingleLinkedList_Normal(t *testing.T) {
	for i := 5; i > 0; i-- {
		cs.AddHead(i)
	}
	cs.Add(5, 111)
	for i := 6; i < 10; i++ {
		cs.AddTail(i)
	}
	t.Log(cs.Get(3))
	t.Log(cs.Set(3, 333))
	t.Log(cs.Get(3))
	t.Log(cs.IsEmpty())
	t.Log(cs.Contains(3))
	t.Log(cs.Contains(3666))

}
