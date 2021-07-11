package main

import "fmt"

func generate(numRows int) [][]int {
	slice:=make([][]int,0)
	if numRows==0 {
		return slice
	}
	slice = append(slice, []int{1})



	if numRows==1 {
		return slice
	}
	slice = append(slice, []int{1,1})
	if numRows==2 {
		return slice
	}

	for i:=2;i<numRows ;i++  {
		if len(slice)==i {
			slice = append(slice, []int{1})
		}
		for t:=1;t< len(slice[i-1]);t++  {
			slice[i] = append(slice[i],slice[i-1][t] + slice[i-1][t-1])
		}
		slice[i] = append(slice[i], 1)

	}
	fmt.Println(slice)
	return slice
}


func main() {
	fmt.Println(generate(0))
}

/*
	总结
	1. 这是一道的简单杨辉三角题
	2. 注意slice[i][t]，要保证slice[i][t]空间已经开了
*/