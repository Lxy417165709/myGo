package main

import (
	"fmt"
	"github.com/Lxy417165709/utils/array"
)

type Int int

func (i Int) Greater(c array.Comparable) bool {
	j := c.(Int)
	return i > j
}

type String string

func (i String) Greater(c array.Comparable) bool {
	j := c.(String)
	return i > j
}

func main() {
	fmt.Println(array.IsExistInSortedArray([]array.Comparable{
		Int(1), Int(2), Int(3),
	}, Int(2)))
	fmt.Println(array.GetFirstGreaterIndex([]array.Comparable{
		String("123"), String("456"), String("789"),
	}, String("456")))
	fmt.Println(array.GetFirstGreaterOrEqualIndex([]array.Comparable{
		String("123"), String("456"), String("789"),
	}, String("456")))
}
