package main

// 二叉树，dfsNumColor
// 力扣：https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/

// 写法一：dfs没有返回值
// var result int
//
// func sumNumbers(root *TreeNode) int {
// 	result = 0
// 	sumNumbersDfs(root, 0)
// 	return result
// }
//
// func sumNumbersDfs(head *TreeNode, demo int) {
// 	if head == nil {
// 		return
// 	}
// 	demo = demo*10 + head.Val
// 	// 叶子结点（这样写是为了防止加两次demo，只有到了叶子结点，值才可以被加上去）
// 	if head.Left == nil && head.Right == nil {
// 		result += demo
// 	}
// 	sumNumbersDfs(head.Left, demo)
// 	sumNumbersDfs(head.Right, demo)
// }

// // 写法二：dfs有返回值
// func sumNumbers(root *TreeNode) int {
// 	return sumNumbersDfs(root, 0)
// }
//
// func sumNumbersDfs(head *TreeNode, demo int) int {
// 	if head == nil {
// 		return 0
// 	}
// 	demo = demo*10 + head.Val
// 	// 叶子结点（这样写是为了防止加两次demo，只有到了叶子结点才可以值才可以被加上去）
// 	if head.Left == nil && head.Right == nil {
// 		return demo
// 	}
//
// 	return sumNumbersDfs(head.Left, demo) + sumNumbersDfs(head.Right, demo)
// }

// func main() {
// 	// root := &TreeNode{
// 	// 	Val: 1,
// 	// 	Left: &TreeNode{
// 	// 		Val:   2,
// 	// 		Left:  nil,
// 	// 		Right: nil,
// 	// 	},
// 	// 	Right: &TreeNode{
// 	// 		Val:   3,
// 	// 		Left:  nil,
// 	// 		Right: nil,
// 	// 	},
// 	// }
// 	root := &TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val:   0,
// 			Left:  nil,
// 			Right: nil,
// 		},
// 	}
//
// 	fmt.Println(sumNumbers(root))
// }
