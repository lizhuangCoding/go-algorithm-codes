package main

// 二叉树，dfs
// 力扣：https://leetcode.cn/problems/path-sum-iii/

// var resultPathSum int
//
// func pathSum(root *TreeNode, targetSum int) int {
// 	resultPathSum = 0
//
// 	dfsPathSum(root, 0, targetSum)
//
// 	return resultPathSum
// }
//
// // 写法错误，原因：路径必须是连续的，不能跳过节点
// // 你的代码允许这样的路径：从节点A开始，走到子节点B。然后"不选择"B为起点，继续走B的子节点
// // 这实际上创建了不连续的路径（跳过了中间的节点）。
// func dfsPathSum(head *TreeNode, val int, targetSum int) {
// 	if val == targetSum {
// 		resultPathSum++
// 	}
//
// 	if head == nil {
// 		return
// 	}
//
// 	// 选择本节点为起点，走左节点
// 	dfsPathSum(head.Left, val+head.Val, targetSum)
// 	// 不选择本节点为起点，走左节点
// 	dfsPathSum(head.Left, 0, targetSum)
//
// 	// 选择本节点为起点，走右节点
// 	dfsPathSum(head.Right, val+head.Val, targetSum)
// 	// 不选择本节点为起点，走右节点
// 	dfsPathSum(head.Right, 0, targetSum)
// }

// 修改：
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := dfsPathSum(root, targetSum)
	res += pathSum(root.Left, targetSum)  // 从左节点开始计算，跳过父节点
	res += pathSum(root.Right, targetSum) // 从右节点开始计算，跳过父节点

	return res
}

func dfsPathSum(head *TreeNode, targetSum int) (res int) {
	if head == nil {
		return
	}

	if head.Val == targetSum {
		res++
	}

	res += dfsPathSum(head.Left, targetSum-head.Val)  // 继续向左走
	res += dfsPathSum(head.Right, targetSum-head.Val) // 继续向右走
	return
}
