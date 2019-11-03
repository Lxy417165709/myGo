package main

import "fmt"


// 在A数组中,问有多少种能组成target的组合。		(A数组无重复元素，每个元素可取无数次)
// 一维写法
var dp [100000]int
func change(A []int,target int) int{
	dp = [100000]int{}
	dp[0] = 1
	for i:=0;i<len(A) ;i++  {
		for t:=0;t<=target;t++  {
			if t-A[i]>=0 {
				dp[t]+=dp[t-A[i]]
			}
		}
	}
	return dp[target]
}

// 在A数组中,问有多少种能组成target的组合。		(A数组无重复元素，每个元素可取无数次)
// 二维写法
// 这种写法耗费很大的空间
var dp2 [500][50000]int
func change2(A []int,target int) int{
	if target==0 {
		return 1
	}
	if len(A)==0 {
		return 0
	}
	dp2 = [500][50000]int{}
	dp2[0][0] = 1
	for i:=0;i<len(A) ;i++  {
		for t:=0;t<=target;t++  {
			if t-A[i]>=0{
				dp2[i][t]+=dp2[i][t-A[i]]
			}
			if i-1>=0 {
				dp2[i][t]+=dp2[i-1][t]
			}
		}
	}
	return dp2[len(A)-1][target]
}

// 相关题目：零钱兑换2
//-------------------------------




// 在A数组中,问有多少种能组成target的排列。		(A数组无重复元素，每个元素可取无数次)
// 一维就可以解决了
var plt_dp[105]int
func plt(A []int,target int) int{
	if target==0 {
		return 1
	}
	if len(A)==0 {
		return 0
	}
	plt_dp = [105]int{}
	plt_dp[0] = 1
	for i:=1;i<=target ;i++  {
		for t:=0;t< len(A);t++  {
			if i-A[t]>=0 {
				plt_dp[i]+=plt_dp[i-A[t]]
			}
		}

	}
	return plt_dp[target]
}

func climbStairs(n int) int {

	return plt([]int{1,2},n)
}

// 相关题目：LeetCode_70_爬楼梯
// ------------


// 在A数组中,问最少几个元素的和为target。		([A数组无重复元素](因为可以无穷取，所以这个有无重复没意义了)，每个元素可取无数次)

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

// 相关题目：LeetCode_279_完全平方数，LeetCode_517_零钱兑换

func main() {
	fmt.Println(climbStairs(5))
}
