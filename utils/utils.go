package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Clear clear the output
func Clear(platform string) {
	if platform == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	} else if platform == "linux" || platform == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	} else {
		fmt.Println("Platform not supported!")
		return
	}
}

// PrintBoard print the resulting game board
func PrintBoard(board [3][3]string) {
	fmt.Println(Purple + "   0 1 2 " + Reset)
	for index, val := range board {
		fmt.Println(Purple+strconv.Itoa(index)+Reset, val)
	}
}

// Check if the game is won
func Check(board [3][3]string, player string) bool {
	for idx, val := range board {
		if val == [3]string{player, player, player} {
			return true
		} else if board[0][idx] == player && board[1][idx] == player && board[2][idx] == player {
			return true
		}
	}

	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	} else if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

// Check2 if the game is finished
func Check2(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}

// IsValid Check if the inputted row and column are valid
func IsValid(board [3][3]string, row int, col int) bool {
	if row > 2 || row < 0 || col > 2 || col < 0 {
		return false
	}
	if board[row][col] != "-" {
		return false
	}
	return true
}
