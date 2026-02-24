package main

// 链表
// 力扣：https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

// 方法一：暴力
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	for headA != nil {
// 		demo := headB // 要使用中间变量，不能直接用headB
// 		for demo != nil {
// 			if demo == headA {
// 				return demo
// 			}
// 			demo = demo.Next
// 		}
//
// 		headA = headA.Next
// 	}
// 	return nil
// }

// 方法二：计算长度
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	lenA, lenB := 0, 0
// 	demoA, demoB := headA, headB
// 	for demoA != nil {
// 		demoA = demoA.Next
// 		lenA++
// 	}
// 	for demoB != nil {
// 		demoB = demoB.Next
// 		lenB++
// 	}
//
// 	demoA, demoB = headA, headB
// 	// 让两条链表到达长度持平的位置
// 	if lenA > lenB {
// 		for i := 1; i <= lenA-lenB; i++ {
// 			demoA = demoA.Next
// 		}
// 	} else {
// 		for i := 1; i <= lenB-lenA; i++ {
// 			demoB = demoB.Next
// 		}
// 	}
//
// 	for demoA != demoB {
// 		demoA = demoA.Next
// 		demoB = demoB.Next
// 	}
// 	return demoA
// }

// 或者：
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	lenA, lenB := 0, 0
// 	demoA, demoB := headA, headB
// 	for demoA != nil {
// 		lenA++
// 		demoA = demoA.Next
// 	}
// 	for demoB != nil {
// 		lenB++
// 		demoB = demoB.Next
// 	}
//
// 	// 假设headA长，如果不是headA长，我们让headA变为长的那一个
// 	demoA, demoB = headA, headB
// 	if lenA < lenB {
// 		lenA, lenB = lenB, lenA
// 		demoA, demoB = demoB, demoA
// 	}
//
// 	for i := 0; i < lenA-lenB; i++ {
// 		demoA = demoA.Next
// 	}
//
// 	for demoA != nil && demoB != nil {
// 		if demoA == demoB {
// 			return demoA
// 		}
// 		demoA = demoA.Next
// 		demoB = demoB.Next
// 	}
//
// 	return nil
// }

// 方法三：哈希集合
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	vis := map[*ListNode]bool{}
// 	for tmp := headA; tmp != nil; tmp = tmp.Next {
// 		vis[tmp] = true
// 	}
//
// 	for tmp := headB; tmp != nil; tmp = tmp.Next {
// 		if vis[tmp] {
// 			return tmp
// 		}
// 	}
// 	return nil
// }

// 方法四：双指针
// 指针 n1 遍历链表A，然后切换到链表B；指针 n2 遍历链表B，然后切换到链表A
// 如果两个链表相交，两个指针最终会在交点相遇；如果不想交，两个指针最终都会指向nil。
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	n1, n2 := headA, headB

	for n1 != n2 {
		if n1 != nil {
			n1 = n1.Next
		} else {
			n1 = headB
		}

		if n2 != nil {
			n2 = n2.Next
		} else {
			n2 = headA
		}
	}

	return n1
}
