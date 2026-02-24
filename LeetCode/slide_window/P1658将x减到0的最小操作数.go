package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/

// 方法一：逆向思维:总元素个数 - 最大选几个数的和可以达到 total-x
func minOperations(nums []int, x int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	// 判断
	if total < x {
		return -1
	}
	if total == x {
		return len(nums)
	}

	target := total - x
	maxLen := -1 // 不能取0，否则如果最后的结果maxLen就是0，那么这两个0怎么区分，这个0到底是初始化赋的值，还是结果就是0
	left := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > target {
			sum -= nums[left]
			left++
		}

		if sum == target && maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}

	if maxLen < 0 {
		return -1
	}
	return len(nums) - maxLen
}

// 方法二：直接双指针
// 首先算出最长的元素和不超过 x 的后缀，然后不断枚举前缀长度，另一个指针指向后缀最左元素，保证前缀+后缀的元素和不超过 x。 答案就是前缀+后缀长度之和的最小值。
// func minOperations(nums []int, x int) int {
// 	s, n := 0, len(nums)
// 	right := n
// 	for right > 0 && s+nums[right-1] <= x { // 计算最长后缀
// 		right--
// 		s += nums[right]
// 	}
// 	if right == 0 && s < x { // 全部移除也无法满足要求
// 		return -1
// 	}
// 	ans := n + 1
// 	if s == x {
// 		ans = n - right
// 	}
// 	for left, num := range nums {
// 		s += num
// 		for ; right < n && s > x; right++ { // 缩小后缀长度
// 			s -= nums[right]
// 		}
// 		if s > x { // 缩小失败，说明前缀过长
// 			break
// 		}
// 		if s == x {
// 			ans = min(ans, left+1+n-right) // 前缀+后缀长度
// 		}
// 	}
// 	if ans > n {
// 		return -1
// 	}
// 	return ans
// }

// func main() {
// 	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))                                                                                        // 2
// 	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))                                                                                        // -1
// 	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))                                                                                   // 5
// 	fmt.Println(minOperations([]int{1, 1}, 3))                                                                                                 // -1
// 	fmt.Println(minOperations([]int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309}, 134365)) // 16
// }
