package leetcode

import (
	"container/heap"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements.
You may return the answer in any order.
*/
func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}

	for _, v := range nums {
		freq[v] = freq[v] + 1
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for k, v := range freq {
		heap.Push(&pq, &Item{value: k, priority: v})
	}

	var res []int
	for len(res) != k {
		v := heap.Pop(&pq).(*Item)
		res = append(res, v.value)
	}

	return res
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want the lowest priority (smallest integer) as the highest priority
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
