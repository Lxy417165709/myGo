package main

import "fmt"

//type arr [][]int
//func (a arr) Len() int{
//	return len(a)
//}
//func (a arr)Less(i,j int)bool{
//	if a[i][0]==a[j][0] {
//		return a[i][0]<a[j][0]
//	}else{
//		return a[i][1]<a[j][1]
//	}
//}
//
func dis( a,b []int) int{
	return (a[0]-b[0]) * (a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

func numberOfBoomerangs(points [][]int) int {

	ans:=0
	for i:=0;i< len(points);i++  {
		m:=make(map[int]int)
		slice:=make([]int,0)
		for t:=0;t<len(points);t++  {
			d:=dis(points[i],points[t])
			if m[d]==0 {
				slice=append(slice, d)
			}
			m[d]++
		}

		for t:=0;t< len(slice);t++  {

			ans+= (m[slice[t]])*(m[slice[t]]-1)
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{

	}))
}
/*
	总结
	1. 这是一道简单题,在做的时候有思路，但是想太多，采用了错误的策略，写了一半才发现。。。
	2. 这题注意点
		外层遍历每个点 i，内层遍历并计算其他点 j 到 i 的距离并通过 Map 保存相等距离的频次
		计算距离公式不用开根号
		计算排列组合公式 n * (n - 1)


*/