package main

import "fmt"

func search_l_r_1(slice [][]int,target int,) int{


	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid][0] == target {
			l = l + 1
		}else{
			if slice[mid][0] > target{
				r = r -1
			}else{
				l = l + 1
			}
		}
	}
	return r
}

func search_l_r_2(slice []int,target int) int{


	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid] == target {
			l = l + 1
		}else{
			if slice[mid] > target{
				r = r -1
			}else{
				l = l + 1
			}
		}
	}
	return r
}

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix)==0 {
		return false
	}else{
		if len(matrix[0])==0 {
			return false
		}
	}
	x,y:=-1,-1

	x = search_l_r_1(matrix,target)
	if x==-1 {
		return false
	}else{
		y = search_l_r_2(matrix[x],target)
		if y==-1 {
			return false
		}
	}
	if matrix[x][y]==target {
		fmt.Println(x,y)
		return true
	}else{
		return false
	}

}

func main() {
	a:=[][]int{{}}
	fmt.Println(searchMatrix(a,-1))

}


/*
	总结
	1. 这题是二分查找的升级版
	2. 在使用自己的二分查找时，注意所搜索的数组一定要非空！
*/