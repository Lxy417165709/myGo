package main

import "fmt"

var mp map[string]bool

func isMatch(s string, p string) bool {
	mp = make(map[string]bool)
	return solve(s, p)
}
func solve(s string, p string) bool {
	if x, ok := mp[p+" "+s]; ok {
		return x
	}

	if p == "" {
		mp[p+" "+s] = s == ""
		return mp[p+" "+s]
	}
	flag := false
	if s != "" {
		if s[0] == p[0] || p[0] == '.' {
			flag = true
		}
	}
	if len(p) >= 2 && p[1] == '*' {
		mp[p+" "+s] = solve(s, p[2:]) || flag && solve(s[1:], p)	// 匹配0次和匹配1次
		return mp[p+" "+s]
	} else {
		mp[p+" "+s] = flag && solve(s[1:], p[1:])	// flag为了短路.
		return mp[p+" "+s]
	}

}
func main() {
	fmt.Println(isMatch("abc", "a*bc"))
}

/*
	总结
	1. 第一次做太麻烦了...
	2. 看了官方题解后，采取了暴力递归的方式做，之后用了备忘录优化了下时间复杂度，但是还是提升不了多少（可能是字符串规模太小了.）
	3. 这题目我要好好理解
*/
