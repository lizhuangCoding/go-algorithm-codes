package main

// 签到：数组

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	for i := 0; i < len(items); i++ {
		if ruleKey == "type" && ruleValue == items[i][0] {
			res++
		} else if ruleKey == "color" && ruleValue == items[i][1] {
			res++
		} else if ruleKey == "name" && ruleValue == items[i][2] {
			res++
		}
	}
	return res
}
