package main

// 链表+二进制
// 力扣：https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/

// 相当于，比如说计算十进制的数：526，你先读取到52，如果发现后面还有一个6，那直接把52乘10加6就可以了
// 不过这个是二进制，所以要 * 2 + 后一位

// 链表的结点总数不超过 30；根据二进制的特性来写
func getDecimalValue(head *ListNode) int {
	sli := make([]int, 0)

	for head != nil {
		for i := 0; i < len(sli); i++ {
			sli[i] *= 2 // 每一个已经读取的二进制位都扩大2倍
		}
		sli = append(sli, head.Val)
		head = head.Next
	}

	res := 0
	for i := 0; i < len(sli); i++ {
		res += sli[i]
	}
	return res
}

// 空间优化：
// func getDecimalValue(head *ListNode) int {
// 	res := 0
// 	for head != nil {
// 		res += res*2 + head.Val
// 		head = head.Next
// 	}
// 	return res
// }
