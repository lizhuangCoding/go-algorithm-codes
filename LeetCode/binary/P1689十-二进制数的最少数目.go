package main

// 找最大值
// 力扣：https://leetcode.cn/problems/partitioning-into-minimum-number-of-deci-binary-numbers/description/

func minPartitions(n string) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result = max(result, int(n[i]-'0'))
	}
	return result
}
