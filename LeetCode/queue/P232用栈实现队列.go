package queue

// 栈和队列
// 力扣：https://leetcode.cn/problems/implement-queue-using-stacks/description/

// 官方：
type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

// type Stack struct {
// 	sli []int
// 	len int
// }
//
// func NewStack() Stack {
// 	return Stack{
// 		sli: make([]int, 0),
// 		len: 0,
// 	}
// }
//
// func (s *Stack) push(n int) bool {
// 	s.sli = append(s.sli, n)
// 	s.len++
// 	return true
// }
//
// func (s *Stack) peek() (int, bool) {
// 	if s.len == 0 {
// 		return 0, false
// 	}
// 	return s.sli[s.len-1], true
// }
//
// func (s *Stack) pop() (int, bool) {
// 	if s.len == 0 {
// 		return 0, false
// 	}
// 	demo := s.sli[s.len-1]
// 	s.sli = s.sli[:s.len-1]
// 	s.len--
// 	return demo, true
// }
//
// func (s *Stack) size() int {
// 	return s.len
// }
//
// func (s *Stack) isEmpty() bool {
// 	return s.len == 0
// }
//
// type MyQueue struct {
// 	s1 Stack
// 	s2 Stack
// }
//
// func Constructor() MyQueue {
// 	return MyQueue{
// 		s1: NewStack(),
// 		s2: NewStack(),
// 	}
// }
//
// func (m *MyQueue) Push(x int) {
// 	m.s1.push(x)
// }
//
// func (m *MyQueue) Pop() int {
// 	if m.s2.isEmpty() {
// 		m.MoveS1ToS2()
// 	}
// 	fmt.Println(m.s1, m.s2)
//
// 	demo, _ := m.s2.pop()
// 	return demo
// }
//
// func (m *MyQueue) Peek() int {
// 	// if m.s2.isEmpty() {
// 	// 	for !m.s1.isEmpty() {
// 	// 		demo, _ := m.s1.pop()
// 	// 		m.s2.push(demo)
// 	// 	}
// 	// }
// 	if m.s2.isEmpty() {
// 		m.MoveS1ToS2()
// 	}
//
// 	demo, _ := m.s2.peek()
// 	return demo
// }
//
// func (m *MyQueue) Empty() bool {
// 	return m.s1.isEmpty() && m.s2.isEmpty()
// }
//
// func (m *MyQueue) MoveS1ToS2() {
// 	// if m.s2.isEmpty() {
// 	for !m.s1.isEmpty() {
// 		demo, _ := m.s1.pop()
// 		m.s2.push(demo)
// 	}
// 	// }
// }
//
// func main() {
// 	var s Stack = NewStack()
// 	fmt.Println(s.peek())
// 	fmt.Println(s.push(111))
// 	fmt.Println(s.push(222))
// 	fmt.Println(s.push(333))
// 	fmt.Println(s.peek())
// 	fmt.Println(s.pop())
// 	fmt.Println(s.size())
// 	fmt.Println(s.isEmpty())
// 	fmt.Println(s.pop())
//
// 	m := Constructor()
// 	m.Push(1)
// 	m.Push(2)
// 	m.Push(3)
// 	m.Push(4)
// 	m.Pop()
// 	m.Push(5)
// 	m.Pop()
// 	m.Pop()
// 	m.Pop()
// 	m.Pop()
// }
