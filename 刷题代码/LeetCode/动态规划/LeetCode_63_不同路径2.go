package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp:=[105][105]int{}

	m,n := len(obstacleGrid[0]),len(obstacleGrid)
	if m*n==0 {
		return 0
	}else{
		if obstacleGrid[0][0]==1 {
			return 0
		}
	}


	dp[0][0] = 1
	for i:=1; i<m;i++  {
		if obstacleGrid[0][i]==1 {
			dp[0][i]=0
		}else{
			dp[0][i]=dp[0][i-1]
		}
	}

	for i:=1; i<n;i++  {
		if obstacleGrid[i][0]==1 {
			dp[i][0]=0
		}else{
			dp[i][0] = dp[i-1][0]
		}
	}

	for i:=1;i<n ;i++  {
		for t:=1;t<m ;t++  {
			if obstacleGrid[i][t]==1 {
				dp[i][t]=0
			}else{
				dp[i][t] = dp[i-1][t] + dp[i][t-1]
			}
		}
	}
	return dp[n-1][m-1]


}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{1}}))
}

/*
	总结
	1. 该题目用到了动态规划
	2. 该题目是上道题的拓展版~
*/