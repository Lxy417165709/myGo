package main

import "fmt"

// 对于方阵可以，转置
func transfer(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < i; t++ {
			matrix[i][t], matrix[t][i] = matrix[t][i], matrix[i][t]
		}
	}
}

// 对于矩阵可以，横向镜像
func mirror(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for t := 0; t < len(matrix[i])/2; t++ {
			matrix[i][t], matrix[i][len(matrix[i])-1-t] = matrix[i][len(matrix[i])-1-t], matrix[i][t]
		}
	}
}

// 顺时针旋转90°
func rotate(matrix [][]int) {
	transfer(matrix)
	mirror(matrix)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

func show(matrix [][]int){
	for i:=0;i< len(matrix);i++  {
		for t:=0;t<len(matrix[i]) ;t++  {
			fmt.Print(matrix[i][t]," ")
		}
		fmt.Println()
	}
}


func main() {
	//m:=[][]int{
	//	{0,1,2,3,4},{5,6,7,8,9},{10,11,12,13,14},{15,16,17,18,19},{20,21,22,23,24},
	//}

}
