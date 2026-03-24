package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/eat-pizzas/description/

func maxWeight(pizzas []int) int64 {
	sort.Slice(pizzas, func(i, j int) bool {
		return pizzas[i] < pizzas[j]
	})

	l := len(pizzas)
	// p1：奇数天吃几个Z，p2：偶数天吃几个Y
	p1, p2 := (l/4+1)/2, l/4/2

	// 奇数天吃披萨
	res := 0
	for i := 1; i <= p1; i++ {
		res += pizzas[l-i]
	}

	// 偶数天吃披萨
	for i := 1; i <= p2; i++ {
		res += pizzas[l-(p1+i*2)]
	}

	return int64(res)
}
