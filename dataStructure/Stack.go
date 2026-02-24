package dataStructure

// 栈，使用切片实现
// 包含方法：IsEmpty Push Pop Peek Len

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (int, bool) { // 当然可以把返回值改为 int 和 error 类型
	if s.IsEmpty() {
		return -1, false
	}

	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	}

	return (*s)[len(*s)-1], true
}

func (s *Stack) Len() int {
	return len(*s)
}
