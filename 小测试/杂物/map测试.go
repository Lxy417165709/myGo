package 杂物

import "fmt"



func change1(innerMap map[int]int){

	for i:=0;i<10;i++{
		innerMap[i] = i
	}
}
func change2(innerMap map[int]int){
	innerMap = deepCopy(innerMap)	// 多了深复制
	for i:=0;i<10;i++{
		innerMap[i] = i
	}
}


func deepCopy(innerMap map[int]int) map[int]int{

	targetMap :=make(map[int]int)

	for k,v := range innerMap{
		targetMap[k] = v
	}
	return targetMap
}

func main(){
	var testMap map[int]int
	testMap = make(map[int]int)
	fmt.Println("进入函数change1前: ", testMap)
	change1(testMap)
	fmt.Println("进入函数change1后: ", testMap)
	// 进入函数change1前:  map[]
	// 进入函数change1后:  map[4:4 7:7 8:8 9:9 5:5 6:6 0:0 1:1 2:2 3:3]


	testMap = make(map[int]int)
	fmt.Println("进入函数change2前: ", testMap)
	change2(testMap)
	fmt.Println("进入函数change2后: ", testMap)
	// 进入函数change2前:  map[]
	// 进入函数change2后:  map[]

}


/*
	总结
	1. 由于map的传递是引用传递，所以函数体内对map的修改会改变原map
	2. 我们可以使用深复制来避免上面的问题，深复制可以保持原map的信息
*/
