package main

func isValid(s string) bool {
	stack := []uint8{}
	mp :=make(map[uint8]int)
	mp['(']=1
	mp['[']=2
	mp['{']=3
	mp[')']=4
	mp[']']=5
	mp['}']=6

	for i:=0;i<len(s);i++{
		// >=4说明不是左括号
		if mp[s[i]]>=4 {
			if len(stack)==0{
				return false
			}
			top:=stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mp[top]!=mp[s[i]]-3{
				return false
			}
		}else{
			stack = append(stack,s[i])
		}
	}
	return len(stack)==0
}

/*
	总结
	1. 这题是很简单的，直接使用栈就可以AC了
	2. 官方评论有用替换思想，将() [] {}替换为""，最后如果s==""则有效
	3. 官方还有使用数字映射的，把左右括号映射为相应的相反数，之后从左到右扫描，再加上栈，也可以AC。
		判断就变成了，栈顶value和现在的字符value和是否为0
*/