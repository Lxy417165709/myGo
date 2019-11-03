package main

import "fmt"
const(
	mod = 1000000007
)
func numTilings(N int) int {
	dp:=make([]int,0)
	dp = append(dp,1)
	dp = append(dp,1)
	dp = append(dp,2)
	for i:=3;i<=N ;i++  {
		dp = append(dp,0)

		for t:=1;t<=2 ;t++  {
			dp[i] = (dp[i] +  dp[i-t])%mod
		}
		for t:=3;t<=i ;t=t+1  {
			dp[i] = (dp[i] +  2*dp[i-t])%mod
		}

	}
	//fmt.Println(dp)
	return dp[N]
}


func main() {
	fmt.Println(numTilings(4))
}
/*
	总结
	1. 这题是一道dp题，因为没有考虑清楚状态转移 ， 比如 4可以这样摆
																	xyyz
																	xxzz

														5可以这样摆
																	xyyzz
																	xxyyz
							所以WA了2次 0.0..不过最终还是AC出来了
	2. 官方有用二维dp做出来的 0.0..我是用1维
	3. 其实可以找规律，如果你能列出正确的数值的话
*/