package main

func main() {
    // Initialize all utilities to 0
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            utilityTable[r][c] = 0.0
            // Initial policy is move right for each state (except walls)
            if !Walls[[2]int{r, c}] {
                policyTable[r][c] = "R"
            } else {
                policyTable[r][c] = "W"
            }
        }
    }

    
    // Run and plot Value Iteration
    ValueIteration()
    PlotValueIteration()

    // Reset utilities
    for r := 0; r < numRows; r++ {
        for c := 0; c < numCols; c++ {
            utilityTable[r][c] = 0.0
        }
    }

    // Reset history
    policyIterationHistory = nil

    // Run and plot Policy Iteration
    PolicyIteration()
    PlotPolicyIteration()
}

