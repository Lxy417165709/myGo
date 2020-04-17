package main

import "fmt"
type person interface {

}
func main() {
	var slice []int
	var itf interface{}
	var p person
	fmt.Println(slice, itf,p)
	fmt.Println(slice == nil, itf == nil,p!=nil)

}
