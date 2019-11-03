package main

import "fmt"

func removeOuterParentheses(S string) string {
	slice:= []int{}
	ans:=""
	b:=0
	for i:=0;i< len(S);i++  {
		if S[i]=='(' {
			slice = append(slice, int(S[i]))
		}else{
			slice = slice[0:len(slice)-1]

			if len(slice)==0{
				ans+=S[b+1:i]
				b=i+1
			}
		}
	}
	fmt.Println(ans)
	return ans
}


func main() {
	fmt.Println(removeOuterParentheses(""))
}

/*
	总结
	1. 该题是一道简单的栈题~
*/