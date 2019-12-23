package main

import "fmt"
func isPowerOfFour(num int) bool {
	if num==0{
		return false
	}
	if num==1{
		return true
	}
	if num%4==0 {
		return isPowerOfFour(num/4)
	}
	return false
}


func main() {
	fmt.Println(isPowerOfFour(0))
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(9))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(81))

}


/*
	总结
	1. 三个题都可以用形似的代码
	2. 要注意考虑数值为0的情况
	3. 位运算也可以解答这题，而且更快
*/