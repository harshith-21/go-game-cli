package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)



func main() {

	args := os.Args

	var length int

	UpdateArrChan := make(chan [][]string)

	if len(args) > 1 {
		length, _ = strconv.Atoi(args[1])
	} else {
		length = 30
	}

	// Create a 2D slice of strings of that length
	arr := make([][]string, length)

	// Create each inner slice (each row) with 'length' columns
	for i := range arr {
		arr[i] = make([]string, 2*length)
	}

	// Create border with "%"
	for i, s := range arr {
		for j, _ := range s {
			if i == 0 || i == length-1 || j == 0 || j == 2*length-1 {
				arr[i][j] = "%"
			} else {
				arr[i][j] = " "
			}
		}
	}

	// One goroutine to refresh and print the screen

	go refreshAndPrint(UpdateArrChan, 300)
	time.Sleep(1000)
	UpdateArrChan <- arr

	// MAIN EVENT LOOP (i think thats what this is called), Simulate dynamic updates to the array
	for {
		// Send the updated array to the channel
		UpdateArrChan <- arr
	}

}

func printarr(arr [][]string) {
	for i, s := range arr {
		for j, _ := range s {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}


// refreshAndPrint clears the screen and prints the array to simulate refreshing.
func refreshAndPrint(UpdateArrChan <- chan [][]string, refreshRateMS int) {

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
		time.Sleep(time.Duration(refreshRateMS) * time.Millisecond)
	}
}

