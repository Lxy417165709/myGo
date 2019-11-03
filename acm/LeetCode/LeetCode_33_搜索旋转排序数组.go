package main

import "fmt"



func search(nums []int, target int) int {

	return solve(nums,0, len(nums)-1,target)

}

func solve(nums []int, l int,r int,target int) int{
	if l>r {
		return -1
	}
	mid:=(l+r)/2
	if nums[mid]==target{
		return mid
	}else{
		if nums[mid]==nums[l] {
			return solve(nums,mid+1,r,target)
		}

		if nums[mid]>nums[l] {
			if nums[mid]>target {
				if target>=nums[l] {	// 这个等号很重要
					return solve(nums,l,mid-1,target)
				}else{
					return solve(nums,mid+1,r,target)
				}
			}else{
				return solve(nums,mid+1,r,target)
			}
		}else{
			if target>nums[mid] {
				if target<nums[l] {
					return solve(nums,mid+1,r,target)
				}else{
					return solve(nums,l,mid-1,target)
				}
			}else{
				return solve(nums,l,mid-1,target)
			}
		}
	}

}




func main() {
	fmt.Println(search([]int{60,70,85,1,3,4},70))
}


/*
	总结
	1. 自己写的时候if else太多了，其实可以用更简洁的代码表达(分3种情况，对mid,target,l全排列就只有3种，再添加小于号，小于等于号)
	2. 看了解答才有思路的~
*/