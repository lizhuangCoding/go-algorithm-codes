package main

// 栈
// 力扣：https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/description/

type MinStack struct {
	sli    []int
	minSli []int
}

func Constructor() MinStack {
	return MinStack{
		sli:    make([]int, 0),
		minSli: make([]int, 0),
	}
}

func (m *MinStack) Push(x int) {
	if len(m.sli) == 0 {
		m.sli = append(m.sli, x)
		m.minSli = append(m.minSli, x)
		return
	}
	m.sli = append(m.sli, x)

	demo := m.minSli[len(m.minSli)-1]
	if x < demo {
		m.minSli = append(m.minSli, x)
	} else {
		m.minSli = append(m.minSli, demo)
	}

}

func (m *MinStack) Pop() {
	m.sli = m.sli[:len(m.sli)-1]
	m.minSli = m.minSli[:len(m.minSli)-1]
}

func (m *MinStack) Top() int {
	return m.sli[len(m.sli)-1]
}

func (m *MinStack) GetMin() int {
	return m.minSli[len(m.minSli)-1]
}
