package main

// 链表
// 力扣：https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/

// 方法一：数组
func pairSum(head *ListNode) int {
	sli := make([]int, 0)
	for head != nil {
		sli = append(sli, head.Val)
		head = head.Next
	}

	res := 0
	for i := 0; i < len(sli)/2; i++ {
		res = max(res, sli[i]+sli[len(sli)-1-i])
	}
	return res
}

// // 方法二：快慢指针
// func pairSum(head *ListNode) int {
//
// }
