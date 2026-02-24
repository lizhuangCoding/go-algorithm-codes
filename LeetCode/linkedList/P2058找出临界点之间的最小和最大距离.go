package main

import "math"

// 链表
// 力扣：https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/

// func nodesBetweenCriticalPoints(head *ListNode) []int {
// 	if head == nil || head.Next == nil || head.Next.Next == nil {
// 		return []int{-1, -1}
// 	}
//
// 	res := make([]int, 0)
// 	result := make([]int, 2)
//
// 	pre := head
// 	cur := head.Next
// 	index := 2
// 	for cur != nil && cur.Next != nil {
// 		if (pre.Val > cur.Val && cur.Val < cur.Next.Val) || (pre.Val < cur.Val && cur.Val > cur.Next.Val) {
// 			res = append(res, index)
// 		}
//
// 		index++
// 		cur = cur.Next
// 		pre = pre.Next
// 	}
//
// 	if len(res) < 2 {
// 		return []int{-1, -1}
// 	}
//
// 	result[0] = math.MaxInt
// 	for i := 0; i < len(res)-1; i++ {
// 		result[0] = min(result[0], res[i+1]-res[i])
// 	}
//
// 	result[1] = res[len(res)-1] - res[0]
// 	return result
// }

// 空间优化：
func nodesBetweenCriticalPoints(head *ListNode) []int {
	a, b, c := head, head.Next, head.Next.Next
	first, last, minDis := 0, 0, math.MaxInt32

	prev := 0                   // 前一个临界点位置
	for i := 1; c != nil; i++ { // 遍历链表，寻找临界点
		if (a.Val < b.Val && b.Val > c.Val) || (a.Val > b.Val && b.Val < c.Val) {
			if first == 0 {
				first = i // 首个临界点位置
			}
			last = i // 最末临界点位置

			if prev > 0 && last-prev < minDis {
				minDis = last - prev // 更新相邻临界点位置之差的最小值
			}
			prev = i
		}
		a, b, c = b, c, c.Next
	}
	if first == last { // 临界点少于两个
		return []int{-1, -1}
	}

	return []int{minDis, last - first}
}
