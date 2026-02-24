package main

// 分组循环，字符串
// 力扣：https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/

func longestBeautifulSubstring(word string) int {
	target := "aeiou"
	index := 0  // 计算到了哪一个元音字母
	result := 0 // 最终的结果
	demo := 0   // 临时变量

	for i := 0; i < len(word); i++ {
		start := i

		for ; i < len(word)-1 && word[i] == word[i+1]; i++ {
		}

		if word[i] == target[index] {
			index++
			demo += i - start + 1

			if index == len(target) {
				result = longestBeautifulSubstringMax(result, demo)
				// 重置为0
				demo = 0
				index = 0
			}
		} else {
			// 发现不符合 aeiou 的顺序，重置之前的计数
			demo = 0
			index = 0

			// 特殊情况处理：不能直接就这样结束了，必须要加下面的代码，因为如果是 aeiaaaaeiou，循环到 aei 后一定要及时把 aaaa 也算上，不能跳过 aaa
			if word[i] == target[0] {
				demo = i - start + 1
				index++
			}

		}
	}
	return result
}

func longestBeautifulSubstringMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	// fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu")) // 13
// 	// fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou"))           // 5
// 	// fmt.Println(longestBeautifulSubstring("a"))                              // 0
// 	// 特殊情况：循环到 aei 后没有算 aaaa
// 	fmt.Println(longestBeautifulSubstring("aeiaaaaeiou")) // 8
// }
