package main

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/

// 枚举左端点，移动右端点
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1

	// 通过从右向左遍历数组找到一个最长的非递减后缀子数组。遍历过程中，right 会停止在第一个不满足非递减关系的位置。循环结束时，arr[right:] 就是数组的非递减后缀子数组。
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 { // arr 已经是非递减数组
		return 0
	}

	// 此时 arr[right-1] > arr[right]
	ans := right // 将 ans 初始化为 right，表示最坏情况下需要删除 arr[:right] 才能得到一个非递减数组。
	// 从左向右遍历找到最优解：使用 left 从数组开头开始遍历，找到从左到右的非递减前缀子数组。
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		// 内部循环：当 right < n && arr[right] < arr[left] 时，将right向右移动，确保当前arr[left]位置可以与arr[right:]的非递减后缀子数组相连。
		for right < n && arr[right] < arr[left] {
			right++
		}

		// 更新ans：right-left-1 表示中间需要删除的子数组长度 arr[left+1:right]。
		if ans > right-left-1 {
			ans = right - left - 1 // 删除 arr[left+1:right]
		}
	}
	return ans
}

// 或者：枚举右端点，移动左端点
// func findLengthOfShortestSubarray(arr []int) int {
// 	n := len(arr)
// 	right := n - 1
// 	for right > 0 && arr[right-1] <= arr[right] {
// 		right--
// 	}
// 	if right == 0 { // arr 已经是非递减数组
// 		return 0
// 	}
// 	// 此时 arr[right-1] > arr[right]
// 	ans := right               // 删除 arr[:right]
// 	for left := 0; ; right++ { // 枚举 right
// 		for right == n || arr[left] <= arr[right] {
// 			if ans > right-left-1 {
// 				ans = right - left - 1 // 删除 arr[left+1:right]
// 			}
// 			if arr[left] > arr[left+1] {
// 				return ans
// 			}
// 			left++
// 		}
// 	}
// }

// func main() {
// 	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})) // 3
// 	fmt.Println(findLengthOfShortestSubarray([]int{2, 3, 4, 5, 9, 8, 6, 7}))  // 2
// 	fmt.Println(findLengthOfShortestSubarray([]int{100, 99, 2, 95, 3, 4, 5})) // 4
// }
