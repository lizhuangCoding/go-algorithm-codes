package main

// 二叉搜索树
// 力扣：https://leetcode.cn/problems/search-in-a-binary-search-tree/

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}
