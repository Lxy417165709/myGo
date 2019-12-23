package main


var mp map[string]bool
func partition(s string) [][]string {
	mp = make(map[string]bool)
	dp := [1050][1050]bool{}

	// 回文串判断
	for length:=1;length<=len(s);length++{
		for begin:=0; begin + length - 1<len(s);begin++{
			end:=begin+length - 1
			dp[begin][end]= (s[begin]==s[end]) && ((length<=2) || dp[begin+1][end-1])
			mp[s[begin:end+1]]=dp[begin][end]
		}
	}


	return solve(s)
}

func solve(s string) [][]string{
	if len(s)==0{
		return [][]string{{}}
	}

	ans := [][]string{}
	for i:=1;i<=len(s);i++{
		tmpS:=s[:i]
		if isHW(tmpS){
			after:=solve(s[i:])

			for t:=0;t<len(after);t++{
				ans = append(ans,[]string{tmpS})
				ans[len(ans)-1] = append(ans[len(ans)-1],after[t]...)
			}
		}
	}

	return ans
}


func isHW(s string)bool{
	return mp[s]
}


/*
	总结
	1. 这个加了回文串记录..但是时间和空间都比原来的大...
*/
