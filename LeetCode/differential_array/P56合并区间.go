package main

import (
	"sort"
)

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/merge-intervals/

// // 错误：
//
//	func merge(intervals [][]int) [][]int {
//		l := 0
//		for i := 0; i < len(intervals); i++ {
//			if intervals[i][1] > l {
//				l = intervals[i][1]
//			}
//		}
//		diff := make([]int, l+2)
//		for i := 0; i < len(intervals); i++ {
//			diff[intervals[i][0]]++
//			diff[intervals[i][1]]--
//		}
//		// fmt.Println(diff)
//
//		result := make([][]int, 0)
//		cnt := 0
//		left, right := 0, 0
//		isHaveLeft, isHaveRight := false, false
//
//		for i := 0; i < len(diff); i++ {
//			cnt += diff[i]
//			if cnt > 0 && !isHaveLeft {
//				isHaveLeft = true
//				left = i
//				isHaveRight = false
//			}
//			if i != 0 && cnt == 0 && !isHaveRight {
//				isHaveRight = true
//				right = i
//				result = append(result, [][]int{{left, right}}...)
//				isHaveLeft = false
//			}
//		}
//		return result
//	}
// func main() {
// 	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,18]]
// 	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1,5]]
// 	fmt.Println(merge([][]int{{1, 4}, {5, 6}}))                    // [[1,4],[5,6]]
// 	fmt.Println(merge([][]int{{1, 4}, {0, 0}}))                    // 输出：[[1,4]]。应该输出：[[0,0],[1,4]]
// }

func merge(intervals [][]int) [][]int {
	// 按照左端点从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	for _, p := range intervals {
		m := len(result)
		if m > 0 && p[0] <= result[m-1][1] { // 可以合并
			// 更新右端点最大值
			if p[1] > result[m-1][1] {
				result[m-1][1] = p[1]
			}
		} else { // 不相交，无法合并
			result = append(result, p) // 新的合并区间
		}
	}
	return result
}
