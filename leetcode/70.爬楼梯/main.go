package main

// Recursive way. Use Memo
func climbStairs1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairs2(n int) int {
	prev, cur := 1, 1 // prev:0, cur:1

	for i := 2; i <= n; i++ {
		tmp := cur
		cur = prev + cur
		prev = tmp
	}

	return cur
}
