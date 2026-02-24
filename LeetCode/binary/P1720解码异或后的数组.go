package main

// 异或
// 力扣：https://leetcode.cn/problems/decode-xored-array/

func decode(encoded []int, first int) []int {
	res := make([]int, 0)
	res = append(res, first)
	for i := 0; i < len(encoded); i++ {
		res = append(res, res[i]^encoded[i])
	}
	return res
}

/*

	encoded = [6,2,7,3], first = 4 输出：[4,2,0,7,4]

	4 ^ x = 6 (x = 2)
	100 ^ x = 110
	x = 110 ^ 100
	所以 x = 4 ^ 6

	fmt.Println(4^6, 2^2, 0^7, 7^3) // 2 0 7 4

*/
