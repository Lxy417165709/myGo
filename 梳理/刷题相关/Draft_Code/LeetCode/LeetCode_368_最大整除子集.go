package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	dp:=make([][]int,0)
	//for i:=0;i<len(nums) ;i++  {
	//	dp[i] = append(dp[i], nums[i])
	//}

	for i:=0;i< len(nums);i++  {
		y:=make([]int,0)
		for t:=0;t<=i-1 ;t++  {
			if nums[i]%nums[t]==0 {
				if len(dp[t])> len(y) {
					y = dp[t]
				}
			}
		}
		dp = append(dp, []int{})
		dp[i] = append(y,nums[i])
	}
	ans:=make([]int,0)
	for i:=0;i< len(dp);i++  {
		if len(ans)< len(dp[i]) {
			ans = dp[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1,2,3,4,9,27}))
}
/*
	总结
	1. 该题我使用了DP，先排序后DP 0.0..DP定义是，以nums[i]元素为尾的最长整除序列
	2. 通过得到的最长整除序列集，通过遍历拿出他们中最长的一条就OK了
*/