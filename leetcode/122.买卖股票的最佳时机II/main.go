package main

/*
Refer: https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/solutions/476997/shou-hua-tu-jie-xiang-jie-duo-chong-jie-fa-mai-mai/?envType=study-plan-v2&envId=top-interview-150
*/

func maxProfit(prices []int) int {
	lenP := len(prices)
	i := 1
	profit := 0

	for i < lenP {
		// 递减
		for i < lenP && prices[i-1] >= prices[i] {
			i++
		}
		// 到达局部波谷
		trough := prices[i-1]

		// 递增
		for i < lenP && prices[i-1] <= prices[i] {
			i++
		}
		// 到达局部波峰
		peak := prices[i-1]
		profit += peak - trough
	}

	return profit
}

func main() {
	prices := []int{3, 1, 2, 0, 4, 3}
	println(maxProfit(prices))
}
