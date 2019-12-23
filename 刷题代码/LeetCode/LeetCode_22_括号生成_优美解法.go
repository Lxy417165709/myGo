package main

var ans []string
func generateParenthesis(n int) []string {
	ans = make([]string,0)
	solve(n,n,"")
	return ans
}
func solve(left int,right int,x string){
	if left>right{
		return
	}
	if left==0 && right==0{
		ans = append(ans,x)
		return
	}
	if right>0{
		solve(left,right-1,x+")")
	}
	if left>0{
		solve(left-1,right,x+"(")
	}
}

/*
	总结
	1. 这是一个优美的回溯解法
	2. 规则，solve函数left表示剩余的左括号数,right表示右括号数
		如果left>right 那么接下来不可能构成合法括号
		如果left==right==0说明已经完成了，直接把括号串加入答案
		如果有剩余的括号，则递归解决 0.0
*/