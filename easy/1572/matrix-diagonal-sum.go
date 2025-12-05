package leetcode

/*
Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.
*/
func diagonalSum(mat [][]int) int {
	left, right := 0, len(mat[0])-1
	sum := 0

	for i := 0; i < len(mat); i++ {
		if left == right {
			sum += mat[i][left]
		} else {
			sum += mat[i][left]
			sum += mat[i][right]
		}

		left++
		right--
	}

	return sum
}
