package main

// 数组，双指针
// 力扣：https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/

// func trainingPlan(actions []int) []int {
// 	left, right := 0, len(actions)-1
//
// 	for left < right {
//
// 		for left < right && actions[left]&1 != 0 {
// 			left++
// 		}
// 		for left < right && actions[right]&1 != 1 {
// 			right--
// 		}
//
// 		if left < right {
// 			actions[left], actions[right] = actions[right], actions[left]
// 		}
// 	}
// 	return actions
// }

// 如果左右两边分割的条件不是奇数和偶数，而是正负呢？
// 具有可扩展性的代码：
func trainingPlan(actions []int) []int {
	return trainingPlanDemo(actions, trainingPlanIsEven)
}

func trainingPlanDemo(actions []int, f func(n int) bool) []int {
	left, right := 0, len(actions)-1

	for left < right {
		for left < right && !f(actions[left]) {
			left++
		}
		for left < right && f(actions[right]) {
			right--
		}
		if left < right {
			actions[left], actions[right] = actions[right], actions[left]
		}
	}
	return actions
}

// 判断是否是偶数（如果判断条件改变了，比如改为左右两边根据正负来分割只需要修改这个函数里面的内容）
func trainingPlanIsEven(n int) bool {
	return n&1 == 0
}

// func main() {
// 	fmt.Println(trainingPlan([]int{1, 2, 3, 4, 5}))
// }
