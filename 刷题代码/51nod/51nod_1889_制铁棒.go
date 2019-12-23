package _1nod

import (
	"fmt"
	"math"
)

// 法1 不用浮点数
//var L [100500]int
//func main() {
//	n,m:=0,0
//	fmt.Scan(&n,&m)
//	a:=0.0
//	for i:=1;i<=n ;i++  {
//		fmt.Scan(&a)
//		L[i] = int(a*100)
//
//	}
//
//
//
//	l:=0
//	r:=10000001
//	out:
//	for t:=1;t<=100 ;t++  {
//		mid:= (l+r)/2
//
//		gs:=0
//		for i:=1;i<=n ;i++  {
//			if mid <1{
//				break out
//			}
//
//			gs+=L[i]/mid
//		}
//		if gs != m {
//			if gs > m {
//				l = mid
//			}else{
//				r = mid
//			}
//		}else{
//			l = mid
//		}
//		//fmt.Println(l,r)
//	}
//	fmt.Printf("%.2f\n",float64(l)/100)
//
//}


// 法2 用浮点数
var L [100500]float64
func main() {
	n,m:=0,0
	fmt.Scan(&n,&m)
	for i:=1;i<=n ;i++  {
		fmt.Scan(&L[i])

	}



	l:=0.0
	r:=100000.01
out:
	for t:=1;t<=100 ;t++  {
		mid:= math.Trunc((l+r)/2 * 1e2)/1e2

		gs:=0
		for i:=1;i<=n ;i++  {
			if mid==0.00{
				break out
			}
			gs+= int(math.Floor(float64(L[i])/mid))
		}
		if gs != m {
			if gs > m {
				l = mid
			}else{
				r = mid
			}
		}else{
			l = mid
		}
		//fmt.Println(l,r)
	}
	fmt.Printf("%.2f\n",float64(l))

}



/*
	总结
	1. 该题要十分注意精度问题!保留2位小数!
	2. 我们可以通过乘以某个10的次方的数，来把小数划掉，比如 0.01*100 = 1
	3. 解题可以通过列方程找思路，逼近思维也是一种二分思维的拓展
*/