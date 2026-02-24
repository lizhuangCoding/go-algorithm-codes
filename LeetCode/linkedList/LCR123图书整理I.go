package main

// 从尾到头遍历链表
// 力扣：https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/

func reverseBookList(head *ListNode) []int {
	sli := make([]int, 0)
	dfs(head, &sli)
	return sli
}

func dfs(head *ListNode, sli *[]int) {
	if head == nil {
		return
	}

	dfs(head.Next, sli)

	// 要传递指针，因为有append扩容，扩容之后指向的底层数组就会变化
	*sli = append(*sli, head.Val)
}
