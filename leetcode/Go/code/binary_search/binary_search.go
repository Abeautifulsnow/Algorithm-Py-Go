package main

import "fmt"

func BinarySearch(orderSlice []int, searchV int) int {
	left, right := 0, len(orderSlice)-1
	searchIdx := 0

	for left <= right {
		mid := (left + right) / 2

		if searchV < orderSlice[mid] {
			right = mid - 1
		}

		if searchV > orderSlice[mid] {
			left = mid + 1
		}

		if searchV == orderSlice[mid] {
			searchIdx = mid
			break
		}
	}

	return searchIdx
}

func main() {
	arr := []int{2, 5, 6, 8, 12, 15, 17, 23, 27, 31, 39, 40, 45, 56, 79, 90}
	fmt.Println(BinarySearch(arr, 12))
	fmt.Println(BinarySearch(arr, 45))
	fmt.Println(BinarySearch(arr, 90))
}
