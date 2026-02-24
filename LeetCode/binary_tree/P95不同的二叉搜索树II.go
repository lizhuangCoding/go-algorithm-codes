package main

// 二叉搜索树（Binary Search Tree, BST）
// 力扣：https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
// 题解：https://leetcode.cn/problems/unique-binary-search-trees-ii/solutions/339143/bu-tong-de-er-cha-sou-suo-shu-ii-by-leetcode-solut/

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesBackTrack(1, n)
}

func generateTreesBackTrack(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := make([]*TreeNode, 0)

	// 枚举可行根节点（i是当前的父节点）
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := generateTreesBackTrack(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := generateTreesBackTrack(i+1, end)

		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, leftNode := range leftTrees {
			for _, rightNode := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = leftNode
				currTree.Right = rightNode
				result = append(result, currTree)
			}
		}
	}
	return result
}
