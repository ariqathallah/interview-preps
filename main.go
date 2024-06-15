package main

import "fmt"

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})

	ms := matchingStrings([]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"})
	fmt.Println(ms)

	tc := timeConversion("07:05:45PM")
	fmt.Println(tc)

	lt := lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1})
	fmt.Println(lt)

	isPangrams := pangrams("We promptly judged antique ivory buckles for the next prize")
	fmt.Println(isPangrams)
}
