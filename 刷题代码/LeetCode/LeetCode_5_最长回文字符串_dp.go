package main

func longestPalindrome(s string) string {

	var dp [1005][1005]int

	mx,str,m:=0,"",len(s)
	for length:=1;length<=m;length++{
		for begin:=0;begin+length-1<m;begin++{
			end:=begin+length-1
			ans:=0
			if s[begin]==s[end]{
				if length<=2{
					ans = length
				}else{
					if dp[begin+1][end-1]==0{
						ans = 0
					}else{
						ans = dp[begin+1][end-1] + 2
					}
				}
			}else{
				ans = 0 // 注意dp数组的含义
			}

			if ans>mx{
				mx = ans
				str = s[begin:end+1]
			}

			dp[begin][end] = ans
		}
	}
	return str
}

/*
	总结
	1. 注意dp数组的含义 dp[begin][end]表示 s[begin:end+1]字符数 (前提是s[begin:end+1]是回文字符串，否则为0)
*/
