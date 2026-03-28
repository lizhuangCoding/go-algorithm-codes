package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/uOAnQW/description/

// 将 cards 从大到小排序后，先贪心的将后 cnt 个数字加起来，若此时 sum 为偶数，直接返回即可。
// 若此时答案为奇数，有两种方案:
// 1.在后面找到一个最大的奇数与前 cnt 个数中最小的偶数进行替换；
// 2.在后面找到一个最大的偶数与前 cnt 个数中最小的奇数进行替换。
// 两种方案选最大值即可。

func maximumScore(cards []int, cnt int) int {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	res := 0
	odd, even := -1, -1
	for i := 0; i < cnt; i++ {
		res += cards[i]
		if cards[i]%2 == 1 {
			odd = cards[i]
		} else {
			even = cards[i]
		}
	}

	if res%2 == 0 {
		return res
	}

	result := 0

	for i := cnt; i < len(cards); i++ {
		if cards[i]%2 == 1 && even != -1 { // 用偶数换奇数
			result = max(result, res-even+cards[i])
		} else if cards[i]%2 == 0 && odd != -1 { // 用奇数换偶数
			result = max(result, res-odd+cards[i])
		}
	}

	return result
}
