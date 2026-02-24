package main

// 前缀和

// 暴力：
// func sumOddLengthSubarrays(arr []int) (sum int) {
// 	n := len(arr)
// 	for start := range arr {
// 		for length := 1; start+length <= n; length += 2 {
// 			for _, v := range arr[start : start+length] {
// 				sum += v
// 			}
// 		}
// 	}
// 	return sum
// }

// 前缀和
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func sumOddLengthSubarrays(arr []int) (sum int) {
	pre := make([]int, len(arr)+1)
	// pre[0] = arr[0]
	for i := 0; i < len(arr); i++ {
		pre[i+1] = pre[i] + arr[i]
	}

	res := 0
	for i := 0; i < len(pre); i++ {
		for j := i + 1; j < len(pre); j += 2 {
			res += pre[j] - pre[i]
		}
	}
	return res
}
