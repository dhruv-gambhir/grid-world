package main

import (
	"fmt"
	"math"
)


// Value iteration function
func ValueIteration() {
	fmt.Println("Value Iteration")

	for {
		var newUtility [numRows][numCols]float64
		var delta float64

		for r := 0; r < numRows; r++ {
			for c := 0; c < numCols; c++ {
				if Walls[[2]int{r, c}] {
					continue
				}

				// Finding max utility value
				bestVal := math.Inf(-1)
				for _, intended := range actions {
					actionVal := 0.0
					totProb := 0.0
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
						totProb += prob
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

		// update utility table
		utilityTable = newUtility

		// for plotting
		valueIterationHistory = append(valueIterationHistory, utilityTable)

		// check if threshold is reached
		if delta < threshold {
			break
		}
	}

	fmt.Println("Value Iteration converged.")
	PrintUtilities()
    policyVI := GetVIPolicy()
    PrintValuePolicy(policyVI)
}


func GetVIPolicy() [numRows][numCols]string {
    var policyVI [numRows][numCols]string

    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            if Walls[[2]int{r, c}] {
                policyVI[r][c] = "[W]"
                continue
            }

            bestAction := "."
            bestVal    := math.Inf(-1)

            // Evaluate Q(s,a) for each action
            for _, intended := range actions {
                var actionValue float64
                var totalProb float64

                for _, actual := range actions {
                    prob := unintendedMoveProb
                    if actual.dr == intended.dr && actual.dc == intended.dc {
                        prob = intendedMoveProb
                    }
                    nr := r + actual.dr
                    nc := c + actual.dc
                    if !IsValid(nr, nc) {
                        nr, nc = r, c
                    }
                    actionValue += prob * (Rewards[nr][nc] + discount*utilityTable[nr][nc])
                    totalProb   += prob
                }
                if totalProb > 0 {
                    actionValue /= totalProb
                }
                if actionValue > bestVal {
                    bestVal    = actionValue
                    bestAction = intended.label
                }
            }
            policyVI[r][c] = bestAction
        }
    }
    return policyVI
}
