package main

import "math"

// 二叉搜索树，中序遍历
// 力扣：https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/

// 时间复杂度：O(n)，其中 n 为二叉搜索树节点的个数。每个节点在中序遍历中都会被访问一次且只会被访问一次，因此总时间复杂度为 O(n)。
// 空间复杂度：O(n)。递归函数的空间复杂度取决于递归的栈深度，而栈深度在二叉搜索树为一条链的情况下会达到 O(n) 级别。
func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1 // pre 用于记录前一个遍历到的节点值

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		// 如果 pre 不为 -1，说明已经遍历过至少一个节点
		// 计算当前节点值与前一个节点值的差值，并更新最小差值 ans
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}

		pre = node.Val

		dfs(node.Right)
	}

	dfs(root)
	return ans
}

/*

例子：
    4
   / \
  2   6
 / \
1   3

中序遍历过程：
初始化：ans = math.MaxInt64，pre = -1。

遍历左子树：
首先遍历到节点 1，由于 pre = -1，不进行差值计算，更新 pre = 1。
接着遍历到节点 2，计算差值 2 - 1 = 1，因为 1 < ans，更新 ans = 1，更新 pre = 2。
然后遍历到节点 3，计算差值 3 - 2 = 1，由于 1 = ans，不更新 ans，更新 pre = 3。

遍历根节点：
遍历到节点 4，计算差值 4 - 3 = 1，由于 1 = ans，不更新 ans，更新 pre = 4。

遍历右子树：
遍历到节点 6，计算差值 6 - 4 = 2，由于 2 > ans，不更新 ans。

最终结果：遍历结束后，最小差值 ans = 1，因此函数返回 1。

*/
