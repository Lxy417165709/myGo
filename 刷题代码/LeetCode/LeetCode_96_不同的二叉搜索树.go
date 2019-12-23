package main
func numTrees(n int) int {
	dp:=make([]int,0)
	dp = append(dp,1)
	for i:=1;i<=n;i++{
		sum:=0
		for t:=0;t<=i-1;t++{
			sum += dp[t]*dp[i-t-1]
		}
		dp = append(dp,sum)
	}
	return dp[n]
}
func main() {

}
/*
	总结
	1. 这题目是一个dp题
	2. 这个dp是一个卡特兰数列 0.0.
*/
