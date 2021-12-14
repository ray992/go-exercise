package main

import "fmt"

func validTicTacToe(board []string) bool {
	var xCount int
	var oCount int
	for _, val := range board {
		for _, c := range val {
			if c == 'X' {
				xCount++
			}else if c == 'O' {
				oCount++
			}
		}
	}
	if oCount != xCount && oCount != xCount -1 {
		return false
	}
	if win(board, 'X') && oCount != xCount - 1 {
		return false
	}
	if win(board, 'O') && oCount != xCount {
		return false
	}
	return true
}

func win(board []string, p byte) bool  {
	for i := 0; i < 3; i++ {
		if p == board[0][i] && p == board[1][i] && p == board[2][i] {
			return true
		}
		if p == board[i][0] && p == board[i][1] && p == board[i][2]  {
			return true
		}
	}
	if p == board[0][0] && p == board[1][1] && p == board[2][2]  {
		return true
	}
	if p == board[0][2] && p == board[1][1] && p == board[2][0]  {
		return true
	}
	return false
}

func main() {
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "}))
	fmt.Println(validTicTacToe([]string{"XOX", " X ", "   "}))
}
