package main

// 二叉树，bfs
// 力扣：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }
//
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
//
// 	queue := make([]*Node, 0)
// 	queue = append(queue, root)
//
// 	for len(queue) > 0 {
// 		l := len(queue)
//
// 		for i := 0; i < l; i++ {
// 			if i != l-1 {
// 				queue[i].Next = queue[i+1]
// 			}
//
// 			if queue[i].Left != nil {
// 				queue = append(queue, queue[i].Left)
// 			}
// 			if queue[i].Right != nil {
// 				queue = append(queue, queue[i].Right)
// 			}
//
// 		}
// 		queue = queue[l:]
// 	}
//
// 	return root
// }
