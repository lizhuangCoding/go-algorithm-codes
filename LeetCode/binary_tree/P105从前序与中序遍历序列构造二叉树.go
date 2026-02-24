package main

// 二叉树
// 力扣：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 官方题解：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solutions/255811/cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9/

// 相同题目：https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	// 找到中序遍历中的根节点位置
	index := 0
	for ; index < len(inorder); index++ {
		if preorder[0] == inorder[index] {
			break
		}
	}

	// 左子树的长度
	leftLen := len(inorder[0:index])
	// 右子树长度
	// rightLen := len(inorder[index+1:])

	root.Left = buildTree(preorder[1:leftLen+1], inorder[0:index])
	root.Right = buildTree(preorder[leftLen+1:], inorder[index+1:])
	// 这样也可以：
	// root.Left = buildTree(preorder[1:index+1], inorder[:index])
	// root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
}
