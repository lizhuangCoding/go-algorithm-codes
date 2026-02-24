package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/

// 进阶题目：P30串联所有单词的子串
// 力扣：https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
func findAnagrams(s string, p string) []int {
	sli1, sli2 := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		sli1[p[i]-'a']++
	}

	result := make([]int, 0)
	for right := 0; right < len(s); right++ {
		sli2[s[right]-'a']++

		if right >= len(p)-1 {
			left := right - (len(p) - 1)
			// 如果两个数组包含的元素相同，那么说明是异位词的子串
			if sli1 == sli2 {
				result = append(result, left)
			}
			sli2[s[left]-'a']--
		}
	}
	return result
}
