package main

// 位运算：>> & <<
// 力扣：https://leetcode.cn/problems/single-number-ii/

// 方法一：使用map
// 时间复杂度：O(n)；空间复杂度：O(n)。
// func singleNumber(nums []int) int {
// 	m := make(map[int]int)
//
// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]]++
// 	}
//
// 	for k, v := range m {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return -1
// }

// 方法二：二进制的特点。答案的第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数。
// 题解：力扣官方
// 时间复杂度：O(nlogC)，其中 n 是数组的长度，C 是元素的数据范围，在本题中 logC=log2^32=32，也就是我们需要遍历第 0∼31 个二进制位。
// 空间复杂度：O(1)。
//
//	func singleNumber(nums []int) int {
//		ans := int32(0) // 注意要用 int32类型，否则处理负数的时候就错了
//
//		for i := 0; i < 32; i++ {
//			total := int32(0)
//			for _, num := range nums {
//				// 统计每一个元素的第 i 个二进制位上是1还是0
//				total += int32(num) >> i & 1
//			}
//			// 判断答案的第 i 个二进制位上是0还是1
//			if total%3 > 0 {
//				ans |= 1 << i
//			}
//		}
//		return int(ans)
//	}

// 方法三：没看懂
// 题解：https://www.cnblogs.com/gzshan/p/12535178.html
// 为了区分出现一次的数字和出现三次的数字，使用两个位掩码：seen_once 和 seen_twice。
// 思路是：仅当 seen_twice 未变时，改变 seen_once；仅当 seen_once 未变时，改变seen_twice。

// 这个代码实现了一个算法，用来在数组 nums 中找到只出现一次的数字，而其他所有数字出现恰好三次。该算法的关键在于通过位运算管理数字的出现次数，使用两个变量 seenOnce 和 seenTwice 来跟踪每个数字的位在不同次数时的状态。
// 通过不断更新 seenOnce 和 seenTwice，可以在不使用额外空间的情况下过滤掉那些出现三次的数字，从而只保留出现一次的数字。
// func singleNumber(nums []int) int {
// 	seenOnce, seenTwice := 0, 0
//
// 	for _, v := range nums {
// 		// 先更新 seenOnce, 再更新 seenTwice
// 		seenOnce = ^seenTwice & (seenOnce ^ v)
// 		seenTwice = ^seenOnce & (seenTwice ^ v)
// 	}
// 	return seenOnce
// }
