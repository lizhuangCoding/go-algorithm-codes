package main

// 二叉树，dfsNumColor
// 力扣：https://leetcode.cn/problems/leaf-similar-trees/description/

// var sli []int
//
// func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
// 	sli = make([]int, 0)
// 	huisuLeafSimilar(root1)
// 	var sli1 []int
// 	sli1 = append(sli1, sli...)
//
// 	sli = make([]int, 0)
// 	huisuLeafSimilar(root2)
//
// 	return slices.Equal(sli, sli1)
// }
//
// func huisuLeafSimilar(root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
//
// 	if root.Left == nil && root.Right == nil {
// 		sli = append(sli, root.Val)
// 		return
// 	}
//
// 	if root.Left != nil {
// 		huisuLeafSimilar(root.Left)
// 	}
//
// 	if root.Right != nil {
// 		huisuLeafSimilar(root.Right)
// 	}
// }

// func main() {
// root1 := &TreeNode{
// 	Val: 3,
// 	Left: &TreeNode{
// 		Val: 5,
// 		Left: &TreeNode{
// 			Val:   6,
// 			Left:  nil,
// 			Right: nil,
// 		},
// 		Right: &TreeNode{
// 			Val: 2,
// 			Left: &TreeNode{
// 				Val:   7,
// 				Left:  nil,
// 				Right: nil,
// 			},
// 			Right: &TreeNode{
// 				Val:   4,
// 				Left:  nil,
// 				Right: nil,
// 			},
// 		},
// 	},
// 	Right: &TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val:   9,
// 			Left:  nil,
// 			Right: nil,
// 		},
// 		Right: &TreeNode{
// 			Val:   8,
// 			Left:  nil,
// 			Right: nil,
// 		},
// 	},
// }
// fmt.Println(leafSimilar(root1, root1))
//
// root1 := &TreeNode{Val: 1}
// fmt.Println(leafSimilar(root1, root1))
// }
