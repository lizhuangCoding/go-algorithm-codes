package main

// 遍历二叉树
// 力扣：https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/

var maxRes int

func maxDepth(root *TreeNode) int {
	maxRes = 0
	if root == nil {
		return 0
	}

	maxDepthDfs(root, 1)

	return maxRes
}

func maxDepthDfs(head *TreeNode, num int) {
	if head.Left == nil && head.Right == nil {
		if maxRes < num {
			maxRes = num
		}
	}

	if head.Left != nil {
		maxDepthDfs(head.Left, num+1)
	}
	if head.Right != nil {
		maxDepthDfs(head.Right, num+1)
	}
}

// 或者：
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return axDepthMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
// }
//
// func axDepthMax(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// // 或者：自底而上
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	lDepth := maxDepth(root.Left)
// 	rDepth := maxDepth(root.Right)
// 	return maxDepthMax(lDepth, rDepth) + 1
// }

// // 方法二：bfs（没看）
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	queue := []*TreeNode{}
// 	queue = append(queue, root)
// 	ans := 0
// 	for len(queue) > 0 {
// 		sz := len(queue)
// 		for sz > 0 {
// 			node := queue[0]
// 			queue = queue[1:]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 			sz--
// 		}
// 		ans++
// 	}
// 	return ans
// }
