package unique

type FirstUnique struct {
	queue []int
	dic   map[int]int
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		queue: []int{},
		dic:   map[int]int{},
	}

	for _, v := range nums {
		fu.Add(v)
	}

	return fu

}

func (this *FirstUnique) ShowFirstUnique() int {
	for len(this.queue) > 0 && this.dic[this.queue[0]] > 1 {
		this.queue = this.queue[1:]
	}

	if len(this.queue) == 0 {
		return -1

	} else {
		return this.queue[0]
	}
}

func (this *FirstUnique) Add(value int) {
	if _, ok := this.dic[value]; ok {
		this.dic[value]++

	} else {
		this.dic[value] = 1
		this.queue = append(this.queue, value)
	}
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */
