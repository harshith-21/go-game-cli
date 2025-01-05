package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"bufio"
	"strings"
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
	Basearr := createArr(rows, cols)

	// Create a channel to update the 2D array
	UpdateArrChan := make(chan [][]string)

	// Create inpput channel to take input from user, map that to directions, move the snake accoringly
	InputChan := make(chan string)

	// One goroutine to refresh and print the screen
	go refreshAndPrint(UpdateArrChan, refreshRate)
	time.Sleep(1000)
	UpdateArrChan <- Basearr

	go getInputViaChan(InputChan)

	// test
	// testSnakeRun(UpdateArrChan, arr)

	// actual
	snake := [][2]int {
		{1,5},
		{1,4},
		{1,3},
		{1,2},
	}

	arr := deepCopyArr(Basearr)

	for {
		select {
		// take input and move snake
		case newDirection := <-InputChan:
			currentDirection = newDirection

		case <-time.After(refreshRate):
			// Move the snake in the current direction
			snake = UpdateSnake(snake, currentDirection)

			// Draw the snake
			for _, body := range snake {
				arr[body[0]][body[1]] = snakeBody
			}
			arr[snake[0][0]][snake[0][1]] = snakeHead

			UpdateArrChan <- arr
			arr = deepCopyArr(Basearr)
		}
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

func getInputViaChan(InputChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())

		if text == "w" && currentDirection != "down" {
			InputChan <- "up"
		} else if text == "s" && currentDirection != "up" {
			InputChan <- "down"
		} else if text == "a" && currentDirection != "right" {
			InputChan <- "left"
		} else if text == "d" && currentDirection != "left" {
			InputChan <- "right"
		}
	}
	close(InputChan)
}

func UpdateSnake(snake [][2]int, currentDirection string) [][2]int {
	// Get direction offsets
	offset := directions[currentDirection]
	head := snake[0]

	// Calculate new head position
	newHead := [2]int{head[0] + offset[0], head[1] + offset[1]}

	// Add new head to the front of the snake
	newSnake := append([][2]int{newHead}, snake...)

	// Remove the last part (tail) to simulate movement
	newSnake = newSnake[:len(newSnake)-1]

	return newSnake
}

func deepCopyArr(src [][]string) [][]string {
	// Create a new 2D slice with the same dimensions
	dst := make([][]string, len(src))
	for i := range src {
		dst[i] = make([]string, len(src[i]))
		copy(dst[i], src[i]) // Copy each row individually
	}
	return dst
}

func testSnakeRun(UpdateArrChan chan [][]string, arr [][]string){
	
	for i := 1; i<2*len(arr)-1 ;i++ {
		time.Sleep(300*time.Millisecond)
		arr[1][i] = "*"
		UpdateArrChan <- arr
	}

}
