package main

// 二叉树，dfsNumColor
// 力扣：https://leetcode.cn/problems/path-sum/description/

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumDfs(root, root.Val, targetSum)
}

func hasPathSumDfs(head *TreeNode, n int, targetSum int) bool {
	// 叶子节点是指没有子节点的节点。
	if head.Left == nil && head.Right == nil {
		return targetSum == n
	}

	if head.Left != nil && hasPathSumDfs(head.Left, n+head.Left.Val, targetSum) {
		return true
	}

	if head.Right != nil && hasPathSumDfs(head.Right, n+head.Right.Val, targetSum) {
		return true
	}
	return false
}

// 或者用减法：
// func hasPathSum(root *TreeNode, targetSum int) bool {
// 	if root == nil {
// 		return false
// 	}
// 	targetSum -= root.Val
// 	if root.Left == root.Right { // root 是叶子
// 		return targetSum == 0
// 	}
// 	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
// }
