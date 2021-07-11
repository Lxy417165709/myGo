package main

import "fmt"

func main() {
	var value = new(int)

	for {
		x := 0
		fmt.Scanf("%d\n", &x)
		switch {
		case x == 1:
			*value--
			fmt.Println(*value)
		case x == 2:
			value = new(int)
		case x == 3:
			fmt.Println("i don't know")
		}

	}
}
