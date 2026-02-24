package main

// 二叉树，bfs
// 力扣：https://leetcode.cn/problems/average-of-levels-in-binary-tree/

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	sli := make([]*TreeNode, 0)
	sli = append(sli, root)

	result := make([]float64, 0)

	for len(sli) != 0 {

		l := len(sli)
		demo := 0

		for i := 0; i < l; i++ {
			demo += sli[i].Val

			if sli[i].Left != nil {
				sli = append(sli, sli[i].Left)
			}
			if sli[i].Right != nil {
				sli = append(sli, sli[i].Right)
			}
		}
		result = append(result, float64(demo)/float64(l))

		sli = sli[l:]
	}

	return result
}
