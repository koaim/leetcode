package leetcode

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := [][]int{[]int{}}

	var currRow int
	var currLength int
	for _, v := range original {
		if currLength == n {
			res = append(res, []int{v})
			currLength = 1
			currRow++
		} else {
			res[currRow] = append(res[currRow], v)
			currLength++
		}
	}

	return res
}
