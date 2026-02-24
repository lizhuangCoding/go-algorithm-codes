package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串，双指针
// 力扣：https://leetcode.cn/problems/compare-version-numbers/description/

// 方法一：分割为切片
func compareVersion(version1 string, version2 string) int {
	sli1 := strings.Split(version1, ".")
	sli2 := strings.Split(version2, ".")

	// 找到最小的长度
	minLen := min(len(sli1), len(sli2))

	for i := 0; i < minLen; i++ {
		n1, _ := strconv.Atoi(sli1[i])
		n2, _ := strconv.Atoi(sli2[i])
		fmt.Println(n1, n2)
		if n1 < n2 {
			return -1
		} else if n1 > n2 {
			return 1
		}
	}

	for i := minLen; i < len(sli1); i++ {
		n1, _ := strconv.Atoi(sli1[i])
		if n1 != 0 {
			return 1
		}
	}

	for i := minLen; i < len(sli2); i++ {
		n2, _ := strconv.Atoi(sli2[i])
		if n2 != 0 {
			return -1
		}
	}

	return 0
}

// 方法二：双指针（没看）
// func compareVersion(version1, version2 string) int {
// 	n, m := len(version1), len(version2)
// 	i, j := 0, 0
//
// 	for i < n || j < m {
// 		x := 0
// 		for ; i < n && version1[i] != '.'; i++ {
// 			x = x*10 + int(version1[i]-'0')
// 		}
//
// 		i++ // 跳过点号
// 		y := 0
// 		for ; j < m && version2[j] != '.'; j++ {
// 			y = y*10 + int(version2[j]-'0')
// 		}
//
// 		j++ // 跳过点号
//
// 		if x > y {
// 			return 1
// 		}
//
// 		if x < y {
// 			return -1
// 		}
//
// 	}
//
// 	return 0
// }
