package main

import (
	"fmt"

)
//var ans int
//var m map[int]int
//func change(amount int, coins []int) int {
//	ans=0
//	m = make(map[int]int)
//	sort.Ints(coins)
//	solve(amount,coins,[]int{})
//
//	for i:=0;i<=500 ;i++  {
//		fmt.Println(i,m[i])
//	}
//
//	return ans
//}
//func search_l_r(slice []int,target int) int{
//	l:=0
//	r:= len(slice)-1
//
//	for ;l<=r ;  {
//		mid := (l+r)/2
//		if slice[mid] == target {
//			l = l + 1
//		}else{
//			if slice[mid] > target{
//				r = r -1
//			}else{
//				l = l + 1
//			}
//		}
//	}
//	return r
//}
//
//func solve(amount int,coins []int,slice []int) {
//
//
//
//	if amount == 0 {
//		m[0]++
//		for i:=0;i< len(slice);i++  {
//			m[slice[i]]++
//		}
//	}
//
//	if m[amount]!=0 {
//		ans+=m[amount]
//		return
//	}
//
//
//
//	//fmt.Println(slice)
//
//	index:=search_l_r(coins,amount)
//	for i:= index;i>=0 ;i--  {
//		//slice = append(slice, amount-coins[i])
//		solve(amount-coins[i],coins[:i+1],slice)
//		// solve(amount-coins[i],coins) 这样的话就变成了跳楼梯了
//	}
//}

var m [100000]int
func change(amount int,coins []int) int{
	m = [100000]int{}
	m[0] = 1
	for i:=0;i<len(coins) ;i++  {
		for t:=0;t<=amount;t++  {
			if t-coins[i]>=0 {
				m[t]=m[t] + m[t-coins[i]]
			}
		}
	}
	//fmt.Println(m)
	return m[amount]
}

func main() {
	fmt.Println(change(500,[]int{1}))
}

/*
	总结
	1. 第一次回溯递归的时候TLE了 0.0.. (自己写的)
	2. 看官方题解发现是完全背包，之后借鉴后写出来AC了
*/