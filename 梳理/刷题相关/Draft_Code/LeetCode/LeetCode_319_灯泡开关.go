package main

import (
	"fmt"
	"math"
)

func get(n int) int{
	ans:=0
	for i:=1;i*i<=n ;i++  {
		if n%i==0 {
			if n/i==i {
				ans++
			}else{
					ans = ans+2
			}
		}
	}
	return ans
}

func bulbSwitch(n int) int {
	//dp:=make([]int,0)
	//dp = append(dp, 0)
	//for i:=1;i<=n ;i++  {
	//	dp = append(dp,dp[i-1]+get(i)%2)
	//}
	//fmt.Println(dp)
	x:=math.Sqrt(float64(n))



	return int(x)
}
func main() {
	fmt.Println( bulbSwitch(999999))
}
/*
	总结
	1. 第一次模拟的时候TLE了
	2. 然后找出了规律
	3. 看上面代码直接提交 0.0.。
	4. 其实考察了数论，可以看官方的解答 0.0. 完全平方数的不同因子为奇数个 0.0.
*/