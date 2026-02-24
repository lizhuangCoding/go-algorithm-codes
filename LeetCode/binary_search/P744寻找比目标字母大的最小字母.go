package main

// 二分查找，找到第一个符合条件的元素
// 力扣：https://leetcode.cn/problems/find-smallest-letter-greater-than-target/description/

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	var res = letters[0]
	for left <= right {
		mid := left + (right-left)>>1

		if letters[mid] == target {
			left = mid + 1
		} else if letters[mid] > target {
			res = letters[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
