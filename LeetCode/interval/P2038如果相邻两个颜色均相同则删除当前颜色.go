package main

// 分组循环，字符串
// 力扣：https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/

func winnerOfGame(colors string) bool {
	nA, nB := 0, 0 // 这个两个人分别最多可以操作多少次

	for i := 0; i < len(colors); i++ {
		start := i
		for ; i < len(colors)-1 && colors[i] == colors[i+1]; i++ {
		}

		// 上面套模版，下面是逻辑
		if i-start+1 >= 3 {
			if colors[start] == 'A' {
				nA += i - start + 1 - 2
			} else if colors[start] == 'B' {
				nB += i - start + 1 - 2
			}
		}

	}
	return nA > nB
}
