package main

import "math"

// 暴力，哈希，数学几何
// 力扣：https://leetcode.cn/problems/max-points-on-a-line/description/

func maxPoints(points [][]int) (ans int) {
	for i, p := range points { // 假设直线一定经过 points[i]
		x, y := p[0], p[1]
		cnt := map[float64]int{}
		for _, q := range points[i+1:] {
			dx, dy := q[0]-x, q[1]-y
			k := math.MaxFloat64
			if dx != 0 {
				k = float64(dy) / float64(dx)
			}
			cnt[k]++
			ans = max(ans, cnt[k]) // 这里没有算上 (x,y) 这个点，最后再加一
		}
	}
	return ans + 1
}

// 时间复杂度：O(n^2),其中 n 是 points 的长度
// 空间复杂度：O(n)
