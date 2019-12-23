package main

import "fmt"

func min(a,b,c int) int{
	if a<b {
		if a<c {
			return a
		}else{
			return c
		}
	}else{
		if b<c {
			return b
		}else{
			return c
		}
	}
}
func maximalSquare(matrix [][]byte) int {
	dp := [1050][1050]int{}
	ans:=0
	for i:=0;i< len(matrix);i++  {
		for t:=0;t<len(matrix[i]) ;t++  {
			if matrix[i][t]=='1' {
				if i-1>=0 && t-1>=0 {
					dp[i][t]=min(dp[i-1][t-1],dp[i-1][t],dp[i][t-1])+1
				}else{
					if i-1>=0 {
						dp[i][t]=dp[i-1][t]+1
					}
					if t-1>=0 {
						dp[i][t]=dp[i][t-1]+1
					}
					if i-1<0 || t-1<0 {
						dp[i][t] = 1
					}
				}
			}
			if ans<dp[i][t] {
				ans = dp[i][t]
			}
		}
	}
	return ans*ans
}

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1','1','1','1','1'},
		{'1','1','1','1','1'},
		{'1','1','1','1','1'},
		{'1','1','1','1','1'},
	}))
}

/*
	总结
	1. 这是一道动态规划题，思路其实我还没完全掌握...
*/