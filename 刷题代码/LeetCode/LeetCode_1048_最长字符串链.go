package main

import "sort"

func longestStrChain(words []string) int {
	if len(words)==0{
		return 0
	}
	sort.Slice(words,func (i,j int) bool{
		return len(words[i])<len(words[j])
	})

	dp := [1050]int{}
	dp[0] = 1
	ans :=1
	for i:=1;i<len(words);i++{
		dp[i] = 1
		for t:=0;t<i;t++{
			if judge(words[t],words[i]){
				dp[i] = max(dp[i],dp[t]+1)
			}
		}
		ans=max(ans,dp[i])
	}
	return ans
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func judge(a,b string) bool{

	if len(b)!=len(a)+1{
		return false
	}
	first,second:=0,0
	wordFlag:=false
	for first<len(a) && second<len(b){
		if a[first]==b[second]{
			first++
			second++
		}else{
			if wordFlag==true{
				return false
			}
			second++
			wordFlag = true
		}
	}
	return true
}

/*
	总结
	1. 我的思路是，使用dp,dp[i]表示以words[i]结尾的最长字符串链长度，所以 dp[i] = dp[0...i-1]+1 (如果可成为链的话)，否则为1
	2. 这里就要一个判断子序列的函数了，我写的是judge,first,second分别指向第一个第二个字符串，之后根据条件，两个指针适当移动，最终就能得出结果了。
	3. 记得先以长度递增排序
*/
