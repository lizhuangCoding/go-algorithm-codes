package main

// 数学：找规律
// 力扣：https://leetcode.cn/problems/nim-game/description/
// 题解：力扣官方：https://leetcode.cn/problems/nim-game/solutions/1003627/nim-you-xi-by-leetcode-solution-95g8/

func canWinNim(n int) bool {
	// n是4的倍数的时候，我总是输
	if n%4 != 0 {
		return true
	}
	return false
}
