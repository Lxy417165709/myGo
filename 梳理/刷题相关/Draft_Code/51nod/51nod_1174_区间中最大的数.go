package _1nod

import (
"fmt"
"math"
)

var m [10050][30]float64
var arr [10050]float64

func main() {
	var n int
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	for i:=0;i<20;i++{
		for t:=0;t<n;t++{
			if i==0{
				m[t][i] = arr[t]
			}else{
				next:=t+int(math.Pow(2,float64(i-1)))
				if next >= n{
					next = t
				}
				m[t][i] = math.Max(m[t][i-1],m[next][i-1])
			}
		}
	}
	for i:=0;i<4;i++{
		for t:=0;t<n;t++{
			fmt.Print(m[t][i]," ")
		}
		fmt.Println()
	}
	var a,b,Q int
	fmt.Scan(&Q)
	for i:=0;i<Q ;i++  {
		fmt.Scan(&a,&b)
		var delta = b - a + 1
		var mx = 0.0
		var w = 0
		for ;delta!=0 ;  {
			if delta%2 == 1{
				mx = math.Max(mx,m[a][w])
				a = a + int(math.Pow(2,float64(w)))
			}
			w++
			delta = delta / 2
		}
		fmt.Println(int(mx))
	}

	//fmt.Println(arr)

}


/*
	总结
	1. 这是一道二分思维的拓展，求区间最大值。
	2. 自己手码的，必须夸夸~！！！
*/