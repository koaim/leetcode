package leetcode

import "container/list"

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

- MinStack() initializes the stack object.
- void push(int val) pushes the element val onto the stack.
- void pop() removes the element on the top of the stack.
- int top() gets the top element of the stack.
- int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.
*/
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
