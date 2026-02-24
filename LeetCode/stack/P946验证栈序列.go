package main

// 栈
// 力扣：https://leetcode.cn/problems/validate-stack-sequences/description/

// func validateStackSequences(pushed []int, popped []int) bool {
// 	index1 := 0
// 	index2 := 0
//
// 	sli := make([]int, 0)
// 	sli = append(sli, pushed[index1])
//
// 	for index1 < len(pushed) && index2 < len(popped) {
// 		// 添加元素
// 		for index1 < len(pushed) && pushed[index1] != popped[index2] {
// 			index1++
// 			if index1 < len(pushed) {
// 				sli = append(sli, pushed[index1])
// 			}
// 		}
//
// 		// 查找可以取出的元素
// 		for index2 < len(popped) && len(sli) >= 1 && sli[len(sli)-1] == popped[index2] {
// 			sli = sli[:len(sli)-1]
// 			index2++
// 		}
// 	}
//
// 	// 判断是否全部都遍历完了
// 	return index2 == len(popped) && len(sli) == 0
// }

// 优化：
func validateStackSequences(pushed, popped []int) bool {
	stack := make([]int, 0)
	index := 0

	// 遍历 pushed，把元素都放到 stack 里面
	for _, x := range pushed {
		stack = append(stack, x)

		// 把当前可以取干净的取干净
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	// 判断是否取干净了
	return len(stack) == 0
}

// func main() {
// 	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
// 	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
// }
