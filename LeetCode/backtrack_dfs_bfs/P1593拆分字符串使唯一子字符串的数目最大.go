package main

// 回溯
// 力扣：https://leetcode.cn/problems/split-a-string-into-the-max-number-of-unique-substrings/description/

// 对于长度为 n 的字符串，有 n−1 个拆分点。从左到右遍历字符串，对于每个拆分点，如果在此拆分之后，
// 新得到的一个非空子字符串 （即拆分点左侧的最后一个被拆分出的非空子字符串）与之前拆分出的非空子字符串都不相同，则当前的拆分点可以进行拆分，
// 然后继续对剩下的部分（即拆分点右侧的部分）进行拆分。
//
// func maxUniqueSplit(s string) int {
// 	// 存放结果
// 	var maxUniqueSplitRes int
// 	m := make(map[string]struct{})
//
// 	// index：索引，拆分的时候起始的位置
// 	// m：存放拆分的结果
// 	var maxUniqueSplitBacktrack func(s string, index int, m map[string]struct{})
// 	maxUniqueSplitBacktrack = func(s string, index int, m map[string]struct{}) {
// 		// 获取最大结果值
// 		if index == len(s) {
// 			maxUniqueSplitRes = maxUniqueSplitMax(maxUniqueSplitRes, len(m))
// 			return
// 		}
//
// 		for i := index + 1; i <= len(s); i++ {
// 			demo := s[index:i]
// 			// 判断子字符串是否重复
// 			_, v := m[demo]
// 			if v { // 重复，跳过
// 				continue
// 			} else {
// 				// 存放demo
// 				m[demo] = struct{}{}
// 				maxUniqueSplitBacktrack(s, i, m)
// 				// 清除demo
// 				delete(m, demo)
// 			}
// 		}
// 	}
//
// 	// 开始调用回溯
// 	maxUniqueSplitBacktrack(s, 0, m)
//
// 	return maxUniqueSplitRes
// }

func maxUniqueSplitMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化，枝剪
func maxUniqueSplit(s string) (ans int) {
	st := map[string]bool{}
	n := len(s)

	var dfs func(int)
	dfs = func(i int) {
		// 剪枝，剩余的字符串就算每一个都拆分也无法比result大了
		if len(st)+n-i <= ans {
			return
		}

		if i >= n {
			ans = maxUniqueSplitMax(ans, len(st))
			return
		}

		for j := i + 1; j <= n; j++ {
			if t := s[i:j]; !st[t] {
				st[t] = true
				dfs(j)
				delete(st, t)
			}
		}
	}
	dfs(0)
	return
}
