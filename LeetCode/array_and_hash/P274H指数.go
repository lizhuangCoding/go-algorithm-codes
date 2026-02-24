package main

// 力扣：https://leetcode.cn/problems/h-index/description/

// 方法一：排序
// 时间复杂度：O(n*logn)
// func hIndex(citations []int) int {
// 	result := 0
//
// 	sort.Slice(citations, func(i, j int) bool {
// 		return citations[i] > citations[j]
// 	})
//
// 	for i := 0; i < len(citations); i++ {
// 		if result+1 <= citations[i] {
// 			result++
// 		} else {
// 			break
// 		}
// 	}
// 	return result
// }

// 方法二：存储内容
// 创建一个大小为 n+1 的计数数组，其中 n 是论文总数
// 统计规则：如果一篇论文的引用次数 c ≥ n，就将其计数在 cnt[n] 中（因为 h 指数最大为 n）
// 从高到低累加计数，找到第一个满足 累计论文数 ≥ 引用次数 的位置
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// func hIndex(citations []int) int {
// 	n := len(citations)
// 	cnt := make([]int, n+1)
// 	for _, c := range citations {
// 		cnt[min(c, n)]++ // 引用次数 > n，等价于引用次数为 n
// 	}
//
// 	s := 0
// 	for i := n; ; i-- { // i=0 的时候，s>=i 一定成立
// 		s += cnt[i]
//
// 		// 引用次数 ≥ i 的论文总数
// 		if s >= i { // 说明有至少 i 篇论文的引用次数至少为 i
// 			return i
// 		}
// 	}
// }
