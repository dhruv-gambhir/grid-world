package main

import (
	"fmt"
)

// IsValid checks if (r, c) is in-bounds and not a wall.
func IsValid(r, c int) bool {
	if r < 0 || r >= numRows || c < 0 || c >= numCols {
		return false
	}
	return !Walls[[2]int{r, c}]
}

// PrintUtilities shows the numeric utility values for each cell,
// printing W for Walls.
func PrintUtilities() {
	fmt.Println("Utility Table:")
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if Walls[[2]int{r, c}] {
				fmt.Printf("   W   ")
			} else {
				fmt.Printf("%6.2f ", utilityTable[r][c])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// printPolicy shows the best action symbol in each cell, or [W] for Walls.
func PrintPolicy() {
	fmt.Println("Policy Table:")
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if Walls[[2]int{r, c}] {
				fmt.Printf("W  ")
			} else {
				a := policyTable[r][c]
				if a == "" {
					a = "."
				}
                switch a {
                    case "U": fmt.Printf("↑  ")
                    case "R": fmt.Printf("→  ")
                    case "L": fmt.Printf("←  ")
                    case "D": fmt.Printf("↓  ")
                }
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
 

// averageUtility computes the average utility across all valid cells
func AverageUtility() float64 {
	var sum float64
	var count int
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if !Walls[[2]int{r, c}] {
				sum += utilityTable[r][c]
				count++
			}
		}
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}
