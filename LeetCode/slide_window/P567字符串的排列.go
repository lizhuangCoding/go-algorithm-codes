package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/permutation-in-string/description/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var arr1, arr2 [30]int
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	if arr1 == arr2 {
		return true
	}

	for right := len(s1); right < len(s2); right++ {
		arr2[s2[right]-'a']++
		arr2[s2[right-len(s1)]-'a']--
		if arr1 == arr2 {
			return true
		}
	}

	return false
}
