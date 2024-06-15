package main

import "strings"

func pangrams(s string) string {
	// Write your code here
	str := strings.ToLower(s)
	m := make(map[rune]int)

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			m[c]++
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		if m[i] == 0 {
			return "not pangram"
		}
	}

	return "pangram"
}
