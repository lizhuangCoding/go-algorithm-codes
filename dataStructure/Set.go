package dataStructure

// Set，使用map实现
// 包含方法：New Add Contains Size Clear Equal IsSubset

// Exists 空结构体
var Exists = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

// New Key可以为任意类型的数据，用空接口来实现
func New(items ...interface{}) *Set {
	// 获取Set的地址
	s := &Set{}
	// 为m分配空间
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

// Add 简化操作可以添加不定个数的元素进入到Set中，用变长参数的特性来实现这个需求即可，
// 因为Map不允许Key值相同，所以不必有排重操作。同时将Value数值指定为空结构体类型。
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

// Contains 查看是否包含某个元素
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

// Clear 清空：通过重新初始化Set来实现
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

// Equal 相等：判断两个Set是否相等，通过循环遍历，将A中的每一个元素查询在B中是否存在，只要有一个不存在，A和B就不相等。
func (s *Set) Equal(other *Set) bool {
	// 如果两者Size不相等，就不用比较了
	if s.Size() != other.Size() {
		return false
	}

	// 迭代查询遍历
	for key := range s.m {
		// 只要有一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// IsSubset 子集：循环遍历判断A是不是B的子集
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，s一定不是other的子集
	if s.Size() > other.Size() {
		return false
	}
	// 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// // Set，使用map实现
// // 包含方法：Add Remove Contains
//
//
// type Set map[int]bool
//
// func NewSet() Set {
// 	return make(Set)
// }
//
// func (s Set) Add(n int) {
// 	s[n] = true
// }
//
// func (s Set) Remove(n int)  {
// 	delete(s, n)
// }
//
// func (s Set) Contains(n int) bool {
// 	_, ok := s[n]
// 	return ok
// }
