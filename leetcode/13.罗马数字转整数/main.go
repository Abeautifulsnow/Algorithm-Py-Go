package main

var mapping = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var max, sum int

	for i := len(s) - 1; i >= 0; i-- {
		preNum := mapping[s[i]]

		if max <= preNum {
			sum += preNum
			max = preNum
		} else {
			sum -= preNum
		}

	}

	return sum
}
