package main

// 单调栈
// 力扣：https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/description/

// 时间复杂度：
// Get_max：时间复杂度为O(1)，为只需要直接访问 maxSli 队列的队首元素。
// Add：均摊时间复杂度为O(1)。虽然在最坏情况下可能需要移除 maxSli 队列中的所有元素，但每个元素最多入队和出队 maxSli 队列一次。
// Remove：时间复杂度为O(1)，只需要移除 sli 和 maxSli 队列的队首元素。
// 空间复杂度：O(n)

// type Checkout struct {
// 	sli    []int // 存储队列中的实际元素
// 	maxSli []int // 辅助队列，用于存储当前队列中的最大值信息。它是一个单调递减队列，队首元素始终是当前队列 sli 中的最大值。
// }
//
// func Constructor() Checkout {
// 	return Checkout{}
// }
//
// func (c *Checkout) Get_max() int {
// 	if len(c.sli) == 0 {
// 		return -1
// 	}
// 	return c.maxSli[0]
// }
//
// func (c *Checkout) Add(value int) {
// 	c.sli = append(c.sli, value)
//
// 	// 从队尾开始检查，如果 maxSli 队列不为空且 value 大于 maxSli 队列的队尾元素，就将 maxSli 队列的队尾元素移除，
// 	// 直到 maxSli 队列为空或者 value 小于等于 maxSli 队列的队尾元素。这样做的目的是保证 maxSli 队列的单调递减性质。
// 	for len(c.maxSli) > 0 && value > c.maxSli[len(c.maxSli)-1] {
// 		c.maxSli = c.maxSli[:len(c.maxSli)-1]
// 	}
//
// 	c.maxSli = append(c.maxSli, value)
// }
//
// func (c *Checkout) Remove() int {
// 	if len(c.sli) == 0 {
// 		return -1
// 	}
//
// 	// 如果 sli 队列的队首元素是否等于 maxSli 队列的队首元素。如果相等，意味着当前队列 sli 中的最大值要被移除，此时需要将 maxSli 队列的队首元素也移除
// 	if c.sli[0] == c.maxSli[0] {
// 		c.maxSli = c.maxSli[1:]
// 	}
//
// 	value := c.sli[0]
// 	c.sli = c.sli[1:]
// 	return value
// }
