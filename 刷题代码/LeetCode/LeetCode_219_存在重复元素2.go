package main

import "fmt"

var m map[int]int

func containsNearbyDuplicate(nums []int, k int) bool {
	m=make(map[int]int)

	if len(nums)==0 {
		return false
	}

	for i:=1;i<= len(nums);i++  {
		if m[nums[i-1]]==0 {
			m[nums[i-1]]=i
		}else{
			if i-m[nums[i-1]]>k {
				m[nums[i-1]]=i
			}else{
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1,2,1},0))
}
/*
	总结
	1. 这是一道简单题
	2. 注意题目是要求两个不同的索引 i 和 j..之前没看清楚把k=0考虑进去，然后返回错误答案true
*/