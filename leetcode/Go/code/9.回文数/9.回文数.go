package main

import "fmt"

func IsPalindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedNumber := 0

	// x = 1221
	// 1: x % 10 = 1, reversedNumber = 0 * 10 + 1 = 1, x = 122
	// 2: x % 10 = 2, reversedNumber = 1 * 10 + 2 = 12, x = 12
	// x > reversedNumber is false, jump out the forloop
	for x > reversedNumber {
		reversedNumber = reversedNumber*10 + x%10
		x /= 10
	}

	// when x = 1221 => reversedNumber = 12, x = 12
	// when x = 12321 => reversedNumber = 123, x = 12
	// So if the length of x is an odd, check x == reversedNumber / 10
	// If the length of x is even, check x == reversedNumber
	return x == reversedNumber || x == reversedNumber/10
}

func main() {
	x := 123
	x1 := 12321
	x2 := 1234321
	fmt.Println(x, ":", IsPalindromeNumber(x))
	fmt.Println(x1, ":", IsPalindromeNumber(x1))
	fmt.Println(x2, ":", IsPalindromeNumber(x2))
}
