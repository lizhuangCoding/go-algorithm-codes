package main

// 图，dfs，bfs
// 力扣：https://leetcode.cn/problems/clone-graph/description/

// type Node struct {
// 	Val       int
// 	Neighbors []*Node
// }

// 错误的写法：
/*

1 -- 2
|    |
4 -- 3

只复制了最表层的节点 1 2 4 ，3没有被复制
*/

// func cloneGraph(node *Node) *Node {
// 	if node == nil {
// 		return nil
// 	}
//
// 	m := make(map[*Node]*Node)
// 	m[node] = &Node{
// 		Val:       node.Val,
// 		Neighbors: make([]*Node, 0),
// 	}
//
// 	for i := 0; i < len(node.Neighbors); i++ {
// 		m[node.Neighbors[i]] = &Node{
// 			Val:       node.Neighbors[i].Val,
// 			Neighbors: make([]*Node, 0),
// 		}
// 	}
//
// 	for k, v := range m {
// 		for i := 0; i < len(k.Neighbors); i++ {
// 			v.Neighbors = append(v.Neighbors, m[k.Neighbors[i]])
// 		}
// 	}
//
// 	return m[node]
// }

// bfs
// func cloneGraph(node *Node) *Node {
// 	if node == nil {
// 		return nil
// 	}
//
// 	// 创建克隆的起始节点
// 	clonedNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
// 	visited := make(map[*Node]*Node)
// 	visited[node] = clonedNode
//
// 	// 使用队列进行BFS
// 	queue := []*Node{node}
//
// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]
//
// 		// 遍历当前节点的所有邻居
// 		for _, neighbor := range current.Neighbors {
// 			// 如果邻居节点还没有被克隆
// 			if _, exists := visited[neighbor]; !exists {
// 				// 创建邻居的克隆
// 				clonedNeighbor := &Node{
// 					Val:       neighbor.Val,
// 					Neighbors: make([]*Node, 0),
// 				}
// 				visited[neighbor] = clonedNeighbor
// 				queue = append(queue, neighbor)
// 			}
//
// 			// 将克隆的邻居添加到当前克隆节点的邻居列表中
// 			visited[current].Neighbors = append(visited[current].Neighbors, visited[neighbor])
// 		}
// 	}
//
// 	return clonedNode
// }
