package main

// 滑动窗口，二进制
// 力扣：https://leetcode.cn/problems/check-if-a-string-contains-all-binary-codes-of-size-k/

// func hasAllCodes(s string, k int) bool {
// 	// need := int(math.Pow(float64(2), float64(k)))
// 	m := make(map[string]bool)
//
// 	for right := k - 1; right < len(s); right++ {
// 		left := right - (k - 1)
// 		m[s[left:right+1]] = true
// 	}
//
// 	// return len(m) == need
// 	return len(m) == 1<<k
// }

// func hasAllCodes(s string, k int) bool {
//     n, set := len(s), make(map[string]bool)
//     if n < k { return false }
//     for i := 0; i <= n - k; i++ {
//         set[s[i:i+k]] = true
//         if len(set) == 1 << k {
//             return true
//         }
//     }
//     return len(set) == 1 << k
// }

// 优化：使用二进制运算+滑动窗口
func hasAllCodes(s string, k int) bool {
	need := 1 << k
	// mask：用来创建长度为 k 的二进制数字掩码，例如 k = 3 时，mask = 111 (7)。
	// curr：用来表示当前子串所对应的整数值。
	visited, mask, curr := make([]bool, need), need-1, 0

	// 检查子串数量是否足够
	if len(s)-k < mask {
		return false
	}

	for i := 0; i < len(s); i++ {
		// 位运算构建二进制组合:
		// curr << 1：左移一位，相当于将现有二进制数的最高位空出来。
		// int(s[i]-'0')：将字符 '0' 或 '1' 转换为数字 0 或 1。
		// |（按位或）：将 s[i] 的值添加到 curr 中。
		// & mask：确保 curr 只保留长度为 k 的位数。超出的部分被屏蔽掉
		curr = ((curr << 1) | int(s[i]-'0')) & mask
		// Build the number associated with the initial substring

		if i < k-1 {
			continue // 由于 curr 需要长度为 k 的二进制串，所以在未达到 k-1 时跳过循环。
		}
		if !visited[curr] {
			need--
			visited[curr] = true
			if need == 0 {
				return true
			}
		}
	}
	return false
}
