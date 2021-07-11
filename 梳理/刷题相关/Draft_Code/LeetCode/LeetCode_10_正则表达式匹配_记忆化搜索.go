package main

var mp map[string] bool
func isMatch(s string, p string) bool {
	mp=make(map[string]bool)
	return solve(s,p)
}
func solve(s string, p string) bool {

	// 记忆化搜索
	key:=s + "|" + p

	if x,ok:=mp[key];ok{
		return x
	}

	if s==p{
		return true
	}
	if p==""{
		return s==""
	}
	end:=len(p)-1
	if s==""{
		if p[end]=='*'{

			return solve(s,p[:end-1])
		}
		return false
	}

	// 接下来是s,p都非空时
	ends,endp:=len(s)-1,len(p)-1

	ans:=false
	if s[ends]==p[endp] || p[endp]=='.'{
		ans = solve(s[:ends],p[:endp])
	}else{
		if p[endp]=='*'{
			if p[endp-1]==s[ends] || p[endp-1]=='.'{
				ans = solve(s,p[:endp]) || solve(s[:ends],p) || solve(s,p[:endp-1])
			}else{
				ans = solve(s,p[:endp-1])
			}
		}
	}

	mp[key]=ans
	return mp[key]
}
/*
	总结
	1. 第一次写的时候也是用记忆化搜索AC的，但是那个时候是看了官方答案的
	2. 这次独立AC了 0.0
*/