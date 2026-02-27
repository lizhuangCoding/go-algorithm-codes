package main

import "slices"

// 二叉树，bfs
// 力扣：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	result := make([][]int, 0)

	for len(queue) > 0 {

		l := len(queue)
		demo := make([]int, 0)

		for i := 0; i < l; i++ {
			node := queue[i]
			demo = append(demo, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[l:]
		result = append(result, demo)
	}

	// 反转结果
	slices.Reverse(result)

	return result
}
