package main

func divRoundDown(numerator, denominator int) int {
	quotient := numerator / denominator
	if numerator%denominator != 0 && (numerator < 0) != (denominator < 0) {
		quotient--
	}
	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) <= 1 || indexDiff < 0 || valueDiff < 0 {
		return false
	}

	allBuckets := make(map[int]int, indexDiff+1)

	// Set the size of bucket.
	bucketSize := valueDiff + 1

	for i, v := range nums {
		// 1. Get the id of bucket
		bucketId := divRoundDown(v, bucketSize)

		// 2. If the bucketId is already existed, return true.
		if _, ok := allBuckets[bucketId]; ok {
			return true
		}

		// 3. Compare to previous bucket
		if _, ok := allBuckets[bucketId-1]; ok && abs(allBuckets[bucketId-1]-v) <= valueDiff {
			return true
		}

		// 4. Compare to next bucket
		if _, ok := allBuckets[bucketId+1]; ok && abs(allBuckets[bucketId+1]-v) <= valueDiff {
			return true
		}

		// 5. Put the bucketId into allBuckets with value v.
		allBuckets[bucketId] = v

		// 6. If the current index is great than indexDiff,
		// then pop the i - indexDiff element in nums which match bucketId in allBuckets.
		if i >= indexDiff {
			delete(allBuckets, divRoundDown(nums[i-indexDiff], bucketSize))
		}
	}

	return false
}

func main() {
	res := containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	println(res)
}
