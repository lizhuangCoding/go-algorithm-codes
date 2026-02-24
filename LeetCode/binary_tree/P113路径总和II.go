package main

// 二叉搜索树
// 力扣：https://leetcode.cn/problems/path-sum-ii/description/
// 相同题目：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/

// var result [][]int
//
// func pathTarget(root *TreeNode, target int) [][]int {
//
// 	result = make([][]int, 0)
//
// 	pathTargetDfs(root, target, 0, make([]int, 0))
//
// 	return result
// }
//
// func pathTargetDfs(head *TreeNode, target int, demo int, res []int) {
// 	if head == nil {
// 		return
// 	}
//
// 	res = append(res, head.Val)
// 	demo += head.Val
//
// 	if head.Left == nil && head.Right == nil {
// 		if demo == target {
// 			temp := make([]int, len(res)) // 创建 res 的副本。避免后续修改影响到 result 中的元素。
// 			copy(temp, res)
// 			result = append(result, temp)
// 		}
// 		return
// 	}
// 	pathTargetDfs(head.Left, target, demo, res)
// 	pathTargetDfs(head.Right, target, demo, res)
// }
