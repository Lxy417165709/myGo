package main

import "sort"

var mp map[int]int

func combinationSum4(nums []int, target int) int {
	mp = make(map[int]int)
	sort.Ints(nums)
	return solve(nums,target)
}

func solve(nums []int,target int) int{

	if x,ok:=mp[target];ok{
		return x
	}
	if len(nums)==0{
		return 0
	}

	if target==0{
		return 1
	}

	end:=search(nums,target)
	ans:=0
	for i:=0;i<end;i++{
		ans+=solve(nums,target-nums[i])
	}
	mp[target]=ans
	return mp[target]
}

// 二分查找，找第一个大于target的数的索引
func search(nums []int,target int) int{
	l,r:=0,len(nums)-1
	for l<=r{
		mid:=(l+r)/2
		if nums[mid]==target{
			l = mid + 1
		}else{
			if nums[mid]>target{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}
	}
	return l
}

/*
	总结
	1. 这题目也是回溯，第一提交的时候没有注意到还可以使用记忆化map，于是TLE,然后加了map后就AC了
	2. 这题目总体思路就是: 我先把数组排序后，然后找第一个比target大的数的索引end，之后再遍历一次 (这是剪枝)，回溯下去就可以AC了 (记得记忆化)。

*/