package main

func isMatch(s string, p string) bool {

	return solve(s, p)
}
func solve(s string, p string) bool {

	mp := make(map[[2]int]bool)
	mp[[2]int{-1,-1}] = true

	// s为空时的判断
	for i:=1;i<len(p) && p[i]=='*';i=i+2{
		mp[[2]int{i,-1}] = true
	}
	for i:=0;i<len(p);i++{
		for t:=0;t<len(s);t++{

			if p[i]=='.' {
				mp[[2]int{i,t}] = mp[[2]int{i-1,t-1}]
				continue
			}

			if p[i]==s[t] {
				mp[[2]int{i,t}] = mp[[2]int{i-1,t-1}]
				continue
			}

			if p[i]=='*'{
				if p[i-1]==s[t] || p[i-1]=='.'{
					mp[[2]int{i,t}] = mp[[2]int{i-2,t}] || mp[[2]int{i-2,t-1}] || mp[[2]int{i,t-1}]
				}else{
					mp[[2]int{i,t}] = mp[[2]int{i-2,t}]
				}
			}

		}
	}

	return mp[[2]int{len(p)-1,len(s)-1}]
}


/*
	总结
	1. 这是没用递归AC的代码
	2. 这题很麻烦的一点就是，S==""时候还要进行判断
	3. 我们可以使用map代替dp数组，这样可以防止数组越界...
	4. 对于该dp的意思是，p[:i+1]与s[:t+1]是否匹配，其中如果i+1==0 || t+1==0 表示空字符串
*/