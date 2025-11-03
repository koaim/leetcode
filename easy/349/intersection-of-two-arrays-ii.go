package leetcode

/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.
*/
func intersection(nums1 []int, nums2 []int) []int {
	numsMap1 := map[int]struct{}{}
	for _, v := range nums1 {
		numsMap1[v] = struct{}{}
	}

	numsMap2 := map[int]struct{}{}
	for _, v := range nums2 {
		numsMap2[v] = struct{}{}
	}

	var res []int
	for k := range numsMap1 {
		_, ok := numsMap2[k]
		if ok {
			res = append(res, k)
		}
	}

	return res
}
