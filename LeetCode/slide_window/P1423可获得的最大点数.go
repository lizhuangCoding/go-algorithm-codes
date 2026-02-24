package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/description/

// 方法一：正向思维，通过修改切片，再找适合的滑动窗口
// func maxScore(cardPoints []int, k int) int {
// 	if k > len(cardPoints) {
// 		return 0
// 	}
//
// 	// 更改切片
// 	// cardPoints = [1,2,3,4,5,6,1], k = 3
// 	// 将 cardPoints 改为 [1,2,3,4,5,6,1,1,2,3,4,5,6,1]，因为必须从两边开始取，所以取中间，再改为： [5,6,1,1,2,3]
// 	cardPoints = append(cardPoints, cardPoints...)
// 	cardPoints = cardPoints[len(cardPoints)/2-k : len(cardPoints)/2+k]
//
// 	// 滑动窗口
// 	maxResult := 0
// 	demo := 0
// 	for right := 0; right < len(cardPoints); right++ {
// 		demo += cardPoints[right]
//
// 		if right >= k-1 {
// 			if maxResult < demo {
// 				maxResult = demo
// 			}
// 			left := right - (k - 1)
// 			demo -= cardPoints[left]
// 		}
// 	}
// 	return maxResult
// }

// 或者，规律：
// 前 k 个数的和。
// 前 k−1 个数以及后 1 个数的和。
// 前 k−2 个数以及后 2 个数的和。
// ……
// 前 2 个数以及后 k−2 个数的和。
// 前 1 个数以及后 k−1 个数的和。
// 后 k 个数的和。
func maxScore(cardPoints []int, k int) int {
	s := 0
	// 取前k个数的和
	for _, v := range cardPoints[:k] {
		s += v
	}
	// fmt.Println("s = ", s)

	ans := s
	// 需要滑动k次，由前k个数的和 滑动为 后k个数的和，看一看那个滑动窗口的值最大
	for i := 1; i <= k; i++ {
		s += cardPoints[len(cardPoints)-i] - cardPoints[k-i]
		// fmt.Printf("s = %d, cardPoints[len(cardPoints)-i] = %d,  cardPoints[k-i] = %d\n", s, cardPoints[len(cardPoints)-i], cardPoints[k-i])
		if ans < s {
			ans = s
		}
	}
	return ans
}

// 方法二：逆向思维，通过求出剩余卡牌点数之和的最小值，来求出拿走卡牌点数之和的最大值。
// func maxScore(cardPoints []int, k int) int {
// 	total := 0
// 	for _, v := range cardPoints {
// 		total += v
// 	}
//
// 	// 滑动窗口大小为 len(cardPoints)-k
// 	windowSize := len(cardPoints) - k
// 	// 选前 n-k 个作为初始值
// 	sum := 0
// 	for _, v := range cardPoints[:windowSize] {
// 		sum += v
// 	}
// 	minSum := sum
//
// 	for i := windowSize; i < len(cardPoints); i++ {
// 		// 滑动窗口每向右移动一格，增加从右侧进入窗口的元素值，并减少从左侧离开窗口的元素值
// 		sum += cardPoints[i] - cardPoints[i-windowSize]
// 		if minSum > sum {
// 			minSum = sum
// 		}
// 	}
//
// 	return total - minSum
// }
