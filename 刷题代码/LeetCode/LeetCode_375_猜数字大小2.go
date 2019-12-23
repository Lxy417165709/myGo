package main

import "fmt"

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
var m map[int]int
func getMoneyAmount(n int) int {
	m = make(map[int]int)
	//for i:=1;i<=30 ;i++  {
	//	fmt.Println(i,solve(1,i))
	//}
	return solve(1,n)
}

func solve(l int,r int) int{
	if x,ok:=m[l<<30|r];ok {
		return x
	}
	if l==r {
		return 0
	}
	if l+1==r {
		return l
	}
	if l+2==r {
		return l+1
	}
	ans:=10000000000
	// i:=(l+r)/2 是优化过的，因为我们并不需要从l+1开始遍历，官方说[l+1，(l+r)/2)的总代价一定不是最小的 0.0，所以我们从(l+r)/2开始就OK了
	for i:=(l+r)/2;i<=r-1 ;i++  {
		x:=max(solve(l,i-1),solve(i+1,r))+i
		ans = min(ans,x)
	}

	m[l<<30|r] = ans
	return m[l<<30|r]
}

func main() {
	fmt.Println(getMoneyAmount(300))
}
/*
	总结
	1. 第一次做的时候没思路，用自己所认为的策略来计算，然后直接错了
	2. 第二次使用分治法 0.0，可以得出正确答案，但是由于重叠子问题，时间复杂度太高了，于是使用map优化，提交AC了
	3. 看了官方题解后，对自己写的分治法做了一些优化(其实现在分治更倾向于从上到下的DP)
	4. 官方有使用DP做出来的 0.0，我是使用递归DP而他是迭代DP

*/