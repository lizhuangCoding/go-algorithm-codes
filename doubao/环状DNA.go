package main

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
// 	fmt.Println(solution("ATCA") == "AATC")
// 	fmt.Println(solution("CGAGTC") == "AGTCCG")
// 	// fmt.Println(solution("TCATGGAGTGCTCCTGGAGGCTGAGTCCATCTCCAGTAG") == "AGGCTGAGTCCATCTCCAGTAGTCATGGAGTGCTCCTGG")
// }
