package main

const (
	size = 100
	inf  = 1000000000
)

func numSquares(n int) int {
	// 平方数数组
	array := [size]int{}

	for i := 1; i <= size; i++ {
		array[i-1] = i * i
	}

	dp := [10500]int{}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = inf
	}

	for t := 0; t < size; t++ {
		for i := array[t]; i <= n; i++ {
			dp[i] = min(dp[i], dp[i-array[t]]+1)
		}
	}
	/*
	这样写就可以省略了array了
	for t := 0; t < size; t++ {
		number := t * t
		for i := number; i <= n; i++ {
			dp[i] = min(dp[i], dp[i - number]+1)
		}
	}
	*/
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
	总结
	1. 这题和零钱兑换1是一样的，思路就是 从array中找出能组成n的最少元素数
	2. array数组其实可以省略，我们在dp的时候再通过t*t求就可以了
*/
