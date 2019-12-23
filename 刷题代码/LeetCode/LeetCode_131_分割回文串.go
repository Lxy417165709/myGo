package main


func partition(s string) [][]string {
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
	if len(s)==0{
		return false
	}
	for i:=0;i<len(s)/2;i++{
		if s[i]!=s[len(s)-1-i]{
			return false
		}
	}
	return true
}

/*
	总结
	1. 第一次做的时候想的太复杂了...
	2. 之后看了官方题解后，用回溯AC了
	3. 这题还可以优化，在回文判断处
*/
