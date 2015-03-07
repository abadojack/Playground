//game of fifteen :
//Player should arrange integers 1 to 15 in ascending order on a 4 * 4 board
package main


import  (
		"fmt"
		"time"
)

var (
	d = 4
	mrow = 0
	scol = 0
	mcol = 0
	srow = 0
	findRow = 0
	findCol = 0
)

var board [4][4]int

func main(){
	var tile int
	//diplay welcome message
	welcome()

	//initialize the board
	initBoard()

	for {
		clearscreen()

		displayBoard()

		if won(){
			fmt.Println("Congratulations, you win")
			break;
		}

		//prompt user to enter tile to move
		fmt.Print("Enter tile to move: ")
		fmt.Scanf("%d", &tile)



		if !move(tile){
			fmt.Println("Illegal move")
			time.Sleep(500 * time.Millisecond)
		}

		time.Sleep(500 * time.Millisecond)
	}
	clearscreen()
}

//Displays 'animated' message to the screen
func welcome(){
	msg := []byte("WELCOME TO THE GAME OF FIFTEEN")

	clearscreen()

	for i := 0; i <= len(msg); i++{
		fmt.Println(string(msg[:i]))
		time.Sleep(150 * time.Millisecond)
		clearscreen()
	}
}

//Clear screen
func clearscreen(){
	fmt.Printf("\033[2J")
	fmt.Printf("\033[%d;%dH",0,0)
}

//Initalize board with tiles numbered 1 to d * d - 1
func initBoard(){
	isOdd := false
	tileCount := d * d - 1
	if tileCount % 2 != 0 {
		isOdd = true
	}

	for row := 0; row < d; row++ {
		for column := 0; column < d; column++ {
			board[row][column] = tileCount
			tileCount--
		}
	}

	if isOdd {
		board[d-1][d-3], board[d-1][d-2] = board[d-1][d-2], board[d-1][d-3]
	}
}

//Display the board in its current state
func displayBoard(){
	for r := 0; r < d; r++ {
		for c := 0; c < d; c++ {
			if board[r][c] < 10{
				fmt.Print(" ")
			}

			if board[r][c] != 0 {
				fmt.Print(board[r][c], " ")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//returns true if tile boarders an empty space else returns false
func move(tile int) bool {
	if findTile(tile) {
		mrow = findRow
		mcol = findCol


		if borderSpace(){
			board[mrow][mcol],  board[srow][scol] = board[srow][scol], board[mrow][mcol]

			return true
		}
	}

	return false
}

//check for empty space around tiles and assign 'em
func borderSpace() bool {
	findTile(0)

	srow = findRow
	scol = findCol

	if scol == mcol {
		if srow == mrow + 1 || srow == mrow - 1 {
			return true
		}
	}

	if srow == mrow {
		if scol == mcol + 1 || scol == mcol - 1{
			return true
		}
	}

	return false
}


/*findTile: find row and column of tile asked by user */
func findTile(tile int) bool{
	for rows := 0; rows < d; rows++{
		for columns := 0; columns < d; columns++{
			if board[rows][columns] == tile{
				findRow = rows
				findCol = columns
				return true
			}
		}
	}

	return false
}

func won() bool {
	tileKey := 1
	count := 1

	for r := 0; r < d; r++ {
		for c := 0; c < d; c++{
			if board[r][c] == tileKey && count == tileKey {
				if tileKey == d * d - 1 {
					return true
				}
				tileKey++
			}
			count++

		}
	}
	return false
}
