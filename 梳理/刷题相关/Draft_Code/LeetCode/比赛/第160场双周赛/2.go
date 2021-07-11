package main

import "fmt"

func pow(x int,n int) int{
	sum:=1
	for	i:=0;i<n;i++{
		sum = sum *x
	}
	return sum
}
func get(n int) []int{

	ans:=[][]int{
		{0},
	}
	for i:=1;i<=n ;i++  {
		fmt.Println(i)
		k:=pow(2,i-1)
		ans = append(ans,make([]int,len(ans[i-1])))
		copy(ans[i],ans[i-1])
		for t:=len(ans[i-1])-1;t>= 0;t--  {
			ans[i] = append(ans[i], ans[i-1][t] + k)

		}

	}
	return ans[n]
}

func circularPermutation(n int, start int) []int {
	sli:=get(n)

	index:=0
	for i:=0;i< len(sli);i++  {
		if sli[i]==start {
			index = i
			break
		}
	}
	ans:=make([]int,0)
	for i:=0;i< len(sli);i++  {
		ans = append(ans, sli[(index+i)%len(sli)])
	}
	return ans
}

func main() {
	fmt.Println(circularPermutation(3,0))
}

/*
	总结
	1. 这道题之前做过，但是比赛的时候好久才想起
	2. 这道题和之前的做了一个变形，之前是规定0开始的，现在规定从x开始，其实也简单 0.0
*/