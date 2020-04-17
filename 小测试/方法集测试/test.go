package main

import (
	"fmt"
	"net/rpc"
)

func foo(par *int) {
	fmt.Print(par)
}
func main(){
	client,_:=rpc.Dial("x","")
	client.Call()
}
