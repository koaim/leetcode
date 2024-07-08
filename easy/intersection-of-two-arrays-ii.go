package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, len(nums1))
	m2 := make(map[int]int, len(nums2))

	for _, v := range nums1 {
		m1[v] += 1
	}

	for _, v := range nums2 {
		m2[v] += 1
	}

	res := make([]int, 0, len(nums1)+len(nums2))
	for k, v := range m1 {
		v1, ok := m2[k]
		if ok {
			r := min(v, v1)
			for i := 0; i < r; i++ {
				res = append(res, k)
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
