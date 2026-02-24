package main

import "sort"

// 并查集（迷迷糊糊的）
// 力扣：https://leetcode.cn/problems/accounts-merge/description/

// 时间复杂度：O(n * logn)，其中n是所有邮箱的数量。主要时间开销在于并查集的查找和合并操作以及邮箱列表的排序操作。
// 空间复杂度：O(n)，主要空间开销在于存储邮箱到索引的映射、邮箱到姓名的映射、并查集的父节点数组以及根节点到邮箱列表的映射。

func accountsMerge(accounts [][]string) (ans [][]string) {
	// 存储邮箱到索引的映射，用于后续并查集操作
	emailToIndex := map[string]int{}
	// 存储邮箱到账户所有者姓名的映射
	emailToName := map[string]string{}

	for _, account := range accounts {
		// 获取账户所有者的姓名
		name := account[0]

		// 遍历该账户下的所有邮箱
		for _, email := range account[1:] {
			// 如果该邮箱还没有对应的索引
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex) // 为该邮箱分配一个唯一的索引
				emailToName[email] = name               // 记录该邮箱对应的账户所有者姓名
			}
		}
	}

	// 初始化并查集的父节点数组，长度为邮箱的数量
	parent := make([]int, len(emailToIndex))
	for i := range parent {
		parent[i] = i
	}

	// 定义并查集的查找函数，用于查找元素所在集合的根节点
	var find func(int) int
	find = func(x int) int {
		// 如果当前节点不是根节点，则递归查找其父节点的根节点，并进行路径压缩
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// 定义并查集的合并函数，用于将两个元素所在的集合合并
	union := func(from, to int) {
		// 找到两个元素所在集合的根节点，并将一个根节点的父节点设置为另一个根节点
		parent[find(from)] = find(to)
	}

	// 合并邮箱
	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]] // 获取该账户第一个邮箱的索引

		// 遍历该账户除第一个邮箱外的其他邮箱
		for _, email := range account[2:] {
			// 将这些邮箱所在的集合合并到第一个邮箱所在的集合中
			union(emailToIndex[email], firstIndex)
		}
	}

	// 存储根节点索引到邮箱列表的映射
	indexToEmails := map[int][]string{}
	// 遍历所有邮箱及其对应的索引
	for email, index := range emailToIndex {
		// 找到该邮箱所在集合的根节点
		index = find(index)
		// 将该邮箱添加到根节点对应的邮箱列表中
		indexToEmails[index] = append(indexToEmails[index], email)
	}

	// 遍历所有根节点对应的邮箱列表
	for _, emails := range indexToEmails {
		// 对邮箱列表进行字典序排序
		sort.Strings(emails)
		// 构造合并后的账户，第一个元素是账户所有者姓名，后续元素是排序后的邮箱列表
		account := append([]string{emailToName[emails[0]]}, emails...)
		// 将合并后的账户添加到结果列表中
		ans = append(ans, account)
	}
	return
}
