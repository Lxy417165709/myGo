package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m:=make(map[int]int)
	for i:=0;i<len(nums) ;i++  {
		m[nums[i]]++
		if m[nums[i]]==2 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(containsDuplicate([]int{
		1,2,3,4,5,6,8,9,0,-1,
	}))
}
/*
	总结
	1. 这题我做法是时空复杂度分别是O(n),O(n)
	2. 排序的话时空复杂度分别是O(nlogn),O(1)
*/