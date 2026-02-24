package main

// 链表
// 力扣：https://leetcode.cn/problems/reverse-linked-list/description/

// 方法一：迭代
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	var temp *ListNode

	for cur != nil {
		temp = cur.Next // 记录cur的下一个节点
		cur.Next = prev // 更改cur的Next要指向的节点
		prev = cur      // 移动 prev
		cur = temp      // 移动cur
	}
	return prev // prev最后会指向头节点
}

// 方法二：递归
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
//
// 	newHead := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return newHead
// }
// 看不懂的话，模拟一下过程：
// 以链表 1 -> 2 -> 3 -> 4 -> 5 -> nil 为例，详细分析递归的执行流程。
//
// 递归深入阶段：
// 每次递归将链表拆分，直到链表末尾：
//
// 第1步：head = 1，递归调用 reverseList(2 -> 3 -> 4 -> 5 -> nil)。
// 第2步：head = 2，递归调用 reverseList(3 -> 4 -> 5 -> nil)。
// 第3步：head = 3，递归调用 reverseList(4 -> 5 -> nil)。
// 第4步：head = 4，递归调用 reverseList(5 -> nil)。
// 第5步：head = 5，递归终止条件满足，返回 head（即 5 -> nil）。
// 递归回溯阶段：
// 从最深层逐步回溯，完成链表的反转：
//
// 第5步（回溯）：
//
// 当前 head = 4。
// head.Next = 5，让 5.Next = 4，即 5 -> 4。
// 切断 4 的旧连接：head.Next = nil，链表变为 5 -> 4 -> nil。
// 返回 newHead = 5。
// 第4步（回溯）：
//
// 当前 head = 3。
// head.Next = 4，让 4.Next = 3，即 4 -> 3。
// 切断 3 的旧连接：head.Next = nil，链表变为 5 -> 4 -> 3 -> nil。
// 返回 newHead = 5。
// 第3步（回溯）：
//
// 当前 head = 2。
// head.Next = 3，让 3.Next = 2，即 3 -> 2。
// 切断 2 的旧连接：head.Next = nil，链表变为 5 -> 4 -> 3 -> 2 -> nil。
// 返回 newHead = 5。
// 第2步（回溯）：
//
// 当前 head = 1。
// head.Next = 2，让 2.Next = 1，即 2 -> 1。
// 切断 1 的旧连接：head.Next = nil，链表变为 5 -> 4 -> 3 -> 2 -> 1 -> nil。
// 返回 newHead = 5。
