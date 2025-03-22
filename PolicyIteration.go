package main

import (
    "fmt"
    "math"
)

// We'll store a slice of utility snapshots for Policy Iteration as well
var policyIterationHistory []([numRows][numCols]float64)

// PolicyIteration runs evaluation/improvement until stable, storing each iteration’s utilities.
func PolicyIteration() {
    fmt.Println("=== POLICY ITERATION ===")

    stable := false
    for !stable {
        // 1) Policy Evaluation
        for {
            var newUtility [numRows][numCols]float64
            var delta float64

            for r := 0; r < numRows; r++ {
                for c := 0; c < numCols; c++ {
                    if Walls[[2]int{r, c}] {
                        continue
                    }

                    // The chosen action is whatever policy says
                    chosen := policyTable[r][c]
                    if chosen == "" || chosen == "W" {
                        // no policy set => do nothing special
                        newUtility[r][c] = utilityTable[r][c]
                        continue
                    }

                    // find the deltas for the chosen action
                    var chosenDr, chosenDc int
                    for _, a := range actions {
                        if a.label == chosen {
                            chosenDr, chosenDc = a.dr, a.dc
                            break
                        }
                    }

                    val := 0.0
                    totProb := 0.0
                    // distribution over actual moves
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
                        totProb += prob
                    }
                    if totProb > 0 {
                        val /= totProb
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

        // 2) Policy Improvement
        stable = true
        for r := 0; r < numRows; r++ {
            for c := 0; c < numCols; c++ {
                if Walls[[2]int{r, c}] {
                    continue
                }

                oldAction := policyTable[r][c]

                bestAction := oldAction
                bestValue  := math.Inf(-1)

                // Evaluate all possible actions
                for _, candidate := range actions {
                    val := 0.0
                    totProb := 0.0

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
                        totProb += prob
                    }
                    if totProb > 0 {
                        val /= totProb
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

        // Another snapshot after improvement
        policyIterationHistory = append(policyIterationHistory, utilityTable)
    }

    fmt.Println("Policy Iteration converged.")
    PrintPolicy()
}

