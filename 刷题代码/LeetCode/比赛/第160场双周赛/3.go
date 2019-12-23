package main

import "fmt"

func maxLength(arr []string) int {
	sli:=make([]int, len(arr))
	ans = 0
	for i:=0;i< len(arr);i++  {
		for t:=0;t< len(arr[i]);t++  {
			x:=1<<uint8(arr[i][t]-'a')
			if sli[i]&x==0 {
				sli[i] = sli[i]|x
			}else{
				sli[i] = -1
				break
			}
		}
	}

	fmt.Println(sli)
	solve(sli,arr,[]int{},0, len(arr))
	return ans
}
var ans int

func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
func solve(sli []int,arr []string,finallySli []int,index int,n int){
	if index==n {
		x:=0
		length:=0
		for i:=0;i< len(finallySli);i++  {
			if finallySli[i]==-1 {
				continue
			}
			if x&finallySli[i]==0 {
				x = x|finallySli[i]
				length+= len(arr[i])
				ans = max(ans,length)
			}else{
				return
			}
		}

		return
	}
	finallySli = append(finallySli, sli[index])	// 表示选
	solve(sli,arr,finallySli,index+1,n)
	finallySli[len(finallySli)-1] = -1	// 表示不选
	solve(sli,arr,finallySli,index+1,n)
}

func main() {
	fmt.Println(maxLength([]string{
		"aa",
	}))
}
/*
	总结
	1. 这题我用dfs暴力搜索 + 位运算判断是否有重复字母 AC了 0.0
	2. x&y==0 表示x与y在同位上没有1  x = x|y 表示把y上的1放入x中
	3. 这题目我应该是做复杂了 0.0..
*/