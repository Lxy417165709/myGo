package main

import "fmt"

func dieSimulator(n int, rollMax []int) int {
	ans = 0
	solve(n,rollMax)
	return ans
}
var ans int
func solve(n int,rollMax []int) {
}

func main() {
	fmt.Println(dieSimulator(2,[]int{
		1,1,2,2,2,3,
	}))
}
