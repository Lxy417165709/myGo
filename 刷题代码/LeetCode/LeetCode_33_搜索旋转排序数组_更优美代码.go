package main

func search(nums []int, target int) int {
	return solve(nums,target)
}

func solve(nums []int,target int) int{
	if len(nums)==0{
		return -1
	}
	l,r:=0,len(nums)-1
	mid := (l+r)/2

	if target==nums[mid]{
		return mid
	}
	if nums[mid]>=nums[0]{
		if target < nums[mid] && target>=nums[0]{
			return solve(nums[:mid],target)
		}else{
			ans:=solve(nums[mid+1:],target)+mid+1
			if ans==mid{
				ans = -1
			}
			return ans
		}
	}else{
		if target > nums[mid] && target<=nums[len(nums)-1]{
			ans:=solve(nums[mid+1:],target)+mid+1
			if ans==mid{
				ans = -1
			}
			return ans
		}else{
			return solve(nums[:mid],target)
		}
	}
}
/*
	总结
	1. 这代码还可以更优美，比如使用迭代

*/