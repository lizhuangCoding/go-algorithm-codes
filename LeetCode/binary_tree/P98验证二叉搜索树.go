package main

import "math"

// 二叉搜索树，中序遍历
// 力扣：https://leetcode.cn/problems/validate-binary-search-tree/
// 中序遍历类似的题目：https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/

// 方法一：前序遍历
// func isValidBST(root *TreeNode) bool {
// 	return isValidBSTDfs(root, math.MinInt64, math.MaxInt64)
// }
//
// func isValidBSTDfs(head *TreeNode, left, right int) bool {
// 	if head == nil {
// 		return true
// 	}
//
// 	return left < head.Val && right > head.Val && isValidBSTDfs(head.Left, left, head.Val) && isValidBSTDfs(head.Right, head.Val, right)
// }

/*

例子：

    5
   / \
  3   7
 / \   \
2   4   8

进入 dfsNumColor 函数处理根节点 5
当前节点 node 为 5，left = math.MinInt，right = math.MaxInt。
节点 5 的值 x = 5，满足 math.MinInt < 5 && 5 < math.MaxInt。
接着递归调用 dfsNumColor(node.Left, left, x)，即 dfsNumColor(节点 3, math.MinInt, 5) 处理左子树；同时递归调用 dfsNumColor(node.Right, x, right)，即 dfsNumColor(节点 7, 5, math.MaxInt) 处理右子树。

处理左子树节点 3
当前节点 node 为 3，left = math.MinInt，right = 5。
节点 3 的值 x = 3，满足 math.MinInt < 3 && 3 < 5。
接着递归调用 dfsNumColor(node.Left, left, x)，即 dfsNumColor(节点 2, math.MinInt, 3) 处理左子树；同时递归调用 dfsNumColor(node.Right, x, right)，即 dfsNumColor(节点 4, 3, 5) 处理右子树。

处理左子树节点 2
当前节点 node 为 2，left = math.MinInt，right = 3。
节点 2 的值 x = 2，满足 math.MinInt < 2 && 2 < 3。
由于节点 2 没有左子树和右子树，递归调用 dfsNumColor(node.Left, left, x) 和 dfsNumColor(node.Right, x, right) 时，传入的节点为 nil，根据 dfsNumColor 函数的逻辑，会返回 true。

处理左子树节点 4
当前节点 node 为 4，left = 3，right = 5。
节点 4 的值 x = 4，满足 3 < 4 && 4 < 5。
由于节点 4 没有左子树和右子树，递归调用 dfsNumColor(node.Left, left, x) 和 dfsNumColor(node.Right, x, right) 时，传入的节点为 nil，会返回 true。

处理右子树节点 7
当前节点 node 为 7，left = 5，right = math.MaxInt。
节点 7 的值 x = 7，满足 5 < 7 && 7 < math.MaxInt。
接着递归调用 dfsNumColor(node.Right, x, right)，即 dfsNumColor(节点 8, 7, math.MaxInt) 处理右子树。

处理右子树节点 8
当前节点 node 为 8，left = 7，right = math.MaxInt。
节点 8 的值 x = 8，满足 7 < 8 && 8 < math.MaxInt。
由于节点 8 没有左子树和右子树，递归调用 dfsNumColor(node.Left, left, x) 和 dfsNumColor(node.Right, x, right) 时，传入的节点为 nil，会返回 true。

*/

// 方法二：中序遍历
// 时间复杂度：O(n)，其中 n 为二叉搜索树的节点个数。
// 空间复杂度：O(n)。最坏情况下，二叉搜索树退化成一条链（注意题目没有保证它是平衡树），因此递归需要 O(n) 的栈空间。
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64

	var dfs func(head *TreeNode) bool
	dfs = func(head *TreeNode) bool {
		if head == nil {
			return true
		}
		// 递归检查左子树
		if !dfs(head.Left) {
			return false
		}

		// 检查当前节点值是否大于前一个节点值，如果前一个节点值>=当前的节点值那么说明不是二叉搜索树
		if pre != math.MinInt64 && pre >= head.Val {
			return false
		}

		// 更新前一个节点值
		pre = head.Val

		// 递归检查右子树
		return dfs(head.Right)
	}

	return dfs(root)
}

// 方法三：后序遍历
// 原理：对于二叉搜索树中的任意节点，其左子树中的所有节点值都小于该节点值，其右子树中的所有节点值都大于该节点值。
// 代码通过深度优先搜索（DFS）的方式遍历二叉树的每个节点，同时记录并比较以每个节点为根的子树中的最小值和最大值，
// 以此来验证整棵树是否满足 BST 的性质。
// func isValidBST(root *TreeNode) bool {
// 	_, mx := isValidBSTDfs(root)
// 	return mx != math.MaxInt
// }
//
// func isValidBSTDfs(node *TreeNode) (int, int) {
// 	if node == nil {
// 		return math.MaxInt, math.MinInt
// 	}
// 	lMin, lMax := isValidBSTDfs(node.Left)
// 	rMin, rMax := isValidBSTDfs(node.Right)
// 	x := node.Val
// 	if x <= lMax || x >= rMin { // 也可以在递归完左子树之后立刻判断，如果发现不是二叉搜索树，就不用递归右子树了
// 		return math.MinInt, math.MaxInt
// 	}
// 	return min(lMin, x), max(rMax, x)
// }

/*

例子：

    5
   / \
  3   7
 / \   \
2   4   8

进入 dfsNumColor 函数处理根节点 5
首先递归调用 dfsNumColor 处理左子树，即 dfsNumColor(节点 3)，以及右子树，即 dfsNumColor(节点 7)。

处理左子树节点 3
对于节点 3，同样要递归调用 dfsNumColor 处理其左子树 dfsNumColor(节点 2) 和右子树 dfsNumColor(节点 4)。

处理左子树节点 2
因为节点 2 没有左右子树，调用 dfsNumColor 时传入 nil，根据代码逻辑，dfsNumColor(nil) 会返回 (math.MaxInt, math.MinInt)。这里 math.MaxInt 作为初始的最小值，math.MinInt 作为初始的最大值，方便后续比较。
所以对于节点 2，lMin = math.MaxInt，lMax = math.MinInt。
节点 2 的值 x = 2，此时返回 (min(math.MaxInt, 2), max(math.MinInt, 2))，即 (2, 2)。

处理左子树节点 4
节点 4 也没有左右子树，dfsNumColor(nil) 返回 (math.MaxInt, math.MinInt)。
节点 4 的值 x = 4，返回 (min(math.MaxInt, 4), max(math.MinInt, 4))，即 (4, 4)。

回到节点 3
得到左子树（节点 2）的结果 lMin = 2，lMax = 2；右子树（节点 4）的结果 rMin = 4，rMax = 4。
节点 3 的值 x = 3，满足 3 > 2 且 3 < 4，即 x > lMax 且 x < rMin。
则返回 (min(2, 3), max(4, 3))，即 (2, 4)。

处理右子树节点 7
节点 7 有右子树 dfsNumColor(节点 8)，节点 8 没有左右子树，dfsNumColor(nil) 返回 (math.MaxInt, math.MinInt)。
节点 8 的值 x = 8，返回 (8, 8)。
对于节点 7，左子树为空，dfsNumColor(nil) 返回 (math.MaxInt, math.MinInt)，右子树（节点 8）结果为 rMin = 8，rMax = 8。
节点 7 的值 x = 7，满足 7 < 8，即 x < rMin。
则返回 (min(math.MaxInt, 7), max(8, 7))，即 (7, 8)。

回到根节点 5
得到左子树（节点 3）的结果 lMin = 2，lMax = 4；右子树（节点 7）的结果 rMin = 7，rMax = 8。
节点 5 的值 x = 5，满足 5 > 4 且 5 < 7，即 x > lMax 且 x < rMin。
则返回 (min(2, 5), max(8, 5))，即 (2, 8)。

回到 isValidBST 函数
isValidBST 函数接收 dfsNumColor(root) 的返回值，_, mx := dfsNumColor(root) 得到 mx = 8。
因为 mx != math.MaxInt，所以 isValidBST 函数返回 true，表明这棵树是一棵二叉搜索树。

*/
