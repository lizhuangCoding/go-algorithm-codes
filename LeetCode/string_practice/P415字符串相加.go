package main

import "strconv"

// 字符串：模拟两个整数相加
// 力扣：https://leetcode.cn/problems/add-strings/description/

// 我们定义两个指针i和j分别指向 num1和num2 的末尾，即最低位，同时定义一个变量 add 维护当前是否有进位，然后从末尾到开头逐位相加即可。
// 你可能会想两个数字位数不同怎么处理，这里我们统一在指针当前下标处于负数的时候返回 0，等价于对位数较短的数字进行了补零操作，这样就可以除去两个数字位数不同情况的处理
func addStrings(num1 string, num2 string) string {
	add := 0 // 是否有进位
	ans := ""

	// 注意判断的条件：add != 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int // 当前索引位置的数字
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		result := x + y + add
		add = result / 10 // 判断是否需要进位

		ans = strconv.Itoa(result%10) + ans
	}
	return ans
}
