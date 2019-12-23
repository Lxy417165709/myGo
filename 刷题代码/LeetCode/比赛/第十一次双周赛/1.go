package main

import "fmt"

func missingNumber(arr []int) int {
	sum:=0
	for i:=0;i< len(arr);i++  {
		sum+=arr[i]
	}
	return (arr[0]+arr[len(arr)-1])*(len(arr)+1)/2 - sum
}
func main() {
	fmt.Println(missingNumber([]int{
		5,7,11,13,
	}))
}
/*
	总结
	1. 这题考查等差数列求和公式 0.0
*/