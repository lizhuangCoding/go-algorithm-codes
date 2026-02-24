package main

// 字符串：签到

func isBalanced(num string) bool {
	var ji, ou uint8 = 0, 0
	for i := 0; i < len(num); i++ {
		demo := num[i] - '0'
		if i%2 == 1 {
			ji += demo
		} else {
			ou += demo
		}
	}
	return ji == ou
}

// func main() {
// 	fmt.Println(isBalanced("1234"))
// 	fmt.Println(isBalanced("24123"))
// }
