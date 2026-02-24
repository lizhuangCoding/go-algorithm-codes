package main

// 链表
// 力扣：https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur != nil && cur.Next != nil && cur.Next.Next != nil {

		if cur.Next.Val == cur.Next.Next.Val { // 后两个节点值相同
			val := cur.Next.Val

			// 把值等于 val 的节点全部删除
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}

		} else {
			// 必须要放在else里面，否则无法检查新的cur.Next是否还是重复的（比如连续多组重复的情况，如 3 3 4 4 的情况就会把 4 4 漏掉）
			cur = cur.Next
		}
	}
	return dummy.Next
}

/*

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]


输入：head = [1,1,1,2,3]
输出：[2,3]

*/
