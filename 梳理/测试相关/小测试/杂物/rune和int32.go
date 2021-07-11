package 杂物

import (
	"fmt"
	"reflect"
	"sort"
)
type newType []int
func main() {
	var a rune = '我'
	var b int32 = '你'

	fmt.Println(a,b,reflect.TypeOf(a),reflect.TypeOf(b))
	// 25105 20320 int32 int32



	//arr1:=make([]rune,0)
	//arr2:=make([]int32,0)
	arr3:=make(newType,0)

	//sort.Ints([]int(arr1))	// 无法转换
	//sort.Ints([]int(arr2))	// 无法转换


	arr3 = append(arr3, 1)
	arr3 = append(arr3, 6)
	arr3 = append(arr3, 5)
	sort.Ints(arr3)	// 可以.
	fmt.Println(arr3)
	// 1 5 6


}
/*
	总结
	1. rune本质是int32
	2. []int32的切片不能转为[]int切片 0.0
*/