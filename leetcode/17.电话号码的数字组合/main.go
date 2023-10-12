package main

import "fmt"

// global declaration
var combinations []string

var phone = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// letterCombinations is entry function.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	combinations = []string{}
	backTrack("", digits, 0)

	return combinations
}

/*
* backTrack
* combination: the all concatenated letters at the moment.
* digits: original digits
* index: the index of digits string which is based from 0.
 */
func backTrack(combination string, digits string, index int) {
	// 当index和原始数字字母表的长度相同时, 不再回溯
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		// 获取对应的数字
		digit := string(digits[index])
		// 获取对应数字的所有字母
		letters := phone[digit]
		// 获取对应数字的所有字母的长度
		lettersCount := len(letters)

		// 取当前数字对应的字母表的第一个字母, 进行回溯
		for i := 0; i < lettersCount; i++ {
			backTrack(combination+string(letters[i]), digits, index+1)
		}
	}
}

func main() {
	fmt.Println(letterCombinations("24"))
}
