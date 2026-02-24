package main

import "math"

// 二叉树，层序遍历
// 力扣：https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/

func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	resultLevel := 1        // 最大和的层级
	nowLevel := 1           // 当前层级
	maxRes := math.MinInt64 // 最大和

	for len(queue) != 0 {
		l := len(queue)

		demoMaxRes := 0
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			demoMaxRes += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if maxRes < demoMaxRes {
			maxRes = demoMaxRes
			resultLevel = nowLevel
		}

		nowLevel++
	}

	return resultLevel
}
