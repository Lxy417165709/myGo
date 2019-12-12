package main

import (
	"fmt"
	"time"
)

func main() {

	arr := []int64{1,2,3}
	for _,e := range arr{
		go func(param int64){
			fmt.Println(param)
		}(e)
	}
	time.Sleep(5000)
	// ok!
}
