package main

import (
	"fmt"
	"strconv"
)

var m map[string]int
func evalRPN(tokens []string) int {

	if len(tokens)==0 {
		return 0
	}
	m = make(map[string]int)
	m["+"] = 1
	m["-"] = 2
	m["*"] = 3
	m["/"] = 4

	stack_num :=make([]int,0)
	for i:=0;i< len(tokens);i++  {
		x:=tokens[i]
		if m[x]==0 {
			num,_:=strconv.Atoi(x)
			stack_num = append(stack_num, num)
		}else{
			num1 :=stack_num[len(stack_num)-1]
			num2 :=stack_num[len(stack_num)-2]
			stack_num = stack_num[:len(stack_num)-2]
			if m[x]==1 {
				stack_num = append(stack_num, num2+num1)
			}
			if m[x]==2 {
				stack_num = append(stack_num, num2-num1)
			}
			if m[x]==3 {
				stack_num = append(stack_num, num2*num1)
			}
			if m[x]==4 {
				stack_num = append(stack_num, num2/num1)
			}
		}
	}

	return stack_num[0]
}



func main() {
	fmt.Println(evalRPN(
		[]string{

		},
	))
}

/*

	总结
	1. 这是一道简单的栈题，之前做过了~ 0.0..
*/