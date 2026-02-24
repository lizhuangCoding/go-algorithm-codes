package main

// 分组循环，字符串，等差数列
// 力扣：https://leetcode.cn/problems/count-number-of-homogenous-substrings/description/

func countHomogenous(s string) int {
	mod := 1000000007
	result := 0

	for i := 0; i < len(s); i++ {
		start := i
		for ; i < len(s)-1 && s[i] == s[i+1]; i++ {
		}

		n := i - start + 1
		// 等差公式
		result += n * (n + 1) / 2
		result %= mod
	}
	return result
}

// 这道题为什么可以用等差数列呢？
/*
	abbcccaa：
	连续的 a 可以贡献1种情况
	连续的 bb 可以贡献2+1=3种："b"出现2次。"bb"出现1次。
	连续的 ccc 可以贡献3+2+1=6种："c"出现3次。"cc"出现2次。"ccc"出现1次。
	连续的 aa 可以贡献3种

	符合等差数列

*/

// 或者：也是等差数列的做法，每一轮加的cnt就相当于 + 1 + 2 + 3 + ...
// func countHomogenous(s string) int {
// 	const mod = 1000000007
// 	cnt := 1
// 	ans := 1
// 	for i, n := 1, len(s); i < n; i++ {
// 		if s[i] == s[i-1] {
// 			cnt++
// 		} else {
// 			cnt = 1
// 		}
// 		ans += cnt % mod
// 	}
// 	return ans % mod
// }
