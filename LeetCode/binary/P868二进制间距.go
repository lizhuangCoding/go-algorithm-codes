package main

// 位运算
// 力扣：https://leetcode.cn/problems/binary-gap/

func binaryGap(n int) int {
	maxRes := 0
	last := -1
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				maxRes = max(maxRes, i-last)
			}
			last = i
		}

		n >>= 1
	}

	return maxRes
}
