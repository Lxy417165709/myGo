package main

import "fmt"

func main(){
	fmt.Println(max(1,2,3))
}
func max(arr ...int) int{
	if len(arr)==0{
		return -1000000000000
	}
	if len(arr) == 1{
		return arr[0]
	}

	a,b := arr[0],max(arr[1:]...)
	if a > b{
		return a
	}
	return b

}
