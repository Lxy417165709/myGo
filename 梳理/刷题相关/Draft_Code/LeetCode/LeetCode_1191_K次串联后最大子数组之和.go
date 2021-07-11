package main

import "fmt"

func max(a,b int) int{
	if	a>b{
		return a
	}
	return b
}
func get(arr []int) int{
	if len(arr)==0 {
		return 0
	}

	dp:=make([]int,0)
	mx:=0
	dp = append(dp, max(0,arr[0]))
	mx = max(mx,dp[0])
	for i:=1;i< len(arr);i++  {
		dp = append(dp, max(0,dp[i-1]+arr[i]))
		mx = max(mx,dp[i])
	}
	return mx
}

func kConcatenationMaxSum(arr []int, k int) int {
	slice:=make([]int, len(arr))
	copy(slice,arr)
	a:=get(slice)
	for i:=0;i< len(arr);i++  {
		slice = append(slice, arr[i])
	}
	b:=get(slice)
	for i:=0;i< len(arr);i++  {
		slice = append(slice, arr[i])
	}
	c:=get(slice)
	fmt.Println(a,b,c)

	if k==1 {
		return a%1000000007
	}
	if k==2 {
		return b%1000000007
	}
	if k==3{
		return c%1000000007

	}



	if a==b {
		return a%1000000007
	}else{
		if a<b && b<c {
			f:=b-a
			ans:=a
			for i:=2;i<=k ;i++  {
				ans+=f
			}
			return ans%1000000007
		}else{
			if k==1 {
				return a%1000000007
			}else{
				return b%1000000007
			}
		}
	}

}

func main() {
	fmt.Println(kConcatenationMaxSum([]int{
		1,1,2,-3,
	},3))
}
/*
	总结
	1. 我这题没做出来(我做出来了！) 0.0.我的想法是获取前3个前缀和，abc通过他们的关系推得k的情况
	2. 其实我做出来了！只是我忘记取模了，我哭了，你呢？..
*/