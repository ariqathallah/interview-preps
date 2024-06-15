package main

func lonelyinteger(a []int32) int32 {
	// Write your code here
	mp := make(map[rune]int)

	for _, n := range a {
		if n > 0 {
			mp[n]++
		}
	}

	for k, v := range mp {
		if v == 1 {
			return k
		}
	}

	return 0
}
