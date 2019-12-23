package main

import (
	"fmt"
)

func min(a, b int) int{
	if a>b {
		return b
	}
	return a
}

func max(a, b int) int{
	if a>b {
		return a
	}
	return b
}

func dis(p1 []int,p2 []int) int{
	a:=(p1[0]-p2[0])*(p1[0]-p2[0])
	b:=(p1[1]-p2[1])*(p1[1]-p2[1])
	return a+b
}
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	slice := make([][]int,0)
	slice = append(slice, p1,p2,p3,p4)

	sliMn:=make([]int,0)
	sliMx:=make([]int,0)

	for i:=0;i<4 ;i++  {
		mn:=10000000000
		mx:=0
		for t:=0;t<4 ;t++  {
			if i==t {
				continue
			}
			x:=dis(slice[i],slice[t])
			mn = min(mn,x)
			mx = max(mx,x)
		}
		sliMn = append(sliMn, mn)
		sliMx = append(sliMx, mx)
	}
	a,b:=sliMn[0],sliMx[0]
	for i:=0;i< len(sliMn);i++  {
		if sliMn[i]!=a || sliMx[i]!=b || sliMn[i]*2!=sliMx[i] {
			return false
		}
	}

	return a!=0
}




func main() {
	fmt.Println(validSquare([]int{-2,2},[]int{2,-2},[]int{2,2},[]int{-2,-2}))
}

/*
	总结
	1. 第一次的策略是获得4个点分别的最小的距离，然后相等且不等于0则为正方形，但是这种
		方法是错误的，比如在同一直线 0.0
	2. 第二次发现诶，正方形的对角线是最长的边，于是再加上获取最长边，然后对角线与
		边的关系是  边*sqrt(2) == 对角线 ，然后判断就OK了
		注意 由于我计算距离没开方 所以最后判断的是  边^2 * 2 == 对角线^2
*/