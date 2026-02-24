package main

// 二叉树，bfs
// 力扣：https://leetcode.cn/problems/binary-tree-right-side-view/description/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	sli := make([]*TreeNode, 0)
	sli = append(sli, root)

	result := make([]int, 0)

	for len(sli) != 0 {
		l := len(sli)

		result = append(result, sli[len(sli)-1].Val)

		for i := 0; i < l; i++ {
			if sli[i].Left != nil {
				sli = append(sli, sli[i].Left)
			}
			if sli[i].Right != nil {
				sli = append(sli, sli[i].Right)
			}
		}

		sli = sli[l:]
	}

	return result
}
