package dataStructure

import "fmt"

// LinkedNode:linkedList
// 包含方法：Append、Display、Insert、Remove、Search

// ListNode 定义链表节点结构
type ListNode struct {
	Val  int       // 节点数据
	Next *ListNode // 指向下一个节点的指针
}

// Append 在链表末尾添加节点
func (node *ListNode) Append(val int) {
	newNode := &ListNode{Val: val, Next: nil} // 创建新节点
	current := node
	for current.Next != nil {
		current = current.Next // 找到链表的最后一个节点
	}
	current.Next = newNode // 将新节点链接到最后一个节点的next上
}

// Display 打印链表
func (node *ListNode) Display() {
	current := node // 从头节点开始遍历链表
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

// Insert 在指定位置插入节点
func (node *ListNode) Insert(index, val int) {
	if index < 0 {
		return
	}
	newNode := &ListNode{Val: val, Next: nil} // 创建新节点
	if index == 0 {
		newNode.Next = node // 将新节点插入头部
		node = newNode
		return
	}
	current := node
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next // 遍历到指定位置的前一个节点
	}
	if current == nil {
		return
	}
	newNode.Next = current.Next // 将新节点链接到当前节点的next上
	current.Next = newNode
}

// Remove 删除指定索引位置的节点
func (node *ListNode) Remove(index int) {
	if index < 0 || node == nil {
		return
	}
	if index == 0 {
		node = node.Next // 删除头节点
		return
	}
	current := node
	for i := 0; current != nil && i < index-1; i++ {
		current = current.Next // 找到要删除节点的前一个节点
	}
	if current == nil || current.Next == nil {
		return
	}
	current.Next = current.Next.Next // 删除节点
}

// // Remove 删除指定值的节点
// func (node *ListNode) Remove(val int) {
// 	if node == nil {
// 		return
// 	}
// 	if node.Val == val {
// 		node = node.Next // 删除头节点
// 		return
// 	}
// 	current := node
// 	for current.Next != nil && current.Next.Val != val {
// 		current = current.Next // 找到要删除节点的前一个节点
// 	}
// 	if current.Next != nil {
// 		current.Next = current.Next.Next // 删除节点
// 	}
// }

// Search 查找指定值的节点
func (node *ListNode) Search(val int) *ListNode {
	current := node
	for current != nil && current.Val != val {
		current = current.Next // 遍历链表，直到找到值为val的节点或者到达链表末尾
	}
	return current
}

// func main() {
// 	// 创建链表头节点
// 	head := &ListNode{Val: 1, Next: nil}
//
// 	// 在链表末尾添加节点
// 	head.Append(2)
// 	head.Append(3)
// 	head.Append(4)
//
// 	// 插入节点
// 	head.Insert(2, 10)
//
// 	// 删除节点
// 	head.Remove(3)
//
// 	// 查找节点
// 	found := head.Search(2)
// 	if found != nil {
// 		fmt.Println("节点 2 存在")
// 	} else {
// 		fmt.Println("节点 2 不存在")
// 	}
//
// 	// 打印链表
// 	fmt.Println("链表内容：")
// 	head.Display()
// }
