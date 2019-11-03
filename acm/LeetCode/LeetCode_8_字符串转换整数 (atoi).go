package main

import (
	"fmt"
)

func myAtoi(str string) int {
	num:=0
	flag:=0
	for i:=0;i< len(str);i++  {
		if str[i]==' ' {
			continue
		}
		if str[i]=='-' || str[i]=='+'|| str[i]>=48 && str[i]<=57 {
			t:=0
			if str[i]=='-' || str[i]=='+'{
				t = i+1
				if str[i]=='-' {
					flag=1
				}
			} else{
				t =i
			}
			for ;t<len(str) && str[t]>=48 && str[t]<=57;t++  {
				num=num*10 + int(str[t])-48

				if flag==1 {
					if num > (1<<31) {
						return -(1<<31)
					}
				}else{
					if num > (1<<31)-1 {
						return (1<<31)-1
					}
				}
			}
			break
		}else{
			break
		}
	}
	if flag==1 {
		num=-num
	}
	return num
}

func main() {
	fmt.Println(myAtoi("-8878     "))
	fmt.Println(myAtoi( "   -42"))
	fmt.Println(myAtoi( "4193 with words"))
	fmt.Println(myAtoi( "words and 987"))
	fmt.Println(myAtoi("2147483245455445545454545445445"))
}

/*
	总结
	1. 这题目考察细心，要考虑多种情况
	2. 字符串转整数的时候，要注意字符串很长的时候，此时整数装不下
*/