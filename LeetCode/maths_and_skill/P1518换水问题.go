package main

// 模拟，数学

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	demo := numBottles // 空水瓶数量
	for demo/numExchange != 0 {
		res += demo / numExchange
		demo = demo/numExchange + demo%numExchange
	}
	return res
}
