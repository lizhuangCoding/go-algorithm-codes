package main

// 链表
// 力扣：https://leetcode.cn/problems/linked-list-components/description/

func numComponents(head *ListNode, nums []int) (ans int) {
	set := make(map[int]int, len(nums))
	for _, v := range nums {
		set[v] = 1
	}

	inSet := false // 前一个节点的状态，是否在集合中
	for ; head != nil; head = head.Next {
		// 如果不在集合中，就重置一下状态
		if _, ok := set[head.Val]; !ok {
			inSet = false
		} else if !inSet { // 如果在集合中，并且前一个节点不在集合中，就++，并且改变前一个节点的状态，避免重复++
			inSet = true
			ans++
		}
	}
	return
}
