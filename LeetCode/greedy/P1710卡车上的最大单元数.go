package main

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-units-on-a-truck/description/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sli := make([]int, 1005)
	for _, v := range boxTypes {
		sli[v[1]] += v[0]
	}

	res := 0
	// i:容量，sli[i]：箱子数量
	for i := len(sli) - 1; i >= 0 && truckSize > 0; i-- {
		if truckSize > sli[i] {
			res += sli[i] * i
		} else {
			res += truckSize * i
		}
		truckSize -= sli[i]
	}
	return res
}
