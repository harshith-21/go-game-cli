package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	rows       = 30
	cols       = 60
	emptyCell  = " "
	wallCell   = "#"
	snakeHead  = "@"
	snakeBody  = "*"
	foodCell   = "*"
	refreshRate = 200 * time.Millisecond 
)

var (
	directions = map[string][2]int{
		"up":    {-1, 0},
		"down":  {1, 0},
		"left":  {0, -1},
		"right": {0, 1},
	}
	currentDirection = "right"
)


func main() {
	// Create a 2D slice of strings of that length
	arr := createArr(rows, cols)

	// Create a channel to update the 2D array
	UpdateArrChan := make(chan [][]string)

	// One goroutine to refresh and print the screen
	go refreshAndPrint(UpdateArrChan, refreshRate)
	time.Sleep(1000)
	UpdateArrChan <- arr

	// test
	// testSnakeRun(UpdateArrChan, arr)

	// actual
	snake := [][2]int {
		{1,5},
		{1,4},
		{1,3},
		{1,2},
	}

	for {
		// take input and move snake

		// update the snake in updateArrchan to print
		for _, body := range snake {
			arr[body[0]][body[1]] = snakeBody
		}
		arr[snake[0][0]][snake[0][1]] = snakeHead
		UpdateArrChan <- arr
	}

}

func createArr(rows, cols int) [][]string {
	arr := make([][]string, rows)
	for i := range arr {
		arr[i] = make([]string, cols)
	}

	// Create border with "%"
	for i, s := range arr {
		for j, _ := range s {
			if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
				arr[i][j] = "#"
			} else {
				arr[i][j] = " "
			}
		}
	}

	return arr
}

// refreshAndPrint clears the screen and prints the array to simulate refreshing.
func refreshAndPrint(UpdateArrChan <- chan [][]string, refreshRateMS time.Duration) {

	for arr := range UpdateArrChan {
		// clears the screen
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		// Print the 2D array
		for _, row := range arr {
			for _, val := range row {
				fmt.Print(val)
			}
			fmt.Println() // Move to the next line after each row
		}

		// Sleep for the refresh rate duration
		time.Sleep(refreshRateMS)
	}
}

func testSnakeRun(UpdateArrChan chan [][]string, arr [][]string){
	
	for i := 1; i<2*len(arr)-1 ;i++ {
		time.Sleep(300*time.Millisecond)
		arr[1][i] = "*"
		UpdateArrChan <- arr
	}

}
