package main

// 由于 i < j，我们可以先排序数组，然后对每个 nums[i]，用二分查找找到满足 lower - nums[i] <= nums[j] <= upper - nums[i] 的 j 的范围。
// 步骤：
// 排序数组（排序后不影响 i < j 的条件，因为 i 和 j 是索引，但排序后可以用双指针优化）。
// 遍历 nums[i]，对每个 i：
// 计算 left = lower - nums[i]（nums[j] 的最小值）
// 计算 right = upper - nums[i]（nums[j] 的最大值）
// 用二分查找找到 j 的范围 [l, r)，left <= nums[j] <= right，并且 j > i。
// 累加符合条件的 j 的数量。

// func countFairPairs(nums []int, lower int, upper int) int64 {
// 	sort.Ints(nums)
// 	n := len(nums)
// 	var res int64
//
// 	for i := 0; i < n; i++ {
// 		left := lower - nums[i]
// 		right := upper - nums[i]
//
// 		// 找到第一个 >= left 的 j（j > i）
// 		l := sort.Search(n, func(j int) bool {
// 			return nums[j] >= left
// 		})
// 		l = max(l, i+1) // 确保 j > i
//
// 		// 找到第一个 > right 的 j（j > i）
// 		r := sort.Search(n, func(j int) bool {
// 			return nums[j] > right
// 		})
// 		r = max(r, i+1) // 确保 j > i
//
// 		res += int64(r - l)
// 	}
// 	return res
// }
//
// func countFairPairs(nums []int, lower int, upper int) int64 {
// 	// 虽然说 0 <= i < j < n，但其实就是从数组中选择两个数即可，所以可以对数组进行排序
// 	slices.Sort(nums)
//
// 	res := 0
//
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > upper {
// 			continue
// 		}
//
// 		// 挑选另一个值
// 		n1 := lower - nums[i] // 另一个值需要 >= n1（左边界）
// 		n2 := upper - nums[i] // 另一个值需要 <= n2（右边界）
//
// 		// 找到左边界
// 		index1 := leftCountFairPairs(nums[i:], n1)
// 		// 找到右边界
// 		index2 := rightCountFairPairs(nums[i:], n2)
//
// 		if index1 == -1 || index2 == -1 {
// 			continue
// 		}
//
// 		if index2 >= index1 {
// 			res += index2 - index1 + 1
// 		}
// 	}
// 	return int64(res)
// }
//
// func leftCountFairPairs(nums []int, target int) int {
// 	index := -1
//
// 	l, r := 0, len(nums)-1
//
// 	for l <= r {
// 		mid := (l + r) >> 1
//
// 		if nums[mid] >= target {
// 			index = mid
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return index
// }
//
// func rightCountFairPairs(nums []int, target int) int {
// 	index := -1
//
// 	l, r := 0, len(nums)-1
//
// 	for l <= r {
// 		mid := (l + r) >> 1
//
// 		if nums[mid] <= target {
// 			index = mid
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return index
// }
//
// func main() {
// 	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6)) // 6
// 	fmt.Println(countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))  // 1
//
// }
