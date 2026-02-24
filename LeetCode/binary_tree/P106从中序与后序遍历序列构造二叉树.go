package main

// 二叉树
// 力扣：https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
// 相似题目：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	if len(postorder) == 0 {
// 		return nil
// 	}
//
// 	last := len(postorder) - 1
// 	root := &TreeNode{Val: postorder[last]}
//
// 	index := 0
// 	for i := 0; i < len(inorder); i++ {
// 		if inorder[i] == postorder[last] {
// 			index = i
// 			break
// 		}
// 	}
//
// 	// 左子树的长度
// 	leftLen := len(inorder[0:index])
//
// 	root.Left = buildTree(inorder[0:index], postorder[:leftLen])
// 	root.Right = buildTree(inorder[index+1:], postorder[leftLen:last])
//
// 	return root
// }
//
// func main() {
// 	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
// 	// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 	// 输出：[3,9,20,null,null,15,7]
// }
