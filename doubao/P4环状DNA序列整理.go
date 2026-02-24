package main

// 滑动窗口
// 青训营：https://www.marscode.cn/practice/qjjjy30l3kv9xk?problem_no=4

// func solution(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	sli := make([]string, 0)
//
// 	for i := 0; i < len(s); i++ {
// 		sli = append(sli, s[i:]+s[:i])
// 	}
// 	sort.Slice(sli, func(i, j int) bool {
// 		return sli[i] < sli[j]
// 	})
//
// 	return sli[0]
// }
//
// func main() {
// 	// You can add more test cases here
// 	fmt.Println(solution("ATCA") == "AATC")
// 	fmt.Println(solution("CGAGTC") == "AGTCCG")
// 	fmt.Println(solution("TCATGGAGTGCTCCTGGAGGCTGAGTCCATCTCCAGTAG") == "AGGCTGAGTCCATCTCCAGTAGTCATGGAGTGCTCCTGG")
// }
