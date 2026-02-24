package main

// 并查集
// 力扣：https://leetcode.cn/problems/properties-graph/description/

// 最后两个案例时间超限
// 时间复杂度为 O(n^2 * m)，其中n是properties数组的长度，m是每个子数组的平均长度。这是因为双重循环遍历properties数组的时间复杂度为 O(n^2)，
// 而每次循环中调用getNum函数的时间复杂度为 O(m)。
// func numberOfComponents(properties [][]int, k int) int {
// 	n := len(properties)
// 	parent := make([]int, n)
// 	for i := 0; i < len(parent); i++ {
// 		parent[i] = i
// 	}
// 	count := n // 连通分量
//
// 	// 查找
// 	var find func(x int) int
// 	find = func(x int) int {
// 		if parent[x] != x {
// 			parent[x] = find(parent[x])
// 		}
// 		return parent[x]
// 	}
//
// 	// 合并
// 	var union func(x, y int)
// 	union = func(x, y int) {
// 		rootX, rootY := find(x), find(y)
//
// 		// 判断x和y的父节点是否相同如果相同就不用合并了（不然count会额外的自减，数据会错误）
// 		if rootX != rootY {
// 			parent[rootY] = rootX
//
// 			count-- // 连通分量--
// 		}
// 	}
//
// 	// 判断“两个数组中共有的不同整数”的数量是否 >= k
// 	var getNum func(x, y []int, k int) bool
// 	getNum = func(x, y []int, k int) bool {
// 		// m1和m2相当于set集合，去除重复的数字
// 		m1, m2 := make(map[int]int), make(map[int]int)
// 		for i := 0; i < len(x); i++ {
// 			m1[x[i]] = 1
// 			m2[y[i]] = 1
// 		}
//
// 		m := make(map[int]int)
// 		for i := range m1 {
// 			m[i]++
// 		}
// 		for i := range m2 {
// 			m[i]++
// 		}
//
// 		ans := 0
// 		for _, v := range m {
// 			if v >= 2 {
// 				ans++
// 			}
// 		}
// 		return ans >= k
// 	}
//
// 	for i := 0; i < len(properties); i++ {
// 		for j := i + 1; j < len(properties); j++ {
// 			// 符合
// 			if getNum(properties[i], properties[j], k) {
// 				// 合并
// 				union(i, j)
// 			}
// 		}
// 	}
// 	return count
// }

// 优化：
// 减少不必要的比较：可以使用哈希表来记录每个整数出现的数组索引，这样就可以直接找出哪些数组可能有共同元素，避免对所有数组对进行比较。
// 优化 getNum 函数：通过更高效的方式统计两个数组中共同不同整数的数量。
// 时间复杂度：优化后的代码时间复杂度为 O(n * m + C)，其中n是properties数组的长度，m是每个子数组的平均长度，C是所有可能的数组对数量。通常情况下，C小于 n^2，因此在大多数情况下，优化后的代码比原代码更高效。
// 空间复杂度：空间复杂度为 O(n * m)，主要用于存储 numToIndices和pairCount哈希表。
func numberOfComponents(properties [][]int, k int) int {
	n := len(properties)
	parent := make([]int, n)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	count := n // 连通分量

	// 查找
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	// 合并
	var union func(x, y int)
	union = func(x, y int) {
		rootX, rootY := find(x), find(y)

		// 判断x和y的父节点是否相同如果相同就不用合并了（不然count会额外的自减，数据会错误）
		if rootX != rootY {
			parent[rootY] = rootX
			count-- // 连通分量--
		}
	}

	// 记录每个整数出现的数组索引
	numToIndices := make(map[int][]int)
	for i, prop := range properties {
		uniqueNums := make(map[int]bool)
		for _, num := range prop {
			if !uniqueNums[num] {
				numToIndices[num] = append(numToIndices[num], i) // 记录：某一个值出现在哪一个子数组中
				uniqueNums[num] = true
			}
		}
	}

	// fmt.Println(numToIndices) // map[1:[0 1] 2:[0 1] 3:[2] 4:[2 3] 5:[3 4] 6:[4] 7:[5]]

	// 统计每个数组对的共同不同整数数量
	pairCount := make(map[[2]int]int)
	for _, indices := range numToIndices {
		for i := 0; i < len(indices); i++ {
			for j := i + 1; j < len(indices); j++ {
				pair := [2]int{indices[i], indices[j]}
				pairCount[pair]++
			}
		}
	}
	// fmt.Println(pairCount) // map[[0 1]:2 [2 3]:1 [3 4]:1]

	// 合并满足条件的数组对
	for pair, commonCount := range pairCount {
		if commonCount >= k {
			union(pair[0], pair[1])
		}
	}

	return count
}

// func main() {
// 	arr := [][]int{{1, 2}, {1, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
// 	fmt.Println(numberOfComponents(arr, 1))
// }

// 或者：
// type uf struct {
// 	fa []int
// 	cc int
// }
//
// func newUnionFind(n int) uf {
// 	fa := make([]int, n)
// 	for i := range fa {
// 		fa[i] = i
// 	}
// 	return uf{fa, n}
// }
//
// func (u uf) find(x int) int {
// 	if u.fa[x] != x {
// 		u.fa[x] = u.find(u.fa[x])
// 	}
// 	return u.fa[x]
// }
//
// func (u *uf) merge(from, to int) {
// 	x, y := u.find(from), u.find(to)
// 	if x == y {
// 		return
// 	}
// 	u.fa[x] = y
// 	u.cc--
// }
//
// // 暴力枚举所有 properties[i] 和 properties[j]，如果交集大小 ≥k，那么用并查集合并 i 和 j。
// // 初始连通块个数 cc=n，每成功合并一次，就把 cc 减一。
// func numberOfComponents(properties [][]int, k int) int {
// 	sets := make([]map[int]bool, len(properties))
// 	for i, a := range properties {
// 		sets[i] = map[int]bool{}
// 		for _, x := range a {
// 			sets[i][x] = true
// 		}
// 	}
//
// 	// fmt.Println(sets) // [map[1:true 2:true] map[1:true 2:true] map[3:true 4:true] map[4:true 5:true] map[5:true 6:true] map[7:true]]
//
// 	u := newUnionFind(len(properties))
// 	for i, a := range sets {
// 		for j, b := range sets[:i] {
//
// 			cnt := 0
// 			for x := range b {
// 				if a[x] {
// 					cnt++
// 				}
// 			}
// 			if cnt >= k {
// 				u.merge(i, j)
// 			}
//
// 		}
// 	}
// 	return u.cc
// }
