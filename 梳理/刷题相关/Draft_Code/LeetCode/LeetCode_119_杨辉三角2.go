package main

import "fmt"

func getRow(rowIndex int) []int {

	slice:=make([]int,0)
	slice = append(slice, 1)

	for i:=1;i<=rowIndex ;i++  {
		length:=len(slice)
		for t:=length-1;t>=1 ;t--  {
			slice[t] = slice[t]+slice[t-1]
		}
		slice = append(slice, 1)
	}
	return slice
	
}
func main() {
	fmt.Println(getRow(33))
}
/*
	总结
	1. 这是一道简答题
	2. 在获取下一行的值的时候，必须从后往前遍历，从前到后遍历的话，由于slice[t] = slice[t]+slice[t-1]
	   这会导致后面的数据不正确
*/