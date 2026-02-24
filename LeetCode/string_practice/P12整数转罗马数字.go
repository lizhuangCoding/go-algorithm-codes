package main

// 字符串
// 力扣：https://leetcode.cn/problems/integer-to-roman/description/

// 方法一：暴力
// func intToRoman(num int) string {
// 	result := ""
//
// 	// 千
// 	if num/1000 != 0 {
// 		for i := 0; i < num/1000; i++ {
// 			result += "M"
// 		}
// 		num = num % 1000
// 	}
//
// 	// 百
// 	if num/100 != 0 {
// 		n := num / 100
//
// 		if n == 4 {
// 			result += "CD"
// 		} else if n == 9 {
// 			result += "CM"
// 		}
//
// 		if n >= 5 && n <= 8 {
// 			result += "D"
// 			n -= 5
// 		}
// 		if n >= 1 && n <= 3 {
//
// 			for i := 0; i < n; i++ {
// 				result += "C"
// 			}
// 		}
// 		num = num % 100
// 	}
//
// 	// 十
// 	if num/10 != 0 {
// 		n := num / 10
//
// 		if n == 4 {
// 			result += "XL"
// 		} else if n == 9 {
// 			result += "XC"
// 		}
//
// 		if n >= 5 && n <= 8 {
// 			result += "L"
// 			n -= 5
// 		}
// 		if n >= 1 && n <= 3 {
// 			for i := 0; i < n; i++ {
// 				result += "X"
// 			}
// 		}
// 		num = num % 10
// 	}
//
// 	// 个
// 	if num == 4 {
// 		result += "IV"
// 	} else if num == 9 {
// 		result += "IX"
// 	}
// 	if num >= 5 && num <= 8 {
// 		result += "V"
// 		num -= 5
// 	}
// 	if num >= 1 && num <= 3 {
// 		for i := 0; i < num; i++ {
// 			result += "I"
// 		}
// 	}
//
// 	return result
// }

// 方法二：模拟
func intToRoman(num int) string {
	sli := []struct {
		k int
		v string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, s := range sli {
		for num >= s.k {
			num -= s.k
			result += s.v
		}
		if num == 0 {
			break
		}
	}
	return result
}

/*
I	1
V	5
X	10
L	50
C	100
D	500
M	1000
*/
// func main() {
// 	fmt.Println(intToRoman(3749)) // MMMDCCXLIX
// 	fmt.Println(intToRoman(58))   // LVIII
// 	fmt.Println(intToRoman(1994)) // MCMXCIV
// }
