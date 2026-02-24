package main

// 二叉树，dfs
// 力扣：https://leetcode.cn/problems/sum-of-left-leaves/

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		}

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return res
}
