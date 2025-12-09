package leetcode

import "math/rand"

/*
Implement the RandomizedSet class:
    - RandomizedSet() Initializes the RandomizedSet object.
    - bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
    - bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
    - int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average O(1) time complexity.
*/

type RandomizedSet struct {
	set map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: map[int]int{},
		arr: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.set[val] = len(this.arr) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	v, ok := this.set[val]
	if !ok {
		return false
	}

	delete(this.set, val)

	last := this.arr[len(this.arr)-1]
	if v != len(this.arr)-1 {
		this.arr[v] = last
		this.set[last] = v
	}
	this.arr = this.arr[:len(this.arr)-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
