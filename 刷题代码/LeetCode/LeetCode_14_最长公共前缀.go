package main

import "fmt"

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}

func longestCommonPrefix(strs []string) string {
	if len(strs)==0 {
		return ""
	}
	mn:=100000000
	for i:=0;i< len(strs);i++  {
		mn = min(mn, len(strs[i]))
	}
	ans:=""

	strFlag:=0
	for i:=0;i<mn;i++  {
		char :=strs[0][i]
		for t:=0;t< len(strs);t++  {
			if strs[t][i]!=char {
				strFlag=1
				break
			}
		}
		if strFlag==0 {
			ans+=string(strs[0][i])
		}else{
			break
		}
	}
	return ans
}


func main() {
	fmt.Println(longestCommonPrefix([]string{
	"","","",
	}))

}
/*
	总结
	1. 这是一道很简单的题目0.0..
	2. 官方有用分治法，二分法做的 0.0.
	3. 字典树也能做 0.0..
*/