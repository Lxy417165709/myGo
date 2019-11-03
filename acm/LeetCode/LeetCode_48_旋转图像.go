package main

import "fmt"

// 对于方阵可以，转置
func transfer(matrix [][]int){
	for i:=0;i< len(matrix);i++  {
		for t:=0;t<i ;t++  {
			matrix[i][t],matrix[t][i] = matrix[t][i],matrix[i][t]
		}
	}
}
// 对于矩阵可以，横向镜像
func mirror(matrix [][]int){
	for i:=0;i< len(matrix);i++  {
		for t:=0;t< len(matrix[i])/2 ;t++  {
			matrix[i][t],matrix[i][len(matrix[i])-1-t] = matrix[i][len(matrix[i])-1-t],matrix[i][t]
		}
	}
}
// 顺时针旋转90°
func rotate(matrix [][]int)  {
	//n:= len(matrix)
	transfer(matrix)
	mirror(matrix)
	for i:=0;i< len(matrix);i++  {
		fmt.Println(matrix[i])
	}
}
func main() {
	rotate([][]int{
	{0},
	})
}

/*
	总结
	1. 第一次做的时候没有思路，只有空间复杂度为O(n^2)的做法
	2. 看了官方题解后使用了官方题解的第一种方法
	3. 官方还有更好的方法，虽然效率提升了一些，但是可读性可能有些差
*/