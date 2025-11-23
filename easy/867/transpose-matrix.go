package leetcode

/*
Given a 2D integer array matrix, return the transpose of matrix.
The transpose of a matrix is the matrix flipped over its main diagonal, switching the matrix's row and column indices.
*/
func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		res[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			res[i][j] = matrix[j][i]
		}
	}

	return res
}
