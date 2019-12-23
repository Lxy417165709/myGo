package main

import (
	"fmt"
	"strconv"
)

func clumsy(N int) int {
	ans:=N
	str:=strconv.Itoa(N)
	i,num:=0,N-1
	for ;num>0 ;i,num = i+1,num-1  {
		op := i%4
		if op==0 {
			ans = ans*num
			str = str + " * " + strconv.Itoa(num)
		}
		if op==1 {
			ans = ans/num
			str = str + " / " + strconv.Itoa(num)
		}
		if op==2 {
			ans = ans+num
			str = str + " + " + strconv.Itoa(num)
		}
		if op==3 {
			y:=num
			str = str + " - " + strconv.Itoa(num)
			num=num-1
			if num>0 {
				i++
				y*=num
				str = str + " * " + strconv.Itoa(num)
				num=num-1
				if num>0 {
					i++
					y/=num
					str = str + " / " + strconv.Itoa(num)
				}
			}
			ans = ans - y

		}
	}

	//fmt.Println(str)
	return ans

}

func main() {
	x:=0
	for ;true ;  {
		fmt.Scan(&x)
		fmt.Println(clumsy(x))
	}
	fmt.Println(clumsy(x))
}

/*
	总结
	1. 该题我是用模拟做的，在模拟的过程中发现了一些规律 比如 12 -11*10/9 == 0
	2. 官方有规律总结
*/



