package main

// 链表，归并排序
// 力扣：https://leetcode.cn/problems/sort-list/description/

// 归并排序
// 时间：O(n*logn)
// 空间：O(n)
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

// 合并
func merge(n1, n2 *ListNode) *ListNode {
	dummyNode := &ListNode{}

	demo := dummyNode
	for n1 != nil && n2 != nil {
		if n1.Val <= n2.Val {
			demo.Next = n1
			n1 = n1.Next
		} else {
			demo.Next = n2
			n2 = n2.Next
		}
		demo = demo.Next
	}

	if n1 != nil {
		demo.Next = n1
	}
	if n2 != nil {
		demo.Next = n2
	}

	return dummyNode.Next
}

/*

例子：
4->2->1->3

第一轮：
mid = slow = 1
左边：[4,1)  右边：[1,nil)

第二轮：
。。。


*/
