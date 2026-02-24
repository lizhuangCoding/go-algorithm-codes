package main

// 简单：模拟
func numberOfAlternatingGroups(colors []int) int {
	colors = append(colors, colors[0:2]...)

	res := 0
	for i := 0; i <= len(colors)-3; i++ {
		if colors[i] == colors[i+2] && colors[i] != colors[i+1] {
			res++
		}
	}
	return res
}
