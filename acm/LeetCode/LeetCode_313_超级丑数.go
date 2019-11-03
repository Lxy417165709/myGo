package main

import (
	"fmt"
)



// TLEd的
//var slice []int
//var m map[int]int	// 标记已是丑数, 否则递归会浪费很多步,因为 2*3 == 3*2
//func nthSuperUglyNumber(n int, primes []int) int {
//	slice = make([]int,0)
//	m = make(map[int]int)
//	sort.Sort(sort.Reverse(sort.IntSlice(primes)))
//	fmt.Println(primes)
//	get(1,primes)
//	sort.Ints(slice)
//	fmt.Println(slice[:n])
//	return slice[n-1]
//}
//
//func get(num int,primes []int){
//
//	// m[num]==0很重要!
//	if num<= 1<<32 && m[num]==0 {
//		m[num]=1
//		slice = append(slice, num)
//		for i:=0;i< len(primes);i++  {
//
//			get(num*primes[i],primes)
//		}
//	}
//
//	return
//
//}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

var dp []int
var m map[int]int
func nthSuperUglyNumber(n int, primes []int) int {
	dp = make([]int,0)
	m = make(map[int]int)
	dp = append(dp, 1)
	for i:=0;i< len(primes);i++  {
		m[primes[i]] = 0
	}
	for i:=1;i<=n-1 ;i++  {
		mn := 300000000000
		for t:=0;t< len(primes);t++  {
			mn = min(mn,primes[t]*dp[m[primes[t]]])
		}
		dp= append(dp, mn)
		for t:=0;t< len(primes);t++  {
			if dp[i]==primes[t]*dp[m[primes[t]]] {
				m[primes[t]]++
			}
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(4,[]int{
		2,
	}))
}

/*
	总结
	1. 第一次写的时候，没有标记，所以递归会重复计算，时间复杂度很大
	2. 之后改进了，标记已读，但提交还是TLE
	3. 我要继续想
	4. 看了官方题解后，他是用的DP，之后也用这个DP做出来了，不过这个思路我还没完全掌握，虽然有动手模拟0.0.
*/