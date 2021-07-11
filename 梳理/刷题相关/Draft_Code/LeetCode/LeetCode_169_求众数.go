package main

import "fmt"

//func majorityElement(nums []int) int {
//	m:=make(map[int]int)
//
//	for i:=0;i< len(nums);i++  {
//		m[nums[i]]++
//		if m[nums[i]]> len(nums)/2 {
//			return nums[i]
//		}
//	}
//	return -1
//}

// 摩尔投票法，空间复杂度O(1)
func majorityElement(nums []int) int {

	candidate,time :=-1,0
	for i:=0;i< len(nums);i++  {
		if time==0 {
			candidate = nums[i]
		}
		if nums[i]==candidate {
			time++
		}else{
			time--
		}
	}
	return candidate
}
func main() {
	fmt.Println(majorityElement([]int{
		2,1,1,1,1,2,1,2,2,
	}))
}
/*
	总结
	1. 这是一道很简单的题目,自己用哈希表AC了。空间复杂度是O(n)，
	2. 之后看官方有摩尔投票法，看了之后写出来了，空间复杂度是O(1)，
    3. 两者时间复杂度都是O(n)
*/