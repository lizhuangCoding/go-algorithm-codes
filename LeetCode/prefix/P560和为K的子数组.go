package main

// 前缀和 + 哈希表（不能直接用滑动窗口，因为这道题不是单调的）
// 力扣：https://leetcode.cn/problems/subarray-sum-equals-k/

// 这题迷迷糊糊的

func subarraySum(nums []int, k int) (res int) {
	// 前缀和
	prefix := make([]int, len(nums)+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + x
	}

	// 哈希表，统计次数
	m := make(map[int]int, len(prefix))
	for _, v := range prefix {
		res += m[v-k]
		m[v]++ // 记录当前前缀和 v 的出现次数，以便在后续遍历中计算以当前索引为右端点的子数组和为 k 的个数
	}
	return
}

// func main() {
// 	fmt.Println(subarraySum([]int{1, 2, 3, -3, 3}, 3))
// }

/*

以 nums = [1,2,3,-3,3], k = 3 为例，进一步说明 m[v]++ 的作用：

前缀和：prefix = [0,1,3,6,3,6]

遍历：
v = 0：m[-3] = 0, res += 0, m[0] = 1.
v = 1：m[-2] = 0, res += 0, m[1] = 1.
v = 3：m[0] = 1, res += 1，m[3] = 1.
v = 6：m[3] = 1, res += 1，m[6] = 1.
v = 3：m[0] = 1, res += 1，m[3] = 2.
v = 6：m[3] = 2, res += 2，m[6] = 2.

*/
