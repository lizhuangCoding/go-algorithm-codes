package main

import (
	"unicode"
)

// 回溯
// 力扣：https://leetcode.cn/problems/letter-case-permutation/description/

// 方法一：回溯（isVisited 在功能上来说可以不写，但是不写提交就会时间超限）
// func letterCasePermutation(s string) []string {
// 	m := make(map[string]struct{})
// 	l := len(s)
// 	isVisited := make([]bool, l)
//
// 	letterCasePermutationDfs(s, s, 0, isVisited, m)
//
// 	result := make([]string, 0)
// 	for k := range m {
// 		result = append(result, k)
// 	}
// 	return result
// }
//
// func letterCasePermutationDfs(s string, demo string, n int, isVisited []bool, m map[string]struct{}) {
// 	m[demo] = struct{}{}
//
// 	for i := n; i < len(s); i++ {
// 		// 可以变为大写
// 		if s[i] >= 'a' && s[i] <= 'z' && !isVisited[i] {
// 			isVisited[i] = true
// 			b := s[i] - 32
// 			letterCasePermutationDfs(s, demo[:i]+string(b)+demo[i+1:], n+1, isVisited, m)
// 			isVisited[i] = false
// 		}
//
// 		// 可以变为小写
// 		if (s[i] >= 'A' && s[i] <= 'Z') && !isVisited[i] {
// 			isVisited[i] = true
// 			b := s[i] + 32
// 			letterCasePermutationDfs(s, demo[:i]+string(b)+demo[i+1:], n+1, isVisited, m)
// 			isVisited[i] = false
// 		}
// 	}
// }

// 回溯的另一种写法:
// 从左往右依次遍历字符，当在进行搜索时，搜索到字符串 s 的第 i 个字符 c 时:
// 如果 c 为一个数字，则我们继续检测下一个字符；
// 如果 c 为一个字母，我们将字符中的第 i 个字符 c 改变大小写形式后，往后继续搜索，完成改写形式的子状态搜索后，我们将 c 进行恢复，继续往后搜索；
// 如果当前完成字符串搜索后，则表示当前的子状态已经搜索完成，该序列为全排列中的一个；
// 使用异或运算符实现字母的大小写转换。
func letterCasePermutation(s string) (ans []string) {
	var dfs func([]byte, int)
	dfs = func(s []byte, pos int) {
		// 跳过数字
		for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
			pos++
		}
		// 终止条件
		if pos == len(s) {
			ans = append(ans, string(s))
			return
		}

		// 不改变当前字符大小写，继续递归调用 dfs 函数处理下一个字符
		dfs(s, pos+1)
		// 大小写转换
		s[pos] ^= 32
		// 对转换大小写后的字符，继续递归调用 dfs 函数处理下一个字符，生成另一种排列组合
		dfs(s, pos+1)
		// 将字符恢复
		s[pos] ^= 32
	}

	dfs([]byte(s), 0)
	return
}

// 方法二：广度优先搜索（没怎么看）
// 题解：https://leetcode.cn/problems/letter-case-permutation/solutions/1934375/zi-mu-da-xiao-xie-quan-pai-lie-by-leetco-cwpx/?envType=problem-list-v2&envId=LBZR6IMZ
