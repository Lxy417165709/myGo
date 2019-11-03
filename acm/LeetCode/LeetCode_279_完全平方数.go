package main

import "fmt"
// 已通过
//func numSquares(n int) int {
//	dp:=[50500]int{}
//	for i:=1;i<=n ;i++  {
//		x:=10000000
//		for t:=1;t*t<=i ;t++  {
//			if dp[i-t*t]<x {
//				x = dp[i-t*t]
//			}
//		}
//		dp[i] = x+1
//	}
//	return dp[n]
//}

func numSquares(n int) int {
	slice:=make([]int,0)
	for i:=0;i<50000 ;i++  {
		slice = append(slice, i*i)
	}
	return lqdh(slice,n)
}

func lqdh(A []int,target int) int{
	dp:=[50500]int{}
	for i:=1;i<=target ;i++  {
		x:=10000000	 // 表示无穷大
		//如果数组有序，其实这还可以优化
		for t:=1;A[t]<=i ;t++  {
			if dp[i-A[t]]<x {
				x = dp[i-A[t]]
			}
		}
		dp[i] = x+1
	}
	return dp[target]


}

func main() {
	fmt.Println(numSquares(0))
}
/*
	总结
	1. 这是一道dp题，其实做法和 零钱兑换1一样~

*/