package main

import "fmt"

func function(slice []int){
	slice = append(slice, 4)
	slice[0] = 5
}
func main() {
	slice1 := []int{1,2,3}
	function(slice1)

	slice2 := make([]int,4)
	copy(slice2,slice1)

	fmt.Print(slice2)
	// 请问输出是什么？
	// A. [1 2 3 4]
	// B. [5 2 3 4]
	// C. [5 2 3 0]
	// D. [1 2 3 0]
}
