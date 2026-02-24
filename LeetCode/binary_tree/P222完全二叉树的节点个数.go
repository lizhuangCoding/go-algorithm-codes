package main

// 二叉树，遍历
// 力扣：https://leetcode.cn/problems/count-complete-tree-nodes/description/

func countNodes(root *TreeNode) int {
	return countNodesDigui(root)
}

func countNodesDigui(head *TreeNode) int {
	if head == nil {
		return 0
	}

	// 这里 return 的值表示的是：head节点并且包括head节点的子节点的总个数
	return countNodesDigui(head.Left) + countNodesDigui(head.Right) + 1

	// 也可以拆分开来：
	// leftNum  :=  countNodesDigui(head.Left)
	//	rightNum :=  countNodesDigui(head.Right)
	//	return leftNum + rightNum + 1
}
