package main

func findTheDifference(s string, t string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return ch
		}
	}
}

func findTheDifference1(s string, t string) byte {
	diff := 0

	for c := range s {
		diff -= int(s[c])
	}

	for c := range t {
		diff += int(t[c])
	}

	return byte(rune(diff))
}
