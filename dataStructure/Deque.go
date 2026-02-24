package dataStructure

import "errors"

// deque双端队列，使用list实现，泛型
// 包含方法：PushFront、PushBack、PopFront、PopBack、Front、Back、Len、IsEmpty

type Deque struct {
	data []int
}

func (d *Deque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
}

func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("队列为空，无法删除")
	}

	result := d.data[0]
	d.data = d.data[1:]

	return result, nil
}

func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("队列为空，无法删除")
	}

	result := d.data[d.Len()-1]
	d.data = d.data[0 : d.Len()-1]

	return result, nil
}

func (d *Deque) Front() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("队列为空")
	}

	return d.data[0], nil
}

func (d *Deque) Back() (int, error) {
	if d.IsEmpty() {
		return -1, errors.New("队列为空")
	}

	return d.data[d.Len()-1], nil
}

func (d *Deque) Len() int {
	return len(d.data)
}

func (d *Deque) IsEmpty() bool {
	return d.Len() == 0
}
