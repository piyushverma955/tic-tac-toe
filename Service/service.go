package Service

import (
	"fmt"
	"tic-tac-toe/Models"
)

func InitializeBoard(board *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}

}

func PrintBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t", board[i][j])
		}
		fmt.Print("\n")
	}
}

func StartGame(player1 string, player2 string, board [3][3]string) {
	firstPlayerTurn := true
	var coordinates Models.Move
	currentPlayer := player1

	for true {
		whoseTurn(firstPlayerTurn, player1, player2, &currentPlayer)
		takeCoodinates(&coordinates)

		if !isValidMove(board, coordinates) {
			fmt.Println("Not a valid move. try again")
			PrintBoard(board)
		} else {

			if firstPlayerTurn {
				board[coordinates.Row][coordinates.Column] = "X"
			} else {
				board[coordinates.Row][coordinates.Column] = "O"
			}
			PrintBoard(board)
			if isWon(board) {
				fmt.Println(currentPlayer + " won")
				break
			} else {
				if isDraw(board) {
					fmt.Println("Its Draw")
					break
				}
				firstPlayerTurn = !firstPlayerTurn
			}
		}

	}

}

func takeCoodinates(coordinates *Models.Move) {
	var row, column int
	fmt.Println("Enter the row")
	fmt.Scanln(&row)
	fmt.Println("Enter the column")
	fmt.Scanln(&column)
	*coordinates = Models.Move{Row: row - 1, Column: column - 1}
}

func whoseTurn(firstPlayerTurn bool, player1 string, player2 string, curentPlayer *string) {
	if firstPlayerTurn {
		fmt.Println(player1 + " turn")
		*curentPlayer = player1
	} else {
		fmt.Println(player2 + " turn")
		*curentPlayer = player2
	}
}

func rowConnected(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[0][i] != "-" {
			return true
		}
	}
	return false
}

func columnConnected(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != "-" {
			return true
		}
	}
	return false
}

func diagonalConnected(board [3][3]string) bool {

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != "-" {
		return true
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[0][2] != "-" {
		return true
	}

	return false
}

func isWon(board [3][3]string) bool {

	return rowConnected(board) || columnConnected(board) || diagonalConnected(board)

}

func isValidMove(board [3][3]string, move Models.Move) bool {
	if move.Row >= 0 && move.Row < 3 && move.Column >= 0 && move.Column < 3 && board[move.Row][move.Column] == "-" {
		return true
	}

	return false
}

func isDraw(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}
