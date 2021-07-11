package main

import "fmt"

func intToRoman(num int) string {
	m := make(map[int]string)
	m[1] = "I"
	m[4] = "IV"
	m[5] = "V"
	m[9] = "IX"
	m[10] = "X"
	m[40] = "XL"
	m[50] = "L"
	m[90] = "XC"
	m[100] = "C"
	m[400] = "CD"
	m[500] = "D"
	m[900] = "CM"
	m[1000] = "M"

	arr:= []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}

	ans:=""
	for _,a := range arr{
		for ;a<=num ;  {
			num = num - a
			ans = ans + m[a]
		}
	}

	return ans
}


func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(1994))
}

/*
	总结
	1. 这题目使用了整合思维
*/