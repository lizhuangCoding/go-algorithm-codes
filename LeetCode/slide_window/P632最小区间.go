package main

import (
	"sort"
)

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/description/

// 思路：
// 首先将 k 组数据升序合并成一组，并记录每个数字所属的组，例如：[[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
// 合并升序后得到：[(0,1),(4,0),(5,2),(9,1),(10,0),(12,1),(15,0),(18,2),(20,1),(22,2),(24,0),(26,0),(30,2)]
// 然后只看所属组的话，那么为 [1,0,2,1,0,1,0,2,1,2,0,0,2]
// 滑动窗口选择同时出现 0，1，2 的区域

func smallestRange(nums [][]int) []int {
	sli := make([][2]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			sli = append(sli, [2]int{nums[i][j], i})
		}
	}
	sort.Slice(sli, func(i, j int) bool {
		return sli[i][0] < sli[j][0]
	})
	// fmt.Println(sli) // [[0 1] [4 0] [5 2] [9 1] [10 0] [12 1] [15 0] [18 2] [20 1] [22 2] [24 0] [26 0] [30 2]]

	resultLeft, resultRight := -1, len(sli)
	left := 0
	m := make(map[int]int)
	for right := 0; right < len(sli); right++ {
		m[sli[right][1]]++

		for len(m) >= len(nums) {
			if (resultLeft == -1 || resultRight == len(sli)) || sli[resultRight][0]-sli[resultLeft][0] > sli[right][0]-sli[left][0] ||
				(sli[resultRight][0]-sli[resultLeft][0] == sli[right][0]-sli[left][0] && sli[resultLeft][0] > sli[left][0]) {
				resultLeft, resultRight = left, right
			}
			m[sli[left][1]]--
			if m[sli[left][1]] == 0 {
				delete(m, sli[left][1])
			}
			left++
		}
	}
	return []int{sli[resultLeft][0], sli[resultRight][0]}
}

// func main() {
// 	fmt.Println(smallestRange([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}))
// 	fmt.Println(smallestRange([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}))
// }
