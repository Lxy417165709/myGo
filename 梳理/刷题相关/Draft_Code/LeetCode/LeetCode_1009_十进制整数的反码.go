package main

import "fmt"

// 自己写的
//func bitwiseComplement(N int) int {
//	str:=""
//	if N==0 {
//		return 1
//	}
//	for ;N!=0 ;  {
//		if N%2==1 {
//			str = "0"+str
//		}else{
//			str ="1"+str
//		}
//		N = N/2
//	}
//	ans:=0
//	for i:= 0;i<len(str) ;i++  {
//		ans = ans*2 + int(str[i]-'0')
//	}
//	return ans
//
//}

// 异或运算法
func bitwiseComplement(N int) int {

	if N==0 {
		return 1
	}

	temp1,temp2 := 1,N
	for ;temp2>0 ;  {
		N^=temp1
		temp1 = temp1<<1
		temp2 = temp2>>1
	}
	return N

}


func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
	fmt.Println(bitwiseComplement(0))
}
/*
	总结
	1. 这是一道简单的位运算，官方有更简洁的写法
	2. 在上面我拷贝可使用异或AC的方法
	3. 官方还有高位差值法，比如6的二进制是 110,那他的反码 001 = 1000(比6多一位) - 110 - 1

*/