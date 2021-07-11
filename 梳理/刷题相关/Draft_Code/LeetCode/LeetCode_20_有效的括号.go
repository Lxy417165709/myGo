package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	list := list.List{}
	m:=make(map[uint8]uint8)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	for i:=0;i< len(s);i++  {
		if s[i]==')' || s[i]==']' ||s[i]=='}'{

			if list.Len()!=0 {
				x := list.Remove(list.Back())
				fmt.Println(x,m[s[i]])
				if x != m[s[i]] {
					return false
				}
			}else{
				return false
			}

		}else{
			list.PushBack(s[i])

		}
	}

	if list.Len()!=0 {
		return false
	}else{
		return true
	}

}


func main() {
	fmt.Println(isValid("]][["))
}


/*
	总结
	1. 该题考查数据结构--栈
	2. list.New()和list.List{}好像没有区别
*/