package main

// 分组循环，数组
// 力扣：https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/description/

func longestAlternatingSubarray(nums []int, threshold int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		start := i

		if nums[start]%2 == 0 && nums[start] <= threshold {

			for ; i < len(nums)-1 && nums[i]%2 != nums[i+1]%2 && nums[i+1] <= threshold; i++ {
			}

			result = longestAlternatingSubarrayMax(result, i-start+1)
		}
	}
	return result
}

func longestAlternatingSubarrayMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(longestAlternatingSubarray([]int{3, 2, 5, 4}, 5)) // 3
// 	fmt.Println(longestAlternatingSubarray([]int{1, 2}, 2))       // 1
// 	fmt.Println(longestAlternatingSubarray([]int{2, 3, 4, 5}, 4)) // 3
// 	fmt.Println(longestAlternatingSubarray([]int{4}, 1))          // 0
// 	fmt.Println(longestAlternatingSubarray([]int{2, 2}, 18))      // 1
// }
