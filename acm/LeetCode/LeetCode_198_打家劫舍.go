package main

import "fmt"

func max(a,b int) int{
	if a>b {
		return a
	}else{
		return b
	}
}
func rob(nums []int) int {


	if len(nums)==0 {
		return 0
	}
	if len(nums)==1 {
		return nums[0]
	}
	if len(nums)==2 {
		ans:=max(nums[0],nums[1])
		return ans
	}

	dp:=make([]int,0)
	dp = append(dp, 0)
	dp = append(dp, nums[0])
	for i:=1;i< len(nums);i++  {
		mx:=0
		for t:=i-1;t>=0 ;t--  {
			mx=max(mx,dp[t])
		}
		dp = append(dp, mx + nums[i])
	}
	mx:=0
	fmt.Println(dp)
	for i:=0;i< len(dp);i++  {
		mx = max(dp[i],mx)
	}
	return mx
}
func main() {
	fmt.Println(rob([]int{
		2,
	}))
}
/*
	总结
	1. 第一次的时候,这个动态规划我写复杂了，因为这个动态规划我对dp的定义是,偷第i家可获得的最大钱数,时间复杂度是O(n^2)
	2. 之后看了官方，他的定义是，因为对于第i家有偷或不偷，所以dp[i]可以表示从[0,i]家中可以偷的的最大数，这样的话时间复杂度就是O(n)
*/