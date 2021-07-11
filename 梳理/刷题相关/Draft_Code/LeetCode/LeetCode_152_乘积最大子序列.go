package main

func maxProduct(nums []int) int {
	dp:=[][]int{}

	dp = append(dp,[]int{nums[0],nums[0]})
	ans:=nums[0]
	for i:=1;i<len(nums);i++{

		dp = append(
			dp,
			[]int{
				max(dp[i-1][0]*nums[i],dp[i-1][1]*nums[i],nums[i]),
				min(dp[i-1][0]*nums[i],dp[i-1][1]*nums[i],nums[i]),
			},
		)
		ans = max(ans,dp[i][0])
	}
	return ans
}

func max(arr ...int) int{
	if len(arr)==1{
		return arr[0]
	}

	a,b:=arr[0],max(arr[1:]...)
	if a>b{
		return a
	}
	return b

}


func min(arr ...int) int{
	if len(arr)==1{
		return arr[0]
	}

	a,b:=arr[0],min(arr[1:]...)
	if a>b{
		return b
	}
	return a

}


/*
// 滚动优化后 (保留dp数组)
func maxProduct(nums []int) int {
    dp:=[][]int{[]int{nums[0],nums[0]}}

    ans:=nums[0]
    fmt.Println(dp[0][0],dp[0][1])
    for i:=1;i<len(nums);i++{
        mx := max(dp[0][0]*nums[i],dp[0][1]*nums[i],nums[i])
        mn := min(dp[0][0]*nums[i],dp[0][1]*nums[i],nums[i])
        dp[0][0] = mx
        dp[0][1] = mn

		// 注意不能合下面那样写，这样的话会导致修改后的dp[0][0]影响到dp[0][1]
		// dp[0][0] = max(dp[0][0]*nums[i],dp[0][1]*nums[i],nums[i])
		// dp[0][1] = min(dp[0][0]*nums[i],dp[0][1]*nums[i],nums[i])

    }
    return ans
}

// 直接剔除了dp数组
func maxProduct(nums []int) int {
    ans,mx,mn:=nums[0],nums[0],nums[0]
    for i:=1;i<len(nums);i++{
        tmp1 := max(mx*nums[i],mn*nums[i],nums[i])
        tmp2 := min(mx*nums[i],mn*nums[i],nums[i])
        mx,mn = tmp1,tmp2
        ans = max(ans,mx)
    }
    return ans
}



*/


/*
	总结
	1. 这题目我是看了官方题解AC的，第一次想到的也是使用dp,只不过思路不清晰，然后就没做出来。
	2. 清楚了思路后，我就设计了个dp数组, dp[i][0]表示以nums[i]为结尾的乘积子串的最大值，而dp[i][1]为最小值
		于是可以得出，
			当选择nums[i]，且不抛弃之前的一段时,
				dp[i][0] = dp[i-1][0]*nums[i]	// 此时nums[i]>=0
				dp[i][0] = dp[i-1][1]*nums[i]	// 此时nums[i]<0，因为如果nums[i]<0,那么如果dp[i-1][1]<0，他们乘积就＞0了。
												   当然此时dp[i-1][0]也可以是负数，但是显然乘最小的负数最好(所得数更大)

				dp[i][1] = dp[i-1][1]*nums[i]	// 此时nums[i]>=0
				dp[i][1] = dp[i-1][0]*nums[i]	// 此时nums[i]<0
			当选择nums[i], 且抛弃之前的一段时，
				dp[i][0] = nums[i]
				dp[i][1] = nums[i]
		于是就可以写出状态转移方程
				dp[i][0] = max(dp[i-1][0]*nums[i],dp[i-1][1]*nums[i],nums[i]),
				dp[i][1] = min(dp[i-1][0]*nums[i],dp[i-1][1]*nums[i],nums[i]),
		当然，此时我们也可以用滚动数组优化(因为每次我们只用到前一位的数据)，将空间复杂度优化为O(1)。

*/
