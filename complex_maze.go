package main
/*
func setupScaledEnvironment(scale int) {
	fmt.Printf("\n\n=== Setting up environment scaled by %d ===\n", scale)

	// 1) set rows, cols
	numRows = 6 * scale
	numCols = 6 * scale

	// 2) Build the big environment
	bigRewards := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		bigRewards[i] = make([]float64, cols)
	}
	bigWalls := make(map[[2]int]bool)

	for tileR := 0; tileR < scale; tileR++ {
		for tileC := 0; tileC < scale; tileC++ {
			for r := 0; r < 6; r++ {
				for c := 0; c < 6; c++ {
					R := tileR*6 + r
					C := tileC*6 + c
					if originalWalls[[2]int{r, c}] {
						// It's a wall
						bigWalls[[2]int{R, C}] = true
						bigRewards[R][C] = 0.0
					} else {
						bigRewards[R][C] = originalRewards[r][c]
					}
				}
			}
		}
	}

	// Assign to globals
	rewards = bigRewards
	walls = bigWalls

	// 3) Initialize utilityTable and policyTable
	utilityTable = make([][]float64, rows)
	policyTable = make([][]string, rows)
	for r := 0; r < rows; r++ {
		utilityTable[r] = make([]float64, cols)
		policyTable[r] = make([]string, cols)
	}

	// 4) For a naive starting policy, let's say everything not a wall is "→"
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !walls[[2]int{r, c}] {
				policyTable[r][c] = "→"
			} else {
				policyTable[r][c] = "[W]"
			}
		}
	}

	// 5) Clear any iteration stats
	viIterationStats = nil
	piIterationStats = nil
}
*/
