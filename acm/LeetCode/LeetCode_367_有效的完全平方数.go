package main

import "fmt"

func isPerfectSquare(num int) bool {
	l:=0.0
	r:=50000.0
	mid:=(l+r)/2
	for i:=0;i<100 ;i++  {
		mid=(l+r)/2
		if mid*mid>float64(num) {
			r = mid
		}else{
			l = mid
		}
	}

	if mid-float64(int64(mid)) <= 1e-8 {
		return true
	}else{
		return false
	}
}

func main() {
	x:=0
	for ;true ;  {
		fmt.Scan(&x)
		fmt.Println(isPerfectSquare(x))
	}
}

/*
	总结
	1. 这是一道简单题，我用二分法做出来了
	2. 官方解法3：公式法
		利用 1+3+5+7+9+…+(2n-1)=n^2，即完全平方数肯定是前n个连续奇数的和

*/