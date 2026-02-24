package main

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/count-tested-devices-after-test-operations/description/

// 方法一：模拟
// func countTestedDevices(batteryPercentages []int) int {
// 	n := len(batteryPercentages)
// 	need := 0
// 	for i := 0; i < n; i++ {
// 		if batteryPercentages[i] > 0 {
// 			need++
// 			for j := i + 1; j < n; j++ {
// 				batteryPercentages[j] = max(batteryPercentages[j] - 1, 0)
// 			}
// 		}
// 	}
// 	return need
// }

// 方法二：差分数组思想
func countTestedDevices(batteryPercentages []int) int {
	// 可以把 sum 和 demo 优化为一个变量，这两个变量的值是相同的
	sum := 0
	demo := 0
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i]-demo > 0 {
			sum++
			demo++
		}
	}
	return sum
}
