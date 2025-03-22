package main

func main() {
    // Initialize all utilities to 0
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            utilityTable[r][c] = 0.0
            // Start with a naive policy (e.g. always "R" if not a wall)
            if !Walls[[2]int{r, c}] {
                policyTable[r][c] = "R"
            } else {
                policyTable[r][c] = "W"
            }
        }
    }

    // 1) Run Value Iteration
    ValueIteration()

    // Plot each cell’s utility across iterations for Value Iteration
    PlotValueIterationPerState()

    // 2) Reset utilities for Policy Iteration
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            utilityTable[r][c] = 0.0
        }
    }

    // Clear history if you want a fresh start
    policyIterationHistory = nil

    // 3) Run Policy Iteration
    PolicyIteration()

    // Plot each cell’s utility across iterations for Policy Iteration
    PlotPolicyIterationPerState()
}

