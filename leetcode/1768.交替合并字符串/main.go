package main

func mergeAlternately(word1 string, word2 string) string {
	word1L, word2L := len(word1), len(word2)

	minL := min(word1L, word2L)

	var newS string

	for i := 0; i < minL; i++ {
		newS += string(word1[i])
		newS += string(word2[i])
	}

	return newS + string(word2[minL:]) + string(word1[minL:])
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func mergeAlternately1(word1 string, word2 string) string {
	word1L, word2L := len(word1), len(word2)

	newS := make([]byte, 0, word1L+word2L)

	i, j := 0, 0

	for i < word1L || j < word2L {
		if i < word1L {
			newS = append(newS, word1[i])
			i++
		}

		if j < word2L {
			newS = append(newS, word2[j])
			j++
		}
	}

	return string(newS)
}
