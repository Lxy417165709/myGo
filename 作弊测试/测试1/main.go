package main

import "fmt"

func main() {
	value := 100
	for {
		x := 0
		fmt.Scanf("%d\n",&x)
		switch  {
		case x==1:
			value--
			fmt.Println(value)
		default:
			fmt.Println("i don't know")
		}

	}
}
