package main

func matchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	mp := make(map[string]int32)

	var result []int32

	for _, str := range strings {
		mp[str]++
	}

	for _, q := range queries {
		_, exists := mp[q]
		if exists {
			result = append(result, mp[q])
		} else {
			result = append(result, 0)
		}
	}

	return result
}
