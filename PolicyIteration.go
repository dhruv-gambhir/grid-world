package main

import (
    "fmt"
    "math"
)

// storing history for plotting
var policyIterationHistory []([numRows][numCols]float64)

// PolicyIteration function
func PolicyIteration() {
    fmt.Println("=== POLICY ITERATION ===")

    stable := false
    for !stable {
        // First we run Policy Evaluation
        for {
            var newUtility [numRows][numCols]float64
            var delta float64

            for r := 0; r < numRows; r++ {
                for c := 0; c < numCols; c++ {
                    if Walls[[2]int{r, c}] {
                        continue
                    }

                    //chosen action from policy iteration for each cell, intiall all move-right
                    chosen := policyTable[r][c]
                    if chosen == "" || chosen == "W" {
                        // empty means no update is necessary
                        newUtility[r][c] = utilityTable[r][c]
                        continue
                    }

                    // finding the deltas for the chosen action
                    var chosenDr, chosenDc int
                    for _, a := range actions {
                        if a.label == chosen {
                            chosenDr, chosenDc = a.dr, a.dc
                            break
                        }
                    }

                    //find probability distribution
                    val := 0.0
                    totalProb := 0.0
                    for _, actual := range actions {
                        prob := unintendedMoveProb
                        if actual.dr == chosenDr && actual.dc == chosenDc {
                            prob = intendedMoveProb
                        }

                        nr, nc := r+actual.dr, c+actual.dc
                        if !IsValid(nr, nc) {
                            nr, nc = r, c
                        }
                        val += prob * (Rewards[nr][nc] + discount*utilityTable[nr][nc])
                        totalProb += prob
                    }
                    if totalProb > 0 {
                        val /= totalProb
                    }

                    newUtility[r][c] = val

                    dLocal := math.Abs(val - utilityTable[r][c])
                    if dLocal > delta {
                        delta = dLocal
                    }
                }
            }

            utilityTable = newUtility
            policyIterationHistory = append(policyIterationHistory, utilityTable)

            if delta < threshold {
                break
            }
        }

        // Next we run Policy Improvement
        stable = true
        for r := 0; r < numRows; r++ {
            for c := 0; c < numCols; c++ {
                if Walls[[2]int{r, c}] {
                    continue
                }

                oldAction := policyTable[r][c]

                bestAction := oldAction
                // best utility value of each cell
                bestValue  := math.Inf(-1)

                // finfing best action
                for _, candidate := range actions {
                    val := 0.0
                    totalProb := 0.0

                    for _, actual := range actions {
                        prob := unintendedMoveProb
                        if actual.dr == candidate.dr && actual.dc == candidate.dc {
                            prob = intendedMoveProb
                        }
                        nr, nc := r+actual.dr, c+actual.dc
                        if !IsValid(nr, nc) {
                            nr, nc = r, c
                        }
                        val += prob * (Rewards[nr][nc] + discount*utilityTable[nr][nc])
                        totalProb += prob
                    }
                    if totalProb > 0 {
                        val /= totalProb
                    }

                    if val > bestValue {
                        bestValue  = val
                        bestAction = candidate.label
                    }
                }

                if bestAction != oldAction {
                    policyTable[r][c] = bestAction
                    stable = false
                }
            }
        }

        // save for plotting
        policyIterationHistory = append(policyIterationHistory, utilityTable)
    }

    fmt.Println("Policy Iteration converged.")
    PrintPolicy()
}

