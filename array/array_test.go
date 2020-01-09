package array

import (
	"fmt"
	"strings"
	"testing"
)

var arr = NewArray(1)

func TestArray_Add(t *testing.T) {
	for i := 5; i > 0; i-- {
		arr.AddHead(i)
	}
	for i := 6; i < 11; i++ {
		arr.AddTail(i)
	}
	arr.AddHead(0)
	arr.Traverse()
	fmt.Printf("数组容量为：%d\n", arr.GetCapacity())
	fmt.Println(arr.IsEmpty(), arr.Contains(44), arr.IndexOf(44))

	fmt.Println(strings.Repeat("-", 200))
	arr.Remove(0)
	arr.Traverse()
	arr.Set(1, 22)
	fmt.Printf("数组容量为：%d\n", arr.GetCapacity())
	fmt.Println(arr.GetLength(), arr.IsEmpty(), arr.Contains(44), arr.IndexOf(44))
	arr.Clear()
	fmt.Println(arr.GetLength(), arr.IsEmpty())
}

func BenchmarkArray_AddHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr.AddHead(i)
	}
}
