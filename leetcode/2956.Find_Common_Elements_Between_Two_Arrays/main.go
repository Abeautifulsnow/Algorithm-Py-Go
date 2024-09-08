package main

func findIntersectionValues1(nums1 []int, nums2 []int) []int {
	n1, n2 := make(map[int]bool, 0), make(map[int]bool, 0)

	for _, v := range nums1 {
		n1[v] = true
	}

	for _, v := range nums2 {
		n2[v] = true
	}

	res := make([]int, 2)
	for _, n := range nums1 {
		if n2[n] {
			res[0]++
		}
	}

	for _, n := range nums2 {
		if n1[n] {
			res[1]++
		}
	}

	return res
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	a, b := 0, 0
	num1Map := make(map[int]int)

	for _, n := range nums1 {
		num1Map[n]++
	}

	for _, n := range nums2 {
		if value, ok := num1Map[n]; ok {
			b++
			a += value
			num1Map[n] = 0 // 防止重复计算
		}
	}

	return []int{a, b}
}
