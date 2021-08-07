package leetcode

func singleNumber(nums []int) int {
	m := map[int]bool{}

	for _, v := range nums {
		_, ok := m[v]
		if ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}

	var r int
	for k := range m {
		r = k
	}

	return r
}
