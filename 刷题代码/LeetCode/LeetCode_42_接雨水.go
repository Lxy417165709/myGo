package main

import "fmt"

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
func trap(height []int) int {
	lmax := make([]int,0)
	rmax := make([]int,0)
	lmax = append(lmax, 0)
	for i:=1;i< len(height);i++  {
		if lmax[i-1]<height[i-1] {
			lmax = append(lmax, height[i-1])
		}else{
			lmax = append(lmax, lmax[i-1])
		}
	}
	rmax = append(rmax, 0)
	for i:= len(height)-2;i>=0;i--  {
		if rmax[len(height)-i-2]<height[i+1] {
			rmax = append(rmax, height[i+1])
		}else{
			rmax = append(rmax, rmax[len(height)-i-2])
		}
	}

	fmt.Println(lmax,rmax)
	ans:=0
	for i:=0;i<len(height) ;i++  {
		x:=min(lmax[i],rmax[len(height)-1-i])-height[i]
		if x>0 {
			ans+=x
		}
	}
	return ans
}
func main() {
	fmt.Println(trap([]int{

	}))
}
/*
	总结
	1. 要把握整体与局部呀，多角度看问题 0.0
	2. 这题我是看答案了之后才会的 0.0.
	3. 看了官方评论后还有一种想法，就是互补法，先求出最大的矩形，之后减去总高度，这样也可以得出来好像 0.0
*/