package main

import "fmt"

func main() {
	slice3 := []int{1,2,3}
	slice4:= make([]int,3)
	copy(slice4, slice3)
	slice3[0] = 5
	fmt.Println("通过等号赋值把slice3[0]改为5后")
	fmt.Println("	slice3的内容: ",slice3)
	fmt.Println("	slice4的内容: ",slice4)

	slice4[1] = 10
	fmt.Println("通过等号赋值把slice4[1]改为10后")
	fmt.Println("	slice3的内容: ",slice3)
	fmt.Println("	slice4的内容: ",slice4)

	// 通过等号赋值把slice3[0]改为5后
	//	slice3的内容:  [5 2 3]
	//	slice4的内容:  [1 2 3]
	// 通过等号赋值把slice4[1]改为10后
	//	slice3的内容:  [5 2 3]
	//	slice4的内容:  [1 10 3]
}
