package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	return solve(s)
}

func solve(s string) string{
	fmt.Println(s)
	stack:=make([]byte,0)
	for i:=0;i<len(s);i++{
		if s[i]!=')'{
			stack = append(stack,s[i])
		}else{

			if len(stack)!=0{
				x:=""
				stack = append(stack,s[i])
				flag:=0
				for len(stack)!=0{
					char:=stack[len(stack)-1]
					x = string(char) + x
					stack = stack[:len(stack)-1]
					if char=='('{
						flag=1
						break
					}
				}
				if flag==1{
					ans:=solve(x[1:len(x)-1])
					stack = append(stack,'(')
					for t:=0;t< len(ans);t++{
						stack = append(stack,ans[t])
					}
					stack = append(stack,')')
				}else{
					for t:=0;t< len(x)-1;t++{
						stack = append(stack,x[t])
					}
				}

			}
		}

	}
	return string(stack)

}

func main() {
	fmt.Println(minRemoveToMakeValid("))ab(cd(ee)bb(wwe))ea)))"))
}
