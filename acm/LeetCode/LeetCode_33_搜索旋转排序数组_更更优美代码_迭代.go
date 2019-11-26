package main

func search(nums []int, target int) int {
	return solve(nums,target)
}

func solve(nums []int,target int) int{

	l,r:=0,len(nums)-1
	for l<=r{
		mid := (l+r)/2
		if target==nums[mid]{
			return mid
		}
		if nums[mid]>=nums[l]{
			if target < nums[mid] && target>=nums[l]{
				r = mid -1
			}else{
				l = mid + 1
			}
		}else{
			if target > nums[mid] && target<=nums[r]{
				l = mid + 1
			}else{
				r = mid -1
			}
		}
	}
	return -1
}

/*
	总结
	1. 已经很优美了 0.0..判断条件异或的话 1.go语言bool不能异或，2. 不便于理解
	2. 注意33这题目，数组内是没有重复元素的，而81是有的 0.0
*/
