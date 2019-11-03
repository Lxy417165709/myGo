package main

import (
	"fmt"
	"strconv"
)
var ans []string
func restoreIpAddresses(s string) []string {

	ans = make([]string,0)	// 注意初始化
	solve(s,3,"")
	return ans
}


func solve(s string,lay int,temp string){

	if lay==0 {
		// 判断前导0
		if len(s)>1 && s[0]=='0' {
			return
		}
		// 防止过长无法转换
		if len(s)>0 && len(s)<4 {
			num,_ := strconv.Atoi(s)
			if num>=0 && num<=255 {
				temp = temp  + s
				ans = append(ans, temp)
			}
		}
		return
	}

	for i:=1;i< len(s);i++  {

		num,_ := strconv.Atoi(s[0:i])
		if num>=0 && num<=255 {
			t := temp

			// 判断前导0
			if len(s[0:i])>1 && s[0]=='0' {
				return
			}

			temp = temp + s[0:i] + "."
			solve(s[i:],lay - 1,temp)
			temp = t
		}else{
			return
		}
	}

}

func main() {
	fmt.Println(restoreIpAddresses("010010"))
}

/*
	总结
	1. 该题用了回溯法
	2. 注意IP地址的每一段不能有前导零
	3. 如果有全局变量，注意在调用函数的时候先初始化，否则该全局变量会带有上次的结果
       比如第一次答案是[1]，第二次答案是[2]，如果没初始化，输出的话是[1,2]
*/