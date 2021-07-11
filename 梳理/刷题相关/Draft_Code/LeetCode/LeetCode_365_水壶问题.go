package main

import "fmt"

func gcd(a,b int) int{
	if	b==0{
		return a
	}
	return gcd(b,a%b)
}

func canMeasureWater(x int, y int, z int) bool {
	if z==0 {
		return true
	}
	if x==0 && y==0 {
		return false
	}

	a:= (z-x) % gcd(x,y)
	b:=(z-y) % gcd(x,y)
	 return (a==0 || b==0) && x+y>=z

}
func main() {
	fmt.Println(canMeasureWater(0,0,0))
}
/*
	总结
	1. 这题我自己想了后，看了算法笔记的拓展欧几里得做出来了(期间WA一次，因为没考虑到0的情况)
	2. 这题不单单考了拓展欧几里得，而且要结合自身情况，比如两个水壶的总容量小于Z时肯定是false
	3. 要看看欧几里得加深理解 0.0.
*/