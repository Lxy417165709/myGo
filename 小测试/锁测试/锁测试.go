package main

import (
	"fmt"
	"sync"
)

func main() {

	mutex2 := sync.Mutex{}
	nowNumber := 0
	//empty := sync.RWMutex{}

	// 写
	for i:=0;;i++{
		go func (x int){
			mutex2.Lock()
			nowNumber = x
			mutex2.Unlock()
		}(i)

		go func (){
			mutex2.Lock()
			fmt.Println(nowNumber)
			mutex2.Unlock()
		}()
	}
	sync.Map{}
}