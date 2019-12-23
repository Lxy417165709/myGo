package main

import "fmt"

func Generator(initialValue int) func()int{

	return func()int{
		initialValue++
		return initialValue
	}
}

func main() {
	generatorA := Generator(0)
	fmt.Println(generatorA())
	fmt.Println(generatorA())
	fmt.Println(generatorA())

	generatorB := Generator(100)

	fmt.Println(generatorB())
	fmt.Println(generatorB())
	fmt.Println(generatorB())

	fmt.Printf("%p %p\n",generatorA,generatorB)

	fmt.Println(generatorA())
	fmt.Println(generatorB())
	fmt.Println(generatorA())
	// 输出
	// 1
	// 2
	// 3
	// 101
	// 102
	// 103
	// 0x48d570 0x48d570
	// 4
	// 104
	// 5
}
/*
	总结
	1. 该生成器指向同一地址的代码段
	2. 生成器之间不会互相影响
	3. 生成器本质是闭包，闭包 = 环境 + 函数
*/