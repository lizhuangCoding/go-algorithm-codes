package main

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/car-pooling/description/

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1010)

	for i := 0; i < len(trips); i++ {
		diff[trips[i][1]] += trips[i][0]
		// 乘客会在trips[i][2]下车，这时候其他人可以上车，所以要 diff[trips[i][2]]-= trips[i][0]，而不是 diff[trips[i][2]+1]-= trips[i][0]
		diff[trips[i][2]] -= trips[i][0]
	}

	cnt := 0
	for i := 0; i < len(diff); i++ {
		cnt += diff[i]
		if cnt > capacity {
			return false
		}
	}
	return true
}
