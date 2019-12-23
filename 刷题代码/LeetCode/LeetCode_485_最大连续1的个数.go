package main

import "fmt"

func max(a,b int) int{
	if a>b {
		return a
	}else{
		return b
	}
}
func findMaxConsecutiveOnes(nums []int) int {
	dp:=make([]int,0)
	if len(nums)==0 {
		return 0
	}
	dp = append(dp, nums[0])
	mx:=dp[0]
	for i:=1;i< len(nums);i++  {
		if nums[i]==1 {
			dp = append(dp, dp[i-1]+1)
		}else{
			dp = append(dp, 0)
		}
		mx = max(dp[i],mx)
	}
	fmt.Println(dp)
	return mx
}
func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{
		0,
	}))
}
/*
	总结
	1. 这是一道简单题，我使用DP的思维做出来了 0.0..
*/