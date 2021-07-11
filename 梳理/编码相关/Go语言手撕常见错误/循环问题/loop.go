package main

import "fmt"

// 目的: 遍历切片 slice
func problemOne() {
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i <= len(slice); i++ { // mark
		fmt.Println(slice[i])
	}

	// panic: runtime error: index out of range
	// 出现问题: 循环变量取值越界
	// 解决方法: i <= len(slice)   ---->   i < len(slice)
}

// 目的: 逆序遍历切片 slice
func problemTwo() {
	slice := []int{1, 2, 3, 4, 5}
	for i := len(slice) - 1; i >= 0; i++ { // mark
		fmt.Println(slice[i])
	}

	// panic: runtime error: index out of range
	// 出现问题: 循环变量取值越界
	// 解决方法: i++   ---->   i--
}

// 目的: 遍历二维切片 slice
func problemThree() {
	slice := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{9, 10, 11, 12},
	}
	for i := 0; i < len(slice); i++ {
		for t := 0; t < len(slice[i]); i++ { // mark
			fmt.Println(slice[i][t])
		}
	}

	// panic: runtime error: index out of range  (是 i 越界了)
	// 出现问题: 循环变量取值越界
	// 解决方法: i++   ---->   t++
}

func main() {
	problemThree()
}
