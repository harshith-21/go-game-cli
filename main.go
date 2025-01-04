package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)



func main() {

	var wg sync.WaitGroup

	args := os.Args

	var length int

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
	wg.Add(1)
	go refreshAndPrint(arr, 300, &wg)

	// Wait for the goroutine to finish (which won't happen unless we add termination logic)
	wg.Wait()

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
func refreshAndPrint(arr [][]string, refreshRateMS int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure this is called when the goroutine finishes

	for {
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