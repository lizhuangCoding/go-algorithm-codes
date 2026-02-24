package main

import (
	"container/heap"
	"math"
)

// 广度优先搜索+优先队列：Dijkstra 最短路
// 力扣：https://leetcode.cn/problems/find-minimum-time-to-reach-last-room-ii/description/

// 方向数组 dirs。上、下、左、右
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 时间复杂度：O(nmlog(nm))，其中 n 和 m 分别为 moveTime 的行数和列数。
// 空间复杂度：O(nm)。
func minTimeToReach(moveTime [][]int) (ans int) {
	n, m := len(moveTime), len(moveTime[0])
	// dis 数组记录了从起点到每个位置的最短时间。初始化时，所有位置的值设为最大整数（表示未访问），除了起点 (0, 0)，其值设为 0。
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0

	// 优先队列 h（最小堆：快速获取堆中的最小元素）
	h := hp{{}}
	for {
		// 使用 heap.Pop 从最小堆中取出当前距离最小的元素。
		// top 是一个 tuple 类型的结构体，包含 dis（距离）、x 和 y（坐标）。
		top := heap.Pop(&h).(tuple)
		i, j := top.x, top.y

		// 如果到达目标位置，返回到达该位置的最短时间top.dis
		if i == n-1 && j == m-1 {
			return top.dis
		}
		// 检查当前从优先队列中取出的状态是否为当前位置的最短路径。如果当前距离 top.dis 大于 dis[i][j]，就跳过这个状态，因为我们已经找到了一条更短路径。
		if top.dis > dis[i][j] {
			continue
		}

		// 更新相邻房间的距离。遍历四个相邻位置 (x, y)，并检查它们是否在网格范围内。
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < m {
				// 计算新的时间 newD
				// 如果当前时间 top.dis 大于 moveTime[x][y]，那么我们无需等待。否则，我们需要等待到 moveTime[x][y] 才能进入房间。
				newD := 0
				if top.dis > moveTime[x][y] {
					newD = top.dis + (i+j)%2 + 1
				} else {
					newD = moveTime[x][y] + (i+j)%2 + 1
				}

				// 更新 dis 数组并加入优先队列。
				// 如果 newD 小于当前记录的到达时间 dis[x][y]，则更新 dis[x][y]，并将新状态 {newD, x, y} 压入堆中。
				if newD < dis[x][y] {
					dis[x][y] = newD
					heap.Push(&h, tuple{newD, x, y})
				}
			}
		}
	}
}

// 定义了 tuple 结构体，包含 dis（当前到达时间）、x 和 y（坐标）。
type tuple struct {
	dis, x, y int
}

// hp 是一个实现了堆接口的 tuple 切片，用来作为优先队列。
type hp []tuple

// Len 返回堆的长度。
func (h hp) Len() int { return len(h) }

// Less 比较两个元素的 dis 字段，构建最小堆。Less方法指定了比较的规则，即优先选择距离较小的元素（最小堆）
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }

// Swap 交换堆中的两个元素。
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 将元素 v 推入堆。
func (h *hp) Push(v any) { *h = append(*h, v.(tuple)) }

// Pop 将堆顶元素弹出，并返回它。
func (h *hp) Pop() (v any) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
