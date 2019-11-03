package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	for i:=len(digits)-2;i>=0 ;i--  {

		digits[i] += digits[i+1]/10
		digits[i+1] = digits[i+1]%10

	}
	if digits[0]==10 {
		digits[0] = digits[0]%10
		digits = append([]int{1},digits...)
	}
	return digits



}

func main() {
	fmt.Println(plusOne([]int{9}))
}

/*
	总结
	1. 该题考查加法进位~
*/