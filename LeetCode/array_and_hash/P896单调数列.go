package main

// 数组
// 力扣：https://leetcode.cn/problems/monotonic-array/description/

func isMonotonic(nums []int) bool {
	// 要么增或者减，要么数组里面的元素全部相等
	return (judgeZengIsMonotonic(nums) && !judgeJianIsMonotonic(nums)) ||
		(!judgeZengIsMonotonic(nums) && judgeJianIsMonotonic(nums)) ||
		(judgeZengIsMonotonic(nums) && judgeJianIsMonotonic(nums))
}

// 判断是否递减
func judgeJianIsMonotonic(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// 判断是否递增
func judgeZengIsMonotonic(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

// 遍历两次
// func isMonotonic(nums []int) bool {
// 	return mergeSort.IntsAreSorted(nums) || mergeSort.IsSorted(mergeSort.Reverse(mergeSort.IntSlice(nums)))
// }

// 遍历一次
// func isMonotonic(nums []int) bool {
// 	inc, dec := true, true
// 	for i := 0; i < len(nums)-1; i++ {
// 		if nums[i] > nums[i+1] {
// 			inc = false
// 		}
// 		if nums[i] < nums[i+1] {
// 			dec = false
// 		}
// 	}
// 	return inc || dec
// }

// func main() {
// 	fmt.Println(isMonotonic([]int{1, 1, 1, 1})) // true
// 	fmt.Println(isMonotonic([]int{6, 5, 4, 4})) // true
// }
