package main

import "strconv"

func timeConversion(s string) string {
	// Write your code here
	aop := string(s[len(s)-2])
	hour := s[0:2]

	if aop == "A" {
		if hour == "12" {
			hour = "00"
		}
	}
	if aop == "P" {
		if hour != "12" {
			hourInt, _ := strconv.Atoi(hour)
			hourInt += 12
			hour = strconv.Itoa(hourInt)
		}
	}

	return hour + string(s[2:len(s)-2])
}
