package 杂物

import "fmt"

func main() {
	fmt.Println((-10)<<1)	// -20
	fmt.Println((-11)>>1)	// -6
	fmt.Println((-10)>>1)	// -5
	fmt.Println((-7)<<1)	// -5
	a := -1<<63 + 1	// a是int64的最小值
	fmt.Println(a)		// -9223372036854775807
	fmt.Println(a<<1)	// 0
	fmt.Println(a>>1)	// -4611686018427387904

	fmt.Println((-1) & (-2))	// -2
	fmt.Println((-1) ^ (-2))	// 1
	fmt.Println((-1) | (-2))	// -1

	fmt.Println(uint(1) & uint(2))	// 0
	fmt.Println(uint(1) ^ uint(2))	// 3
	fmt.Println(uint(1) | uint(2))	// 3


	var i uint8  = 20
	fmt.Println(^i,^20)	// 235 -21

}
/*
	总结
	1. 负数左右还是乘2 (前提是没溢出)
	2. 负数右移是除2，并取floor。 ((-11)>>1 == -6 而不是 == -5)
	3. go语言int,uint位运算建立在补码基础上
*/
