package main

import "fmt"



func main(){
	fmt.Println('\U0000c3a4')
	fmt.Println('\U000000e4')
	fmt.Println('\uc3a4')
	fmt.Println('\u00e4')

	fmt.Println("\U0000c3a4")
	fmt.Println("\U000000e4")
	fmt.Println("\uc3a4")
	fmt.Println("\u00e4")
	// 50084
	// 228
	// 50084
	// 228
	// 쎤
	// ä
	// 쎤
	// ä

	// 总结
	// P17 有错误

}
