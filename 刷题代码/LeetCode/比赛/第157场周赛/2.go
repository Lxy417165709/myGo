package main

import "fmt"

func max(a,b int) int{
	if a>b {
		return a

	}
	return b
}
func longestSubsequence(arr []int, difference int) int {
	m:=make(map[int]int)
	ans:=0
	for i:=0;i< len(arr);i++  {
		m[arr[i]]=max(m[arr[i]],m[arr[i]-difference]+1)
		ans=max(m[arr[i]],ans)
	}
	return ans

}

func main() {
fmt.Println(longestSubsequence([]int{
	1,
},-2000))
}
/*
	总结
	1. 这道题也有点dp的思维 0.0 我用map存放以x结尾且元素相差等于differens的最长长度
*/