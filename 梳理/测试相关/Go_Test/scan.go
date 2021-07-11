package main

import (
	"bufio"
	"fmt"
	"os"
)

//func main() {
//	n := 0
//	for {
//		count, _ := fmt.Scan(&n)
//		if count == 0 {
//			break
//		}
//		slice := make([]int, 0)
//		for i := 0; i < n; i++ {
//			num := 0
//			fmt.Scan(&num)
//			slice = append(slice, num)
//		}
//		fmt.Println(slice)
//	}
//
//}

func main()  {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input,err := inputReader.ReadString('\n')
		if err != nil{
			return
		}
		fmt.Println(input[:len(input)-1],len(input))
	}
}
