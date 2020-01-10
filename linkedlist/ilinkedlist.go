package linkedlist

// ILinkedList 链表公共接口定义
type ILinkedList interface {
	// 常规操作
	GetSize() int  // 元素的数量
	IsEmpty() bool // 是否为空

	// 节点添加
	AddHead(element interface{})
	AddTail(element interface{})
	Add(index int, element interface{}) // 添加元素到最后面

	// 节点删除
	Remove(index int) interface{} // 删除index位置对应的元素
	Clear()                       // 清除所有元素

	// 修改节点
	Set(index int, element interface{}) interface{} // 设置index位置的元素

	// 数据查询
	Contains(element interface{}) bool // 是否包含某个元素
	Get(index int) interface{}         //获取index位置的元素
	IndexOf(element interface{}) int   // 查看元素的位置
}

const (
	//ElementNotFound 找不到该元素
	ElementNotFound = -1
)
