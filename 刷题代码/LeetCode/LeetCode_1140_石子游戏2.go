package main

import "fmt"

func max(a,b int) int{
	if	a>b{
		return a
	}
	return b
}

var m map[int]int
var sum []int
func stoneGameII(piles []int) int {
	m=make(map[int]int)
	sum=make([]int,1)
	for i:=0;i< len(piles);i++  {
		sum = append(sum,piles[i]+sum[i])
	}

	return solve(piles,0,1)
}

func solve(piles []int,begin int,M int) int{
	// 由于下面循环判断有 begin+i<=len(piles)，所以这个if不用判断了 0.0，加上也OK，但是代码永远不会执行这个if
	//if begin>= len(piles) {
	//	return 0
	//}
	if x,ok:=m[begin<<20+M];ok {
		return x
	}

	ans:=0
	// 为了防止越界，我加上了begin+i<=len(piles) 0.0
	for i:=1;begin+i<=len(piles) && i<=2*M ;i++  {
		rest:=sum[len(piles)]-sum[begin+i]
		own:=sum[begin+i]-sum[begin]
		ans = max(ans,own+(rest-solve(piles,begin+i,max(M,i))))
	}
	m[begin<<20+M]=ans
	return ans
}

func main() {
	fmt.Println(stoneGameII([]int{
		1,2,5,8,8,8,
	}))
}
/*
	总结
	1. 第一次想到用回溯写，但是由于方程什么的没想好，所以总没有想要的答案，最难的在于他们都足够聪明
	2. 看了官方的题解后，知道了足够聪明应该如何写，然后就写出来了 0.0.
	3. 我通过记忆化搜索+前缀和 写出来了,AC了 0.0.
	4. 我学会了足够聪明的写法了！！我太牛皮了！！！
	5. 这类博弈题，dfs时候可以不需要分清两人的身份
*/