package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	word := "ABCCED"

	fmt.Println(exist(board, word))

}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				rs := dfs(board, i, j, 0, word)

				if rs {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j, pos int, word string) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' {
		return false
	}

	if word[pos] != board[i][j] {
		return false
	}

	if len(word)-1 == pos {
		return true
	}

	pos++
	tmp := board[i][j]
	board[i][j] = '*'

	dfsResult :=
		dfs(board, i-1, j, pos, word) ||
			dfs(board, i+1, j, pos, word) ||
			dfs(board, i, j+1, pos, word) ||
			dfs(board, i, j-1, pos, word)
	board[i][j] = tmp

	return dfsResult
}
