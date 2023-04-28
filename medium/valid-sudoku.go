package leetcode

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
*/
func isValidSudoku(board [][]byte) bool {
	lines := [9]map[byte]int{}
	columns := [9]map[byte]int{}
	boxes := map[int]map[byte]int{}

	for i := 0; i < 9; i++ {
		lines[i] = map[byte]int{}

		for j := 0; j < 9; j++ {
			column := columns[j]
			if column == nil {
				columns[j] = map[byte]int{}
			}

			value := board[i][j]

			if string(value) == "." {
				continue
			} else {
				if val := lines[i][value]; val == 0 {
					lines[i][value] = 1
				} else {
					return false
				}

				if val := columns[j][value]; val == 0 {
					columns[j][value] = 1
				} else {
					return false
				}

				box := calcBox(i, j)
				if _, ok := boxes[box]; !ok {
					boxes[box] = map[byte]int{}
				}

				if val := boxes[box][value]; val == 0 {
					boxes[box][value] = 1
				} else {
					return false
				}
			}
		}
	}

	return true
}

func calcBox(i, j int) int {
	if i < 3 {
		return 3 * (j / 3)
	} else if i < 6 {
		if j < 3 {
			return 1
		} else if j < 6 {
			return 4
		} else {
			return 7
		}
	} else {
		if j < 3 {
			return 2
		} else if j < 6 {
			return 5
		} else {
			return 8
		}
	}
}
