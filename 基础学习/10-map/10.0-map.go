package main

import "fmt"

func main(){
	// 定义map
	var map1 map[string]string		// 这个没有空间
	var map2 = make(map[string]string)	// 申请空间
	fmt.Println(map1,map2)
	// map[] map[]
	// map1["1"] = "李学悦"		-> 报错,因为没空间
	map2["1"] = "李学悦"
	map2["2"] = "哇哈哈"
	fmt.Println(map1,map2)
	// map[1:李学悦 2:哇哈哈]

	//delete函数
	delete(map2,"1")
	fmt.Println(map2)
	// map2[1]= 		->空


	fmt.Println("map2[1]=",map2["1"])
	// map[2:哇哈哈]
}

// map读入写出和python字典一样