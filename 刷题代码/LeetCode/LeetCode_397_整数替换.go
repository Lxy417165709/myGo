package main

import "fmt"

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
// 内存超出了 0.0..
func integerReplacement(n int) int {
	//dp:=make([]int,2)
	//for i:=2;i<=n ;i++  {
	//	if i&1==0 {
	//		dp = append(dp, dp[i/2]+1)
	//	}else{
	//		dp = append(dp, min(dp[(i-1)/2],dp[(i+1)/2])+2)
	//	}
	//}
	//fmt.Println(dp)
	//return dp[n]

	return solve2(n)
}
// 递归写
func solve(n int) int{
	if n==1 {
		return 0
	}

	if n&1==0 {
		return solve(n/2)+1
	}else{
		return min(solve((n+1)/2),solve((n-1)/2))+2
	}

}

func solve2(n int) int{
	ans:=0
	for n!=1{
		if n&1==0 {
			n>>=1
		}else{
			if n&3==3 {
				if n==3 {
					n = n-1
				}else{
					n = n+1
				}
			}else{
				n = n-1
			}
		}
		ans+=1
	}
	return ans
}

func main() {
	fmt.Println(integerReplacement(100000000))
}
/*
	总结
	1. 第一次用dp迭代写，然后内存超出了 0.0，数据太大了  时间和空间复杂度都是O(n)
	2. 第二次用递归写，然后AC了  但是还是用到了logn的空间0.0..
	3. 官方有时间O(logn)空间O(1)的算法 0.0 我也按照思路实现了，不过要特殊判断3 0.0..
	4. 注意 n&2表示的不是取最后2位比特 0.0.n&3才是
*/