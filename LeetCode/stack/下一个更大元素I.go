package main

// 单调栈
// LeetCode：https://leetcode.cn/problems/next-greater-element-i/description/
// 求下一个更大的元素，先想单调栈

// 暴力：
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	var ans = make([]int, len(nums1))

	// 全部初始化为 -1
	for i := 0; i < len(nums1); i++ {
		ans[i] = -1
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {

			// 能够找到 nums1[i] == nums2[j] 的数
			if nums1[i] == nums2[j] {
				for z := j + 1; z < len(nums2); z++ {
					// 找到了第一个符合条件的值后就退出
					if nums2[j] < nums2[z] {
						ans[i] = nums2[z] // 赋值
						break
					}
				}
				break
			}

		}
	}
	return ans
	// 或者：
	/*
	 m, n := len(nums1), len(nums2)
	    res := make([]int, m)
	    for i, num := range nums1 {
	        j := 0
	        for j < n && nums2[j] != num {
	            j++
	        }
	        k := j + 1
	        for k < n && nums2[k] < nums2[j] {
	            k++
	        }
	        if k < n {
	            res[i] = nums2[k]
	        } else {
	            res[i] = -1
	        }
	    }
	    return res
	*/
}

// monotonic_stack(官方题解)
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := []int{} // []int{} 表示创建一个空的整型切片

	// nums2 = []int {2,5,3,6,8,4,7,1}
	// 从后向前遍历
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			m[num] = -1 // len(stack) == 0 的时候说明后面的元素都比前面的元素小，一定为-1
		} else {
			m[num] = stack[len(stack)-1]
		}

		stack = append(stack, num) // 入栈
	}

	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		res = append(res, m[nums1[i]])
	}

	return res
}
