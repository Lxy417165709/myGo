package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	y:=math.Sqrt(2*float64(n)+0.25)-0.5
	return int(y)
}

func main() {
	x:=0
	for ;true ;  {
		fmt.Scan(&x)
		fmt.Println(arrangeCoins(x))
	}

}
/*
	总结
	1. 该题我使用了数学公式，等差数列求和
*/