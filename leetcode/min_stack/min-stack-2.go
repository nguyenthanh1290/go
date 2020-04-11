package stack

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)

	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)

	} else {
		this.minStack = append(this.minStack, min(x, this.minStack[len(this.minStack)-1]))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	val := this.stack[len(this.stack)-1]
	return val
}

func (this *MinStack) GetMin() int {
	min := this.minStack[len(this.minStack)-1]
	return min
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
