package main

// 哈希，桶排序
// 力扣：https://leetcode.cn/problems/top-k-frequent-elements/?envType=study-plan-v2&envId=top-100-liked

func topKFrequent(nums []int, k int) []int {
	maxRes := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		maxRes = max(maxRes, m[nums[i]])
	}

	demo := make([][]int, maxRes+1)
	for k, v := range m {
		demo[v] = append(demo[v], k)
	}

	result := make([]int, 0)
	for i := maxRes; i >= 0; i-- {
		if len(result) < k {
			if len(demo[i]) != 0 {
				result = append(result, demo[i]...)
			}
		} else {
			break
		}
	}

	return result
}
