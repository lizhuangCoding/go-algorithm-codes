package main

// 二叉树
// 力扣：https://leetcode.cn/problems/symmetric-tree/description/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricDigui(root.Left, root.Right)
}

func isSymmetricDigui(head1, head2 *TreeNode) bool {
	if head1 == nil || head2 == nil {
		return head1 == head2
	}

	if head1.Val == head2.Val {
		return isSymmetricDigui(head1.Left, head2.Right) && isSymmetricDigui(head1.Right, head2.Left)
	}
	return false
	// 可以简化：
	// return head1.Val == head2.Val && isSymmetricDigui(head1.Left, head2.Right) && isSymmetricDigui(head1.Right, head2.Left)
}

// 时间复杂度：这里遍历了这棵树，渐进时间复杂度为 O(n)。
// 空间复杂度：这里的空间复杂度和递归使用的栈空间有关，这里递归层数不超过 n，故渐进空间复杂度为 O(n)。
