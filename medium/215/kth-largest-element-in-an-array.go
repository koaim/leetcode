package leetcode

import "container/heap"

func findKthLargest(nums []int, k int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for _, v := range nums {
		heap.Push(&pq, &Item{value: v, priority: v})
	}

	res := 0
	for pq.Len() != len(nums)-k {
		res = heap.Pop(&pq).(*Item).value
	}

	return res
}

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
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
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
