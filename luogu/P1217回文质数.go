package luogu

import (
	"fmt"
	"math"
	"strconv"
)

// 回文质数
// 洛谷：https://www.luogu.com.cn/problem/P1217

// 1.偶数肯定不是质数。
// 2.偶数位数回文数（除11）必定不是质数。如果一个回文素数的位数是偶数，则它的奇数位上的数字和与偶数位上的数字和必然相等；
// 根据数的整除性理论，容易判断这样的数肯定能被11整除，所以它就不可能是素数。

func main() {
	start, end := 0, 0
	fmt.Scanf("%d %d", &start, &end)

	// 偶数位数回文数（除11）必定不是质数，所以也不用判断 >10000000 的数了
	if end > 10000000 {
		end = 10000000
	}

	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if judge2(s) && judge(i) {
			fmt.Println(i)
		}

	}

}

// 判断质数
func judge(n int) bool {

	// for i := 2; i*i <= n; i++ {
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 判断回文数
func judge2(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}
