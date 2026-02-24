package main

// 数组
// 力扣：https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/

// 方法一：暴力，如果数值是大数怎么办？
func countNumbers(cnt int) []int {
	res := make([]int, 0)
	n := 1
	for i := 0; i < cnt; i++ {
		n *= 10
	}

	for i := 1; i < n; i++ {
		res = append(res, i)
	}
	return res
}

// （在这道题目中应该不会出现大数问题，因为函数的返回值是[]int，所以不会出现大数问题）
// 方法二：如果结果存在大数，用数组，递归：
// 时间复杂度：O(n×10^n)，其中 n 是数字的位数。
// 空间复杂度：O(n)，主要由递归栈空间和存储数字的数组空间决定。
// func countNumbers(n int) {
// 	if n <= 0 {
// 		return
// 	}
//
// 	number := make([]byte, n)
// 	countNumbersPrint1ToMaxOfNDigitsRecursively(number, 0)
// }
//
// // 递归打印函数
// func countNumbersPrint1ToMaxOfNDigitsRecursively(number []byte, index int) {
// 	if index == len(number) {
// 		countNumbersPrintNumber(number)
// 		return
// 	}
//
// 	// 数字只能取0到9，所以遍历10次
// 	for i := 0; i < 10; i++ {
// 		number[index] = byte(i + '0')
// 		countNumbersPrint1ToMaxOfNDigitsRecursively(number, index+1)
// 	}
// }
//
// // 打印数字函数
// func countNumbersPrintNumber(number []byte) {
// 	// 跳过前导0
// 	start := 0
// 	for start < len(number)-1 && number[start] == '0' {
// 		start++
// 	}
// 	for i := start; i < len(number); i++ {
// 		fmt.Printf("%c", number[i])
// 	}
// 	fmt.Println()
// }

// 方法三：用长度为n的字符数组，从后向前开始遍历字符数组，每一次循环加减一次（要考虑进位的问题）终止条件是长度为n的字符数组溢出，即全部元素都是999999...

// // print1ToMaxOfNDigits 函数用于从 1 打印到 n 位最大数
// func countNumbers(n int) {
// 	if n <= 0 {
// 		return
// 	}
//
// 	// 初始化字符数组
// 	num := make([]byte, n)
// 	for i := 0; i < n; i++ {
// 		num[i] = '0'
// 	}
//
// 	// 循环加一操作，直到数组溢出
// 	for !increment(num) {
// 		printNumber(num)
// 	}
// }
//
// // increment 函数用于对字符数组表示的数字进行加一操作，并处理进位
// func increment(num []byte) bool {
// 	// 标记是否有进位
// 	isOverflow := false
// 	// 进位值，初始为 1 表示加一操作
// 	nTakeOver := 0
// 	// 数组长度
// 	nLength := len(num)
//
// 	// 从后向前遍历字符数组
// 	for i := nLength - 1; i >= 0; i-- {
// 		// 当前位数字加上进位值
// 		nSum := int(num[i]-'0') + nTakeOver
// 		// 如果是最后一位，再加 1
// 		if i == nLength-1 {
// 			nSum++
// 		}
//
// 		// 如果相加结果大于等于 10，产生进位
// 		if nSum >= 10 {
// 			if i == 0 {
// 				// 如果是最高位产生进位，说明数组溢出
// 				isOverflow = true
// 			} else {
// 				// 进位值设为 1
// 				nTakeOver = 1
// 				// 当前位数字置为 0
// 				num[i] = '0'
// 			}
// 		} else {
// 			// 没有进位，进位值置为 0
// 			nTakeOver = 0
// 			// 当前位数字更新为相加结果
// 			num[i] = byte(nSum + '0')
// 			break
// 		}
// 	}
//
// 	return isOverflow
// }
//
// // printNumber 函数用于打印字符数组表示的数字，跳过前导 0
// func printNumber(num []byte) {
// 	isBeginning0 := true
// 	nLength := len(num)
//
// 	for i := 0; i < nLength; i++ {
// 		if isBeginning0 && num[i] != '0' {
// 			isBeginning0 = false
// 		}
//
// 		if !isBeginning0 {
// 			fmt.Printf("%c", num[i])
// 		}
// 	}
//
// 	if !isBeginning0 {
// 		fmt.Println()
// 	}
// }
