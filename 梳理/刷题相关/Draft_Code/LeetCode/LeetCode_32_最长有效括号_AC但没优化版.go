package main
func longestValidParentheses(s string) int {
	stack :=make([]string,0)
	mxLength:=0
	for i:=0;i<len(s);i++{
		if s[i]==')'{
			tmpStr:=""
			for len(stack)!=0{
				top:=stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if len(top)==1{
					// 可以保证是"("
					tmpStr = "(" + tmpStr + ")"
					stack = append(stack,tmpStr)
					break
				}else{
					tmpStr = top + tmpStr
				}
			}
			tmpStr = ""
			for len(stack)!=0{
				top:=stack[len(stack)-1]
				if len(top)>=2{
					tmpStr = top + tmpStr
					stack = stack[:len(stack)-1]
				}else{
					break
				}
			}
			// fmt.Println(stack)
			if tmpStr != ""{
				stack = append(stack,tmpStr)
			}

			mxLength = max(mxLength,len(tmpStr))
		}else{
			stack = append(stack,string(s[i]))
		}
	}
	return mxLength

}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
/*
	总结
	1. 这题目我的思路是，顺序扫描字符串，遇到（直接加入栈，遇到)则进行判断，操作是找到第一个(，合并为一个合法括号字符串
		之后再加入栈，再把这个合法字符串与其前面的合法字符串合并，计算出总合法字符串的长度值。之后继续顺序扫描字符串。
	2. 这个方法太复杂了...
*/