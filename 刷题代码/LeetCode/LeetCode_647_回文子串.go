package main

func countSubstrings(s string) int {
	dp :=[1050][1050]bool{}
	ans:=0
	for length:=1;length<=len(s);length++{
		for begin:=0;begin+length-1<len(s);begin++{
			end:=begin+length-1
			if s[begin]==s[end]{
				if length<=2 || dp[begin+1][end-1]{
					dp[begin][end]=true
					ans++
				}
			}
		}
	}
	return ans
}
/*
	总结
	1. dp[begin][end], 表示s[begin:end+1]是否为回文子串，
		于是 当dp[begin][end]为true时，我们就找到了一个回文子串
		而当s[begin]==s[end] 且(length <=2 || dp[begin+1][end-1] == true)时，dp[begin][end]=true,
*/
