package main

import (
	"fmt"
	"sort"
)

var slice []int

func nthUglyNumber(n int) int {
	slice:=make([]int,0)
	for i:=1;i<=1<<32 ;i=i*5  {
		for t:=1;i*t<=1<<32 ;t=t*3  {
			for j:=1;i*t*j<=1<<32 ;j=j*2  {
				slice = append(slice, i*j*t)
			}
		}
	}
	sort.Ints(slice)
	return slice[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(1690))
}

/*
	总结
	1. 该题我使用打表方式AC
	2. 官方有使用DP的 0.0..
*/