package main

import "fmt"

func main(){
	// 1.一维数组创建
	// 1.0 数组创建方式1
	var array1 [3]int
	for i:=0;i<len(array1);i++{
		fmt.Println(array1[i])
	}
	// 0
	// 0
	// 0

	// 1.1 数组创建方式2
	var array2 = [...]string{"老李","老白","老黑","老鬼"}
	for index,value := range(array2){
		fmt.Printf("%d %s\n",index,value)
	}
	// 0 老李
	// 1 老白
	// 2 老黑
	// 3 老鬼

	// 2.二维数组创建
	myArr := [3][4]int {
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
	}
	fmt.Println(myArr)
	fmt.Println(myArr[0])
	// [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
	// [1 2 3 4]
}

/*
	总结
	1. 数组的长度在声明的时候已经固定了
	2. go语言数组的操作和C++一样
	3. 二维数组中,每个元素都是一维数组 (n维数组中,每个元素都是(n-1)维数组)
*/
