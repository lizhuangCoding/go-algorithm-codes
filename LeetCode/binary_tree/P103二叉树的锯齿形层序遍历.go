package main

import (
	"slices"
)

// 二叉树S形遍历（层序遍历）
// 力扣：https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isReverse := false

	for len(queue) != 0 {
		demo := make([]int, 0) // 存储值

		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			demo = append(demo, node.Val)

			// 添加子节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]

		// 反转
		if isReverse {
			slices.Reverse(demo)
		}
		result = append(result, demo)
		isReverse = !isReverse
	}

	return result
}

// func main() {
// 	root := &TreeNode{
// 		Val:  3,
// 		Left: &TreeNode{Val: 9},
// 		Right: &TreeNode{
// 			Val:   20,
// 			Left:  &TreeNode{Val: 15},
// 			Right: &TreeNode{Val: 7},
// 		},
// 	}
// 	fmt.Println(zigzagLevelOrder(root))
// }
