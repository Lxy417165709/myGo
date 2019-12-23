package main
func minRemoveToMakeValid(s string) string {

	stack := make([]string,0)

	for i:=0;i<len(s);i++{
		tmp:=string(s[i])
		stack  = append(stack,tmp)

		if tmp==")"{
			str:=""
			flag:=0
			for len(stack)!=0{
				top:=stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				str=top+str
				if top=="("{
					flag=1
					break
				}
			}

			if flag==1{
				stack = append(stack,str)	// 说明有对应的')'包围字符串，于是直接入栈
			}else{
				stack = append(stack,str[:len(str)-1])	// 说明没有对应的')'包围字符串，于是把'('剔除后加入栈

			}
		}

	}
	ans:=""
	for i:=len(stack)-1;i>=0;i--{
		if stack[i]!="(" && stack[i]!=")"{
			ans = stack[i]+ans
		}

	}
	return ans

}
func main() {

}


/*
	总结
	1. 这题我是使用栈做的，每次遇到 ')'就进行字符串判断，否则直接加入栈
	   最后把栈的元素都输出就可以了
*/