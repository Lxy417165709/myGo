package main

import "fmt"

var helloworld = "helloworld"

func LoopAdd(cnt, v0, step int) int {
	result := 0
	for i := 0; i < cnt; i++ {
		result += v0
		v0 += step
	}
	return result
}
func main() {
	fmt.Println(LoopAdd(100, 1, 1))
}
