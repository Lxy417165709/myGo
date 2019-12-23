package main

import "fmt"

func hammingDistance(x int, y int) int {
	ans:=0
	for ;x!=0 || y!=0 ;  {
		wx:=x&1
		wy:=y&1
		x>>=1
		y>>=1
		if wx!=wy {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(hammingDistance(16,12))
}
/*
	总结
	1. 这是一道简单题 0.0.  x&1 等于 x%2
*/