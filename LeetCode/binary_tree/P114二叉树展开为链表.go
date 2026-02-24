package main

// 二叉树
// 力扣：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
// 官方题解：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/solutions/356853/er-cha-shu-zhan-kai-wei-lian-biao-by-leetcode-solu/

// // 方法一：使用额外空间
// // 时间复杂度：O(n)；
// // 空间复杂度：O(n)
// func flatten(root *TreeNode) {
// 	list := preorderFlatten(root)
// 	for i := 0; i < len(list)-1; i++ {
// 		prev, curr := list[i], list[i+1]
// 		prev.Left = nil
// 		prev.Right = curr
// 	}
// }
//
// // 记录前序遍历结果
// func preorderFlatten(root *TreeNode) []*TreeNode {
// 	list := make([]*TreeNode, 0)
//
// 	if root == nil {
// 		return nil
// 	}
//
// 	list = append(list, root)
// 	list = append(list, preorderFlatten(root.Left)...)
// 	list = append(list, preorderFlatten(root.Right)...)
//
// 	return list
// }

// 方法二：寻找前驱节点（不是太理解）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

// func main() {
// 	root := &TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val: 2,
// 			Left: &TreeNode{
// 				Val:   3,
// 				Left:  nil,
// 				Right: nil,
// 			},
// 			Right: &TreeNode{
// 				Val:   4,
// 				Left:  nil,
// 				Right: nil,
// 			},
// 		},
// 		Right: &TreeNode{
// 			Val:  5,
// 			Left: nil,
// 			Right: &TreeNode{
// 				Val:   6,
// 				Left:  nil,
// 				Right: nil,
// 			},
// 		},
// 	}
// }
