package main
func rob(nums []int) int {

	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}

	dp:=make([]int,2)

	// 偷第一间，则最后一间不能偷
	dp[1] = nums[0]
	for i:=2;i<=len(nums)-1;i++{
		dp = append(dp,max(dp[i-2]+nums[i-1],dp[i-1]))
	}
	result1:=dp[len(nums)-1]

	// 不偷第一间，则最后一间能偷
	dp = make([]int,3)
	dp[2] = nums[1]
	for i:=3;i<=len(nums);i++{
		dp = append(dp,max(dp[i-2]+nums[i-1],dp[i-1]))
	}
	result2:=dp[len(nums)]

	return max(result1,result2)
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

/*
	总结
	1. 这是上一题的拓展，这次是成环。
		所以我们可以假设2种情况，一种是偷第一家，则最后一家不能偷
		一种是不偷第一家，则最后一家可以偷，然后其他的就普通DP就oK了
*/
