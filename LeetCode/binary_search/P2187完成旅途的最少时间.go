package main

// 二分查找：对答案进行二分查找
// 力扣：https://leetcode.cn/problems/minimum-time-to-complete-trips/description/

func minimumTime(time []int, totalTrips int) int64 {
	maxTime := 0 // 找到最大时间，计算r边界的时候要用到
	for _, v := range time {
		maxTime = max(maxTime, v)
	}

	res := 0
	l := 1 // 最短时间1
	// time = [1,2,3], totalTrips = 5。最大需要的时间为： 5除以长度3，并向上取整=2，每辆车都要跑两次，然后2 * 一次车的最长时间3=6，所以最长时间为6
	r := (totalTrips + len(time) - 1) / len(time) * maxTime

	for l <= r {
		mid := (l + r) >> 1 // 现在有mid分钟，看一看mid分钟内可以走多少辆车

		total := 0
		for i := 0; i < len(time); i++ {
			total += mid / time[i]
		}

		if total >= totalTrips {
			// 如果是 > ：说明可以走的车的数量太多了，足够了，但是要找最少时间，所以继续往左边缩短
			res = mid
			r = mid - 1
		} else { //  可以走的车的数量太少了
			l = mid + 1
		}

	}
	return int64(res)
}
