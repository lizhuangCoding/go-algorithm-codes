package main

// 数组
// 力扣：https://leetcode.cn/problems/product-of-array-except-self/
// 题解：https://leetcode.cn/problems/product-of-array-except-self/solutions/272369/chu-zi-shen-yi-wai-shu-zu-de-cheng-ji-by-leetcode-/

// 方法一：左右乘积列表（可以看题解的图片）
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// L 和 R 分别表示左右两侧的乘积列表
	L, R := make([]int, n), make([]int, n)

	// L[i] 为索引 i 左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}

	// R[i] 为索引 i 右侧所有元素的乘积
	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
	R[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	// fmt.Println(L, R) // [1 1 2 6] [24 12 4 1]

	result := make([]int, n)
	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < len(result); i++ {
		result[i] = L[i] * R[i]
	}
	return result
}

// func main() {
// 	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) // 输出: [24,12,8,6]
// }

// 方法二：空间复杂度 O(1) 的方法
// 时间复杂度：O(N)
// 空间复杂度：O(N)
// func productExceptSelf(nums []int) []int {
// 	length := len(nums)
// 	answer := make([]int, length)
//
// 	// answer[i] 表示索引 i 左侧所有元素的乘积
// 	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
// 	answer[0] = 1
// 	for i := 1; i < length; i++ {
// 		answer[i] = nums[i-1] * answer[i-1]
// 	}
//
// 	// R 为右侧所有元素的乘积
// 	// 刚开始右边没有元素，所以 R = 1
// 	R := 1
// 	for i := length - 1; i >= 0; i-- {
// 		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
// 		answer[i] = answer[i] * R
// 		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
// 		R *= nums[i]
// 	}
// 	return answer
// }
