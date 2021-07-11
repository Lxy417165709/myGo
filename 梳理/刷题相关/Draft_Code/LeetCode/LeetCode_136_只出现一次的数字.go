package main

import "fmt"

func singleNumber(nums []int) int {
	ans:=0
	for i:=0;i< len(nums);i++  {
		ans ^= nums[i]
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{0,0,8,-5,-5,6,6,8,0}))
}
/*
	总结
	1. 该题目考察异或，a^b^a = b
*/