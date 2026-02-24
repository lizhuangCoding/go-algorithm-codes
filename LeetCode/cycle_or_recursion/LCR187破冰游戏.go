package cycle_or_recursion

// 递归/约瑟夫问题（不太会）
// 力扣：https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/description/

// 方法一：模拟
// func iceBreakingGame(num int, target int) int {
// 	nums := make([]int, 0)
// 	for i := 0; i < num; i++ {
// 		nums = append(nums, i)
// 	}
//
// 	index := 0
// 	for len(nums) > 1 {
// 		index = (index + target - 1) % len(nums)
// 		nums = append(nums[:index], nums[index+1:]...)
// 	}
// 	return nums[0]
// }

// 方法二：数学+递归（约瑟夫问题）
func iceBreakingGame(num int, target int) int {
	return diguiIceBreakingGame(num, target)
}

func diguiIceBreakingGame(num int, target int) int {
	if num == 1 {
		return 0
	}
	x := diguiIceBreakingGame(num-1, target)
	return (target + x) % num
}

// 假设 num = 5, target = 2
// 调用 f(5, 2):
// 进入 f(4, 2)
// 进入 f(3, 2)
// 进入 f(2, 2)
// 进入 f(1, 2), 返回 0
// 计算 (2+0) % 2 -> 0
// 计算 (2+0) % 3 -> 2
// 计算 (2+2) % 4 -> 0
// 计算 (2+0) % 5 -> 2

// 假设 num = 7, target = 4
// 调用 f(7, 4)
// 进入 f(6, 2)
// 进入 f(5, 2)
// 进入 f(4, 2)
// 进入 f(3, 2)
// 进入 f(2, 2)
// 进入 f(1, 2), 返回 0
// 计算 (4+0) % 2 -> 0
// 计算 (4+0) % 3 -> 1
// 计算 (4+1) % 4 -> 1
// 计算 (4+1) % 5 -> 0
// 计算 (4+0) % 6 -> 4
// 计算 (4+4) % 7 -> 1

// 方法三：数学 + 迭代（约瑟夫问题）（没明白）
// func iceBreakingGame(num int, target int) int {
// 	survivor := 0
// 	for i := 2; i <= num; i++ {
// 		survivor = (survivor + target) % i
// 	}
// 	return survivor
// }
