package main

func longestValidParentheses(s string) int {
	dp := make([]int, len(s)+1)

	mx := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
			} else {
				idx := i - dp[i] - 1
				// idx >=0判断很重要
				if idx >= 0 && s[idx] == '(' {
					dp[i+1] = i - idx + 1 + dp[idx]
				} else {
					dp[i+1] = 0
				}
			}
		} else {
			dp[i+1] = 0
		}
		mx = max(mx, dp[i+1])
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
/*
	总结
	1. 这是看了官方题解之后，用dp做出来的
	2. dp[i] 表示以s[i-1]结尾的有效括号长度，dp[0]表示空字符串的有效括号长度
*/