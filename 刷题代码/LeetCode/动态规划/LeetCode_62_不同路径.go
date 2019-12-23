package main

import "fmt"

var arr [105][105]int
func uniquePaths(m int, n int) int {
	arr = [105][105]int{}

	for i:=1;i<=m ;i++  {
		arr[1][i] =  1
	}
	for i:=1;i<=n ;i++  {
		arr[i][1] =  1
	}
	for i:=2;i<=m ;i++  {
		for t:=2;t<=n ;t++  {
			arr[t][i] = arr[t-1][i] + arr[t][i-1]
		}
	}
	return arr[n][m]


}

func main() {
	fmt.Println(uniquePaths(100,100))
}


/*
	总结
	1. 该题目考察动态规划 dp[i][t] = dp[i-1][t] + dp[i][t-1]
	2. 官方O(1)算法：因为机器人到目标一共要走m+n-2步，所以只需要在这m+n-2步中
       选择m-1步向右走就可以了(同理n-1步向下走)，公式是 C(m+n-2)(m-1) == C(m+n-2)(n-1)
*/