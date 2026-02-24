package main

// 二叉树（注意空指针情况的判断）
// 力扣：https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/description/

// 这句话是什么意思？
// 注意，空树 不会是以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
// 解释：如果a和b都是nil，返回false；如果a不为nil，b为nil，也返回nil

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	result := false

	if A != nil && B != nil {
		if A.Val == B.Val {
			result = isSubStructureJudge(A, B)
		}
		if !result && A.Left != nil { // 这个判断就有点多余了：A.Left != nil ， A.Right != nil
			result = isSubStructure(A.Left, B)
		}
		if !result && A.Right != nil {
			result = isSubStructure(A.Right, B)
		}
	}

	return result
}

func isSubStructureJudge(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSubStructureJudge(a.Left, b.Left) && isSubStructureJudge(a.Right, b.Right)
}
