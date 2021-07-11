package main

import "fmt"

func trailingZeroes(n int) int {
	ans:=0
	for ;n!=0 ; {
		fmt.Println(n)
		ans+=n/5
		n = n/5
	}
	return ans
}


func main() {
	fmt.Println(trailingZeroes(0))
}
/*
	总结
	1. 要计算n!阶乘后的零，只需要计算1-n之中有多少个因子5，25的话是算2个5的因为 25 = 5 * 5
*/