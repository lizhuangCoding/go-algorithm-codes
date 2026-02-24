package main

// 链表
// 力扣：https://leetcode.cn/problems/design-linked-list/

type MyLinkedList struct {
	size int
	head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, &ListNode{}}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	if index < 0 {
		index = 0
	}

	l.size++
	prev := l.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	toAdd := &ListNode{val, prev.Next}
	prev.Next = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}

	l.size--
	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}
