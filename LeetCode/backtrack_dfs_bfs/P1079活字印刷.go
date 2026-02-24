package main

// 回溯
// 力扣：https://leetcode.cn/problems/letter-tile-possibilities/description/

func numTilePossibilities(tiles string) int {
	result := make(map[string]struct{})

	l := len(tiles)
	isVisited := make([]bool, l)

	for i := 0; i < l; i++ {
		if !isVisited[i] {
			isVisited[i] = true
			numTilePossibilitiesDfe(tiles, 1, isVisited, string(tiles[i]), result)
			isVisited[i] = false
		}
	}
	return len(result)
}

func numTilePossibilitiesDfe(tiles string, n int, isVisited []bool, demo string, result map[string]struct{}) {
	result[demo] = struct{}{}

	if n == len(tiles) {
		return
	}

	for i := 0; i < len(tiles); i++ {
		if !isVisited[i] {
			isVisited[i] = true
			numTilePossibilitiesDfe(tiles, n+1, isVisited, demo+string(tiles[i]), result)
			isVisited[i] = false
		}
	}
}

// 计数 + 回溯
// 我们先用一个哈希表或数组 cnt 统计每个字母出现的次数。
// 接下来定义一个函数 dfs(cnt)，表示当前剩余字母的计数为 cnt 时，能够组成的不同序列的个数。
// 在函数 dfs(cnt) 中，我们枚举 cnt 中每个大于 0 的值 cnt[i]，将 cnt[i] 减 1 表示使用了这个字母，序列个数加 1，
// 然后进行下一层搜索，在搜索结束后，累加返回的序列个数，然后将 cnt[i] 加 1（回溯，恢复现场）。最后返回序列个数。
// func numTilePossibilities(tiles string) int {
// 	cnt := [26]int{}
// 	for _, c := range tiles {
// 		cnt[c-'A']++
// 	}
//
// 	var dfs func(cnt [26]int) int
// 	dfs = func(cnt [26]int) (res int) {
// 		for i, x := range cnt {
// 			if x > 0 {
// 				res++
// 				cnt[i]--
// 				res += dfs(cnt)
// 				cnt[i]++
// 			}
// 		}
// 		return
// 	}
//
// 	return dfs(cnt)
// }

// func main() {
// 	fmt.Println(numTilePossibilities("AAB"))    // 8
// 	fmt.Println(numTilePossibilities("AAABBC")) // 188
// 	fmt.Println(numTilePossibilities("V"))      // 1
// }
