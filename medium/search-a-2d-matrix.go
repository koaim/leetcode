package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	return search(matrix, 0, len(matrix)*len(matrix[0]), target) != -1
}

func search(matrix [][]int, from, to, target int) int {
	if from == to {
		return -1
	}

	n := from + (to-from)/2
	i := n / len(matrix[0])
	j := n % len(matrix[0])

	v := matrix[i][j]
	if v == target {
		return n
	} else if v < target {
		return search(matrix, n+1, to, target)
	} else {
		return search(matrix, from, n, target)
	}
}
