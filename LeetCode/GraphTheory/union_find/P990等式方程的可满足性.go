package main

// 并查集，判断连通性问题
// 力扣：https://leetcode.cn/problems/satisfiability-of-equality-equations/description/

func equationsPossible(equations []string) bool {
	parent := make([]int, 'z'-'a'+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

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
		parent[rootY] = rootX
	}

	// 判断x和y的父节点是否相同（判断这两个节点是否连通）
	var isSame func(x, y int) bool
	isSame = func(x, y int) bool {
		return find(x) == find(y)
	}

	// sli1存储==，sli2存储!=
	sli1, sli2 := make([]int, 0), make([]int, 0)
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '=' {
			sli1 = append(sli1, int(equations[i][0]-'a'))
			sli1 = append(sli1, int(equations[i][3]-'a'))
		} else {
			sli2 = append(sli2, int(equations[i][0]-'a'))
			sli2 = append(sli2, int(equations[i][3]-'a'))
		}
	}

	// 开始合并
	for i := 0; i < len(sli1); i += 2 {
		union(sli1[i], sli1[i+1])
	}

	// 开始判断是否存在连通性问题，如果已经连通了，那么就是冲突的，就返回false
	for i := 0; i < len(sli2); i += 2 {
		if isSame(sli2[i], sli2[i+1]) { // 如果已经连通了，就返回false
			return false
		}
	}

	return true
}

// 或者：
// func equationsPossible(equations []string) bool {
//    parent := make([]int, 26)
//    for i := 0; i < 26; i++ {
//        parent[i] = i
//    }
//    for _, str := range equations {
//        if str[1] == '=' {
//            index1 := int(str[0] - 'a')
//            index2 := int(str[3] - 'a')
//            union(parent, index1, index2)
//        }
//    }
//
//    for _, str := range equations {
//        if str[1] == '!' {
//            index1 := int(str[0] - 'a')
//            index2 := int(str[3] - 'a')
//            if find(parent, index1) == find(parent, index2) {
//                return false
//            }
//        }
//    }
//    return true
// }
//
// func union(parent []int, index1, index2 int) {
//    parent[find(parent, index1)] = find(parent, index2)
// }
//
// func find(parent []int, index int) int {
//    for parent[index] != index {
//        parent[index] = parent[parent[index]]
//        index = parent[index]
//    }
//    return index
// }
