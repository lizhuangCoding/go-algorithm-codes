package main

// 链表
// 力扣：https://leetcode.cn/problems/split-linked-list-in-parts/description/

// 遍历模拟：
func splitListToParts(head *ListNode, k int) []*ListNode {
	if k == 1 {
		return []*ListNode{head}
	}

	// 计算链表总长度
	l := 0
	demo := head
	for demo != nil {
		demo = demo.Next
		l++
	}

	result := make([]*ListNode, 0)
	begin := head

	for k >= 1 {
		// 计算出每一轮循环中，本次应该分配的节点数：n
		n := l / k
		// 如果有多余的节点，多分配一个
		if l%k != 0 {
			n++
		}

		demo := begin
		for i := 0; i < n-1; i++ {
			demo = demo.Next
		}

		// 使用临时变量 demo2 作为下一轮循环的begin，防止出现空指针错误
		// 输入：head = [1,2,3], k = 5
		// 输出：[[1],[2],[3],[],[]]
		var demo2 *ListNode = nil
		if demo != nil {
			demo2 = demo.Next
			demo.Next = nil
		}

		result = append(result, begin)
		begin = demo2

		l -= n
		k--
	}

	return result
}

// 优化：
// func splitListToParts(head *ListNode, k int) []*ListNode {
// 	n := 0
// 	for node := head; node != nil; node = node.Next {
// 		n++
// 	}
//
// 	quotient, remainder := n/k, n%k
//
// 	parts := make([]*ListNode, k)
// 	for i, curr := 0, head; i < k && curr != nil; i++ {
// 		parts[i] = curr
//
// 		partSize := quotient
// 		if i < remainder {
// 			partSize++
// 		}
// 		for j := 1; j < partSize; j++ {
// 			curr = curr.Next
// 		}
// 		curr, curr.Next = curr.Next, nil
// 	}
// 	return parts
// }
