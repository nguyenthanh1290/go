// package stack

// type elem struct {
// 	val        int
// 	minElemVal int
// }

// type MinStack struct {
// 	stack []elem
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	return MinStack{
// 		stack: make([]elem, 0),
// 	}
// }

// func (s *MinStack) Push(x int) {
// 	e := elem{val: x}

// 	if len(s.stack) == 0 {
// 		e.minElemVal = x

// 	} else if x < s.GetMin() {
// 		e.minElemVal = x

// 	} else {
// 		e.minElemVal = s.GetMin()
// 	}

// 	s.stack = append(s.stack, e)
// }

// func (s *MinStack) Pop() {
// 	if len(s.stack) > 0 {
// 		s.stack = s.stack[:len(s.stack)-1]
// 	}
// }

// func (s *MinStack) Top() int {
// 	if len(s.stack) == 0 {
// 		return 0
// 	}

// 	return s.stack[len(s.stack)-1].val

// }

// func (s *MinStack) GetMin() int {
// 	if len(s.stack) == 0 {
// 		return 0
// 	}

// 	return s.stack[len(s.stack)-1].minElemVal
// }
