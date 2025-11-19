package main

import (
	"fmt"
	"time"
)

var GameScreen = [3][3]byte{}

func main() {

	Grettings()
	step := 1

	for {

		PlayerCall(step)
		row, column := TakeInput()
		
		if step % 2 != 0 {
			GameScreen[row][column] = 'X'
		}else {
			GameScreen[row][column] = 'O'
		}

		PrintGameBoard()
		GameOver := GameStutus()

		if GameOver {
			time.Sleep(500*time.Millisecond)
			fmt.Println("Congratulation!")
			fmt.Println("You Win...!")
			fmt.Println()
			break
		}
		if step == 9 {
			fmt.Println("\nNo possible turn left...")
			fmt.Println("Game Over!")
			fmt.Println("It is a draw!")
			break
		}

		step++
	}

	fmt.Println("\nThank you for playing...")
	fmt.Println()
}

func Grettings() {
	fmt.Println("\nWelcome to the tic-tac-toe game!")
	fmt.Println("It is a two player game.")
	fmt.Println("You win if you hold any  sequential position (3) before the other party")
	fmt.Println()
}

func PrintGameBoard(){
	time.Sleep(800*time.Millisecond)

	fmt.Println("\n - - -")
	for i:=0;i<3;i++{
		fmt.Print("|")
		for j:=0;j<3;j++{
			if GameScreen[i][j] >= 'A' && GameScreen[i][j] <= 'Z'{
				fmt.Printf("%c|", GameScreen[i][j])
			}else{
				fmt.Printf(" |")
			}
		}
		fmt.Println("\n - - -")
	}
	fmt.Println()
}

func GameStutus() bool {

	result := 0
	if GameScreen[1][1] != 0 && GameScreen[0][0] == GameScreen[1][1] && GameScreen[1][1] ==  GameScreen[2][2] {
		result = 1
	}else if GameScreen[1][1] != 0 && GameScreen[0][2] == GameScreen[1][1] && GameScreen[1][1] == GameScreen[2][0] {
		result = 1
	}else if GameScreen[0][0] != 0 && GameScreen[0][0] == GameScreen[0][1] && GameScreen[0][1] == GameScreen[0][2] {
		result = 1
	}else if GameScreen[1][1] != 0 && GameScreen[1][0] == GameScreen[1][1] && GameScreen[1][1] == GameScreen[1][2] {
		result = 1
	}else if GameScreen[2][0] != 0 && GameScreen[2][0] == GameScreen[2][1] && GameScreen[2][1] == GameScreen[2][2] {
		result = 1
	}else if GameScreen[0][0] != 0 && GameScreen[0][0] == GameScreen[1][0] && GameScreen[1][0] == GameScreen[2][0] {
		result = 1
	}else if GameScreen[1][1] != 0 && GameScreen[0][1] == GameScreen[1][1] && GameScreen[1][1] == GameScreen[2][1] {
		result = 1
	}else if GameScreen[0][2] != 0 && GameScreen[0][2] == GameScreen[1][2] && GameScreen[1][2] == GameScreen[2][2] {
		result = 1
	}

	return result == 1
}

func PlayerCall(step int){
	time.Sleep(800*time.Millisecond)

	if step % 2 != 0 {
		fmt.Println("It is the 1st player's turn.\nSymbol : X ")
	} else {
		fmt.Println("It is the 2nd player's turn.\nSymbol : O ")
	}
	fmt.Println()

	time.Sleep(500*time.Millisecond)
}

func TakeInput() (int, int){

	var row, column int
	for {

		fmt.Println("Enter the position: ")
		fmt.Scan(&row)
		fmt.Scan(&column)
		
		row--
		column--
		if InputValidation(row, column) {
			break
		}
		fmt.Println("\nInvalid Position...")
		fmt.Println("Please try again...")
		fmt.Println()
	}
	
	return row, column
}

func InputValidation (row, column int) bool {

	return GameScreen[row][column] == 0 

}