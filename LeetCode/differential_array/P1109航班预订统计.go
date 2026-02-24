package main

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/corporate-flight-bookings/description/

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+10)
	for i := 0; i < len(bookings); i++ {
		// -1 的目的是让diff的索引从0开始（因为start是从1开始的，所以要-1）
		diff[bookings[i][0]-1] += bookings[i][2]
		diff[bookings[i][1]+1-1] -= bookings[i][2]
	}

	result := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += diff[i]
		result[i] = cnt
	}
	return result
}

// func main() {
// 	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))             // [10,55,45,25,25]
// 	fmt.Println(corpFlightBookings([][]int{{2, 3, 30}, {2, 3, 45}, {2, 3, 15}, {1, 3, 15}}, 4)) // [15 105 105 0]
// }
