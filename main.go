package main

import (
)

func main() {
	// Initialize utilities to 0. 
	// Also set initial policy to "R" for non-wall states.
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			utilityTable[r][c] = 0.0
			if !Walls[[2]int{r, c}] {
				policyTable[r][c] = "R"
			} else {
				policyTable[r][c] = "W"
			}
		}
	}

	//Value Iteration
	ValueIteration()
	PlotValueIteration()



	// Reset utilities
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			utilityTable[r][c] = 0.0
			if !Walls[[2]int{r, c}] {
				policyTable[r][c] = "R"
			} else {
				policyTable[r][c] = "W"
			}
		}
	}
	// Clear any old PI history
	policyIterationHistory = nil

	// 3) Policy Iteration
	PolicyIteration()
	PlotPolicyIteration()
}
