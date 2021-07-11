package main

import "strconv"

var mp map[string]int
func numDecodings(s string) int {
	mp = make(map[string]int)
	return solve(s)
}


func solve(s string) int{

	if x,ok:=mp[s];ok{
		return x
	}
	ans:=0
	if len(s)==0{
		ans=1
	}
	if len(s)==1{
		if s!="0"{
			ans=1
		}else{
			ans=0
		}
	}

	if len(s)>=2{
		if s[0]=='0'{
			return 0
		}
		ans=solve(s[1:])
		num,_:=strconv.Atoi(s[:2])
		if num<=26 && num>0{
			ans+=solve(s[2:])
		}

	}

	mp[s]=ans
	return mp[s]
}


/*
	总结
	1. 该题我使用了记忆化搜索AC
	2. 其实这题目和斐波那契数列差不多，也可以用dp解决，并且只用到了O(1)空间复杂度
*/
