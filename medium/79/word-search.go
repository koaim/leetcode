package leetcode

/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. 
The same letter cell may not be used more than once.
*/
func exist(board [][]byte, word string) bool {
	wordRunes := []byte(word)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == wordRunes[0] {

				res := dfs(board, i, j, wordRunes, 0)
				if res {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j int, word []byte, k int) bool {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 || board[i][j] == '-' {
		return false
	}

	if word[k] != board[i][j] {
		return false
	}

	if len(word)-1 == k {
		return true
	}

	k++
	val := board[i][j]
	board[i][j] = '-'
	res := dfs(board, i+1, j, word, k) || dfs(board, i, j+1, word, k) || dfs(board, i-1, j, word, k) || dfs(board, i, j-1, word, k)
	board[i][j] = val

	return res
}
