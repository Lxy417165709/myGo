package main

import "fmt"

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

// 法一 (错误)
//func lengthOfLIS(nums []int) int {
//	if len(nums)==0 {
//		return 0
//	}
//	mx:=0
//	for i:=0;i< len(nums);i++  {
//		mx = max(mx,solve(nums,i,1,[]int{nums[i]}))
//	}
//	return mx
//}
//func solve(nums []int,begin int,ans int,slice []int) int{
//
//	index:=-1
//	base:=nums[begin]
//	value:=100000000
//
//	for t:=begin;t< len(nums);t++  {
//		if nums[t]>base && nums[t]<value {
//			value = nums[t]
//			index = t
//		}
//	}
//	if index==-1 {
//		fmt.Println(slice)
//		return ans
//	}else{
//		slice = append(slice, value)
//		return solve(nums,index,ans+1,slice)
//	}
//
//}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

var dp [1005000]int

func lengthOfLIS(nums []int) int {

	if len(nums)==0 {
		return 0
	}
	dp = [1005000]int{}
	dp[0] = 1
	for i:=1;i< len(nums);i++  {
		mx := 0
		for t:=i-1;t>=0 ;t--  {
			if nums[t]<nums[i] {
				mx = max(dp[t],mx)
			}
		}
		dp[i] = mx + 1
	}
	ans:=0
	for i:=0;i<len(nums) ;i++  {
		ans = max(dp[i],ans)
	}
	return ans
}




func main() {
	fmt.Println(lengthOfLIS([]int{0,5,1,2,3}))
}
/*
	总结
	1. 第一次的时候，我递归向数组后面寻找最接近A[i]的，但是这有错误
	   比如 1,3,6,7,9,4,10,5,6 答案是6，如果是按第一种，答案是5，
       因为[6,7,9]被错过了
	2. 第二种方法使用的是dp,dp数组的含义是：以A[i]结尾的最长上升子序列的长度，没用二分查找优化,时间复杂度O(n)
	3. 官方有第三种nlogn的方法 0.0...
*/