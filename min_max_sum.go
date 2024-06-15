package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var max int32 = arr[0]
	var min int32 = arr[0]
	var sum int64 = 0

	for i := 0; i < 5; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
		sum += int64(arr[i])
	}

	var maxSum int64 = sum - int64(min)
	var minSum int64 = sum - int64(max)

	fmt.Printf("%d %d", minSum, maxSum)
}
