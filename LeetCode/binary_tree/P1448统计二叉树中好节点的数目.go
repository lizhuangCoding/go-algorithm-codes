package main

// 二叉树，dfs深度遍历
// 力扣：https://leetcode.cn/problems/count-good-nodes-in-binary-tree/

var result int

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result = 1

	if root.Left != nil {
		dfsGoodNodes(root.Left, root.Val)
	}

	if root.Right != nil {
		dfsGoodNodes(root.Right, root.Val)
	}

	return result
}

// prevVal 上一个节点的值
func dfsGoodNodes(root *TreeNode, maxVal int) {
	if root == nil {
		return
	}

	if root.Val >= maxVal {
		result++
	}
	maxVal = max(maxVal, root.Val)

	if root.Left != nil {
		dfsGoodNodes(root.Left, maxVal)
	}

	if root.Right != nil {
		dfsGoodNodes(root.Right, maxVal)
	}

}
