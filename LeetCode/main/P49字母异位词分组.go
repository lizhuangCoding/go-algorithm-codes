package main

// 哈希，字符串，排序
// 力扣：https://leetcode.cn/problems/group-anagrams/description/

// func groupAnagrams(strs []string) [][]string {
// 	m := make(map[string][]string)
//
// 	for i := 0; i < len(strs); i++ {
// 		sort.Slice(strs[i], func(x, y int) bool {
// 			return strs[i][x] < strs[i][y]
// 		})
// 	}
//
// 	for i := 0; i < len(strs); i++ {
// 		if m[strs[i]] != nil {
// 			m[strs[i]] = append(m[strs[i]], str)
// 		}f
// 	}
//
// }
