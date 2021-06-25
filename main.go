package main

import (
	"fmt"
	"tic-tac-toe/Service"
)

var board [3][3]string
var player1, player2 string

func main() {
	fmt.Print("Enter First Player name : ")
	fmt.Scanln(&player1)
	fmt.Print("Enter First Player name : ")
	fmt.Scanln(&player2)
	Service.InitializeBoard(&board)
	Service.PrintBoard(board)
	Service.StartGame(player1, player2, board)
}
