package main

import "fmt"

func main(){
	//var a bool = true
	var b int = 10
	// b = int(a) 这个无法转换

	var c float64 = float64(b)
	fmt.Printf("%f\n",c)
	// 10.000000

	c = 12.5
	b = int(c)
	fmt.Printf("%d\n",b)
	// 12
	b = 15

	//var s = string(c)	// 这个不能转
	var s = string(b)
	fmt.Printf("转换后: %s\n",s)
	// 转换后: 	->空
}
