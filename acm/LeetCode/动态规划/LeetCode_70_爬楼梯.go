package main

import "fmt"

var m map[int]int
func climbStairs(n int) int {
	m = make(map[int]int)
	m[1] = 1
	m[2] = 2
	return solve(n)
}

func solve(n int) int{
	if m[n]==0 {
		m[n] = solve(n-1) + solve(n-2)
	}
	return m[n]
}
func main() {
	fmt.Println(climbStairs(1),climbStairs(2),climbStairs(28),climbStairs(29),climbStairs(30))
}

/*
	总结
	1. 爬楼梯是最简单的动态规划~
*/