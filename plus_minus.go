package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	positive := 0
	negative := 0
	zero := 0
	size := len(arr)

	for _, v := range arr {
		if v > 0 {
			positive += 1
		} else if v < 0 {
			negative += 1
		} else {
			zero += 1
		}
	}

	fmt.Printf("%.6f\n", float32(positive)/float32(size))
	fmt.Printf("%.6f\n", float32(negative)/float32(size))
	fmt.Printf("%.6f\n", float32(zero)/float32(size))
}
