package main

func longestValidParentheses(s string) int {
	mxLength := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			mxLength = max(mxLength, right*2)
		} else {
			if right > left {
				right, left = 0, 0
			}
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right {
			mxLength = max(mxLength, left*2)
		} else {
			if right < left {
				right, left = 0, 0
			}
		}
	}
	return mxLength
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	总结
	1. 这个思路是看博客学到的
	2. 思路是: 从左到右遍历字符串，记录当前的左括号数与右括号数，此时如果左==右，则计算长度，如果左<右则将左右清0
				从上面这次遍历的话，我们可以保证右括号可以找到最大的匹配长度。而为了左括号可以找到最大匹配长度，
				我们需要逆序的遍历字符串，操作和上面差不多，只是计算长度方式有点差别，以及右>左时清0。然后取两次遍历的最大值。
*/
