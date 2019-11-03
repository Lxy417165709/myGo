package main

import "fmt"


var mp map[int]int

func solve(n int,m int) int{

	if x,ok:=mp[n<<20+m];ok {
		return x
	}


	if n<m {
		n,m=m,n
	}

	if m==0 {
		mp[n<<20+m] = 0
		return 0
	}

	if n==m {
		mp[n<<20+m] = 1
		return 1
	}else{
		mn:=100000
		for i:=1;i<=m ;i++  {
			mn = min(mn,solve(n-i,m) + solve(m,i))
		}
		mp[n<<20+m] = mn

		return mn
	}
}
func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
func tilingRectangle(n int, m int) int {
	mp = make(map[int]int)
	return solve(n,m)
}
func main() {
	fmt.Println(tilingRectangle(11,13))
}
/*
	总结
	1. 这个dp有错误，看题目的第3个案例
*/