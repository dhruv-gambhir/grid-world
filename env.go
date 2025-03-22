package main

// Environment constants
const (
    numRows            = 6
    numCols            = 6
    discount           = 0.99
    threshold          = 1e-6
    intendedMoveProb   = 0.8
    unintendedMoveProb = 0.1

    rewardWhite = -0.05
    rewardGreen = 1.0
    rewardBrown = -1.0
)

// Actions (Up, Down, Left, Right)
var actions = []struct {
    dr, dc int
    label  string
}{
    {-1, 0, "U"},
    {1, 0,  "D"},
    {0, -1, "L"},
    {0, 1,  "R"},
}

// Global utility table & policy
var utilityTable [numRows][numCols]float64
var policyTable  [numRows][numCols]string
