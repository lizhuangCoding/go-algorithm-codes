package main

// 单调栈 + 循环数组
// LeetCode：https://leetcode.cn/problems/next-greater-element-ii/description/

// 我们可以使用单调栈解决本题。单调栈中保存的是下标，从栈底到栈顶的下标在数组 nums 中对应的值是单调不升的。
// 每次我们移动到数组中的一个新的位置 iii，我们就将当前单调栈中所有对应值小于 nums[i] 的下标弹出单调栈，这些值的下一个更大元素即为 nums[i]（证明很简单：如果有更靠前的更大元素，那么这些位置将被提前弹出栈）。随后我们将位置 i 入栈。
// 但是注意到只遍历一次序列是不够的，例如序列 [2,3,1]，最后单调栈中将剩余 [3,1]，其中元素 [1] 的下一个更大元素还是不知道的。
// 一个朴素的思想是，我们可以把这个循环数组「拉直」，即复制该序列的前 n−1n-1n−1 个元素拼接在原序列的后面。这样我们就可以将这个新序列当作普通序列，用上文的方法来处理。
// 而在本题中，我们不需要显性地将该循环数组「拉直」，而只需要在处理时对下标取模即可。

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for k, _ := range ans {
		ans[k] = -1
	}

	stack := []int{} // stack 里面存的是下标

	for i := 0; i < 2*n-1; i++ {
		// 举个例子：3 2 1 4
		// stack 里面的元素，从左到右依次是栈底到栈顶
		// 0
		// 0 1
		// 0 1 2
		// 3
		// 3 0
		// 3 0 1
		// 3 0 1 2
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return ans
}

func nextGreaterElements2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	stack := []int{}

	for i := 0; i < 2*n-1; i++ {
		// 3 2 1 4
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return ans
}
