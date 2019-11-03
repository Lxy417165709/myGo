package main

import "fmt"

func searchRange(nums []int, target int) []int {

	//return []int{minIndex(nums,target),maxIndex(nums,target)}
	return []int{index(nums,target,0),index(nums,target,1)}
}

// 这个函数可以减少代码量
func index(nums []int, target int,flag int) int{
	l:=0
	r:=len(nums)-1
	mid := 0

	place1 :=-1
	for ; l<=r;   {
		mid = (l+r)/2
		if nums[mid] == target {
			place1 = mid
			if flag==1 {
				l = mid+1
			}else{
				r = mid-1
			}
		}
		if nums[mid] >target {
			r = mid-1
		}
		if nums[mid] <target {
			l = mid+1
		}
	}
	return place1
}


func maxIndex(nums []int, target int) int{
	l:=0
	r:=len(nums)-1
	mid := 0

	place1 :=-1
	for ; l<=r;   {
		mid = (l+r)/2
		if nums[mid] == target {
			place1 = mid
			l = mid+1
		}
		if nums[mid] >target {
			r = mid-1
		}
		if nums[mid] <target {
			l = mid+1
		}
	}
	return place1
}

func minIndex(nums []int, target int) int{
	l:=0
	r:=len(nums)-1
	mid := 0

	place1 :=-1
	for ; l<=r;   {
		mid = (l+r)/2
		if nums[mid] == target {
			place1 = mid
			r = mid-1	// 和max只有这修改了
		}
		if nums[mid] >target {
			r = mid-1
		}
		if nums[mid] <target {
			l = mid+1
		}
	}
	return place1
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10},6))
}

/*
	总结
	1. 这题考查的是二分查找
*/