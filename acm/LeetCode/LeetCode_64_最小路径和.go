package main

import "fmt"

var dp [1005][1005]int

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp = [1005][1005]int{}
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for t := 0; t < m; t++ {

			if i==0 && t==0 {
				dp[i][t] = grid[i][t]
				continue
			}
			x, y := 100000000, 100000000	// 表示无穷大
			if i-1 >= 0 {
				x = dp[i-1][t]
			}
			if t-1 >= 0 {
				y = dp[i][t-1]
			}
			dp[i][t] = min(x, y) + grid[i][t]
		}
	}
	return dp[n-1][m-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1,3,1,9},{1,5,1,9},{4,2,1,2},
	}))
}
/*
	总结
	1. 这是一道简单的DP题
	2. 其实二维数组可以优化为一维数组的~
*/