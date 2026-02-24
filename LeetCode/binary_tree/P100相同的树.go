package main

// 递归
// 力扣：https://leetcode.cn/problems/same-tree/description/

// 时间复杂度：O(min(n,m))，其中 n 为 p 的节点个数，m 为 q 的节点个数。
// 空间复杂度：O(min(n,m))。最坏情况下，二叉树退化成一条链，递归需要 O(min(n,m)) 的栈空间。
func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // 必须都是 nil
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// // 自己写的代码：有几个案例没有通过，为什么会错误
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	return isSameTreeDigui(p, q)
// }
//
// func isSameTreeDigui(p *TreeNode, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	}
//
// 	if p == nil || q == nil {
// 		return false
// 	}
//
// 	if p.Val != q.Val {
// 		return false
// 	}
//
// 	// 错误原因：
// 	// 只在 p.Left 或 p.Right 非空时进行递归比较，而没有考虑 p.Left 为空但 q.Left 不为空，或者 p.Right 为空但 q.Right 不为空的情况。
// 	// 例如，当 p 的左子树为空，q 的左子树不为空时，代码没有正确处理这种差异，直接跳过了对 q.Left 的检查，从而可能导致误判。
// 	// if p.Left != nil {
// 	// 	return isSameTreeDigui(p.Left, q.Left)
// 	// }
// 	//
// 	// if p.Right != nil {
// 	// 	return isSameTreeDigui(p.Right, q.Right)
// 	// }
// 	// return true
//
// 	// 应该改为：
// 	return isSameTreeDigui(p.Left, q.Left) && isSameTreeDigui(p.Right, q.Right)
// }
