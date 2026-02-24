package main

// 数论
// 力扣：https://leetcode.cn/problems/count-primes/
// 题解：力扣官方：方法二：埃氏筛(以及优化方法)：https://leetcode.cn/problems/count-primes/solutions/507273/ji-shu-zhi-shu-by-leetcode-solution/
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	res := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			res++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// for i := 0; i < len(isPrime); i++ {
	// 	fmt.Print(i, isPrime[i], "     ")
	// }
	return res
}

// // 优化：用从x*x开始标记
// func countPrimes(n int) int {
// 	isPrimes := make([]bool, n)
// 	for i := range isPrimes {
// 		isPrimes[i] = true
// 	}
//
// 	// i 不需要遍历到 n，而只需要到 sqrt(n) 即可。
// 	// 也就是说，如果在 [2,sqrt(n)] 这个区间之内没有发现可整除因子，就可以直接断定 n 是素数了，因为在区间 [sqrt(n),n] 也一定不会发现可整除因子。
// 	for i := 2; i*i < n; i++ {
// 		if isPrimes[i] {
//
// 			// 考虑这样一个事实：如果 x 是质数，那么大于 x 的 x 的倍数 2x,3x,… 一定不是质数
// 			// for j := 2 * i; j < n; j += i {
// 			//     isPrime[j] = false
// 			// }
//
// 			// 上面的逻辑还可以继续优化
// 			// 还可以继续优化，对于一个质数 x，按上面的写法我们从 2x 开始标记其实是冗余的，应该直接从 x⋅x 开始标记，因为 2x,3x,… 这些数一定在 x 之前就被其他数的倍数标记过了，
// 			// 例如 2 的所有倍数，3 的所有倍数等。
// 			for j := i * i; j < n; j += i {
// 				isPrimes[j] = false
// 			}
// 		}
// 	}
//
// 	count := 0
// 	for i := 2; i < n; i++ {
// 		if isPrimes[i] {
// 			count++
// 		}
// 	}
// 	return count
// }
