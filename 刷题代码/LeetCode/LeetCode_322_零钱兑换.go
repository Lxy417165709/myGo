package 双指针

import "fmt"

const (
	inf = 1000000000
)

func coinChange(coins []int, amount int) int {

	if len(coins) == 0 {
		return -1
	}

	dp := [10500]int{}
	// dp[i]表示凑成金额i所需最少硬币个数
	// 则 dp[0] == 0
	// 则 dp[i] = min(dp[i], dp[i-coins[t]] + 1)  ( t∈ [0, len(coins)-1])

	for i := 1; i <= amount; i++ {
		dp[i] = inf
	}
	for i := 1; i <= amount; i++ {
		for t := 0; t < len(coins); t++ {
			if coins[t] <= i {
				dp[i] = min(dp[i], dp[i-coins[t]]+1)
			}
		}
	}

	// 硬币放在外部循环也可以，因为我们不关心排列还是组合，因为不管是排列还是组合，dp[i]的前驱状态都是完整的
	//for t:=0;t<len(coins);t++{
	//	for i:=1;i<=amount;i++{
	//		if coins[t]<=i{
	//			dp[i] = min(dp[i],dp[i-coins[t]]+1)
	//		}
	//	}
	//}

	if dp[amount] == inf {
		return -1
	}
	return dp[amount]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 这题之前做过，然后现在是重做了一次。
	2. 本来我是使用二维数组写的，但是太麻烦也太费空间了，于是该用一维数组。
	3. 第一次写的使用用了切片，然后发现时空效率都慢于使用数组...

*/
