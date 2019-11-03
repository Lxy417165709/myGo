package main

import "fmt"

func probabilityOfHeads(prob []float64, target int) float64 {
	dp:=[1050][1050]float64{}

	dp[0][0] = 1
	for i:=1;i<=target ;i++  {
		dp[0][i] = 0
	}
	for i:=1;i<= len(prob);i++  {
		dp[i][0] = dp[i-1][0]*(1-prob[i-1])
	}
	for i:=1;i<= len(prob);i++  {
		for t:=1;t<=target ;t++  {
			dp[i][t] = dp[i-1][t-1]*prob[i-1] + dp[i-1][t]*(1-prob[i-1])
		}
	}
	return dp[len(prob)][target]

}

func main() {
	fmt.Println(probabilityOfHeads([]float64{
		0,0,0,
	},1))
}

/*
	总结
	1. 这题我用的是dp dp[n][t] 表示前n个正面次数为t的概率(n不是索引喔,是长度)
			则可以知道 dp[n][t] = dp[n-1][t-1]*pros[n-1]+dp[n-1][t]*(1-pros[n-1])
						边界条件上面有写了 0.0..
*/
