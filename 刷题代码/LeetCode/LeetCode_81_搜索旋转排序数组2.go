package main

func search(nums []int, target int) bool {
	return solve(nums,target)
}

func solve(nums []int,target int) bool{
	if len(nums)==0{
		return false
	}
	l,r:=0,len(nums)-1
	mid := (l+r)/2

	if target==nums[mid]{
		return true
	}

	// 因为数组有重复元素，当相等时，我们无法判断递增递减规则，所以要左右扫一次
	if nums[mid]==nums[l]{
		return solve(nums[:mid],target) || solve(nums[mid+1:],target)
	}

	if nums[mid]>nums[l]{
		if target < nums[mid] && target>=nums[l]{
			return solve(nums[:mid],target)
		}else{
			return solve(nums[mid+1:],target)
		}
	}else{
		if target > nums[mid] && target<=nums[r]{
			return solve(nums[mid+1:],target)
		}else{
			return solve(nums[:mid],target)
		}
	}
}
/*
	总结
	1. 这题和 33 差不多，多了的一步已经写在上面了 0.0
*/