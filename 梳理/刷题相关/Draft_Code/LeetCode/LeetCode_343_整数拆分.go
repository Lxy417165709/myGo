package main

import "fmt"

func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
func integerBreak(n int) int {
	dp:=[100]int{}
	dp[0],dp[1] = 0,1

	for i:=2;i<=n ;i++  {
		mx:=0
		for j:=1;j<=i-1 ;j++  {
			temp := max((i-j)*j,dp[i-j]*j)
			mx = max(mx,temp)
		}
		dp[i] = mx
	}
	return dp[n]
}

func main() {
	n:=0
	for ;true ;  {
		fmt.Scan(&n)
		fmt.Println(integerBreak(n))
	}
}
/*
	总结
	1. 这是一道中等的DP题
	2. 模拟之后我居然想得到DP..看来最近的做题有些效果呀~
*/