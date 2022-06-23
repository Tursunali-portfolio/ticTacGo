package main

import (
	"TicTac/utils"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(utils.Green + "Welcome to the Tic Tac Toe game!" + utils.Reset)

	board := [3][3]string{}
	for i := 0; i < len(board); i++ {
		board[i] = [3]string{"-", "-", "-"}
	}

	player := "x"

	for {
		utils.PrintBoard(board)

		fmt.Println(utils.Cyan+"\nYour turn player: "+utils.Reset, utils.Yellow+player+utils.Reset)

		var row, col int

		fmt.Print(utils.Blue + "enter row number: " + utils.Reset)
		_, err := fmt.Scan(&row)
		if err != nil {
			utils.Clear(runtime.GOOS)
			fmt.Println(utils.Red + "Enter correct number!" + utils.Reset)
			continue
		}
		fmt.Print(utils.Green + "enter column number: " + utils.Reset)
		_, err = fmt.Scan(&col)
		if err != nil {
			utils.Clear(runtime.GOOS)
			fmt.Println(utils.Red + "Enter correct number!" + utils.Reset)
		}

		if !utils.IsValid(board, row, col) {
			utils.Clear(runtime.GOOS)
			fmt.Println(utils.Red + "Enter correct row and column number!" + utils.Reset)
			continue
		}

		board[row][col] = player

		if utils.Check(board, player) {
			utils.Clear(runtime.GOOS)
			utils.PrintBoard(board)
			fmt.Printf(utils.Yellow+"Congratulations player %s has won the game!\n"+utils.Reset, player)
			break
		}

		if utils.Check2(board) {
			fmt.Println(utils.Cyan + "DRAW play again!" + utils.Reset)
			break
		}

		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}
		utils.Clear(runtime.GOOS)
	}
}
