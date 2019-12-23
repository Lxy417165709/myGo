package main

import "fmt"

var s []int
var m []int
var mp map[int]int
func maximizeSweetness(sweetness []int, K int) int {
	s=make([]int,0)
	m=make([]int,0)
	mp=make(map[int]int)
	s = append(s, sweetness[0])
	m = append(m,sweetness[0])
	for i:=1;i< len(sweetness);i++  {
		s = append(s,s[i-1]+sweetness[i])
		m = append(m,min(m[i-1],sweetness[i]))
	}

	return solve(len(sweetness)-1,K)
}

func getSum(a,b int) int{
	if a==0 {
		return s[b]
	}else{
		return s[b] - s[a-1]
	}
}



func solve(n int,K int) int{

	if K==0 {
		mp[n*15000 + K]  = getSum(0,n)
		return mp[n*15000 + K]
	}
	if n==K {
		mp[n*15000 + K] = m[n]
		return m[n]
	}
	if x,ok:=mp[n*15000 + K];ok {
		return x
	}
	ans:=0
	for i:=n-1;i>=K-1 ;i--  {
		gg:=min(solve(i,K-1),getSum(i+1,n))
		ans=max(gg,ans)
	}
	mp[n*15000 + K] = ans
	return ans
}

func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
func main() {
	fmt.Println(maximizeSweetness([]int{
		1,2,3,4,5,6,
	},3))
}
/*
	总结
	1. 超时了 0.0..
*/