package main

func minCostClimbingStairs(cost []int) int {
	lenCost := len(cost)
	dp := make([]int, lenCost+1)

	for i := 2; i <= lenCost; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[lenCost]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
