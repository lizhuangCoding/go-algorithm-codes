package main

// 二叉树
// 力扣：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/

// 方法一：bfs，暴力（时间超限）
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil || p == nil || q == nil {
// 		return nil
// 	}
// 	result := root
//
// 	sli := make([]*TreeNode, 0)
// 	sli = append(sli, root)
//
// 	for len(sli) != 0 {
// 		l := len(sli)
// 		for i := 0; i < l; i++ {
// 			if lowestCommonAncestorJudge(sli[i], p.Val, q.Val) {
// 				result = sli[i]
//
// 				if sli[i].Left != nil {
// 					sli = append(sli, sli[i].Left)
// 				}
// 				if sli[i].Right != nil {
// 					sli = append(sli, sli[i].Right)
// 				}
// 			}
// 		}
// 		sli = sli[l:]
// 	}
//
// 	return result
// }
//
// // 判断该子树是否包含两个值
// func lowestCommonAncestorJudge(head *TreeNode, p, q int) bool {
// 	isP, isQ := false, false
//
// 	sli := make([]*TreeNode, 0)
// 	sli = append(sli, head)
//
// 	for len(sli) != 0 {
// 		if isP && isQ {
// 			return true
// 		}
// 		l := len(sli)
//
// 		for i := 0; i < l; i++ {
// 			if sli[i].Val == p {
// 				isP = true
// 			}
// 			if sli[i].Val == q {
// 				isQ = true
// 			}
//
// 			if sli[i].Left != nil {
// 				sli = append(sli, sli[i].Left)
// 			}
// 			if sli[i].Right != nil {
// 				sli = append(sli, sli[i].Right)
// 			}
// 		}
// 		sli = sli[l:]
// 	}
//
// 	return isP && isQ
// }

// 方法二：递归（还是不太理解）
// 题解：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solutions/2023872/fen-lei-tao-lun-luan-ru-ma-yi-ge-shi-pin-2r95/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

// 例子：
//        3
//       / \
//      5   1
//     /|   |\
//    6 2   0 8
//      |\
//      7 4
//
// 现在要找出节点p = 5和节点q = 1的最近公共祖先。
// 从根节点3开始调用lowestCommonAncestor函数。
// 根节点3的值既不等于5也不等于1，所以递归调用左子树和右子树。
// 在左子树中，根节点是5，其值等于p的值，所以左子树的递归调用返回5。
// 在右子树中，根节点是1，其值等于q的值，所以右子树的递归调用返回1。
// 由于左子树和右子树的递归调用都返回了非空节点，所以当前根节点3就是节点p和q的最近公共祖先，最终返回3。

// 例2：
//         3
//       / \
//      5   1
//     /|   |\
//    6 2   0 8
//      |\
//      7 4
//
// 现在要查找节点 p = 7 和节点 q = 4 的最近公共祖先。
// 从根节点 3 开始调用 lowestCommonAncestor 函数。
// 根节点 3 的值既不等于 7 也不等于 4，所以递归调用左子树和右子树。
// 右子树：根节点为 1，对其左右子树递归查找，最终会发现右子树中没有 p 和 q 节点，所以右子树的递归结果 right 为 nil。
// 左子树：根节点为 5，继续对其左右子树递归查找。
// 左子树：根节点为 6，同样没有 p 和 q 节点，结果为 nil。
// 右子树：根节点为 2，继续递归，会发现 7 和 4 分别在其左右子树中，最终会返回 2 作为这部分的 LCA。
// 回到根节点 5 的递归调用，此时 left 为 nil（左子树 6 那部分的结果），right 为 2（右子树 2 那部分的结果），根据 if left == nil { return right } 这行代码，会返回 right，也就是 2。
// 再回到根节点 3 的递归调用，此时 left 为 2，right 为 nil，最终会返回 left，也就是 2，说明节点 7 和 4 的最近公共祖先是 2。
