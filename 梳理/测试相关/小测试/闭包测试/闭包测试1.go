package main

import "fmt"

func adder(initialNumber int) func() int{
	fmt.Println("闭包外 initialNumber 地址:", &initialNumber)
	return func() int{
		initialNumber++
		fmt.Println("闭包内 initialNumber 地址:", &initialNumber)
		return initialNumber
	}
}



func main() {

	adder1 := adder(1)
	// 闭包外 initialNumber 地址: 0xc042056080

	fmt.Println(adder1())
	fmt.Println(adder1())
	// 闭包内 initialNumber 地址: 0xc042056080
	// 2
	// 闭包内 initialNumber 地址: 0xc042056080
	// 3

	adder2 := adder(1)
	// 闭包外 initialNumber 地址: 0xc0420560a8

	fmt.Println(adder2())
	fmt.Println(adder2())
	// 闭包内 initialNumber 地址: 0xc0420560a8
	// 2
	// 闭包内 initialNumber 地址: 0xc0420560a8
	// 3


	initialNumber := 0
	fmt.Println("main函数 initialNumber 地址:", &initialNumber)
	// main函数 initialNumber 地址: 0xc0420560c0

	adder3 := adder(1)
	// 闭包外 initialNumber 地址: 0xc0420560c8

	fmt.Println(adder3())
	// 闭包内 initialNumber 地址: 0xc0420560c8
	// 2
	fmt.Println(adder3())
	// 闭包内 initialNumber 地址: 0xc0420560c8
	// 3
}

/*
	总结
	1. 由结果可以看出，闭包引用的变量 就是 产生闭包函数的变量，而不是重新在闭包内部再创建一个名字与闭包外相同的变量。
*/