package leetcode

/*
Given two n x n binary matrices mat and target, return true if it is possible to make mat equal to target by rotating mat in 90-degree increments, or false otherwise.
*/
func findRotation(mat [][]int, target [][]int) bool {
	if equal(mat, target) {
		return true
	}

	for i := 0; i < 3; i++ {
		rotate(mat)
		if equal(mat, target) {
			return true
		}
	}

	return false
}

func rotate(matrix [][]int) {
	n := len(matrix[0])

	for i := 0; i < n; i++ {
		//транспонирую матрицу
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

		//разворачиваю каждую строку
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

func equal(mat [][]int, target [][]int) bool {
	for i := 0; i < len(mat[0]); i++ {
		for j := 0; j < len(target[0]); j++ {
			if mat[i][j] != target[i][j] {
				return false
			}
		}
	}

	return true
}
