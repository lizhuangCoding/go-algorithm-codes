package main

// 数学
// 力扣：https://leetcode.cn/problems/add-digits/

// 方法一：模拟
// func addDigits(num int) int {
// 	for num >= 10 {
// 		sum := 0
// 		for ; num > 0; num /= 10 {
// 			sum += num % 10
// 		}
// 		num = sum
// 	}
// 	return num
// }

// 方法二：数学
// 题解：力扣官方：https://leetcode.cn/problems/add-digits/solutions/1301157/ge-wei-xiang-jia-by-leetcode-solution-u4kj/
// func addDigits(num int) int {
// 	return (num-1)%9 + 1
// }
// 没看懂，方法和下面同理

// 可以举几个例子试一试
func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}
	return num % 9
}
