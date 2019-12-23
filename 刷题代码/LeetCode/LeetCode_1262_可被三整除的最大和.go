package main
func maxSumDivThree(nums []int) int {
	if len(nums)==0{
		return 0
	}
	dp:=[40050][3]int{}
	dp[0][0]=0
	dp[0][1]=-1
	dp[0][2]=-1
	dp[0][nums[0]%3]=nums[0]

	for i:=1;i<len(nums);i++{
		x:=nums[i]%3
		for t:=0;t<=2;t++{
			if dp[i-1][(3+t-x)%3]!=-1{
				dp[i][t] = max(dp[i-1][t],dp[i-1][(3+t-x)%3] + nums[i])
			}else{
				dp[i][t] = dp[i-1][t]
			}

		}
	}
	return dp[len(nums)-1][0]
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

/*
	总结
	1. 这题刚刚开始想用搜索，然后发现不知道怎么写
	2. 于是我用DP，dp[i][t]表示 以nums[0..i]能组成的mod 3 == t的最大值
	3. 注意初始化条件，dp[0][0]=0 因为就算不选任何数，0mod3==0依旧成立
		而dp[0][1]需要-1表示无法组成，因为不选任何数时，0mod3!=1
		同理，2也是这样的
		然后用第一个元素mod3看看，到底能组成谁，然后再赋值给对应的dp就可以了
		比如第一个元素是5，那么dp[0][5%3==2] = 5
*/