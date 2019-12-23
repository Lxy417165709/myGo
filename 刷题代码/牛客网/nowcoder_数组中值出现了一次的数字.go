
package main

import (
"fmt"
)

func main(){

	n,num:=0,0
	fmt.Scan(&n)
	slice := make([]int,0)
	x:=0
	for i:=0;i<n;i++{
		fmt.Scan(&num)
		slice =append(slice,num)
		x = x^num

	}
	lowbit := x & (-x)
	num1,num2:=0,0
	for i:=0;i<n;i++{
		if slice[i]&lowbit != 0{
			num1^=slice[i]
		}else{
			num2^=slice[i]
		}
	}
	if num1>num2{
		num1,num2  = num2,num1
	}
	fmt.Println(num1,num2)

}
/*
	总结
	1. 这题目涉及位运算，涉及知识点 x^x^y == y（2个相同的数异或为0）
		以及   x&(-x)== k,   k 为x最低位的1对应的值
	2. 思路是：假设2个数是num1,num2, 则数组输入后,我们可以得到 x=num1^num2
		之后获取x的最低位（两个数在这位上一定不同），之后再扫描一次数组
       把该位是0的和该位是1的分开，最后就可以得到num1,num2了
*/