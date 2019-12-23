package main

import "fmt"

func max(a,b int) int{
	if a>b {
		return a

	}
	return b
}

func canJump(nums []int) bool {


	dp:=make([]int,0)
	dp = append(dp, 1)
	for i:=1;i<= len(nums);i++  {
		dp = append(dp, max(dp[i-1]-1,nums[i-1]))
		if dp[i]==0 && i!=len(nums) {
			return false
		}
	}
	return true
	
}


func main() {
	fmt.Println(canJump([]int{
		1,
	}))
}
/*
	总结
	1. 这是一道中等DP题，我对dp的定义是，dp[i] 表示走到第i个位置还能跳的最大步数
	2. 第一次错是因为 [0,1]这种情况，他在位置1的时候不能跳了，所以应该是false,但是我没有考虑第1个位置
       我for循环直接考虑第二个位置了 0.0..
	3. 解决方案我是加了一个第零个位置，dp[0]=1

*/