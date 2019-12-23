package main

import "fmt"

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	dx:=[]int{1,-1,-1,1,1,-1,-1,1}
	dy:=[]int{1,-1,1,-1,1,-1,-1,1}
	dz:=[]int{1,1,1,1,-1,-1,-1,-1,}
	arr3:=make([]int,0)
	for i:=0;i< len(arr1);i++  {
		arr3 = append(arr3,i)
	}
	ans:=0
	for i:=0;i< len(dx);i++  {
		mx:=-10000000000
		mn:=10000000000
		for t:=0;t<len(arr1) ;t++  {
			mx=max(mx,arr1[t]*dx[i]+arr2[t]*dy[i]+arr3[t]*dz[i])
			mn=min(mn,arr1[t]*dx[i]+arr2[t]*dy[i]+arr3[t]*dz[i])
		}
		ans = max(ans,mx-mn)
	}
	return ans
}
func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}
func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}
func main() {
	fmt.Println( maxAbsValExpr([]int{
		1,-2,
	},[]int{
		8,8,
	}))
}
/*
	总结
	1. 看题目没有思路，想了双指针，动态规划，分治什么的..也有想过数学的，不过没太到位
	2. 看了官方题解后，解题就是运用去绝对值的后，数学表达式的多样性。
	3. 更规范的说法叫曼哈顿距离 0.0.
	4. 还是有些模糊，希望可以再做道这样的题 0.0.
*/