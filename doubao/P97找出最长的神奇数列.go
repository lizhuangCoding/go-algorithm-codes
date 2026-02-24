package main

// 双指针
// 青训营：https://www.marscode.cn/practice/qjjjy30l3kv9xk?problem_no=97

// func solution(inp string) string {
// 	// result := ""
// 	// left := 0
// 	// for right := 0; right < len(inp)-1; right++ {
// 	// 	if inp[right] == inp[right+1] {
// 	// 		if len(result) < len(inp[left:right+1]) {
// 	// 			result = inp[left : right+1]
// 	// 		}
// 	// 		left = right + 1
// 	// 	}
// 	// }
//
// 	result := ""
// 	// demo := string(inp[0])
// 	demo := ""
// 	left := 0
// 	for right := 0; right < len(inp); right++ {
// 		demo += string(inp[right])
// 		if right >= 1 {
// 			if inp[right-1] == inp[right] {
//
// 			}
// 		}
//
// 	}
//
// 	if len(result) == 0 {
// 		result = inp
// 	}
// 	return result
// }

// func main() {
// 	// Add your test cases here
// 	// fmt.Println(solution("0101011101") == "010101")
// 	// fmt.Println(solution("11") == "1")
// 	// fmt.Println(solution("100") == "10")
// 	// fmt.Println(solution("110") == "10")
// 	// fmt.Println(solution("010101110101010") == "10101010")
//
// 	fmt.Println(solution("0101011101"))
// 	fmt.Println(solution("11"))
// 	fmt.Println(solution("100"))
// 	fmt.Println(solution("110"))
// 	fmt.Println(solution("010101110101010"))
//
// }
