package main

// 字符串
// 青训营：https://www.marscode.cn/practice/qjjjy30l3kv9xk?problem_no=100

// func solution(inp string) string {
// 	for i := 1; i <= len(inp)/2; i++ {
// 		// 如果可以整除
// 		if len(inp)%i == 0 {
// 			s := inp[:i]
// 			demo := strings.Repeat(s, len(inp)/i)
// 			if demo == inp {
// 				return s
// 			}
// 		}
// 	}
// 	return ""
// }
//
// func main() {
// 	fmt.Println(solution("abcabcabcabc") == "abc")
// 	fmt.Println(solution("abababab") == "ab")
// 	fmt.Println(solution("abcd") == "")
// 	fmt.Println(solution("aaaa") == "a")
// 	fmt.Println(solution("xyzxyzxyz") == "xyz")
// }
