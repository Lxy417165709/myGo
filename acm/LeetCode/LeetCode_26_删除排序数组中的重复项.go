package main

import "fmt"

func removeDuplicates(nums []int) int {
	index:=-1
	num:=0
	for i:=0 ;i< len(nums);i++  {
		if nums[i]!=num || index==-1 {
			index++
			nums[index] = nums[i]
			num = nums[index]
		}
	}
	//fmt.Println(nums)
	return index+1
}
func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

/*
	总结
	1. 注意审题，题目说了这是一个排序数组
*/