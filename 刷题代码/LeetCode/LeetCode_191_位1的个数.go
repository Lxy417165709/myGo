package main

import "fmt"

func hammingWeight(num uint32) int {
	i:=0

	for ;num!=0 ;i++  {
		num = num &(num-1)
	}
	return i

}
func main() {
	//fmt.Println(hammingWeight(9))
	//fmt.Println(hammingWeight(2))
	//fmt.Println(hammingWeight(10))
	fmt.Println(hammingWeight(87678677))
}

/*
	总结
	1. num = num &(num-1) 可以把num的最低位去掉 比如1010去掉后变1000
	2. 这是一道简单的位运算题

*/
