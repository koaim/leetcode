package leetcode

func setZeroes(matrix [][]int) {
	var columns []int
	var rows []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				columns = append(columns, j)
				rows = append(rows, i)
			}
		}
	}

	for _, v := range columns {
		setZeroColumn(matrix, v)
	}
	for _, v := range rows {
		setZeroRow(matrix, v)
	}
}

func setZeroRow(matrix [][]int, row int) {
	matrix[row] = make([]int, len(matrix[0]))
}

func setZeroColumn(matrix [][]int, column int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == column {
				matrix[i][j] = 0
			}
		}
	}
}
