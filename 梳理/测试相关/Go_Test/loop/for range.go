package main

import "fmt"

func main() {

    s := []int{1,2,3}
    for i,v:= range s{
        fmt.Printf("%p %p %p %d %d %d\n",&i,&v,&s[i],i,v,s[i])
    }
    for i:=0;i<len(s);i++{
        fmt.Printf("%p %p %d %d\n",&i,&s[i],i,s[i])
    }

    // 0xc0000a0090 0xc0000a0098 0xc00009e140 0 1 1
    // 0xc0000a0090 0xc0000a0098 0xc00009e148 1 2 2
    // 0xc0000a0090 0xc0000a0098 0xc00009e150 2 3 3
    // 0xc0000a0100 0xc00009e140 0 1
    // 0xc0000a0100 0xc00009e148 1 2
    // 0xc0000a0100 0xc00009e150 2 3
}
