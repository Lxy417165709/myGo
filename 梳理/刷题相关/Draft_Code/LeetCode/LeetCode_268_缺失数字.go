package main

import "fmt"

func missingNumber(nums []int) int {
	n:= len(nums)
	sum:=n*(n+1) /2
	sum1:=0
	for i:=0;i<n ;i++  {
		sum1+=nums[i]
	}
	return sum-sum1
}

func main() {
	fmt.Println(missingNumber([]int{
		9,6,4,2,3,5,7,0,1,
	}))
}
/*
	总结
	1. 该题我用等差数列求和解决了
	2. 官方有用位运算解决的 0.0.。
*/