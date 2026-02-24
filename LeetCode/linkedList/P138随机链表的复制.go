package main

// type Node struct {
// 	Val    int
// 	Next   *Node
// 	Random *Node
// }

// 链表
// 力扣：https://leetcode.cn/problems/copy-list-with-random-pointer/description/

// 方法一：哈希+暴力
// func copyRandomList(head *Node) *Node {
// 	if head == nil {
// 		return nil
// 	}
//
// 	// 创建映射：原节点 -> 新节点
// 	m := make(map[*Node]*Node)
//
// 	// 第一遍遍历：创建所有新节点，建立映射
// 	demo := head
// 	for demo != nil {
// 		m[demo] = &Node{Val: demo.Val}
// 		demo = demo.Next
// 	}
//
// 	// 第二遍遍历：连接Next和Random指针
// 	demo = head
// 	for demo != nil {
// 		// 连接Next
// 		if demo.Next != nil {
// 			m[demo].Next = m[demo.Next]
// 		}
// 		// 连接Random
// 		if demo.Random != nil {
// 			m[demo].Random = m[demo.Random]
// 		}
//
// 		demo = demo.Next
// 	}
//
// 	return m[head]
// }

// 方法二：
// 把原节点：1->2->3->4 改为 1->11->2->22->3->33->4->44，然后添加Random的指向，最后得到结果：11->22->33->44
// 题解：https://leetcode.cn/problems/copy-list-with-random-pointer/solutions/2993775/bu-yong-ha-xi-biao-de-zuo-fa-pythonjavac-nzdo/?envType=study-plan-v2&envId=top-interview-150
// 时间复杂度：O(n)，其中 n 是链表的长度。
// 空间复杂度：O(1)。返回值不计入。
// func copyRandomList(head *Node) *Node {
// 	// 复制每个节点，把新节点直接插到原节点的后面
// 	for cur := head; cur != nil; cur = cur.Next.Next {
// 		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
// 	}
//
// 	// 遍历交错链表中的原链表节点
// 	for cur := head; cur != nil; cur = cur.Next.Next {
// 		if cur.Random != nil {
// 			// 要复制的 random 是 cur.Random 的下一个节点
// 			cur.Next.Random = cur.Random.Next
// 		}
// 	}
//
// 	// 把交错链表分离成两个链表
// 	dummy := Node{}
// 	tail := &dummy
// 	for cur := head; cur != nil; cur, tail = cur.Next, tail.Next {
// 		clone := cur.Next     // 新节点
// 		tail.Next = clone     // 把新节点插在 tail 的后面，构建新的链表
// 		cur.Next = clone.Next // 恢复原节点的 next
// 	}
//
// 	return dummy.Next
// }
