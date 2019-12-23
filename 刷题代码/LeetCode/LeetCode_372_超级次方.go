package main

import "fmt"

func superPow(a int, b []int) int {

	ans:=1
	for i:=len(b)-1;i>=0;i--  {
		num:=1
		// 该位数字是多少
		for t:=0;t<b[i] ;t++  {
			num = num*a%1337
		}
		ans = ans * num%1337

		// 进位
		d:=1
		for t:=0;t<10 ;t++  {
			d = d * a%1337
		}
		a = d%1337

	}
	return ans
}

func main() {
	fmt.Println(superPow(133786,[]int{1,0}))
}
/*
	总结
	1. 该题我的解法： 2的[1,3,3,7]次方，等于2^1000 * (2^100)^3 + (2^10)^3 + (2^1)^7
	2. 按照上面的推导，先求出2^1，之后累乘多9个2^1就变成了2^10,同理可获取2^100

*/