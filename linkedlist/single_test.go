package linkedlist

import (
	"log"
	"testing"
)

var s = NewSingleLinkedList()

func Test_singleLinkedList_Add(t *testing.T) {
	for i := 4; i > 0; i-- {
		s.AddHead(i)
	}

	for i := 5; i < 10; i++ {
		s.AddTail(i)
	}
	s.Set(7, 99999)
	traverseSingleLinkedList(s.(*singleLinkedList))
	log.Println(s.GetSize(), s.Get(6), s.Contains(8))
}
func TestSingleLinkedList_Remove(t *testing.T) {
	s.AddTail(11)
	traverseSingleLinkedList(s.(*singleLinkedList))
	s.Remove(0)
}

//func TestSingleLinkedList_Remove(t *testing.T) {
//	//s.Remove(1) // panic : out of range
//
//	// insert element
//	for i := 1; i < 10; i++ {
//		s.AddTail(i)
//	}
//	traverseSingleLinkedList(s.(*singleLinkedList))
//	// 第一个节点
//	if s.Remove(0) != 1 {
//		log.Fatal("remove index=0 is failed")
//	}
//	// 最后一个节点
//	if s.Remove(s.GetSize()-1) != 9 {
//		log.Fatal("remove index=", s.GetSize()-1, " is failed")
//	}
//	if s.Remove(3) != 5 {
//		log.Fatal("remove index=3 is failed")
//	}
//	traverseSingleLinkedList(s.(*singleLinkedList))
//}
func TestSingleLinkedList_Clear(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		s.AddHead(i)
	}
	log.Println("before clear:", s.GetSize(), s.IsEmpty(), s.IndexOf(2))
	s.Clear()
	log.Println("after clear:", s.GetSize(), s.IsEmpty())
}
func BenchmarkSingleLinkedList_AddHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.AddTail(i)
	}
}
