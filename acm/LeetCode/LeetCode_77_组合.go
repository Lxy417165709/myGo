package main

import "fmt"

var ans [][]int
func combine(n int, k int) [][]int {
	ans =make([][]int,0)
	solve(1,n,k,[]int{})
	return ans
}

func solve(x int,n int,k int,slice []int){

	if k==0 {
		temp:=make([]int, len(slice))
		copy(temp,slice)
		ans = append(ans, temp)
		return
	}
	for i:=x;i<=n ;i++  {
		slice = append(slice, i)
		solve(i+1,n,k-1,slice)
		slice = slice[:len(slice)-1]
	}

}

func main() {
	fmt.Println(combine(10,3))
}

/*
	总结
	1. 该题考查回溯递归~
	2. 官方题解还提供了字典序法 ..我还没看~
	3. 2法时间空间复杂度都一样
*/