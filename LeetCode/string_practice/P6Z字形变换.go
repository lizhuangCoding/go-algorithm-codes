package main

// 字符串，二维数组
// 力扣：https://leetcode.cn/problems/zigzag-conversion/

// 模拟
// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s
// 	}
// 	var arr [1001][1001]byte
// 	index := 0
// 	arr[1][1] = s[index]
// 	index++
//
// 	i, j := 1, 1
// 	for index < len(s) {
// 		// 从上当下
// 		for k := 0; k < numRows-1 && index < len(s); k++ {
// 			i++
// 			arr[i][j] = s[index]
// 			index++
// 		}
// 		// 左下角到右上角
// 		for k := 0; k < numRows-1 && index < len(s); k++ {
// 			i--
// 			j++
// 			arr[i][j] = s[index]
// 			index++
// 		}
// 	}
//
// 	result := ""
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr[i]); j++ {
// 			if arr[i][j] != 0 {
// 				result += string(arr[i][j])
// 			}
// 		}
// 	}
// 	return result
// }

// 优化：
// 时间复杂度：O(r⋅n)，其中 r=numRows，n 为字符串 s 的长度。时间主要消耗在矩阵的创建和遍历上，矩阵的行数为 r，列数可以视为 O(n)。
// 空间复杂度：O(r⋅n)。矩阵需要 O(r⋅n) 的空间。
// func convert(s string, numRows int) string {
// 	n := len(s)
// 	row := numRows // 有多少行
// 	if row == 1 || row >= n {
// 		return s
// 	}
//
// 	// 一个周期是：向下填写 r 个字符，然后向右上继续填写 r−2 个字符，最后回到第一行
// 	t := row*2 - 2 // 一个周期有几个元素
// 	// (n + t - 1) / t： 有几个周期。t-1的原因：不满一个周期也要算上
// 	col := (n + t - 1) / t * (row - 1) // 有多少列
//
// 	arr := make([][]byte, row)
// 	for i := range arr {
// 		arr[i] = make([]byte, col)
// 	}
//
// 	x, y := 0, 0
// 	for i, ch := range s {
// 		arr[x][y] = byte(ch)
// 		if i%t < row-1 {
// 			x++ // 向下移动
// 		} else {
// 			x--
// 			y++ // 向右上移动
// 		}
// 	}
//
// 	ans := make([]byte, 0, n)
// 	for _, row := range arr {
// 		for _, ch := range row {
// 			if ch > 0 {
// 				ans = append(ans, ch)
// 			}
// 		}
// 	}
// 	return string(ans)
// }

// 优化：压缩矩阵空间
// 注意到每次往矩阵的某一行添加字符时，都会添加到该行上一个字符的右侧，且最后组成答案时只会用到每行的非空字符。 因此我们可以将矩阵的每行初始化为一个空列表，每次向某一行添加字符时，添加到该行的列表末尾即可。
// 时间复杂度：O(n)。 空间复杂度：O(n)。压缩后的矩阵需要 O(n) 的空间。
func convert(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}

	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}

	result := ""
	for _, v := range mat {
		result += string(v)
	}
	return result
}

// 或者：
// func convert(s string, numRows int) string {
// 	if numRows < 2 {
// 		return s
// 	}
// 	i, flag := 0, -1
// 	ans := make([]string, numRows)
// 	for _, char := range s {
// 		ans[i] += string(char)
// 		if i == 0 || i == numRows-1 {
// 			flag = -flag
// 		}
// 		i += flag
// 	}
// 	return strings.Join(ans, "")
// }
