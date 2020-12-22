/*
No.1
题目：两数之和
难度：简单

给定一个整数数组``nums``和一个目标值``target``，请你在该数组中找出和为目标值的那``两个``整数，
并返回他们的数组下标。你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中
同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
package main

import "fmt"

func twoSumM1(nums []int, target int) []int {
	lenNums := len(nums)
	for i, v := range nums {
		for j := i + 1; j < lenNums; j++ {
			if nums[j]+v == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumM2(nums []int, target int) []int {
	lenNums := len(nums)
	storeMap := make(map[int]int, lenNums)

	for i, v := range nums {
		anotherNum := target - v
		if j, ok := storeMap[anotherNum]; ok {
			return []int{j, i}
		} else {
			storeMap[v] = i
		}
	}

	return []int{}
}

func main() {
	nums := []int{4, 4, 4, 4}
	target := 8
	result1 := twoSumM1(nums, target)
	result2 := twoSumM2(nums, target)
	fmt.Println("Result1:", result1)
	fmt.Println("Result2:", result2)
}
