package leetcode

/*
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.
*/

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
