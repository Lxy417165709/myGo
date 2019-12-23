package 二分

import "fmt"

func mySqrt(x int) int {
	var t1 int64= 0
	var t2 int64 = 2147483648
	var mid int64= 0
	var a = int64(x)
	for i:=0;i<100 ;i++  {
		mid=(t1+t2)/2
		if mid*mid==a {
			break
		}else{
			if mid*mid > a {
				t2 = mid
			}else{
				t1 = mid
			}
		}
	}
	return int(mid)
}

func main() {
	fmt.Println(mySqrt(2<<31))
	fmt.Println(mySqrt(80))
}
/*
	总结
	1. 该题解答用到了2分思维
	2. 官方题解提供了2分法以及牛顿迭代法

*/