package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	return solve(nums)
}



func mx(a,b,c int) int{
	if a>b {
		if a>c {
			return a
		}else{
			return c
		}
	}else{
		if b>c {
			return b
		}else{
			return c
		}
	}
}


// 返回该数组中(连续的子序列的和)  --> 最大
func solve(nums []int) int{
	x:=len(nums)
	max_left,max_right :=0,0
	if len(nums)==1 {
		return nums[0]
	}else{
		max_left = solve(nums[:x/2])
		max_right = solve(nums[x/2:])
	}
	max_l,max_r :=nums[x/2-1],nums[x/2]
	ans:=0
	for i:=x/2 - 1;i>=0 ;i--  {
		ans+=nums[i]
		if max_l<ans {
			max_l = ans
		}
	}
	ans=0
	for i:=x/2 ;i<x ;i++  {
		ans+=nums[i]
		if max_r<ans {
			max_r = ans
		}
	}
	return  mx(max_left,max_right,max_l+max_r)
}

func main() {
	fmt.Println(maxSubArray([]int{5,-5,15,5,-100,-20,200}))
}

/*
	总结
	1. 暴力的话可以利用前缀和，然后2个获取[i,j]之间的值，取最大，复杂度O(n^2)
	2. DP的话从左到右扫,如果dp[i-1] + A[i]<0的话 就放弃，从i+1开始继续扫描,取这些子段的最大值,复杂度(n)
	3. 分治的话可以分为3种情况, 最大子序列在数组左半边,中间,右半边。左右半边直接递归,得max_left,max_right，获取中间的话只需要从中间分别向
       左右两边拓展，得到max_l,max_r, 最终结果就是 max_left,max_right,max_l+max_r中的最大值
	4. 分治的逻辑要符合题意，比如题目说至少要有一个元素，所以如果当nums的长度为1时，即使他是负数，也要返回他，而不是返回0

*/