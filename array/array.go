package array

import (
	"fmt"
	_ "runtime/pprof"
	"strings"
)

// Array 数组定义
type Array struct {
	data   []int
	length int
}

const (
	// DefaultArrayCapacity 数组默认容量
	DefaultArrayCapacity = 10
	// ElementNotFound 元素未找到
	ElementNotFound = -1
)

// NewArray 初始化数组
func NewArray(capacity int) *Array {
	if capacity == 0 || capacity < DefaultArrayCapacity {
		capacity = DefaultArrayCapacity
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

// Traverse 遍历数组
func (a *Array) Traverse() {
	str := ""
	for i := 0; i < a.length; i++ {
		str += fmt.Sprintf("%v->", a.data[i])
	}
	fmt.Printf("数组长度为：%d,数组结构为：%v\n", a.length, strings.TrimRight(str, "->"))
}

// GetLength 获取长度
func (a *Array) GetLength() int {
	return a.length
}

// GetCapacity 获取容量
func (a *Array) GetCapacity() int {
	return cap(a.data)
}

// IsEmpty 数组是否为空
func (a *Array) IsEmpty() bool {
	return a.length == 0
}

// Add 新增元素
// TODO:新增元素
func (a *Array) Add(index int, element int) {
	// 边界检查
	a.rangeCheckForAdd(index)
	// 扩容
	a.ensure()
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	// 添加
	a.data[index] = element
	a.length++
}

// AddHead 新增头元素
func (a *Array) AddHead(element int) {
	a.Add(0, element)
}

// AddTail 新增尾元素
func (a *Array) AddTail(element int) {
	a.Add(a.length, element)
}

// Remove 删除元素
// TODO:删除元素
func (a *Array) Remove(index int) {
	a.rangeCheck(index)

	for i := index; i < a.length; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	a.trim()

}

// Clear 清空数组
func (a *Array) Clear() {
	a.length = 0
	a.data = nil
}

// Set 修改元素
func (a *Array) Set(index int, element int) int {
	ele := a.getNode(index)
	a.data[index] = element
	return ele
}

// Get 查询元素
func (a *Array) Get(index int) int {
	return a.getNode(index)
}

// IndexOf 根据元素查询下标
func (a *Array) IndexOf(element int) int {
	for i := 0; i < a.length; i++ {
		if a.data[i] == element {
			return i
		}
	}
	return ElementNotFound
}

// Contains 判断数组是否包含某个元素
func (a *Array) Contains(element int) bool {
	return a.IndexOf(element) != ElementNotFound
}

func (a *Array) getNode(index int) int {
	a.rangeCheck(index)
	var element int
	for i := 0; i < index; i++ {
		element = a.data[i]
	}
	return element
}

func (a *Array) rangeCheck(index int) {
	if 0 > index || index >= a.length {
		panic("Out Of Range")
	}
}

func (a *Array) rangeCheckForAdd(index int) {
	if index < 0 || index > a.length {
		panic("Out Of Range")
	}
}

// 数组扩容
func (a *Array) ensure() {
	capacity := cap(a.data)
	if a.length == capacity {
		oldArr := a.data
		capacity = capacity << 1
		a.data = make([]int, capacity, capacity)
		copy(a.data, oldArr)
	}
}

// 数组缩容
func (a *Array) trim() {
	capacity := cap(a.data)
	fmt.Println(a.length, capacity>>1)
	if a.length <= capacity>>1 && capacity > DefaultArrayCapacity {
		oldArr := a.data
		capacity = capacity >> 1
		a.data = make([]int, capacity, capacity)
		copy(a.data, oldArr)
	}
}
