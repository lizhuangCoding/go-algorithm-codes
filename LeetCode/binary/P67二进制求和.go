package main

import (
	"math"
	"strconv"
)

// 二进制
// LeetCode：https://leetcode.cn/problems/add-binary/description/

// 使用二进制的计算方式：末尾对齐，逐位相加
func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	result := ""

	count := 0
	for i := 0; i < int(math.Max(float64(lenA), float64(lenB))); i++ {
		if i < lenA {
			count += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			count += int(b[lenB-i-1] - '0')
		}
		result = strconv.Itoa(count%2) + result
		count /= 2
	}

	if count != 0 {
		result = "1" + result
	}

	return result
}

// 或者
func addBinary2(a string, b string) string {
	if len(a) < len(b) {
		return addBinary(b, a)
	}
	difference := len(a) - len(b)

	for i := 0; i < difference; i++ {
		b = "0" + b
	}
	rst := ""
	carry := "0"
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == "0" {
				rst = "0" + rst
				carry = "1"
			} else {
				rst = "1" + rst
			}
		} else if a[i] == '0' && b[i] == '0' {
			if carry == "0" {
				rst = "0" + rst
			} else {
				rst = "1" + rst
				carry = "0"
			}
		} else {
			if carry == "0" {
				rst = "1" + rst
			} else {
				rst = "0" + rst
				carry = "1"
			}
		}
	}
	if carry == "1" {
		rst = "1" + rst
	}
	return rst
}
