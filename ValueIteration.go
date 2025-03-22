package main

import (
    "fmt"
    "math"
)

// We'll store a slice of utility snapshots, one per iteration
var valueIterationHistory []([numRows][numCols]float64)

// ValueIteration runs until convergence, storing each iteration’s utilities.
func ValueIteration() {
    fmt.Println("=== VALUE ITERATION ===")

    for {
        var newUtility [numRows][numCols]float64
        var delta float64

        for r := 0; r < numRows; r++ {
            for c := 0; c < numCols; c++ {
                if Walls[[2]int{r, c}] {
                    continue
                }

                // Find best action
                bestVal := math.Inf(-1)
                for _, intended := range actions {
                    actionVal := 0.0
                    totProb   := 0.0
                    // distribution over actual moves
                    for _, actual := range actions {
                        prob := unintendedMoveProb
                        if actual.dr == intended.dr && actual.dc == intended.dc {
                            prob = intendedMoveProb
                        }
                        nr, nc := r+actual.dr, c+actual.dc
                        if !IsValid(nr, nc) {
                            nr, nc = r, c
                        }

                        actionVal += prob * (Rewards[nr][nc] + discount*utilityTable[nr][nc])
                        totProb   += prob
                    }
                    if totProb > 0 {
                        actionVal /= totProb
                    }

                    if actionVal > bestVal {
                        bestVal = actionVal
                    }
                }

                newUtility[r][c] = bestVal
                diff := math.Abs(bestVal - utilityTable[r][c])
                if diff > delta {
                    delta = diff
                }
            }
        }

        utilityTable = newUtility
        // append a snapshot of the entire table
        valueIterationHistory = append(valueIterationHistory, utilityTable)

        if delta < threshold {
            break
        }
    }

    fmt.Println("Value Iteration converged.")
    PrintUtilities()
}

