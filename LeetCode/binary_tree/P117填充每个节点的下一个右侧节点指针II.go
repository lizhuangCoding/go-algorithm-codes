package main

// 二叉树，bfs，链表
// 力扣：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

// 方法一：bfs
// 时间复杂度：O(n)，其中 n 为二叉树的节点个数。
// 空间复杂度：O(n)。
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
//
// 	sli := make([]*Node, 0)
// 	sli = append(sli, root)
//
// 	for len(sli) != 0 {
// 		l := len(sli)
//
// 		for i := 0; i < l; i++ {
// 			if i < l-1 {
// 				sli[i].Next = sli[i+1]
// 			}
//
// 			if sli[i].Left != nil {
// 				sli = append(sli, sli[i].Left)
// 			}
// 			if sli[i].Right != nil {
// 				sli = append(sli, sli[i].Right)
// 			}
// 		}
// 		sli = sli[l:]
//
// 	}
// 	return root
// }

// 方法二：BFS+链表（原理：在方法一的基础上，因为方法一用的是一个切片存储每一次大循环的每一个节点，但是我们可以利用链表的特性）
// 时间复杂度：O(n)，其中 n 为二叉树的节点个数。
// 空间复杂度：O(1)。只用到若干额外变量。
// func connect(root *Node) *Node {
// 	cur := root
//
// 	for cur != nil {
// 		dummy := &Node{} // 下一层链表的头节点的前一个虚拟节点
// 		nxt := dummy     // 用nxt为这个虚拟链表赋值
//
// 		for cur != nil {
// 			if cur.Left != nil {
// 				nxt.Next = cur.Left
// 				nxt = nxt.Next
// 			}
// 			if cur.Right != nil {
// 				nxt.Next = cur.Right
// 				nxt = nxt.Next
// 			}
// 			cur = cur.Next
// 		}
// 		// 更改 cur 为下一层的链表
// 		cur = dummy.Next
// 	}
// 	return root
// }
