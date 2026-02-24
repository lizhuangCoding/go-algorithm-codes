package main

// 二叉树
// 力扣：https://leetcode.cn/problems/invert-binary-tree/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTreeDfs(root)
	return root
}

func invertTreeDfs(head *TreeNode) {
	if head == nil {
		return
	}

	// 翻转
	head.Left, head.Right = head.Right, head.Left

	invertTreeDfs(head.Left)
	invertTreeDfs(head.Right)
}
