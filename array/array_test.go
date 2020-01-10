package array

import (
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
	t.Logf("数组容量为：%d\n", arr.GetCapacity())
	t.Log(arr.IsEmpty(), arr.Contains(44), arr.IndexOf(44))

	arr.Remove(0)
	arr.Traverse()
	arr.Set(1, 22)
	t.Logf("数组容量为：%d\n", arr.GetCapacity())
	t.Log(arr.GetLength(), arr.IsEmpty(), arr.Contains(44), arr.IndexOf(44))
	arr.Clear()
	t.Log(arr.GetLength(), arr.IsEmpty())
}

func BenchmarkArray_AddHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr.AddHead(i)
	}
}
