package main

var Rewards = [numRows][numCols]float64{
	{rewardGreen, 0.0, rewardGreen, rewardWhite, rewardWhite, rewardGreen},
	{rewardWhite, rewardBrown, rewardWhite, rewardGreen, 0.0, rewardBrown},
	{rewardWhite, rewardWhite, rewardBrown, rewardWhite, rewardGreen, rewardWhite},
	{rewardWhite, rewardWhite, rewardWhite, rewardBrown, rewardWhite, rewardGreen},
	{rewardWhite, 0.0, 0.0, 0.0, rewardBrown, rewardWhite},
	{rewardWhite, rewardWhite, rewardWhite, rewardWhite, rewardWhite, rewardWhite},
}

// Walls are tracked in a map so we can easily test if a cell is a wall.
var Walls = map[[2]int]bool{
	{0, 1}: true,
	{1, 4}: true,
	{4, 1}: true,
	{4, 2}: true,
	{4, 3}: true,
}
