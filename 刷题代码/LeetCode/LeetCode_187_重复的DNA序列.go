package main

func findRepeatedDnaSequences(s string) []string {
	mp:=make(map[string]int)
	ans:=make([]string,0)
	for end:=9;end<len(s);end++{
		son:=s[end-9:end+1]
		if mp[son]==1{
			ans = append(ans,son)
		}
		mp[son]++
	}
	return ans
}

/*
	总结
	1. 这题可以使用哈希，通过哈希判断该串是否在之前出现过，之前出现过则加入答案，为了不产生重复
		当出现次数>=3时，我们不把该串加入答案(因为在它出现第二次时已经加过了)
*/