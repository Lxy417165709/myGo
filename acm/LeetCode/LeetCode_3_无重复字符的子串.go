package main




func lengthOfLongestSubstring(s string) int {
	l:=0
	ans:=0
	mp :=make(map[uint8] int)
	for r:=0;r<len(s);r++{
		if mp[s[r]]==0{
			ans = max(ans,r-l+1)
			mp[s[r]]=1
		}else{
			mp[s[r]]++
			for mp[s[l]]==1{
				mp[s[l]] = 0
				l++
			}
			mp[s[l]] = 1
			l++

		}
	}
	return ans
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func main() {



}
/*
	总结
	1. 我是采用尺取法做的，每次遇到有相同的字符，则左指针一直右移，直到没有重复
       而右指针一直前移
*/