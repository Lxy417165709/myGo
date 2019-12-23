package main

import "fmt"

func balancedString(s string) int {

	mp := make(map[uint8]int)
	mp['Q']=0
	mp['W']=1
	mp['E']=2
	mp['R']=3

	arr :=make([]int,4)
	arrJudge :=make([]int,4)


	for i:=0;i<len(s);i++{
		arr[mp[s[i]]]++
	}
	x:=len(s)/4


	flag:=0
	for i:=0;i<len(arr);i++{
		arr[i]-=x
		if arr[i]!=0{
			flag=1
		}
	}
	if flag==0{
		return 0
	}


	fmt.Println(arr)
	ans:=1000000000
	l:=0
	for r:=0;r<len(s);r++{
		arrJudge[mp[s[r]]]++

		for l<=r && judge(arr,arrJudge){
			ans = min(ans,r-l+1)
			arrJudge[mp[s[l]]]--
			l++
		}
	}
	return ans
}
func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
func judge(a []int, b []int) bool{
	for i:=0;i<len(a);i++{
		if a[i]>0{
			if a[i]>b[i]{
				return false
			}
		}
	}
	return true
}

func main() {

}

/*
	总结
	1. 这题目使用了滑动窗口算法
	2. 代码还可以更优美 0.0
*/