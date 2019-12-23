package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func reverse(x int) int {
	arr := [20]int{}
	flag := 0

	if x<0 {
		flag=1
		x = int(math.Abs(float64(x)))
	}
	var num int64 = 0
	i:=0
	for ;x!=0 ;i++  {
		arr[i] = x%10
		x = x/10
	}
	for t:=0;t<i ;t++  {
		num = num*10 + int64(arr[t])
	}
	if flag==1 {
		num = -num
	}
	if num >= -(1<<31) && num<=(1<<31-1) {
		return int(num)
	}else{
		return 0
	}

}



func main() {
	//fmt.Println(reverse(-0))
	rand.Float64()
	rand.int
	fmt.Println(strconv.Itoa(-123))
	fmt.Println(strconv.Atoi("-123"))
}

/*

	总结
	1. 这题很简单，但是要注意int32的整数范围
	2. 有时候，取绝对值可以避免减少考虑情况，比如这题目的负数，负号在前面不好转入数组
	3. 2^x == 2<<x == 1<<(x-1)

*/