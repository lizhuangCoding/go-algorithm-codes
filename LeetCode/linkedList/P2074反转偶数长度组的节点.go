package main

// 链表
// 力扣：https://leetcode.cn/problems/reverse-nodes-in-even-length-groups/description/

// 方法一：反转
// func reverseEvenLengthGroups(head *ListNode) *ListNode {
// 	dummy := &ListNode{Next: head}
// 	prev := dummy
// 	cur := head
// 	groupNum := 1
//
// 	for cur != nil {
// 		count := 0 // 计算当前组实际长度，判断是否需要反转
// 		temp := cur
// 		for i := 0; i < groupNum && temp != nil; i++ {
// 			temp = temp.Next
// 			count++
// 		}
//
// 		// 需要反转
// 		if count%2 == 0 {
// 			pre0 := prev // left-1 的位置
//
// 			// 反转 count 个节点
// 			var pre *ListNode
// 			start := cur
// 			for i := 0; i < count; i++ {
// 				demo := cur.Next
// 				cur.Next = pre
// 				pre = cur
// 				cur = demo
// 			}
//
// 			pre0.Next.Next = cur // 原头（现尾）连接到下一组
// 			pre0.Next = pre      // 上一组连接到反转后的头
// 			prev = start         // prev 更新为反转后的尾
// 		} else {
// 			// 不反转，跳过当前组
// 			for i := 0; i < count; i++ {
// 				prev = cur
// 				cur = cur.Next
// 			}
// 		}
//
// 		groupNum++
// 	}
//
// 	return dummy.Next
// }

// 方法二：交换元素值
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	var nodes []*ListNode

	node, size := head, 1
	for node != nil {
		nodes = append(nodes, node)

		if len(nodes) == size || node.Next == nil { // 统计到 size 个节点，或到达链表末尾
			n := len(nodes)
			// 有偶数个节点
			if n%2 == 0 {
				// 交换元素值
				for i := 0; i < n/2; i++ {
					nodes[i].Val, nodes[n-1-i].Val = nodes[n-1-i].Val, nodes[i].Val // 直接交换元素值
				}
			}

			nodes = nil // 清空切片
			size++
		}

		node = node.Next
	}
	return head
}
