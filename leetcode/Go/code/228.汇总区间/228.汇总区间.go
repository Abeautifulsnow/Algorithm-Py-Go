package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	l, r := 0, 1
	res := make([]string, 0)

	for l < len(nums) {
		for r < len(nums) && nums[r] == nums[r-1]+1 {
			r += 1
		}

		if l == r-1 {
			res = append(res, strconv.Itoa(nums[l]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[l], nums[r-1]))
		}

		l = r
		r += 1
	}

	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}
