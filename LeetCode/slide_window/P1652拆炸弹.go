package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/defuse-the-bomb/description/

func decrypt(code []int, k int) []int {
	arr := make([]int, len(code))
	if k == 0 {
		return arr
	}
	// 新的code数组长度为：原长度+k的绝对值-1
	// k=3: [5,7,1,4] -> [7,1,4,5,7,1]
	if k > 0 {
		code = append(code[1:], code[:k]...)
	}
	// k=-1: [2,4,9,3] -> [9,3,2,4,9]
	if k < 0 {
		k = -k
		code = append(code[len(code)-k:], code[:len(code)-1]...)
	}

	slide := 0
	for i := 0; i < len(code); i++ {
		slide += code[i]
		if i >= k-1 {
			left := i - k + 1
			arr[left] = slide
			slide -= code[left]
		}
	}
	return arr
}

// 或者：
// func decrypt(code []int, k int) []int {
// 	n := len(code)
// 	ans := make([]int, n)
// 	if k == 0 {
// 		return ans
// 	}
// 	code = append(code, code...)
// 	l, r := 1, k
// 	if k < 0 {
// 		l, r = n+k, n-1
// 	}
// 	sum := 0
// 	for _, v := range code[l : r+1] {
// 		sum += v
// 	}
// 	for i := range ans {
// 		ans[i] = sum
// 		sum -= code[l]
// 		sum += code[r+1]
// 		l, r = l+1, r+1
// 	}
// 	return ans
// }
