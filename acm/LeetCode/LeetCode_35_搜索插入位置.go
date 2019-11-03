package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l:=0
	r:= len(nums)-1

	for ;l<=r ;  {
		var mid int = (l+r)/2

		if nums[mid] == target {
			return mid

		}else {
			if nums[mid]>target {
				r = r - 1
			}else{
				l = l + 1
			}
		}
	}
	return l
}


func main() {
	fmt.Println(searchInsert([]int{1,3,5,6},8))
}

/*
	总结
	1. 该题目考察二分查找
*/