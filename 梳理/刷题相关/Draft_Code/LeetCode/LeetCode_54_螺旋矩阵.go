package main

import "fmt"

func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0 {
		return []int{}
	}
	height:= len(matrix)
	width:=len(matrix[0])
	mn :=min(height,width)

	for i:=1;i<=mn-1 ;i++  {
		x0,x1,y0,y1:=0,width-i-1,i,height-1
		// 转y
		for t:=y0;t<=(y1+y0)/2 ;t++  {
			for j:=x0;j<=x1 ;j++  {
				matrix[t][j],matrix[y1+y0-t][j] = matrix[y1+y0-t][j],matrix[t][j]
			}
		}
		// 转x
		fmt.Println(y0,y1)

		for t:=x0;t<=(x1+x0)/2 ;t++  {
			for j:=y0;j<=y1 ;j++  {
				matrix[j][t],matrix[j][x1+x0-t] = matrix[j][x1+x0-t],matrix[j][t]
			}
		}


	}
	ans:=make([]int,0)
	for d:=0;d<mn ;d++  {
		for i:=0;i<width-d ;i++  {
			ans = append(ans, matrix[d][i])
		}
		for i:=1+d;i<height ;i++  {
			ans = append(ans, matrix[i][width-1-d])
		}
	}
	return ans
}


func main() {
	fmt.Println(spiralOrder([][]int{

	}))
}
/*
	总结
	1. 我先对矩阵的左下角部分水平垂直翻转,之后再输出，  比如
	   1 2 3
       4 5 6
       7 8 9   我先对  4 5
					   7 8 翻转，得到 8 7
									  5 4 ,之后对 5 翻转得 5 ，所以矩阵变为 1 2 3
																			8 7 6
																		    5 4 9，之后再把第一行倒数第一列、第二行倒数第三列、第二行倒数第三列分别加入ans切片
*/