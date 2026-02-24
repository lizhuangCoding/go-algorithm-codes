package dataStructure

import (
	"fmt"
	"sync"
)

// 单链表
// 包含方法：Init、Append、Insert、Delete、Get、Display

// SingleObject 节点数据
type SingleObject interface{}

// SingleNode 单链表节点
type SingleNode struct {
	Data SingleObject
	Next *SingleNode
}

// SingleList 单链表
type SingleList struct {
	mutex *sync.RWMutex
	Head  *SingleNode
	Tail  *SingleNode
	Size  uint
}

// Init 初始化
func (list *SingleList) Init() {
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	list.mutex = new(sync.RWMutex)
}

// Append 新增节点，链表节点的新增分为两种，append：在链表后面追加节点；insert：在指定位置插入节点。 新增时，若为第一个节点需特殊处理一下。
// 添加节点到链表尾部
func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1
		return true
	}

	tail := list.Tail
	tail.Next = node
	list.Tail = node
	list.Size += 1
	return true
}

// Insert 插入节点到指定位置
func (list *SingleList) Insert(index uint, node *SingleNode) bool {
	if node == nil || index > list.Size {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Size += 1
		return true
	}
	var i uint
	ptr := list.Head
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	ptr.Next = node
	node.Next = next
	list.Size += 1
	return true
}

// Delete 删除指定位置的节点，如果指定的位置是链表的头部或尾部，都需要特殊处理下
func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		head := list.Head.Next
		list.Head = head
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size -= 1
		return true
	}

	ptr := list.Head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next

	ptr.Next = next.Next
	if index == list.Size-1 {
		list.Tail = ptr
	}
	list.Size -= 1
	return true
}

// Get 查询节点，获取指定位置的节点，不存在则返回nil
func (list *SingleList) Get(index uint) *SingleNode {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return nil
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()

	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 0; i < index; i++ {
		node = node.Next
	}
	return node
}

// Display 打印链表
func (list *SingleList) Display() {
	if list == nil {
		fmt.Println("this single list is nil")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this single list size is %d \n", list.Size)
	ptr := list.Head
	var i uint
	for i = 0; i < list.Size; i++ {
		fmt.Printf("No%3d data is %v\n", i+1, ptr.Data)
		ptr = ptr.Next
	}
}
