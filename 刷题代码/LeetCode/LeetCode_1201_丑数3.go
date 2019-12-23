package main

import "sort"

func nthUglyNumber(n int, a int, b int, c int) int {
	tmp:=[]int{a,b,c}
	sort.Ints(tmp)

	// dp:=make([]int,0)
	// dp = append(dp,1)
	ans,cnt:=tmp[0],1

	points:=[]int{1,1,1}
	// ans:=1
	for cnt<n{
		mn:=2000000000000000000
		idx:=0
		for t:=0;t<len(tmp);t++{
			x := tmp[t]*points[t]
			if mn>x{
				mn = x
				idx = t
			}
		}
		// fmt.Println(idx,i,mn,dp[i])
		// if mn!=dp[i]{
		// dp = append(dp,mn)
		// }
		// fmt.Println(dp)
		if mn!=ans{
			ans = mn
			cnt++
		}
		points[idx]++
	}

	return ans
}
func main() {

}
/*
	总结
	1. 这题如果按照动态规划做的话肯定超时，上面的代码就是动态规划的
	2. 要正确做出这题需要用到数论知识。。
*/