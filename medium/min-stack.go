package leetcode

import "container/list"

type MinStack struct {
	list *list.List
	min  int
}

type kv struct {
	val     int
	prevMin int
}

func Constructor() MinStack {
	return MinStack{list.New(), 0}
}

func (this *MinStack) Push(val int) {
	if this.list.Len() == 0 {
		this.min = val
		this.list.PushBack(kv{val, 0})
	} else if val < this.min {
		this.list.PushBack(kv{val, this.min})
		this.min = val
	} else {
		this.list.PushBack(kv{val, this.min})
	}
}

func (this *MinStack) Pop() {
	top := this.list.Back()

	if top.Value.(kv).val == this.min {
		this.min = top.Value.(kv).prevMin
	}

	this.list.Remove(top)
}

func (this *MinStack) Top() int {
	return this.list.Back().Value.(kv).val
}

func (this *MinStack) GetMin() int {
	return this.min
}
