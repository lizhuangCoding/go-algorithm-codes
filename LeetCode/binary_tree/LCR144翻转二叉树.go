package main

// 二叉树
// 力扣：https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/

func flipTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left

	flipTree(root.Left)
	flipTree(root.Right)

	return root
}
