package main

import "strconv"

// dfsNumColor，bfs
// 力扣：https://leetcode.cn/problems/binary-tree-paths/description/?envType=problem-list-v2&envId=LBZR6IMZ

// 方法一：bfs
// var result []string
//
// func binaryTreePaths(root *TreeNode) []string {
// 	result = make([]string, 0)
// 	binaryTreePathsDfs(root, "")
// 	return result
// }
//
// func binaryTreePathsDfs(heat *TreeNode, path string) {
// 	if heat != nil {
// 		pathDemo := path
// 		pathDemo += strconv.Itoa(heat.Val)
//
// 		if heat.Left == nil && heat.Right == nil {
// 			result = append(result, pathDemo)
// 		} else {
// 			pathDemo += "->"
// 			binaryTreePathsDfs(heat.Left, pathDemo)
// 			binaryTreePathsDfs(heat.Right, pathDemo)
// 		}
//
// 	}
// }

// 方法二：bfs
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	var nodeQueue []*TreeNode
	var pathQueue []string
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}
	return paths
}
