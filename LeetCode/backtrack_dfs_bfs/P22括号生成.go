package main

// 回溯
// 力扣：https://leetcode.cn/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	generateParenthesisDfs(n, 0, 0, "", &res)
	return res
}

// n:括号的对数。now:已经存在的括号数量。leftCount:已经存在的左括号数量。
func generateParenthesisDfs(n, now, leftCount int, demo string, res *[]string) {
	if now == n*2 {
		*res = append(*res, demo)
		return
	}

	// 加左括号
	if leftCount < n {
		generateParenthesisDfs(n, now+1, leftCount+1, demo+"(", res)
	}

	// 加右括号
	if now-leftCount < leftCount {
		generateParenthesisDfs(n, now+1, leftCount, demo+")", res)
	}
}
