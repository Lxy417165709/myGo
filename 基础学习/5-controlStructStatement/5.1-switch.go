package main

import "fmt"

func main(){
	a := 5
	// 这个相当于c++的switch语句,每个case语句添加一个break
	switch  a {
		case 5:
			fmt.Println("a == 5")
		case 4:
			fmt.Println("a == 4")
		default:
			fmt.Println("a != 4 and a != 5")


	}
	b := 9
	// 这个才相当于c++的switch语句
	switch b {
		case 10:
			fmt.Println("b大于等于10")
			fallthrough
		case 9:
			fmt.Println("b大于等于9")
			fallthrough
		case 8:
			fmt.Println("b大于等于8")
			fallthrough
		case 7:
			fmt.Println("b大于等于7")
			fallthrough
		case 6:
			fmt.Println("b大于等于6")
	}

}