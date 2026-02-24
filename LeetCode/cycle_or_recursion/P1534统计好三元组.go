package cycle_or_recursion

// 签到：循环

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if absCountGoodTriplets(arr[i]-arr[j]) <= a && absCountGoodTriplets(arr[j]-arr[k]) <= b && absCountGoodTriplets(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}
	return res
}

func absCountGoodTriplets(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
