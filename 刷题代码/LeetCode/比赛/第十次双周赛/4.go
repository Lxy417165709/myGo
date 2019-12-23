package main

import "fmt"
func max(a,b int) int{
	if a>b {
		return a
	}
	return b
}
func pd(left string,right string) int{
	right = rev(right)
	fmt.Println(left,right)
	dp:=[1050][1050]int{
	}
	for i:=0;i< len(left);i++  {
		for t:=0;t< len(right);t++  {
			if left[i]==right[t] {
				dp[i+1][t+1] = dp[i][t]+1
			}else{
				dp[i+1][t+1] = max(dp[i][t+1],dp[i+1][t])
			}
		}
	}
	return dp[len(left)][len(right)]

}

func rev(s string) string{
	k:=[]byte(s)
	for i:=0;i< len(k)/2;i++  {
		k[i],k[len(k)-1-i] = k[len(k)-1-i],k[i]
	}
	return string(k)
}

func isValidPalindrome(s string, k int) bool {
	for i:=0;i< len(s);i++  {
		x:=pd(s[:i],s[i+1:])
		if k>=len(s)-2*x-1 {
			return true
		}
	}
	for i:=1;i< len(s);i++  {
		if s[i]==s[i-1] {
			x:=pd(s[:i-1],s[i+1:])
			if k>=len(s)-2*x-2 {
				return true
			}
		}

	}


	return false
}
func main() {

	fmt.Println(isValidPalindrome("abcdaaeab",3))
}
