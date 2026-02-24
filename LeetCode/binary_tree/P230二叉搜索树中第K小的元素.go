package main

// 二叉搜索树
// 力扣：https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/

// 中序遍历：记录答案
func kthSmallest(root *TreeNode, k int) int {
	result := 0

	var dfs func(head *TreeNode)
	dfs = func(head *TreeNode) {
		if head == nil {
			return
		}

		dfs(head.Left)

		k--
		if k == 0 {
			result = head.Val
		}

		dfs(head.Right)
	}

	dfs(root)
	return result
}

// 优化：不记录答案 + 提前返回
// func kthSmallest(root *TreeNode, k int) int {
// 	var dfsNumColor func(*TreeNode) int
// 	dfsNumColor = func(node *TreeNode) int {
// 		if node == nil {
// 			return -1 // 题目保证节点值非负，用 -1 表示没有找到
// 		}
//
// 		leftRes := dfsNumColor(node.Left)
// 		if leftRes != -1 { // 答案在左子树中
// 			return leftRes
// 		}
//
// 		k--
// 		if k == 0 { // 答案就是当前节点
// 			return node.Val
// 		}
//
// 		return dfsNumColor(node.Right) // 右子树会返回答案或者 -1
// 	}
//
// 	return dfsNumColor(root)
// }

/*

例子：kthSmallest(root, 3)
    5
   / \
  3   7
 / \   \
2   4   8


进入 dfsNumColor 函数处理根节点 5
首先递归调用 dfsNumColor 处理左子树，即 dfsNumColor(节点 3)。

处理左子树节点 3
继续递归调用 dfsNumColor 处理其左子树 dfsNumColor(节点 2)。

处理左子树节点 2
节点 2 没有左子树，调用 dfsNumColor(nil)，返回 -1。
此时 leftRes = -1，不满足 leftRes != -1 的条件。
k 从 3 减为 2，k 不等于 0，继续递归调用 dfsNumColor(节点 2 的右子树)，由于右子树为空，返回 -1。

回到节点 3
leftRes = -1，不满足 leftRes != -1 的条件。
k 从 2 减为 1，k 不等于 0，继续递归调用 dfsNumColor(节点 3 的右子树)，即 dfsNumColor(节点 4)。

处理节点 4
节点 4 没有左子树，调用 dfsNumColor(nil)，返回 -1。
此时 leftRes = -1，不满足 leftRes != -1 的条件。
k 从 1 减为 0，满足 k == 0 的条件，返回当前节点的值 4。

回到节点 3
由于 dfsNumColor(节点 4) 返回了 4，即 leftRes = 4，满足 leftRes != -1 的条件，返回 4。

回到根节点 5
由于 dfsNumColor(节点 3) 返回了 4，即 leftRes = 4，满足 leftRes != -1 的条件，返回 4。

回到 kthSmallest 函数
kthSmallest 函数接收 dfsNumColor(root) 的返回值 4，最终返回 4。



*/
